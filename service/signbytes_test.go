package service

import (
	"context"
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/gnolang/gno/tm2/pkg/crypto/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	api_gen "github.com/gnolang/gnonative/v4/api/gen/go"
)

// gno's integration test seed (gno.land/pkg/integration DefaultAccount_Seed).
const testMnemonic = "source bonus chronic canvas draft south burst lottery vacant surface solve popular case indicate oppose farm nothing bullet exhibit title speed wink action roast"

func newTestService(t *testing.T) *gnoNativeService {
	t.Helper()
	return &gnoNativeService{
		logger:       zap.NewNop(),
		keybase:      keys.NewInMemory(),
		userAccounts: map[string]*userAccount{},
		chainID:      "dev", // SetPassword validates the signer, which needs a chain ID
	}
}

// activated creates key "sess" with password "pw", activates it and, if setPassword, sets the password.
func activated(t *testing.T, s *gnoNativeService, setPassword bool) keys.Info {
	t.Helper()
	ctx := context.Background()
	info, err := s.keybase.CreateAccount("sess", testMnemonic, "", "pw", 0, 0)
	require.NoError(t, err)
	_, err = s.ActivateAccount(ctx, connect.NewRequest(&api_gen.ActivateAccountRequest{NameOrBech32: "sess"}))
	require.NoError(t, err)
	if setPassword {
		_, err = s.SetPassword(ctx, connect.NewRequest(&api_gen.SetPasswordRequest{Password: "pw", Address: info.GetAddress().Bytes()}))
		require.NoError(t, err)
	}
	return info
}

func TestSignBytes(t *testing.T) {
	s := newTestService(t)
	info := activated(t, s, true)
	data := []byte("gnoconnect-session-v1\nhost=connect\nstate=ab+c=")

	res, err := s.SignBytes(context.Background(), connect.NewRequest(&api_gen.SignBytesRequest{
		Address: info.GetAddress().Bytes(),
		Data:    data,
	}))
	require.NoError(t, err)
	assert.Len(t, res.Msg.Signature, 64)
	assert.Equal(t, info.GetPubKey().Bytes(), res.Msg.PubKey)
	assert.True(t, info.GetPubKey().VerifyBytes(data, res.Msg.Signature), "signature must verify over the raw data")
	assert.False(t, info.GetPubKey().VerifyBytes([]byte("other"), res.Msg.Signature))
}

func TestSignBytesEmptyData(t *testing.T) {
	s := newTestService(t)
	info := activated(t, s, true)

	res, err := s.SignBytes(context.Background(), connect.NewRequest(&api_gen.SignBytesRequest{
		Address: info.GetAddress().Bytes(),
	}))
	require.NoError(t, err)
	assert.True(t, info.GetPubKey().VerifyBytes([]byte{}, res.Msg.Signature))
}

func TestSignBytesNoActiveAccount(t *testing.T) {
	s := newTestService(t)
	info, err := s.keybase.CreateAccount("sess", testMnemonic, "", "pw", 0, 0)
	require.NoError(t, err)

	_, err = s.SignBytes(context.Background(), connect.NewRequest(&api_gen.SignBytesRequest{
		Address: info.GetAddress().Bytes(),
		Data:    []byte("x"),
	}))
	require.Error(t, err)
	assert.True(t, errors.Is(err, api_gen.ErrCode_ErrNoActiveAccount), "got %v", err)
}

func TestSignBytesPasswordUnset(t *testing.T) {
	s := newTestService(t)
	info := activated(t, s, false)

	_, err := s.SignBytes(context.Background(), connect.NewRequest(&api_gen.SignBytesRequest{
		Address: info.GetAddress().Bytes(),
		Data:    []byte("x"),
	}))
	require.Error(t, err)
	assert.True(t, errors.Is(err, api_gen.ErrCode_ErrDecryptionFailed), "got %v", err)
}
