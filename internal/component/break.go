package component

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) BreakButton() spot.Component {
	if g.dayOff {
		return nil
	}
	return &ui.Button{
		X:      410,
		Y:      920,
		Width:  180,
		Height: 25,
		Title:  "Break",
		OnClick: func() {
			g.env.Break()
			g.stats.Empty()
			g.setDayOff(true)
		},
	}
}
