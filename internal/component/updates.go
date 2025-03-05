package component

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) Updates() spot.Component {
	if g.dayOff {
		return nil
	}
	return &ui.Label{
		X:        800,
		Y:        360,
		Width:    200,
		Height:   600,
		Value:    g.env.GetUpdates(),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
