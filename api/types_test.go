package api

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

// allMessageTypes lists every exported wire struct. The reflect audit
// (TestFieldTagConventions) walks it to guard the protojson dialect. Add new
// message types here.
var allMessageTypes = []reflect.Type{
	reflect.TypeOf(SetRemoteRequest{}),
	reflect.TypeOf(SetRemoteResponse{}),
	reflect.TypeOf(GetRemoteRequest{}),
	reflect.TypeOf(GetRemoteResponse{}),
	reflect.TypeOf(SetChainIDRequest{}),
	reflect.TypeOf(SetChainIDResponse{}),
	reflect.TypeOf(GetChainIDRequest{}),
	reflect.TypeOf(GetChainIDResponse{}),
	reflect.TypeOf(SetPasswordRequest{}),
	reflect.TypeOf(SetPasswordResponse{}),
	reflect.TypeOf(RenameKeyRequest{}),
	reflect.TypeOf(RenameKeyResponse{}),
	reflect.TypeOf(RotatePasswordRequest{}),
	reflect.TypeOf(RotatePasswordResponse{}),
	reflect.TypeOf(GenerateRecoveryPhraseRequest{}),
	reflect.TypeOf(GenerateRecoveryPhraseResponse{}),
	reflect.TypeOf(KeyInfo{}),
	reflect.TypeOf(Coin{}),
	reflect.TypeOf(BaseAccount{}),
	reflect.TypeOf(SessionAccount{}),
	reflect.TypeOf(ListKeyInfoRequest{}),
	reflect.TypeOf(ListKeyInfoResponse{}),
	reflect.TypeOf(GetKeyInfoByNameRequest{}),
	reflect.TypeOf(HasKeyByNameRequest{}),
	reflect.TypeOf(HasKeyByNameResponse{}),
	reflect.TypeOf(HasKeyByAddressRequest{}),
	reflect.TypeOf(HasKeyByAddressResponse{}),
	reflect.TypeOf(HasKeyByNameOrAddressRequest{}),
	reflect.TypeOf(HasKeyByNameOrAddressResponse{}),
	reflect.TypeOf(GetKeyInfoByNameResponse{}),
	reflect.TypeOf(GetKeyInfoByAddressRequest{}),
	reflect.TypeOf(GetKeyInfoByAddressResponse{}),
	reflect.TypeOf(GetKeyInfoByNameOrAddressRequest{}),
	reflect.TypeOf(GetKeyInfoByNameOrAddressResponse{}),
	reflect.TypeOf(CreateAccountRequest{}),
	reflect.TypeOf(CreateAccountResponse{}),
	reflect.TypeOf(CreateLedgerRequest{}),
	reflect.TypeOf(CreateLedgerResponse{}),
	reflect.TypeOf(ActivateAccountRequest{}),
	reflect.TypeOf(ActivateAccountResponse{}),
	reflect.TypeOf(GetActivatedAccountRequest{}),
	reflect.TypeOf(GetActivatedAccountResponse{}),
	reflect.TypeOf(QueryAccountRequest{}),
	reflect.TypeOf(QueryAccountResponse{}),
	reflect.TypeOf(QuerySessionAccountRequest{}),
	reflect.TypeOf(QuerySessionAccountResponse{}),
	reflect.TypeOf(DeleteAccountRequest{}),
	reflect.TypeOf(DeleteAccountResponse{}),
	reflect.TypeOf(QueryRequest{}),
	reflect.TypeOf(QueryResponse{}),
	reflect.TypeOf(RenderRequest{}),
	reflect.TypeOf(RenderResponse{}),
	reflect.TypeOf(QEvalRequest{}),
	reflect.TypeOf(QEvalResponse{}),
	reflect.TypeOf(MsgCall{}),
	reflect.TypeOf(CallRequest{}),
	reflect.TypeOf(CallResponse{}),
	reflect.TypeOf(MsgSend{}),
	reflect.TypeOf(SendRequest{}),
	reflect.TypeOf(SendResponse{}),
	reflect.TypeOf(MsgRun{}),
	reflect.TypeOf(RunRequest{}),
	reflect.TypeOf(RunResponse{}),
	reflect.TypeOf(MakeCallTxRequest{}),
	reflect.TypeOf(MakeSendTxRequest{}),
	reflect.TypeOf(MakeRunTxRequest{}),
	reflect.TypeOf(MakeTxResponse{}),
	reflect.TypeOf(SignTxRequest{}),
	reflect.TypeOf(SignTxResponse{}),
	reflect.TypeOf(MsgCreateSession{}),
	reflect.TypeOf(CreateSessionRequest{}),
	reflect.TypeOf(CreateSessionResponse{}),
	reflect.TypeOf(MsgRevokeSession{}),
	reflect.TypeOf(RevokeSessionRequest{}),
	reflect.TypeOf(RevokeSessionResponse{}),
	reflect.TypeOf(RevokeAllSessionsRequest{}),
	reflect.TypeOf(RevokeAllSessionsResponse{}),
	reflect.TypeOf(EstimateGasRequest{}),
	reflect.TypeOf(EstimateGasResponse{}),
	reflect.TypeOf(EstimateTxFeesRequest{}),
	reflect.TypeOf(EstimateTxFeesResponse{}),
	reflect.TypeOf(BroadcastTxCommitRequest{}),
	reflect.TypeOf(BroadcastTxCommitResponse{}),
	reflect.TypeOf(AddressToBech32Request{}),
	reflect.TypeOf(AddressToBech32Response{}),
	reflect.TypeOf(AddressFromBech32Request{}),
	reflect.TypeOf(AddressFromBech32Response{}),
	reflect.TypeOf(AddressFromMnemonicRequest{}),
	reflect.TypeOf(AddressFromMnemonicResponse{}),
	reflect.TypeOf(ValidateMnemonicWordRequest{}),
	reflect.TypeOf(ValidateMnemonicWordResponse{}),
	reflect.TypeOf(ValidateMnemonicPhraseRequest{}),
	reflect.TypeOf(ValidateMnemonicPhraseResponse{}),
	reflect.TypeOf(PubKeyBytesFromBech32Request{}),
	reflect.TypeOf(PubKeyBytesFromBech32Response{}),
	reflect.TypeOf(HelloRequest{}),
	reflect.TypeOf(HelloResponse{}),
	reflect.TypeOf(HelloStreamRequest{}),
	reflect.TypeOf(HelloStreamResponse{}),
}

// TestFieldTagConventions asserts the protojson dialect over every wire struct:
// every field has a json tag with omitempty, every 64-bit int has ",string",
// no yaml tags remain, and no field is a non-pointer struct (omitempty is a
// no-op on value structs).
func TestFieldTagConventions(t *testing.T) {
	for _, typ := range allMessageTypes {
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			where := typ.Name() + "." + f.Name

			if yaml := f.Tag.Get("yaml"); yaml != "" {
				t.Errorf("%s: unexpected yaml tag %q (should be dropped)", where, yaml)
			}

			tag := f.Tag.Get("json")
			if tag == "" || tag == "-" {
				t.Errorf("%s: missing json tag", where)
				continue
			}
			parts := strings.Split(tag, ",")
			name, opts := parts[0], parts[1:]
			if name == "" {
				t.Errorf("%s: json tag has no field name", where)
			}
			if !contains(opts, "omitempty") {
				t.Errorf("%s: json tag %q missing omitempty", where, tag)
			}

			switch f.Type.Kind() {
			case reflect.Int64, reflect.Uint64:
				if !contains(opts, "string") {
					t.Errorf("%s: 64-bit int json tag %q missing ,string", where, tag)
				}
			case reflect.Struct:
				t.Errorf("%s: non-pointer struct field (must be a pointer so omitempty works)", where)
			}
		}
	}
}

func contains(ss []string, want string) bool {
	for _, s := range ss {
		if s == want {
			return true
		}
	}
	return false
}

// TestGoldenMarshal pins the exact JSON output for representative populated and
// zero-value structs. These literals are the wire contract the TS layer relies on.
func TestGoldenMarshal(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{
			name: "CallRequest",
			in: CallRequest{
				GasFee:        "1000000ugnot",
				GasWanted:     200000,
				Memo:          "memo",
				SignerAddress: []byte{0x01, 0x02, 0x03},
				Msgs: []*MsgCall{{
					PackagePath: "gno.land/r/demo/boards",
					Fnc:         "CreateReply",
					Args:        []string{"1", "2"},
				}},
			},
			want: `{"gasFee":"1000000ugnot","gasWanted":"200000","memo":"memo","signerAddress":"AQID","Msgs":[{"packagePath":"gno.land/r/demo/boards","fnc":"CreateReply","args":["1","2"]}]}`,
		},
		{
			name: "EstimateTxFeesResponse",
			in: EstimateTxFeesResponse{
				TxJSON:       "{}",
				GasWanted:    300000,
				GasFee:       &Coin{Denom: "ugnot", Amount: 500},
				StorageDelta: -10,
				StorageFee:   []*Coin{{Denom: "ugnot", Amount: 20}},
				TotalFee:     &Coin{Denom: "ugnot", Amount: 520},
			},
			want: `{"txJson":"{}","gasWanted":"300000","gasFee":{"denom":"ugnot","amount":"500"},"storageDelta":"-10","storageFee":[{"denom":"ugnot","amount":"20"}],"TotalFee":{"denom":"ugnot","amount":"520"}}`,
		},
		{
			name: "ListKeyInfoResponse",
			in: ListKeyInfoResponse{Keys: []*KeyInfo{{
				Type:    1,
				Name:    "alice",
				PubKey:  []byte{0xAA},
				Address: []byte{0xBB},
			}}},
			want: `{"key_info":[{"type":1,"name":"alice","pubKey":"qg==","address":"uw=="}]}`,
		},
		{
			name: "SignTxRequest",
			in: SignTxRequest{
				TxJSON:         "{}",
				Address:        []byte{0x01},
				AccountNumber:  5,
				SequenceNumber: 7,
			},
			want: `{"txJson":"{}","address":"AQ==","accountNumber":"5","sequenceNumber":"7"}`,
		},
		{
			name: "HelloResponse",
			in:   HelloResponse{Greeting: "hi"},
			want: `{"Greeting":"hi"}`,
		},
		{name: "CallRequest zero", in: CallRequest{}, want: `{}`},
		{name: "EstimateTxFeesResponse zero", in: EstimateTxFeesResponse{}, want: `{}`},
		{name: "ListKeyInfoResponse zero", in: ListKeyInfoResponse{}, want: `{}`},
		{name: "SignTxRequest zero", in: SignTxRequest{}, want: `{}`},
		{name: "HelloResponse zero", in: HelloResponse{}, want: `{}`},
		{name: "KeyInfo zero", in: KeyInfo{}, want: `{}`},
		{name: "BaseAccount zero", in: BaseAccount{}, want: `{}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.in)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if string(got) != tt.want {
				t.Errorf("\n got: %s\nwant: %s", got, tt.want)
			}
		})
	}
}

// TestUnmarshalFixtures decodes payloads shaped like what the TS client sends,
// asserting the ",string" 64-bit and base64 []byte handling.
func TestUnmarshalFixtures(t *testing.T) {
	t.Run("CallRequest", func(t *testing.T) {
		var req CallRequest
		payload := `{"gasFee":"1ugnot","gasWanted":"200000","signerAddress":"AQID","Msgs":[{"packagePath":"p","fnc":"F","args":["a"]}]}`
		if err := json.Unmarshal([]byte(payload), &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if req.GasWanted != 200000 {
			t.Errorf("GasWanted = %d, want 200000", req.GasWanted)
		}
		if !reflect.DeepEqual(req.SignerAddress, []byte{1, 2, 3}) {
			t.Errorf("SignerAddress = %v, want [1 2 3]", req.SignerAddress)
		}
		if len(req.Msgs) != 1 || req.Msgs[0].Fnc != "F" {
			t.Errorf("Msgs decoded wrong: %+v", req.Msgs)
		}
	})

	t.Run("SignTxRequest string ints", func(t *testing.T) {
		var req SignTxRequest
		payload := `{"txJson":"{}","accountNumber":"5","sequenceNumber":"7"}`
		if err := json.Unmarshal([]byte(payload), &req); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if req.AccountNumber != 5 || req.SequenceNumber != 7 {
			t.Errorf("got account=%d seq=%d, want 5/7", req.AccountNumber, req.SequenceNumber)
		}
	})

	// ",string" fields reject bare numbers (protojson accepted both; documented
	// in the v4->v5 migration notes).
	t.Run("bare number rejected", func(t *testing.T) {
		var req SignTxRequest
		if err := json.Unmarshal([]byte(`{"accountNumber":5}`), &req); err == nil {
			t.Errorf("expected error unmarshalling bare number into ,string field")
		}
	})
}
