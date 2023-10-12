// @generated by protoc-gen-connect-es v1.1.2
// @generated from file rpc.proto (package land.gno.gnomobile.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CallRequest, CallResponse, DeleteAccountRequest, DeleteAccountResponse, GenerateRecoveryPhraseRequest, GenerateRecoveryPhraseResponse, QueryRequest, QueryResponse, SetChainIDRequest, SetChainIDResponse, SetPasswordRequest, SetPasswordResponse, SetRemoteRequest, SetRemoteResponse } from "./gnomobiletypes_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { CreateAccountRequest, CreateAccountResponse, GetActiveAccountRequest, GetActiveAccountResponse, HelloRequest, HelloResponse, ListKeyInfoRequest, ListKeyInfoResponse, SelectAccountRequest, SelectAccountResponse } from "./rpc_pb.js";

/**
 * GnomobileService is the service to interact with the Gno blockchain
 *
 * @generated from service land.gno.gnomobile.v1.GnomobileService
 */
export const GnomobileService = {
  typeName: "land.gno.gnomobile.v1.GnomobileService",
  methods: {
    /**
     * Set the connection addresse for the remote node. If you don't call this,
     * the default is "127.0.0.1:26657"
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.SetRemote
     */
    setRemote: {
      name: "SetRemote",
      I: SetRemoteRequest,
      O: SetRemoteResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Set the chain ID for the remote node. If you don't call this, the default
     * is "dev"
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.SetChainID
     */
    setChainID: {
      name: "SetChainID",
      I: SetChainIDRequest,
      O: SetChainIDResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Set the password for the account in the keybase, used for later operations
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.SetPassword
     */
    setPassword: {
      name: "SetPassword",
      I: SetPasswordRequest,
      O: SetPasswordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Generate a recovery phrase of BIP39 mnemonic words using entropy from the
     * crypto library random number generator. This can be used as the mnemonic in
     * CreateAccount.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.GenerateRecoveryPhrase
     */
    generateRecoveryPhrase: {
      name: "GenerateRecoveryPhrase",
      I: GenerateRecoveryPhraseRequest,
      O: GenerateRecoveryPhraseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Get the keys informations in the keybase
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.ListKeyInfo
     */
    listKeyInfo: {
      name: "ListKeyInfo",
      I: ListKeyInfoRequest,
      O: ListKeyInfoResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Create a new account the keybase using the name an password specified by
     * SetAccount
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.CreateAccount
     */
    createAccount: {
      name: "CreateAccount",
      I: CreateAccountRequest,
      O: CreateAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * SelectAccount selects the active account to use for later operations
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.SelectAccount
     */
    selectAccount: {
      name: "SelectAccount",
      I: SelectAccountRequest,
      O: SelectAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetActiveAccount gets the active account which was set by SelectAccount.
     * If there is no active account, then return ErrNoActiveAccount.
     * (To check if there is an active account, use ListKeyInfo and check the
     * length of the result.)
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.GetActiveAccount
     */
    getActiveAccount: {
      name: "GetActiveAccount",
      I: GetActiveAccountRequest,
      O: GetActiveAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteAccount deletes the account with the given name, using the password to
     * ensure access. If the account doesn't exist, then return ErrCryptoKeyNotFound.
     * If the password is wrong, then return ErrDecryptionFailed.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.DeleteAccount
     */
    deleteAccount: {
      name: "DeleteAccount",
      I: DeleteAccountRequest,
      O: DeleteAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Make an ABCI query to the remote node.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.Query
     */
    query: {
      name: "Query",
      I: QueryRequest,
      O: QueryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Call a specific realm function.
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.Call
     */
    call: {
      name: "Call",
      I: CallRequest,
      O: CallResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Hello is for debug purposes
     *
     * @generated from rpc land.gno.gnomobile.v1.GnomobileService.Hello
     */
    hello: {
      name: "Hello",
      I: HelloRequest,
      O: HelloResponse,
      kind: MethodKind.Unary,
    },
  }
};

