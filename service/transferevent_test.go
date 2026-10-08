package service

import (
	"testing"

	"github.com/gnolang/gno/tm2/pkg/amino"
	abci "github.com/gnolang/gno/tm2/pkg/bft/abci/types"
	"github.com/stretchr/testify/require"
)

// gno.land mainnet emits bank.TransferEvent for every ugnot movement (gno #6120). A transaction result
// carrying one must decode, or every call that moves coins is reported as failed although it landed.
func TestDecodeBankTransferEvent(t *testing.T) {
	const js = `{"@type":"/bank.TransferEvent","from":"g1jg8mtutu9khhfwc4nxmuhcpftf0pajdhfvsqf5","to":"g1jyh59tr2t4yjmwajjax9mye369gkkujv7tegkk","coins":"1000ugnot"}`
	var ev abci.Event
	require.NoError(t, amino.UnmarshalJSON([]byte(js), &ev))
	require.NotNil(t, ev)
}
