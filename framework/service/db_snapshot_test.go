package gnonative

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// The native DB has no point-in-time view; like goleveldb and boltdb it refuses snapshots
// (only a node's multistore asks for one, not the keybase).
func TestNativeDBRefusesSnapshots(t *testing.T) {
	snap, err := (&db{}).NewSnapshot()
	require.Error(t, err)
	require.Nil(t, snap)
}
