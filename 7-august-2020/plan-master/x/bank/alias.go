package bank

import (
	"github.com/Shushsa/plan/x/bank/keeper"
)

var (
	NewKeeper = keeper.NewKeeper
)

type (
	Keeper = keeper.Keeper
)
