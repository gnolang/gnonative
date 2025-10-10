package land.gno.gnonative

import android.content.Context
import android.content.SharedPreferences
import android.os.Build
import android.security.keystore.KeyGenParameterSpec
import android.security.keystore.KeyProperties
import android.util.Base64
import gnolang.gno.gnonative.NativeDB
import java.io.ByteArrayOutputStream
import java.nio.ByteBuffer
import java.security.KeyStore
import javax.crypto.Cipher
import javax.crypto.KeyGenerator
import javax.crypto.SecretKey
import javax.crypto.spec.GCMParameterSpec
import kotlin.math.min
import androidx.core.content.edit

class NativeDBManager(
  context: Context,
  private val prefsName: String = "gnonative_secure_db",
  private val keyAlias: String = "gnonative_aes_key"
) : NativeDB {

  // -------- storage / index --------
  private val prefs: SharedPreferences =
    context.getSharedPreferences(prefsName, Context.MODE_PRIVATE)
  private val entryPrefix = "kv:"   // entryPrefix + hexKey -> Base64(encrypted blob)
  private val idxKey = "__idx__"    // CSV of hex keys in ascending order

  // -------- crypto --------
  private val ks: KeyStore = KeyStore.getInstance(ANDROID_KEYSTORE).apply { load(null) }

  private val lock = Any()

  init {
    ensureAesKey()
    if (!prefs.contains(idxKey)) prefs.edit { putString(idxKey, "") }
  }

  // ========== NativeDB implementation ==========

  override fun delete(p0: ByteArray?) {
    val key = requireKey(p0)
    val hex = hex(key)
    synchronized(lock) {
      val idx = loadIndexAsc().toMutableList()
      val pos = lowerBound(idx, hex)
      if (pos < idx.size && idx[pos] == hex) {
        idx.removeAt(pos)
        saveIndexAsc(idx)
      }
      prefs.edit { remove("$entryPrefix$hex") }
    }
  }

  override fun deleteSync(p0: ByteArray?) {
    delete(p0)
  }

  override fun get(p0: ByteArray?): ByteArray {
    val key = requireKey(p0)
    val hex = hex(key)
    val b64 = synchronized(lock) { prefs.getString("$entryPrefix$hex", null) }
      ?: return ByteArray(0) // gomobile generated non-null return -> use empty on miss
    val blob = Base64.decode(b64, Base64.NO_WRAP)
    return decrypt(blob) ?: throw Exception("Failed to decrypt value for key: $hex")
  }

  override fun has(p0: ByteArray?): Boolean {
    val key = requireKey(p0)
    val hex = hex(key)
    return synchronized(lock) { prefs.contains("$entryPrefix$hex") }
  }

  override fun scanChunk(
    p0: ByteArray?,   // start
    p1: ByteArray?,   // end
    p2: ByteArray?,   // seekKey
    p3: Long,         // limit
    p4: Boolean       // reverse
  ): ByteArray {
    val limit = if (p3 < 0) 0 else min(p3, Int.MAX_VALUE.toLong()).toInt()
    return synchronized(lock) {
      val asc = loadIndexAsc() // ascending hex keys
      val startHex = p0?.let { hex(it) }
      val endHex   = p1?.let { hex(it) }
      val seekHex  = p2?.let { hex(it) }

      val loBase = startHex?.let { lowerBound(asc, it) } ?: 0
      val hiBase = endHex?.let { lowerBound(asc, it) } ?: asc.size
      var slice: List<String> = if (hiBase <= loBase) emptyList() else asc.subList(loBase, hiBase)

      // seek positioning & direction
      slice = if (!p4) {
        val from = seekHex?.let { upperBound(slice, it) } ?: 0
        if (from >= slice.size) emptyList() else slice.subList(from, slice.size)
      } else {
        val positioned = if (seekHex != null) {
          val idx = upperBound(slice, seekHex) - 1
          if (idx < 0) emptyList() else slice.subList(0, idx + 1)
        } else slice
        positioned.asReversed()
      }

      val page = if (limit == 0) emptyList() else slice.take(limit)
      val hasMore = page.isNotEmpty() && page.size < slice.size
      val nextSeekHex = if (hasMore) page.last() else null

      // materialize kv pairs in traversal order
      val pairs = ArrayList<Pair<ByteArray, ByteArray>>(page.size)
      for (h in page) {
        val b64 = prefs.getString("$entryPrefix$h", null) ?: continue
        val blob = Base64.decode(b64, Base64.NO_WRAP)
        val v = decrypt(blob) ?: throw Exception("Failed to decrypt value for key: $h")
        pairs += (unhex(h) to v)
      }

      // flags(1) | count(u32 BE) | [kLen k vLen v]* | nextSeekLen(u32 BE) | nextSeek
      encodeChunkBlobBE(pairs, nextSeekHex?.let { unhex(it) }, hasMore)
    }
  }

  override fun set(p0: ByteArray?, p1: ByteArray?) {
    val key = requireKey(p0)
    val value = requireValue(p1)
    val hex = hex(key)
    val enc = encrypt(value)
    val b64 = Base64.encodeToString(enc, Base64.NO_WRAP)
    synchronized(lock) {
      val idx = loadIndexAsc().toMutableList()
      val pos = lowerBound(idx, hex)
      if (pos == idx.size || idx[pos] != hex) {
        idx.add(pos, hex)
        saveIndexAsc(idx)
      }
      prefs.edit { putString("$entryPrefix$hex", b64) }
    }
  }

  override fun setSync(p0: ByteArray?, p1: ByteArray?) {
    set(p0, p1)
  }

  // ========== helpers ==========

  private fun requireKey(b: ByteArray?): ByteArray {
    require(!(b == null || b.isEmpty())) { "key must not be null/empty" }
    return b
  }
  private fun requireValue(b: ByteArray?): ByteArray {
    require(b != null) { "value must not be null" }
    return b
  }

  // ----- index (csv of hex keys, ascending) -----
  private fun loadIndexAsc(): List<String> {
    val csv = prefs.getString(idxKey, "") ?: ""
    return if (csv.isEmpty()) emptyList() else csv.split(',').filter { it.isNotEmpty() }
  }
  private fun saveIndexAsc(keys: List<String>) {
    prefs.edit { putString(idxKey, if (keys.isEmpty()) "" else keys.joinToString(",")) }
  }

  private fun lowerBound(list: List<String>, key: String): Int {
    var lo = 0; var hi = list.size
    while (lo < hi) {
      val mid = (lo + hi) ushr 1
      if (list[mid] < key) lo = mid + 1 else hi = mid
    }
    return lo
  }
  private fun upperBound(list: List<String>, key: String): Int {
    var lo = 0; var hi = list.size
    while (lo < hi) {
      val mid = (lo + hi) ushr 1
      if (list[mid] <= key) lo = mid + 1 else hi = mid
    }
    return lo
  }

  // crypto AES/GCM, StrongBox preferred
  private fun ensureAesKey() {
    if (getAesKey() != null) return

    val kg = KeyGenerator.getInstance(KeyProperties.KEY_ALGORITHM_AES, ANDROID_KEYSTORE)
    val base = KeyGenParameterSpec.Builder(
      keyAlias, KeyProperties.PURPOSE_ENCRYPT or KeyProperties.PURPOSE_DECRYPT
    )
      .setBlockModes(KeyProperties.BLOCK_MODE_GCM)
      .setEncryptionPaddings(KeyProperties.ENCRYPTION_PADDING_NONE)
      .setKeySize(256)
      .setRandomizedEncryptionRequired(true)

    try {
      if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.P) {
        base.setIsStrongBoxBacked(true)
      }
      kg.init(base.build())
      kg.generateKey()
      return
    } catch (_: Throwable) {
      // fall back below without StrongBox
    }

    kg.init(
      KeyGenParameterSpec.Builder(
        keyAlias, KeyProperties.PURPOSE_ENCRYPT or KeyProperties.PURPOSE_DECRYPT
      )
        .setBlockModes(KeyProperties.BLOCK_MODE_GCM)
        .setEncryptionPaddings(KeyProperties.ENCRYPTION_PADDING_NONE)
        .setKeySize(256)
        .setRandomizedEncryptionRequired(true)
        .build()
    )
    kg.generateKey()
  }

  private fun getAesKey(): SecretKey? {
    val e = ks.getEntry(keyAlias, null) as? KeyStore.SecretKeyEntry
    return e?.secretKey
  }

  private fun encrypt(plain: ByteArray): ByteArray {
    val key = getAesKey() ?: error("AES key missing")

    val c = Cipher.getInstance(AES_GCM)
    c.init(Cipher.ENCRYPT_MODE, key)

    val iv = c.iv
    val ct = c.doFinal(plain)

    // payload: [version=1][ivLen][iv][ct]
    val out = ByteArray(1 + 1 + iv.size + ct.size)
    var i = 0
    out[i++] = 1
    out[i++] = iv.size.toByte()
    System.arraycopy(iv, 0, out, i, iv.size); i += iv.size
    System.arraycopy(ct, 0, out, i, ct.size)
    return out
  }

  private fun decrypt(blob: ByteArray?): ByteArray? {
    if (blob == null || blob.size < 1 + 1 + 12) return null // iv is usually 12 bytes
    var i = 0
    val ver = blob[i++]
    require(ver.toInt() == 1) { "bad payload version=$ver" }
    val ivLen = blob[i++].toInt() and 0xFF
    require(ivLen in 12..32) { "bad iv length" }
    require(blob.size >= 1 + 1 + ivLen + 1) { "short blob" }
    val iv = ByteArray(ivLen)
    System.arraycopy(blob, i, iv, 0, ivLen); i += ivLen
    val ct = ByteArray(blob.size - i)
    System.arraycopy(blob, i, ct, 0, ct.size)

    val key = getAesKey() ?: error("AES key missing")
    val c = Cipher.getInstance(AES_GCM)
    c.init(Cipher.DECRYPT_MODE, key, GCMParameterSpec(128, iv))
    return c.doFinal(ct)
  }

  // chunk framing (match Go format)
  private fun encodeChunkBlobBE(
    entries: List<Pair<ByteArray, ByteArray>>,
    nextSeek: ByteArray?,
    hasMore: Boolean
  ): ByteArray {
    val bos = ByteArrayOutputStream()

    // flags (bit0 = hasMore)
    bos.write(if (hasMore) 0x01 else 0x00)

    // count (u32 BE)
    bos.write(u32be(entries.size))

    // entries
    for ((k, v) in entries) {
      bos.write(u32be(k.size)); bos.write(k)
      bos.write(u32be(v.size)); bos.write(v)
    }

    // nextSeek
    val ns = nextSeek ?: ByteArray(0)
    bos.write(u32be(ns.size))
    if (ns.isNotEmpty()) bos.write(ns)

    return bos.toByteArray()
  }

  // ----- utils -----
  private fun u32be(n: Int): ByteArray {
    val bb = ByteBuffer.allocate(4)
    bb.putInt(n) // big-endian by default
    return bb.array()
  }

  private fun hex(b: ByteArray): String {
    val out = CharArray(b.size * 2)
    val h = "0123456789abcdef".toCharArray()
    var i = 0
    for (v in b) {
      val x = v.toInt() and 0xFF
      out[i++] = h[x ushr 4]; out[i++] = h[x and 0x0F]
    }
    return String(out)
  }

  private fun unhex(s: String): ByteArray {
    require(s.length % 2 == 0) { "odd hex length" }
    val out = ByteArray(s.length / 2)
    var i = 0; var j = 0
    while (i < s.length) {
      val hi = Character.digit(s[i++], 16)
      val lo = Character.digit(s[i++], 16)
      out[j++] = ((hi shl 4) or lo).toByte()
    }
    return out
  }

  companion object {
    private const val ANDROID_KEYSTORE = "AndroidKeyStore"
    private const val AES_GCM = "AES/GCM/NoPadding"
  }
}
