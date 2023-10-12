// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: rpc.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

public enum Land_Gno_Gnomobile_V1_KeyType: SwiftProtobuf.Enum {
  public typealias RawValue = Int
  case typeLocal // = 0
  case typeLedger // = 1
  case typeOffline // = 2
  case typeMulti // = 3
  case UNRECOGNIZED(Int)

  public init() {
    self = .typeLocal
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .typeLocal
    case 1: self = .typeLedger
    case 2: self = .typeOffline
    case 3: self = .typeMulti
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .typeLocal: return 0
    case .typeLedger: return 1
    case .typeOffline: return 2
    case .typeMulti: return 3
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Land_Gno_Gnomobile_V1_KeyType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Land_Gno_Gnomobile_V1_KeyType] = [
    .typeLocal,
    .typeLedger,
    .typeOffline,
    .typeMulti,
  ]
}

#endif  // swift(>=4.2)

///----------------
/// Special errors
///----------------
public enum Land_Gno_Gnomobile_V1_ErrCode: SwiftProtobuf.Enum {
  public typealias RawValue = Int

  /// default value, should never be set manually
  case undefined // = 0

  /// indicates that you plan to create an error later
  case todo // = 1

  /// indicates that a method is not implemented yet
  case errNotImplemented // = 2

  /// indicates an unknown error (without Code), i.e. in gRPC
  case errInternal // = 3
  case errInvalidInput // = 100
  case errBridgeInterrupted // = 101
  case errMissingInput // = 102
  case errSerialization // = 103
  case errDeserialization // = 104
  case errCryptoKeyTypeUnknown // = 105
  case errCryptoKeyNotFound // = 106
  case errNoActiveAccount // = 107
  case errRunGrpcserver // = 108
  case errDecryptionFailed // = 109
  case UNRECOGNIZED(Int)

  public init() {
    self = .undefined
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .undefined
    case 1: self = .todo
    case 2: self = .errNotImplemented
    case 3: self = .errInternal
    case 100: self = .errInvalidInput
    case 101: self = .errBridgeInterrupted
    case 102: self = .errMissingInput
    case 103: self = .errSerialization
    case 104: self = .errDeserialization
    case 105: self = .errCryptoKeyTypeUnknown
    case 106: self = .errCryptoKeyNotFound
    case 107: self = .errNoActiveAccount
    case 108: self = .errRunGrpcserver
    case 109: self = .errDecryptionFailed
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .undefined: return 0
    case .todo: return 1
    case .errNotImplemented: return 2
    case .errInternal: return 3
    case .errInvalidInput: return 100
    case .errBridgeInterrupted: return 101
    case .errMissingInput: return 102
    case .errSerialization: return 103
    case .errDeserialization: return 104
    case .errCryptoKeyTypeUnknown: return 105
    case .errCryptoKeyNotFound: return 106
    case .errNoActiveAccount: return 107
    case .errRunGrpcserver: return 108
    case .errDecryptionFailed: return 109
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Land_Gno_Gnomobile_V1_ErrCode: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static let allCases: [Land_Gno_Gnomobile_V1_ErrCode] = [
    .undefined,
    .todo,
    .errNotImplemented,
    .errInternal,
    .errInvalidInput,
    .errBridgeInterrupted,
    .errMissingInput,
    .errSerialization,
    .errDeserialization,
    .errCryptoKeyTypeUnknown,
    .errCryptoKeyNotFound,
    .errNoActiveAccount,
    .errRunGrpcserver,
    .errDecryptionFailed,
  ]
}

#endif  // swift(>=4.2)

public struct Land_Gno_Gnomobile_V1_KeyInfo {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var type: Land_Gno_Gnomobile_V1_KeyType = .typeLocal

  public var name: String = String()

  public var pubKey: Data = Data()

  public var address: Data = Data()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Land_Gno_Gnomobile_V1_ListKeyInfoRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Land_Gno_Gnomobile_V1_ListKeyInfoResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var keys: [Land_Gno_Gnomobile_V1_KeyInfo] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Land_Gno_Gnomobile_V1_CreateAccountRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var nameOrBech32: String = String()

  public var mnemonic: String = String()

  public var bip39Passwd: String = String()

  public var password: String = String()

  public var account: UInt32 = 0

  public var index: UInt32 = 0

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Land_Gno_Gnomobile_V1_CreateAccountResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var key: Land_Gno_Gnomobile_V1_KeyInfo {
    get {return _key ?? Land_Gno_Gnomobile_V1_KeyInfo()}
    set {_key = newValue}
  }
  /// Returns true if `key` has been explicitly set.
  public var hasKey: Bool {return self._key != nil}
  /// Clears the value of `key`. Subsequent reads from it will return its default value.
  public mutating func clearKey() {self._key = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _key: Land_Gno_Gnomobile_V1_KeyInfo? = nil
}

public struct Land_Gno_Gnomobile_V1_SelectAccountRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var nameOrBech32: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Land_Gno_Gnomobile_V1_SelectAccountResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var key: Land_Gno_Gnomobile_V1_KeyInfo {
    get {return _key ?? Land_Gno_Gnomobile_V1_KeyInfo()}
    set {_key = newValue}
  }
  /// Returns true if `key` has been explicitly set.
  public var hasKey: Bool {return self._key != nil}
  /// Clears the value of `key`. Subsequent reads from it will return its default value.
  public mutating func clearKey() {self._key = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _key: Land_Gno_Gnomobile_V1_KeyInfo? = nil
}

public struct Land_Gno_Gnomobile_V1_GetActiveAccountRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Land_Gno_Gnomobile_V1_GetActiveAccountResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var key: Land_Gno_Gnomobile_V1_KeyInfo {
    get {return _key ?? Land_Gno_Gnomobile_V1_KeyInfo()}
    set {_key = newValue}
  }
  /// Returns true if `key` has been explicitly set.
  public var hasKey: Bool {return self._key != nil}
  /// Clears the value of `key`. Subsequent reads from it will return its default value.
  public mutating func clearKey() {self._key = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _key: Land_Gno_Gnomobile_V1_KeyInfo? = nil
}

public struct Land_Gno_Gnomobile_V1_HelloRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var name: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Land_Gno_Gnomobile_V1_HelloResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var greeting: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Land_Gno_Gnomobile_V1_ErrDetails {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var codes: [Land_Gno_Gnomobile_V1_ErrCode] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

#if swift(>=5.5) && canImport(_Concurrency)
extension Land_Gno_Gnomobile_V1_KeyType: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_ErrCode: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_KeyInfo: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_ListKeyInfoRequest: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_ListKeyInfoResponse: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_CreateAccountRequest: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_CreateAccountResponse: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_SelectAccountRequest: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_SelectAccountResponse: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_GetActiveAccountRequest: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_GetActiveAccountResponse: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_HelloRequest: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_HelloResponse: @unchecked Sendable {}
extension Land_Gno_Gnomobile_V1_ErrDetails: @unchecked Sendable {}
#endif  // swift(>=5.5) && canImport(_Concurrency)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "land.gno.gnomobile.v1"

extension Land_Gno_Gnomobile_V1_KeyType: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "TypeLocal"),
    1: .same(proto: "TypeLedger"),
    2: .same(proto: "TypeOffline"),
    3: .same(proto: "TypeMulti"),
  ]
}

extension Land_Gno_Gnomobile_V1_ErrCode: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "Undefined"),
    1: .same(proto: "TODO"),
    2: .same(proto: "ErrNotImplemented"),
    3: .same(proto: "ErrInternal"),
    100: .same(proto: "ErrInvalidInput"),
    101: .same(proto: "ErrBridgeInterrupted"),
    102: .same(proto: "ErrMissingInput"),
    103: .same(proto: "ErrSerialization"),
    104: .same(proto: "ErrDeserialization"),
    105: .same(proto: "ErrCryptoKeyTypeUnknown"),
    106: .same(proto: "ErrCryptoKeyNotFound"),
    107: .same(proto: "ErrNoActiveAccount"),
    108: .same(proto: "ErrRunGRPCServer"),
    109: .same(proto: "ErrDecryptionFailed"),
  ]
}

extension Land_Gno_Gnomobile_V1_KeyInfo: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".KeyInfo"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "type"),
    2: .same(proto: "name"),
    3: .standard(proto: "pub_key"),
    4: .same(proto: "address"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.type) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.name) }()
      case 3: try { try decoder.decodeSingularBytesField(value: &self.pubKey) }()
      case 4: try { try decoder.decodeSingularBytesField(value: &self.address) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.type != .typeLocal {
      try visitor.visitSingularEnumField(value: self.type, fieldNumber: 1)
    }
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 2)
    }
    if !self.pubKey.isEmpty {
      try visitor.visitSingularBytesField(value: self.pubKey, fieldNumber: 3)
    }
    if !self.address.isEmpty {
      try visitor.visitSingularBytesField(value: self.address, fieldNumber: 4)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_KeyInfo, rhs: Land_Gno_Gnomobile_V1_KeyInfo) -> Bool {
    if lhs.type != rhs.type {return false}
    if lhs.name != rhs.name {return false}
    if lhs.pubKey != rhs.pubKey {return false}
    if lhs.address != rhs.address {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_ListKeyInfoRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ListKeyInfoRequest"
  public static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_ListKeyInfoRequest, rhs: Land_Gno_Gnomobile_V1_ListKeyInfoRequest) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_ListKeyInfoResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ListKeyInfoResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "keys"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeRepeatedMessageField(value: &self.keys) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.keys.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.keys, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_ListKeyInfoResponse, rhs: Land_Gno_Gnomobile_V1_ListKeyInfoResponse) -> Bool {
    if lhs.keys != rhs.keys {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_CreateAccountRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CreateAccountRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "name_or_bech32"),
    2: .same(proto: "mnemonic"),
    3: .standard(proto: "bip39_passwd"),
    4: .same(proto: "password"),
    5: .same(proto: "account"),
    6: .same(proto: "index"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.nameOrBech32) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.mnemonic) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.bip39Passwd) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self.password) }()
      case 5: try { try decoder.decodeSingularUInt32Field(value: &self.account) }()
      case 6: try { try decoder.decodeSingularUInt32Field(value: &self.index) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.nameOrBech32.isEmpty {
      try visitor.visitSingularStringField(value: self.nameOrBech32, fieldNumber: 1)
    }
    if !self.mnemonic.isEmpty {
      try visitor.visitSingularStringField(value: self.mnemonic, fieldNumber: 2)
    }
    if !self.bip39Passwd.isEmpty {
      try visitor.visitSingularStringField(value: self.bip39Passwd, fieldNumber: 3)
    }
    if !self.password.isEmpty {
      try visitor.visitSingularStringField(value: self.password, fieldNumber: 4)
    }
    if self.account != 0 {
      try visitor.visitSingularUInt32Field(value: self.account, fieldNumber: 5)
    }
    if self.index != 0 {
      try visitor.visitSingularUInt32Field(value: self.index, fieldNumber: 6)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_CreateAccountRequest, rhs: Land_Gno_Gnomobile_V1_CreateAccountRequest) -> Bool {
    if lhs.nameOrBech32 != rhs.nameOrBech32 {return false}
    if lhs.mnemonic != rhs.mnemonic {return false}
    if lhs.bip39Passwd != rhs.bip39Passwd {return false}
    if lhs.password != rhs.password {return false}
    if lhs.account != rhs.account {return false}
    if lhs.index != rhs.index {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_CreateAccountResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".CreateAccountResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "key"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._key) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._key {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_CreateAccountResponse, rhs: Land_Gno_Gnomobile_V1_CreateAccountResponse) -> Bool {
    if lhs._key != rhs._key {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_SelectAccountRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".SelectAccountRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "name_or_bech32"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.nameOrBech32) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.nameOrBech32.isEmpty {
      try visitor.visitSingularStringField(value: self.nameOrBech32, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_SelectAccountRequest, rhs: Land_Gno_Gnomobile_V1_SelectAccountRequest) -> Bool {
    if lhs.nameOrBech32 != rhs.nameOrBech32 {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_SelectAccountResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".SelectAccountResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "key"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._key) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._key {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_SelectAccountResponse, rhs: Land_Gno_Gnomobile_V1_SelectAccountResponse) -> Bool {
    if lhs._key != rhs._key {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_GetActiveAccountRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GetActiveAccountRequest"
  public static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_GetActiveAccountRequest, rhs: Land_Gno_Gnomobile_V1_GetActiveAccountRequest) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_GetActiveAccountResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".GetActiveAccountResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "key"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._key) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._key {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_GetActiveAccountResponse, rhs: Land_Gno_Gnomobile_V1_GetActiveAccountResponse) -> Bool {
    if lhs._key != rhs._key {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_HelloRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".HelloRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "name"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.name) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.name.isEmpty {
      try visitor.visitSingularStringField(value: self.name, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_HelloRequest, rhs: Land_Gno_Gnomobile_V1_HelloRequest) -> Bool {
    if lhs.name != rhs.name {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_HelloResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".HelloResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "greeting"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.greeting) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.greeting.isEmpty {
      try visitor.visitSingularStringField(value: self.greeting, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_HelloResponse, rhs: Land_Gno_Gnomobile_V1_HelloResponse) -> Bool {
    if lhs.greeting != rhs.greeting {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Land_Gno_Gnomobile_V1_ErrDetails: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".ErrDetails"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "codes"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeRepeatedEnumField(value: &self.codes) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.codes.isEmpty {
      try visitor.visitPackedEnumField(value: self.codes, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Land_Gno_Gnomobile_V1_ErrDetails, rhs: Land_Gno_Gnomobile_V1_ErrDetails) -> Bool {
    if lhs.codes != rhs.codes {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
