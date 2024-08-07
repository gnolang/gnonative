// @generated by protoc-gen-es v2.0.0
// @generated from file rpc.proto (package land.gno.gnonative.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { AddressFromBech32RequestSchema, AddressFromBech32ResponseSchema, AddressFromMnemonicRequestSchema, AddressFromMnemonicResponseSchema, AddressToBech32RequestSchema, AddressToBech32ResponseSchema, BroadcastTxCommitRequestSchema, BroadcastTxCommitResponseSchema, CallRequestSchema, CallResponseSchema, CreateAccountRequestSchema, CreateAccountResponseSchema, DeleteAccountRequestSchema, DeleteAccountResponseSchema, GenerateRecoveryPhraseRequestSchema, GenerateRecoveryPhraseResponseSchema, GetActiveAccountRequestSchema, GetActiveAccountResponseSchema, GetChainIDRequestSchema, GetChainIDResponseSchema, GetKeyInfoByAddressRequestSchema, GetKeyInfoByAddressResponseSchema, GetKeyInfoByNameOrAddressRequestSchema, GetKeyInfoByNameOrAddressResponseSchema, GetKeyInfoByNameRequestSchema, GetKeyInfoByNameResponseSchema, GetRemoteRequestSchema, GetRemoteResponseSchema, HasKeyByAddressRequestSchema, HasKeyByAddressResponseSchema, HasKeyByNameOrAddressRequestSchema, HasKeyByNameOrAddressResponseSchema, HasKeyByNameRequestSchema, HasKeyByNameResponseSchema, HelloRequestSchema, HelloResponseSchema, HelloStreamRequestSchema, HelloStreamResponseSchema, ListKeyInfoRequestSchema, ListKeyInfoResponseSchema, MakeTxResponseSchema, QEvalRequestSchema, QEvalResponseSchema, QueryAccountRequestSchema, QueryAccountResponseSchema, QueryRequestSchema, QueryResponseSchema, RenderRequestSchema, RenderResponseSchema, RunRequestSchema, RunResponseSchema, SelectAccountRequestSchema, SelectAccountResponseSchema, SendRequestSchema, SendResponseSchema, SetChainIDRequestSchema, SetChainIDResponseSchema, SetPasswordRequestSchema, SetPasswordResponseSchema, SetRemoteRequestSchema, SetRemoteResponseSchema, SignTxRequestSchema, SignTxResponseSchema } from "./gnonativetypes_pb";

/**
 * Describes the file rpc.proto.
 */
export declare const file_rpc: GenFile;

/**
 * @generated from message land.gno.gnonative.v1.ErrDetails
 */
export declare type ErrDetails = Message<"land.gno.gnonative.v1.ErrDetails"> & {
  /**
   * @generated from field: repeated land.gno.gnonative.v1.ErrCode codes = 1;
   */
  codes: ErrCode[];
};

/**
 * Describes the message land.gno.gnonative.v1.ErrDetails.
 * Use `create(ErrDetailsSchema)` to create a new message.
 */
export declare const ErrDetailsSchema: GenMessage<ErrDetails>;

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
export enum ErrCode {
  /**
   * Undefined is the default value. It should never be set manually
   *
   * @generated from enum value: Undefined = 0;
   */
  Undefined = 0,

  /**
   * TODO indicates that you plan to create an error later
   *
   * @generated from enum value: TODO = 1;
   */
  TODO = 1,

  /**
   * ErrNotImplemented indicates that a method is not implemented yet
   *
   * @generated from enum value: ErrNotImplemented = 2;
   */
  ErrNotImplemented = 2,

  /**
   * ErrInternal indicates an unknown error (without Code), i.e. in gRPC
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
   * @generated from enum value: ErrInitService = 105;
   */
  ErrInitService = 105,

  /**
   * @generated from enum value: ErrSetRemote = 106;
   */
  ErrSetRemote = 106,

  /**
   * @generated from enum value: ErrCryptoKeyTypeUnknown = 150;
   */
  ErrCryptoKeyTypeUnknown = 150,

  /**
   * ErrCryptoKeyNotFound indicates that the doesn't exist in the keybase
   *
   * @generated from enum value: ErrCryptoKeyNotFound = 151;
   */
  ErrCryptoKeyNotFound = 151,

  /**
   * ErrNoActiveAccount indicates that no active account has been set with SelectAccount
   *
   * @generated from enum value: ErrNoActiveAccount = 152;
   */
  ErrNoActiveAccount = 152,

  /**
   * @generated from enum value: ErrRunGRPCServer = 153;
   */
  ErrRunGRPCServer = 153,

  /**
   * ErrDecryptionFailed indicates a decryption failure including a wrong password
   *
   * @generated from enum value: ErrDecryptionFailed = 154;
   */
  ErrDecryptionFailed = 154,

  /**
   * @generated from enum value: ErrTxDecode = 200;
   */
  ErrTxDecode = 200,

  /**
   * @generated from enum value: ErrInvalidSequence = 201;
   */
  ErrInvalidSequence = 201,

  /**
   * @generated from enum value: ErrUnauthorized = 202;
   */
  ErrUnauthorized = 202,

  /**
   * ErrInsufficientFunds indicates that there are insufficient funds to pay for fees
   *
   * @generated from enum value: ErrInsufficientFunds = 203;
   */
  ErrInsufficientFunds = 203,

  /**
   * ErrUnknownRequest indicates that the path of a realm function call is unrecognized
   *
   * @generated from enum value: ErrUnknownRequest = 204;
   */
  ErrUnknownRequest = 204,

  /**
   * ErrInvalidAddress indicates that an account address is blank or the bech32 can't be decoded
   *
   * @generated from enum value: ErrInvalidAddress = 205;
   */
  ErrInvalidAddress = 205,

  /**
   * ErrUnknownAddress indicates that the address is unknown on the blockchain
   *
   * @generated from enum value: ErrUnknownAddress = 206;
   */
  ErrUnknownAddress = 206,

  /**
   * ErrInvalidPubKey indicates that the public key was not found or has an invalid algorithm or format
   *
   * @generated from enum value: ErrInvalidPubKey = 207;
   */
  ErrInvalidPubKey = 207,

  /**
   * ErrInsufficientCoins indicates that the transaction has insufficient account funds to send
   *
   * @generated from enum value: ErrInsufficientCoins = 208;
   */
  ErrInsufficientCoins = 208,

  /**
   * ErrInvalidCoins indicates that the transaction Coins are not sorted, or don't have a
   * positive amount, or the coin Denom contains upper case characters
   *
   * @generated from enum value: ErrInvalidCoins = 209;
   */
  ErrInvalidCoins = 209,

  /**
   * ErrInvalidGasWanted indicates that the transaction gas wanted is too large or otherwise invalid
   *
   * @generated from enum value: ErrInvalidGasWanted = 210;
   */
  ErrInvalidGasWanted = 210,

  /**
   * ErrOutOfGas indicates that the transaction doesn't have enough gas
   *
   * @generated from enum value: ErrOutOfGas = 211;
   */
  ErrOutOfGas = 211,

  /**
   * ErrMemoTooLarge indicates that the transaction memo is too large
   *
   * @generated from enum value: ErrMemoTooLarge = 212;
   */
  ErrMemoTooLarge = 212,

  /**
   * ErrInsufficientFee indicates that the gas fee is insufficient
   *
   * @generated from enum value: ErrInsufficientFee = 213;
   */
  ErrInsufficientFee = 213,

  /**
   * ErrTooManySignatures indicates that the transaction has too many signatures
   *
   * @generated from enum value: ErrTooManySignatures = 214;
   */
  ErrTooManySignatures = 214,

  /**
   * ErrNoSignatures indicates that the transaction has no signatures
   *
   * @generated from enum value: ErrNoSignatures = 215;
   */
  ErrNoSignatures = 215,

  /**
   * ErrGasOverflow indicates that an action results in a gas consumption unsigned integer overflow
   *
   * @generated from enum value: ErrGasOverflow = 216;
   */
  ErrGasOverflow = 216,

  /**
   * ErrInvalidPkgPath indicates that the package path is not recognized.
   *
   * @generated from enum value: ErrInvalidPkgPath = 217;
   */
  ErrInvalidPkgPath = 217,

  /**
   * @generated from enum value: ErrInvalidStmt = 218;
   */
  ErrInvalidStmt = 218,

  /**
   * @generated from enum value: ErrInvalidExpr = 219;
   */
  ErrInvalidExpr = 219,
}

/**
 * Describes the enum land.gno.gnonative.v1.ErrCode.
 */
export declare const ErrCodeSchema: GenEnum<ErrCode>;

/**
 * GnoNativeService is the service to interact with the Gno blockchain
 *
 * @generated from service land.gno.gnonative.v1.GnoNativeService
 */
export declare const GnoNativeService: GenService<{
  /**
   * Set the connection address for the remote node. If you don't call this,
   * the default is "127.0.0.1:26657"
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.SetRemote
   */
  setRemote: {
    methodKind: "unary";
    input: typeof SetRemoteRequestSchema;
    output: typeof SetRemoteResponseSchema;
  },
  /**
   * Get the connection address for the remote node. The response is either
   * the initial default, or the value which was set with SetRemote
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.GetRemote
   */
  getRemote: {
    methodKind: "unary";
    input: typeof GetRemoteRequestSchema;
    output: typeof GetRemoteResponseSchema;
  },
  /**
   * Set the chain ID for the remote node. If you don't call this, the default
   * is "dev"
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.SetChainID
   */
  setChainID: {
    methodKind: "unary";
    input: typeof SetChainIDRequestSchema;
    output: typeof SetChainIDResponseSchema;
  },
  /**
   * Get the chain ID for the remote node. The response is either
   * the initial default, or the value which was set with SetChainID
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.GetChainID
   */
  getChainID: {
    methodKind: "unary";
    input: typeof GetChainIDRequestSchema;
    output: typeof GetChainIDResponseSchema;
  },
  /**
   * Generate a recovery phrase of BIP39 mnemonic words using entropy from the
   * crypto library random number generator. This can be used as the mnemonic in
   * CreateAccount.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.GenerateRecoveryPhrase
   */
  generateRecoveryPhrase: {
    methodKind: "unary";
    input: typeof GenerateRecoveryPhraseRequestSchema;
    output: typeof GenerateRecoveryPhraseResponseSchema;
  },
  /**
   * Get the information for all keys in the keybase
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.ListKeyInfo
   */
  listKeyInfo: {
    methodKind: "unary";
    input: typeof ListKeyInfoRequestSchema;
    output: typeof ListKeyInfoResponseSchema;
  },
  /**
   * Check for the key in the keybase with the given name.
   * In the response, set has true if the keybase has the key.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.HasKeyByName
   */
  hasKeyByName: {
    methodKind: "unary";
    input: typeof HasKeyByNameRequestSchema;
    output: typeof HasKeyByNameResponseSchema;
  },
  /**
   * Check for the key in the keybase with the given address.
   * In the response, set has true if the keybase has the key.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.HasKeyByAddress
   */
  hasKeyByAddress: {
    methodKind: "unary";
    input: typeof HasKeyByAddressRequestSchema;
    output: typeof HasKeyByAddressResponseSchema;
  },
  /**
   * Check for the key in the keybase with the given name or bech32 string address.
   * In the response, set has true if the keybase has the key.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.HasKeyByNameOrAddress
   */
  hasKeyByNameOrAddress: {
    methodKind: "unary";
    input: typeof HasKeyByNameOrAddressRequestSchema;
    output: typeof HasKeyByNameOrAddressResponseSchema;
  },
  /**
   * Get the information for the key in the keybase with the given name.
   * If the key doesn't exist, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrCryptoKeyNotFound.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByName
   */
  getKeyInfoByName: {
    methodKind: "unary";
    input: typeof GetKeyInfoByNameRequestSchema;
    output: typeof GetKeyInfoByNameResponseSchema;
  },
  /**
   * Get the information for the key in the keybase with the given address.
   * If the key doesn't exist, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrCryptoKeyNotFound.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByAddress
   */
  getKeyInfoByAddress: {
    methodKind: "unary";
    input: typeof GetKeyInfoByAddressRequestSchema;
    output: typeof GetKeyInfoByAddressResponseSchema;
  },
  /**
   * Get the information for the key in the keybase with the given name or bech32 string address.
   * If the key doesn't exist, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrCryptoKeyNotFound.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.GetKeyInfoByNameOrAddress
   */
  getKeyInfoByNameOrAddress: {
    methodKind: "unary";
    input: typeof GetKeyInfoByNameOrAddressRequestSchema;
    output: typeof GetKeyInfoByNameOrAddressResponseSchema;
  },
  /**
   * Create a new account in the keybase using the name and password specified by SetAccount.
   * If an account with the same name already exists in the keybase,
   * this replaces it. (If you don't want to replace it, then it's your responsibility
   * to use GetKeyInfoByName to check if it exists before calling this.)
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.CreateAccount
   */
  createAccount: {
    methodKind: "unary";
    input: typeof CreateAccountRequestSchema;
    output: typeof CreateAccountResponseSchema;
  },
  /**
   * SelectAccount selects the active account to use for later operations. If the response has_password is
   * false, then you should set the password before using a method which needs it.
   * If the key doesn't exist, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrCryptoKeyNotFound.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.SelectAccount
   */
  selectAccount: {
    methodKind: "unary";
    input: typeof SelectAccountRequestSchema;
    output: typeof SelectAccountResponseSchema;
  },
  /**
   * Set the password for the active account in the keybase, used for later operations.
   * If no active account has been set with SelectAccount, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrNoActiveAccount.
   * If the password is wrong, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrDecryptionFailed.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.SetPassword
   */
  setPassword: {
    methodKind: "unary";
    input: typeof SetPasswordRequestSchema;
    output: typeof SetPasswordResponseSchema;
  },
  /**
   * GetActiveAccount gets the active account which was set by SelectAccount.
   * If no active account has been set with SelectAccount, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrNoActiveAccount.
   * (To check if there is an active account, use ListKeyInfo and check the
   * length of the result.)
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.GetActiveAccount
   */
  getActiveAccount: {
    methodKind: "unary";
    input: typeof GetActiveAccountRequestSchema;
    output: typeof GetActiveAccountResponseSchema;
  },
  /**
   * QueryAccount retrieves account information from the blockchain for a given
   * address.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.QueryAccount
   */
  queryAccount: {
    methodKind: "unary";
    input: typeof QueryAccountRequestSchema;
    output: typeof QueryAccountResponseSchema;
  },
  /**
   * DeleteAccount deletes the account with the given name, using the password
   * to ensure access. However, if skip_password is true, then ignore the
   * password.
   * If the account doesn't exist, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrCryptoKeyNotFound.
   * If the password is wrong, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrDecryptionFailed.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.DeleteAccount
   */
  deleteAccount: {
    methodKind: "unary";
    input: typeof DeleteAccountRequestSchema;
    output: typeof DeleteAccountResponseSchema;
  },
  /**
   * Make an ABCI query to the remote node.
   * If the request path is unrecognized, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrUnknownRequest.
   * If the request data has a package path that is unrecognized, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrInvalidPkgPath.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.Query
   */
  query: {
    methodKind: "unary";
    input: typeof QueryRequestSchema;
    output: typeof QueryResponseSchema;
  },
  /**
   * Render calls the Render function for package_path with optional args. The
   * package path should include the prefix like "gno.land/". This is similar to
   * using a browser URL <nodeURL>/<pkgPath>:<args> where <pkgPath> doesn't have
   * the prefix like "gno.land/".
   * If the request package_path is unrecognized, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrInvalidPkgPath.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.Render
   */
  render: {
    methodKind: "unary";
    input: typeof RenderRequestSchema;
    output: typeof RenderResponseSchema;
  },
  /**
   * QEval evaluates the given expression with the realm code at package_path.
   * The package path should include the prefix like "gno.land/". The expression
   * is usually a function call like "GetBoardIDFromName(\"testboard\")". The
   * return value is a typed expression like
   * "(1 gno.land/r/demo/boards.BoardID)\n(true bool)".
   * If the request package_path is unrecognized, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrInvalidPkgPath.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.QEval
   */
  qEval: {
    methodKind: "unary";
    input: typeof QEvalRequestSchema;
    output: typeof QEvalResponseSchema;
  },
  /**
   * Call a specific realm function.
   * If no active account has been set with SelectAccount, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrNoActiveAccount.
   * If the password is wrong, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrDecryptionFailed.
   * If the path of a realm function call is unrecognized, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrUnknownRequest.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.Call
   */
  call: {
    methodKind: "server_streaming";
    input: typeof CallRequestSchema;
    output: typeof CallResponseSchema;
  },
  /**
   * Send currency from the active account to an account on the blockchain.
   * If no active account has been set with SelectAccount, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrNoActiveAccount.
   * If the password is wrong, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrDecryptionFailed.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.Send
   */
  send: {
    methodKind: "server_streaming";
    input: typeof SendRequestSchema;
    output: typeof SendResponseSchema;
  },
  /**
   * Temporarily load the code in package on the blockchain and run main() which can
   * call realm functions and use println() to output to the "console".
   * This returns the "console" output.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.Run
   */
  run: {
    methodKind: "server_streaming";
    input: typeof RunRequestSchema;
    output: typeof RunResponseSchema;
  },
  /**
   * Make an unsigned transaction to call a specific realm function.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.MakeCallTx
   */
  makeCallTx: {
    methodKind: "unary";
    input: typeof CallRequestSchema;
    output: typeof MakeTxResponseSchema;
  },
  /**
   * Make an unsigned transaction to send currency to an account on the blockchain.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.MakeSendTx
   */
  makeSendTx: {
    methodKind: "unary";
    input: typeof SendRequestSchema;
    output: typeof MakeTxResponseSchema;
  },
  /**
   * Make an unsigned transaction to temporarily load the code in package on the blockchain and run main().
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.MakeRunTx
   */
  makeRunTx: {
    methodKind: "unary";
    input: typeof RunRequestSchema;
    output: typeof MakeTxResponseSchema;
  },
  /**
   * Sign the transaction using the active account.
   * If no active account has been set with SelectAccount, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrNoActiveAccount.
   * If the password is wrong, return [ErrCode](#land.gno.gnonative.v1.ErrCode).ErrDecryptionFailed.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.SignTx
   */
  signTx: {
    methodKind: "unary";
    input: typeof SignTxRequestSchema;
    output: typeof SignTxResponseSchema;
  },
  /**
   * Broadcast the signed transaction to the blockchain configured in GetRemote and return a stream result.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.BroadcastTxCommit
   */
  broadcastTxCommit: {
    methodKind: "server_streaming";
    input: typeof BroadcastTxCommitRequestSchema;
    output: typeof BroadcastTxCommitResponseSchema;
  },
  /**
   * Convert a byte array address to a bech32 string address.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.AddressToBech32
   */
  addressToBech32: {
    methodKind: "unary";
    input: typeof AddressToBech32RequestSchema;
    output: typeof AddressToBech32ResponseSchema;
  },
  /**
   * Convert a bech32 string address to a byte array address.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.AddressFromBech32
   */
  addressFromBech32: {
    methodKind: "unary";
    input: typeof AddressFromBech32RequestSchema;
    output: typeof AddressFromBech32ResponseSchema;
  },
  /**
   * Convert a mnemonic (as in CreateAccount) to a byte array address.
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.AddressFromMnemonic
   */
  addressFromMnemonic: {
    methodKind: "unary";
    input: typeof AddressFromMnemonicRequestSchema;
    output: typeof AddressFromMnemonicResponseSchema;
  },
  /**
   * Hello is for debug purposes
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.Hello
   */
  hello: {
    methodKind: "unary";
    input: typeof HelloRequestSchema;
    output: typeof HelloResponseSchema;
  },
  /**
   * HelloStream is for debug purposes
   *
   * @generated from rpc land.gno.gnonative.v1.GnoNativeService.HelloStream
   */
  helloStream: {
    methodKind: "server_streaming";
    input: typeof HelloStreamRequestSchema;
    output: typeof HelloStreamResponseSchema;
  },
}>;

