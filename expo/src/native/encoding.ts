// Base64 <-> bytes/string helpers for the /native path. On the wire, protobuf `bytes` fields are
// base64 strings and the bridge resolves base64(protojson) payloads. These use base64-js and
// fast-text-encoding, which are already dependencies of the gRPC path too.
import { fromByteArray, toByteArray } from 'base64-js';

import 'fast-text-encoding';

// bytesToBase64 encodes a Uint8Array as a base64 string (for protobuf `bytes` request fields).
export function bytesToBase64(bytes: Uint8Array): string {
  return fromByteArray(bytes);
}

// base64ToBytes decodes a base64 string (a protobuf `bytes` response field) into a Uint8Array.
export function base64ToBytes(b64: string): Uint8Array {
  return toByteArray(b64);
}

// base64ToString decodes a base64 string into a UTF-8 string. Used to turn the base64(protojson)
// bridge payload into the JSON text before parsing.
export function base64ToString(b64: string): string {
  return new TextDecoder().decode(toByteArray(b64));
}
