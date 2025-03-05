package component

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) Stats() spot.Component {
	if g.dayOff {
		return nil
	}
	return &ui.Label{
		X:        400,
		Y:        700,
		Width:    200,
		Height:   200,
		Value:    g.stats.String(),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
