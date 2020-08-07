package types

import 	(
	sdk "github.com/cosmos/cosmos-sdk/types"
	planTypes "github.com/Shushsa/plan/x/plancoin/types"
)


type Emission struct {
	Current sdk.Int `json:"current"` // Текущая эмиссия
	Threshold sdk.Int `json:"threshold"` // Порог, после которого парамайнинг перестает работать
}

// Достигнут ли порог
func(e Emission) IsThresholdReached() bool {
	return e.Current.GTE(e.Threshold)
}

// Starting emission
func NewEmission() Emission {
	return Emission{
		Current: sdk.NewIntWithDecimal(planTypes.INITIAL, planTypes.POINTS),
		Threshold: sdk.NewIntWithDecimal(7000000000, planTypes.POINTS),
	}
}

func (e Emission) String() string {
	return e.Current.String()
}
