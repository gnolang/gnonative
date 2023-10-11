// @generated by protoc-gen-es v1.3.3
// @generated from file rpc.proto (package land.gno.gnomobile.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum land.gno.gnomobile.v1.KeyType
 */
export declare enum KeyType {
  /**
   * @generated from enum value: TypeLocal = 0;
   */
  TypeLocal = 0,

  /**
   * @generated from enum value: TypeLedger = 1;
   */
  TypeLedger = 1,

  /**
   * @generated from enum value: TypeOffline = 2;
   */
  TypeOffline = 2,

  /**
   * @generated from enum value: TypeMulti = 3;
   */
  TypeMulti = 3,
}

/**
 * ----------------
 * Special errors
 * ----------------
 *
 * @generated from enum land.gno.gnomobile.v1.ErrCode
 */
export declare enum ErrCode {
  /**
   * default value, should never be set manually
   *
   * @generated from enum value: Undefined = 0;
   */
  Undefined = 0,

  /**
   * indicates that you plan to create an error later
   *
   * @generated from enum value: TODO = 1;
   */
  TODO = 1,

  /**
   * indicates that a method is not implemented yet
   *
   * @generated from enum value: ErrNotImplemented = 2;
   */
  ErrNotImplemented = 2,

  /**
   * indicates an unknown error (without Code), i.e. in gRPC
   *
   * @generated from enum value: ErrInternal = 3;
   */
  ErrInternal = 3,

  /**
   * @generated from enum value: ErrInvalidInput = 100;
   */
  ErrInvalidInput = 100,

  /**
   * @generated from enum value: ErrBridgeInterrupted = 101;
   */
  ErrBridgeInterrupted = 101,

  /**
   * @generated from enum value: ErrMissingInput = 102;
   */
  ErrMissingInput = 102,

  /**
   * @generated from enum value: ErrSerialization = 103;
   */
  ErrSerialization = 103,

  /**
   * @generated from enum value: ErrDeserialization = 104;
   */
  ErrDeserialization = 104,

  /**
   * @generated from enum value: ErrCryptoKeyTypeUnknown = 105;
   */
  ErrCryptoKeyTypeUnknown = 105,

  /**
   * @generated from enum value: ErrCryptoKeyNotFound = 106;
   */
  ErrCryptoKeyNotFound = 106,

  /**
   * @generated from enum value: ErrNoActiveAccount = 107;
   */
  ErrNoActiveAccount = 107,

  /**
   * @generated from enum value: ErrRunGRPCServer = 108;
   */
  ErrRunGRPCServer = 108,
}

/**
 * @generated from message land.gno.gnomobile.v1.KeyInfo
 */
export declare class KeyInfo extends Message<KeyInfo> {
  /**
   * @generated from field: land.gno.gnomobile.v1.KeyType type = 1;
   */
  type: KeyType;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: bytes pub_key = 3;
   */
  pubKey: Uint8Array;

  /**
   * @generated from field: bytes address = 4;
   */
  address: Uint8Array;

  constructor(data?: PartialMessage<KeyInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.KeyInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): KeyInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): KeyInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): KeyInfo;

  static equals(a: KeyInfo | PlainMessage<KeyInfo> | undefined, b: KeyInfo | PlainMessage<KeyInfo> | undefined): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.ListKeyInfoRequest
 */
export declare class ListKeyInfoRequest extends Message<ListKeyInfoRequest> {
  constructor(data?: PartialMessage<ListKeyInfoRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.ListKeyInfoRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListKeyInfoRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListKeyInfoRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListKeyInfoRequest;

  static equals(
    a: ListKeyInfoRequest | PlainMessage<ListKeyInfoRequest> | undefined,
    b: ListKeyInfoRequest | PlainMessage<ListKeyInfoRequest> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.ListKeyInfoResponse
 */
export declare class ListKeyInfoResponse extends Message<ListKeyInfoResponse> {
  /**
   * @generated from field: repeated land.gno.gnomobile.v1.KeyInfo keys = 1;
   */
  keys: KeyInfo[];

  constructor(data?: PartialMessage<ListKeyInfoResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.ListKeyInfoResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListKeyInfoResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListKeyInfoResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListKeyInfoResponse;

  static equals(
    a: ListKeyInfoResponse | PlainMessage<ListKeyInfoResponse> | undefined,
    b: ListKeyInfoResponse | PlainMessage<ListKeyInfoResponse> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.CreateAccountRequest
 */
export declare class CreateAccountRequest extends Message<CreateAccountRequest> {
  /**
   * @generated from field: string name_or_bech32 = 1;
   */
  nameOrBech32: string;

  /**
   * @generated from field: string mnemonic = 2;
   */
  mnemonic: string;

  /**
   * @generated from field: string bip39_passwd = 3;
   */
  bip39Passwd: string;

  /**
   * @generated from field: string password = 4;
   */
  password: string;

  /**
   * @generated from field: uint32 account = 5;
   */
  account: number;

  /**
   * @generated from field: uint32 index = 6;
   */
  index: number;

  constructor(data?: PartialMessage<CreateAccountRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.CreateAccountRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAccountRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAccountRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAccountRequest;

  static equals(
    a: CreateAccountRequest | PlainMessage<CreateAccountRequest> | undefined,
    b: CreateAccountRequest | PlainMessage<CreateAccountRequest> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.CreateAccountResponse
 */
export declare class CreateAccountResponse extends Message<CreateAccountResponse> {
  /**
   * @generated from field: land.gno.gnomobile.v1.KeyInfo key = 1;
   */
  key?: KeyInfo;

  constructor(data?: PartialMessage<CreateAccountResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.CreateAccountResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateAccountResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateAccountResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateAccountResponse;

  static equals(
    a: CreateAccountResponse | PlainMessage<CreateAccountResponse> | undefined,
    b: CreateAccountResponse | PlainMessage<CreateAccountResponse> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.SelectAccountRequest
 */
export declare class SelectAccountRequest extends Message<SelectAccountRequest> {
  /**
   * @generated from field: string name_or_bech32 = 1;
   */
  nameOrBech32: string;

  constructor(data?: PartialMessage<SelectAccountRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.SelectAccountRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SelectAccountRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SelectAccountRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SelectAccountRequest;

  static equals(
    a: SelectAccountRequest | PlainMessage<SelectAccountRequest> | undefined,
    b: SelectAccountRequest | PlainMessage<SelectAccountRequest> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.SelectAccountResponse
 */
export declare class SelectAccountResponse extends Message<SelectAccountResponse> {
  /**
   * @generated from field: land.gno.gnomobile.v1.KeyInfo key = 1;
   */
  key?: KeyInfo;

  constructor(data?: PartialMessage<SelectAccountResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.SelectAccountResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SelectAccountResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SelectAccountResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SelectAccountResponse;

  static equals(
    a: SelectAccountResponse | PlainMessage<SelectAccountResponse> | undefined,
    b: SelectAccountResponse | PlainMessage<SelectAccountResponse> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.GetActiveAccountRequest
 */
export declare class GetActiveAccountRequest extends Message<GetActiveAccountRequest> {
  constructor(data?: PartialMessage<GetActiveAccountRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.GetActiveAccountRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetActiveAccountRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetActiveAccountRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetActiveAccountRequest;

  static equals(
    a: GetActiveAccountRequest | PlainMessage<GetActiveAccountRequest> | undefined,
    b: GetActiveAccountRequest | PlainMessage<GetActiveAccountRequest> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.GetActiveAccountResponse
 */
export declare class GetActiveAccountResponse extends Message<GetActiveAccountResponse> {
  /**
   * @generated from field: land.gno.gnomobile.v1.KeyInfo key = 1;
   */
  key?: KeyInfo;

  constructor(data?: PartialMessage<GetActiveAccountResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.GetActiveAccountResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetActiveAccountResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetActiveAccountResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetActiveAccountResponse;

  static equals(
    a: GetActiveAccountResponse | PlainMessage<GetActiveAccountResponse> | undefined,
    b: GetActiveAccountResponse | PlainMessage<GetActiveAccountResponse> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.HelloRequest
 */
export declare class HelloRequest extends Message<HelloRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<HelloRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.HelloRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HelloRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HelloRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HelloRequest;

  static equals(
    a: HelloRequest | PlainMessage<HelloRequest> | undefined,
    b: HelloRequest | PlainMessage<HelloRequest> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.HelloResponse
 */
export declare class HelloResponse extends Message<HelloResponse> {
  /**
   * @generated from field: string greeting = 1;
   */
  greeting: string;

  constructor(data?: PartialMessage<HelloResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.HelloResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HelloResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HelloResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HelloResponse;

  static equals(
    a: HelloResponse | PlainMessage<HelloResponse> | undefined,
    b: HelloResponse | PlainMessage<HelloResponse> | undefined,
  ): boolean;
}

/**
 * @generated from message land.gno.gnomobile.v1.ErrDetails
 */
export declare class ErrDetails extends Message<ErrDetails> {
  /**
   * @generated from field: repeated land.gno.gnomobile.v1.ErrCode codes = 1;
   */
  codes: ErrCode[];

  constructor(data?: PartialMessage<ErrDetails>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "land.gno.gnomobile.v1.ErrDetails";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ErrDetails;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ErrDetails;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ErrDetails;

  static equals(a: ErrDetails | PlainMessage<ErrDetails> | undefined, b: ErrDetails | PlainMessage<ErrDetails> | undefined): boolean;
}
