package component

import (
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) Table() spot.Component {
	if g.dayOff {
		return nil
	}
	return &ui.Label{
		X:        300,
		Y:        360,
		Width:    200,
		Height:   internal.MaxQueueCapacity * 40,
		Value:    g.env.GetInfo(),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
