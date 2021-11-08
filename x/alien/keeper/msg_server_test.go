package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/alien-chain/alien/testutil/keeper"
	"github.com/alien-chain/alien/x/alien/keeper"
	"github.com/alien-chain/alien/x/alien/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AlienKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
