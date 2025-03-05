package component

import (
	"errors"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func SkipMonthButton(dayOff bool, env *entities.Environment, updated func(struct{})) spot.Component {
	if dayOff || env != nil && env.IsOver() {
		return nil
	}

	return &ui.Button{
		X:      340,
		Y:      280,
		Width:  100,
		Height: 25,
		Title:  "Skip to the end",
		OnClick: func() {
			var err error
			for err = env.SkipDay(); err == nil; err = env.SkipDay() {
			}
			if errors.Is(err, entities.ErrEndOfSimulation) {

			} else if err != nil {
				panic(err)
			}
			updated(struct{}{})
		},
	}
}
