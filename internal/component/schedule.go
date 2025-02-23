package component

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func Schedule(dayOff bool, env *entities.Environment) spot.Component {
	if dayOff {
		return nil
	}
	return &ui.Label{
		X:      200,
		Y:      310,
		Width:  600,
		Height: 40,
		Value: fmt.Sprintf("%s\t%s\nwork time: %s\tlunch time: %s",
			env.GetDay(), env.GetTime(), env.GetWorkTime(), env.GetLunchTime()),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
