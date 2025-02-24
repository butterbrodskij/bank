package component

import (
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func Updates(dayOff bool, bank *entities.BankBranch) spot.Component {
	if dayOff {
		return nil
	}
	return &ui.Label{
		X:        700,
		Y:        360,
		Width:    200,
		Height:   600,
		Value:    bank.GetUpdates(),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
