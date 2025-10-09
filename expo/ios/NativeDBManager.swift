//
//  NativeDBManager.swift
//  Pods
//
//  Created by Rémi BARBERO on 25/09/2025.
//

import Foundation
import Security
import GnoCore

public class NativeDBManager: NSObject, GnoGnonativeNativeDBProtocol {
    public static var shared: NativeDBManager = NativeDBManager()
    
    // MARK: - Private Properties
    private let service: String
    private let accessGroup: String?
    
    // MARK: - Initialization
    init(service: String = Bundle.main.bundleIdentifier ?? "GnoNativeService", accessGroup: String? = nil) {
        self.service = service
        self.accessGroup = accessGroup
    }
    
    // MARK: - Public Interface Implementation
    
    public func get(_ key: Data?) throws -> Data {
        guard let key = key, !key.isEmpty else {
            throw NativeDBError.invalidArgument(description: "Key must not be nil or empty.")
        }
        let account = keyToAccount(key)
        
        var query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrService as String: service,
            kSecAttrAccount as String: account,
            kSecReturnData as String: true,
            kSecMatchLimit as String: kSecMatchLimitOne
        ]
        
        if let accessGroup = accessGroup {
            query[kSecAttrAccessGroup as String] = accessGroup
        }
        
        var result: AnyObject?
        let status = SecItemCopyMatching(query as CFDictionary, &result)
        
        switch status {
        case errSecSuccess:
            return (result as? Data) ?? Data()
        case errSecItemNotFound:
            // This is not an error; it's the expected result for a missing key.
            return Data()
        default:
            throw NativeDBError.keychainError(status: status, message: "Failed to get item.")
        }
    }
    
    public func delete(_ key: Data?) throws {
        try self.deleteSync(key)
    }
    
    public func deleteSync(_ key: Data?) throws {
        guard let key = key, !key.isEmpty else {
            throw NativeDBError.invalidArgument(description: "Key must not be nil or empty.")
        }
        
        let account = keyToAccount(key)
        
        var query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrService as String: service,
            kSecAttrAccount as String: account
        ]
        
        if let accessGroup = accessGroup {
            query[kSecAttrAccessGroup as String] = accessGroup
        }
        
        let status = SecItemDelete(query as CFDictionary)
        
        // Deleting a non-existent item is not considered an error.
        if status != errSecSuccess && status != errSecItemNotFound {
            throw NativeDBError.keychainError(status: status, message: "Failed to delete item.")
        }
    }
    
    public func has(_ key: Data?, ret0_ returnPointer: UnsafeMutablePointer<ObjCBool>?) throws {
        // 1. Ensure the return pointer provided by the caller is valid.
        guard let returnPointer = returnPointer else {
            throw NativeDBError.invalidArgument(description: "Return pointer must not be nil.")
        }
        
        // 2. Validate the input key.
        guard let key = key, !key.isEmpty else {
            throw NativeDBError.invalidArgument(description: "Key must not be nil or empty.")
        }
        
        let account = keyToAccount(key)
        
        var query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrService as String: service,
            kSecAttrAccount as String: account
        ]
        if let accessGroup = accessGroup {
            query[kSecAttrAccessGroup as String] = accessGroup
        }
        
        let status = SecItemCopyMatching(query as CFDictionary, nil)
        
        // 3. Handle the result from the Keychain.
        switch status {
        case errSecSuccess:
            // Key exists. Write `true` to the pointer's memory location.
            returnPointer.pointee = ObjCBool(true)
        case errSecItemNotFound:
            // Key does not exist. Write `false` to the pointer's memory location.
            returnPointer.pointee = ObjCBool(false)
        default:
            // A real keychain error occurred. Throw an error.
            throw NativeDBError.keychainError(status: status, message: "Failed to check for item existence.")
        }
    }
    
    public func set(_ key: Data?, value: Data?) throws {
        try self.setSync(key, value: value)
    }
    
    public func setSync(_ key: Data?, value: Data?) throws {
        guard let key = key, !key.isEmpty else {
            throw NativeDBError.invalidArgument(description: "Key must not be nil or empty.")
        }
        guard let value = value else {
            throw NativeDBError.invalidArgument(description: "Value must not be nil.")
        }
        
        let account = keyToAccount(key)
        
        // First, try to update existing item
        var query: [String: Any] = [
            kSecClass as String: kSecClassGenericPassword,
            kSecAttrService as String: service,
            kSecAttrAccount as String: account
        ]
        
        if let accessGroup = accessGroup {
            query[kSecAttrAccessGroup as String] = accessGroup
        }
        
        let attributes: [String: Any] = [
            kSecValueData as String: value
        ]
        
        let updateStatus = SecItemUpdate(query as CFDictionary, attributes as CFDictionary)
        
        switch updateStatus {
        case errSecSuccess:
            // Update was successful.
            return
        case errSecItemNotFound:
            // Item doesn't exist, so add it.
            var newItem = query
            newItem[kSecValueData as String] = value
            newItem[kSecAttrAccessible as String] = kSecAttrAccessibleWhenUnlockedThisDeviceOnly
            
            let addStatus = SecItemAdd(newItem as CFDictionary, nil)
            if addStatus != errSecSuccess {
                throw NativeDBError.keychainError(status: addStatus, message: "Failed to add new item.")
            }
        default:
            // Another error occurred during update.
            throw NativeDBError.keychainError(status: updateStatus, message: "Failed to update item.")
        }
    }
    
    public func scanChunk(_ start: Data?, end: Data?, seekKey: Data?, limit: Int, reverse: Bool) throws -> Data {
        // 1) fetch all items for this service
        var query: [String: Any] = [
            kSecClass as String:             kSecClassGenericPassword,
            kSecAttrService as String:       service,
            kSecMatchLimit as String:        kSecMatchLimitAll,
            kSecReturnAttributes as String:  true,
            kSecReturnData as String:        true,
        ]
        
        if let accessGroup = accessGroup {
            query[kSecAttrAccessGroup as String] = accessGroup
        }
        var result: CFTypeRef?
        let status = SecItemCopyMatching(query as CFDictionary, &result)
        
        var pairs: [(key: Data, val: Data)] = []
        if status == errSecSuccess, let items = result as? [[String: Any]] {
            for item in items {
                guard let account = item[kSecAttrAccount as String] as? String,
                      let keyBytes = accountToKey(account),
                      let val  = item[kSecValueData as String] as? Data else { continue }
                if inRange(keyBytes, start: start, end: end) {
                    pairs.append((keyBytes, val))
                }
            }
        }
        
        // 2) sort
        if reverse {
            pairs.sort { a, b in
                // descending: a > b
                return lt(b.key, a.key)
            }
        } else {
            pairs.sort { a, b in
                // ascending: a < b
                return lt(a.key, b.key)
            }
        }
        
        // 3) apply seekKey (exclusive)
        if let sk = seekKey, !sk.isEmpty {
            if reverse {
                // keep items with key < seekKey
                let idx = pairs.firstIndex(where: { lt($0.key, sk) }) ?? pairs.count
                // pairs are descending, so drop while key >= seekKey
                pairs = Array(pairs[idx...])
            } else {
                // keep items with key > seekKey
                let idx = pairs.lastIndex(where: { lte($0.key, sk) }) ?? -1
                let startIdx = idx + 1
                pairs = (startIdx < pairs.count) ? Array(pairs[startIdx...]) : []
            }
        }
        
        // 4) limit
        let lim = max(0, Int(limit))
        let chunk = (lim > 0 && lim < pairs.count) ? Array(pairs.prefix(lim)) : pairs
        let hasMore = chunk.count < pairs.count
        let nextSeek = chunk.last?.key ?? Data()
        
        // 6) Frame the blob
        var blob = Data(capacity: 1 + 4) // will grow as needed
        var flags: UInt8 = 0
        if hasMore { flags |= 0x01 }
        blob.append(&flags, count: 1)
        
        var countBE = UInt32(chunk.count).bigEndian
        withUnsafeBytes(of: &countBE) { blob.append($0.bindMemory(to: UInt8.self)) }
        
        for (k, v) in chunk {
            var klen = UInt32(k.count).bigEndian
            var vlen = UInt32(v.count).bigEndian
            withUnsafeBytes(of: &klen) { blob.append($0.bindMemory(to: UInt8.self)) }
            blob.append(k)
            withUnsafeBytes(of: &vlen) { blob.append($0.bindMemory(to: UInt8.self)) }
            blob.append(v)
        }
        
        var nlen = UInt32(nextSeek.count).bigEndian
        withUnsafeBytes(of: &nlen) { blob.append($0.bindMemory(to: UInt8.self)) }
        blob.append(nextSeek)
        
        return blob
    }
    
    // MARK: - Utility Methods
    
    private func keyToAccount(_ key: Data) -> String {
        return String(data: key, encoding: .utf8) ?? key.base64EncodedString()
    }
    
    private func accountToKey(_ account: String) -> Data? {
        // Try UTF-8 first
        if let utf8Data = account.data(using: .utf8) {
            // Check if this was originally a base64 string by trying to decode it
            if let base64Data = Data(base64Encoded: account), base64Data != utf8Data {
                // This account was base64 encoded, return the decoded data
                return base64Data
            } else {
                // This was a UTF-8 string, return the UTF-8 data
                return utf8Data
            }
        }
        
        // Fallback: try base64 decoding
        return Data(base64Encoded: account)
    }
    
    // --- byte-wise comparisons on decoded keys ---
    @inline(__always)
    private func lt(_ a: Data, _ b: Data) -> Bool {
        a.lexicographicallyPrecedes(b)
    }
    @inline(__always)
    private func gte(_ a: Data, _ b: Data) -> Bool { !lt(a, b) }
    @inline(__always)
    private func lte(_ a: Data, _ b: Data) -> Bool { !lt(b, a) }
    @inline(__always)
    private func inRange(_ k: Data, start: Data?, end: Data?) -> Bool {
        if let s = start, lt(k, s) { return false }    // k >= s
        if let e = end, !lt(k, e) { return false }     // k < e
        return true
    }
}

public enum NativeDBError: Error, LocalizedError {
    case invalidArgument(description: String)
    case keychainError(status: OSStatus, message: String)
    
    public var errorDescription: String? {
        switch self {
        case .invalidArgument(let description):
            return "Invalid Argument: \(description)"
        case .keychainError(let status, let message):
            let statusMessage = SecCopyErrorMessageString(status, nil) as String? ?? "Unknown error"
            return "Keychain Error: \(message) (status: \(status), reason: \(statusMessage))"
        }
    }
}
