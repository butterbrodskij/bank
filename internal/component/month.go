package component

import (
	"errors"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) SkipMonthButton() spot.Component {
	if g.dayOff || g.env != nil && g.env.IsOver() {
		return nil
	}

	return &ui.Button{
		X:      340,
		Y:      280,
		Width:  150,
		Height: 25,
		Title:  "Skip to the end",
		OnClick: func() {
			var err error
			for err = g.env.SkipDay(); err == nil; err = g.env.SkipDay() {
			}
			if errors.Is(err, entities.ErrEndOfSimulation) {

			} else if err != nil {
				panic(err)
			}
			g.updated(struct{}{})
		},
	}
}
