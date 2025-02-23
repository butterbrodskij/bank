package component

import (
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func StartButton(dayOff bool, workers, queueCapacity, applicationInterval, servingDurationLeft, servingDurationRight,
	profitRangeLeft, profitRangeRight, modelingStep, lunchDuration int, distribution string,
	setDayOff func(next bool), env *entities.Environment) spot.Component {
	if !dayOff {
		return &ui.Label{
			X:        10,
			Y:        280,
			Width:    100,
			Height:   25,
			Value:    "Started",
			FontSize: 15,
			Align:    ui.LabelAlignmentCenter,
		}
	}
	return &ui.Button{
		X:      10,
		Y:      280,
		Width:  100,
		Height: 25,
		Title:  "Start",
		OnClick: func() {
			_, err := env.Update(workers, queueCapacity, applicationInterval, servingDurationLeft,
				servingDurationRight, profitRangeLeft, profitRangeRight, modelingStep, lunchDuration, distribution)
			if err != nil {
				panic(err)
			}
			setDayOff(false)
		},
	}
}
