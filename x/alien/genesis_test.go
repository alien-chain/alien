package alien_test

import (
	"testing"

	keepertest "github.com/alien-chain/alien/testutil/keeper"
	"github.com/alien-chain/alien/x/alien"
	"github.com/alien-chain/alien/x/alien/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AlienKeeper(t)
	alien.InitGenesis(ctx, *k, genesisState)
	got := alien.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
