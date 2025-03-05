package component

import (
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func BreakButton(dayOff bool, setDayOff func(next bool), env *entities.Environment, stats *entities.Statistics) spot.Component {
	if dayOff {
		return nil
	}
	return &ui.Button{
		X:      410,
		Y:      920,
		Width:  180,
		Height: 25,
		Title:  "Break",
		OnClick: func() {
			env.Break()
			stats.Empty()
			setDayOff(true)
		},
	}
}
