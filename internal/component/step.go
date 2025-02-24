package component

import (
	"errors"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func StepButton(dayOff bool, env *entities.Environment, updated func(struct{})) spot.Component {
	if dayOff || env != nil && env.IsOver() {
		return nil
	}

	return &ui.Button{
		X:      120,
		Y:      280,
		Width:  100,
		Height: 25,
		Title:  "Step",
		OnClick: func() {
			err := env.Step()
			if errors.Is(err, entities.ErrEndOfSimulation) {

			} else if err != nil {
				panic(err)
			}
			updated(struct{}{})
		},
	}
}
