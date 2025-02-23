package component

import (
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func Table(dayOff bool, bank *entities.BankBranch) spot.Component {
	if dayOff {
		return nil
	}
	return &ui.Label{
		X:        400,
		Y:        360,
		Width:    200,
		Height:   internal.MaxQueueCapacity * 40,
		Value:    bank.GetChanges(),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
