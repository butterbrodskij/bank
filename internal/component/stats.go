package component

import (
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func Stats(dayOff bool, stats *entities.Statistics) spot.Component {
	if dayOff {
		return nil
	}
	return &ui.Label{
		X:        400,
		Y:        700,
		Width:    200,
		Height:   200,
		Value:    stats.String(),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
