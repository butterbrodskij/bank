package component

import (
	"fmt"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) Schedule() spot.Component {
	if g.dayOff {
		return nil
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
