package component

import (
	"fmt"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) Schedule() spot.Component {
	if g.errMessage != "" {
		return &ui.Label{
			X:        200,
			Y:        310,
			Width:    600,
			Height:   40,
			Value:    fmt.Sprintf("can not start simulation: %s", g.errMessage),
			FontSize: 16,
			Align:    ui.LabelAlignmentCenter,
		}
	}
	if g.dayOff {
		return nil
	}
	if g.env.IsOver() {
		return &ui.Label{
			X:        200,
			Y:        310,
			Width:    600,
			Height:   40,
			Value:    "Simulation Finished\n",
			FontSize: 16,
			Align:    ui.LabelAlignmentCenter,
		}
	}
	return &ui.Label{
		X:      200,
		Y:      310,
		Width:  600,
		Height: 40,
		Value: fmt.Sprintf("%s\t%s\nwork time: %s\tlunch time: %s",
			g.env.GetDay(), g.env.GetTime(), g.env.GetWorkTime(), g.env.GetLunchTime()),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
