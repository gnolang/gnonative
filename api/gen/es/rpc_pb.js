// @generated by protoc-gen-es v1.10.0
// @generated from file rpc.proto (package land.gno.gnonative.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * The ErrCode enum defines errors for gRPC API functions. These are converted
 * from the Go error types returned by gnoclient.
 *
 * ----------------
 * Special errors
 * ----------------
 *
 * @generated from enum land.gno.gnonative.v1.ErrCode
 */
export const ErrCode = /*@__PURE__*/ proto3.makeEnum(
  "land.gno.gnonative.v1.ErrCode",
  [
    {no: 0, name: "Undefined"},
    {no: 1, name: "TODO"},
    {no: 2, name: "ErrNotImplemented"},
    {no: 3, name: "ErrInternal"},
    {no: 100, name: "ErrInvalidInput"},
    {no: 101, name: "ErrBridgeInterrupted"},
    {no: 102, name: "ErrMissingInput"},
    {no: 103, name: "ErrSerialization"},
    {no: 104, name: "ErrDeserialization"},
    {no: 105, name: "ErrInitService"},
    {no: 106, name: "ErrSetRemote"},
    {no: 150, name: "ErrCryptoKeyTypeUnknown"},
    {no: 151, name: "ErrCryptoKeyNotFound"},
    {no: 152, name: "ErrNoActiveAccount"},
    {no: 153, name: "ErrRunGRPCServer"},
    {no: 154, name: "ErrDecryptionFailed"},
    {no: 200, name: "ErrTxDecode"},
    {no: 201, name: "ErrInvalidSequence"},
    {no: 202, name: "ErrUnauthorized"},
    {no: 203, name: "ErrInsufficientFunds"},
    {no: 204, name: "ErrUnknownRequest"},
    {no: 205, name: "ErrInvalidAddress"},
    {no: 206, name: "ErrUnknownAddress"},
    {no: 207, name: "ErrInvalidPubKey"},
    {no: 208, name: "ErrInsufficientCoins"},
    {no: 209, name: "ErrInvalidCoins"},
    {no: 210, name: "ErrInvalidGasWanted"},
    {no: 211, name: "ErrOutOfGas"},
    {no: 212, name: "ErrMemoTooLarge"},
    {no: 213, name: "ErrInsufficientFee"},
    {no: 214, name: "ErrTooManySignatures"},
    {no: 215, name: "ErrNoSignatures"},
    {no: 216, name: "ErrGasOverflow"},
    {no: 217, name: "ErrInvalidPkgPath"},
    {no: 218, name: "ErrInvalidStmt"},
    {no: 219, name: "ErrInvalidExpr"},
  ],
);

/**
 * @generated from message land.gno.gnonative.v1.ErrDetails
 */
export const ErrDetails = /*@__PURE__*/ proto3.makeMessageType(
  "land.gno.gnonative.v1.ErrDetails",
  () => [
    { no: 1, name: "codes", kind: "enum", T: proto3.getEnumType(ErrCode), repeated: true },
  ],
);

