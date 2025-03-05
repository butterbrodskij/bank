package component

import (
	"errors"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) SkipDayButton() spot.Component {
	if g.dayOff || g.env != nil && g.env.IsOver() {
		return nil
	}

	return &ui.Button{
		X:      230,
		Y:      280,
		Width:  100,
		Height: 25,
		Title:  "Skip day",
		OnClick: func() {
			err := g.env.SkipDay()
			if errors.Is(err, entities.ErrEndOfSimulation) {

			} else if err != nil {
				panic(err)
			}
			g.updated(struct{}{})
		},
	}
}
