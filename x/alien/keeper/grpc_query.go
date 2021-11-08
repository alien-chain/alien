package keeper

import (
	"github.com/alien-chain/alien/x/alien/types"
)

var _ types.QueryServer = Keeper{}
