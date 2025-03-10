package component

import (
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func (g *Graphics) StartButton() spot.Component {
	if g.env != nil && g.env.IsOver() {
		return &ui.Label{
			X:        10,
			Y:        280,
			Width:    100,
			Height:   25,
			Value:    "Finished",
			FontSize: 15,
			Align:    ui.LabelAlignmentCenter,
		}
	}
	if !g.dayOff {
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
			if !internal.ValidateWorkers(g.workers) || !internal.ValidateServingDuration(g.servingDurationLeft, g.servingDurationRight) ||
				!internal.ValidateProfitRange(g.profitRangeLeft, g.profitRangeRight) || !internal.ValidateLunchDuration(g.lunchDuration) ||
				!internal.ValidateModelingStep(g.modelingStep) || !internal.ValidateRequestInterval(g.requestInterval) ||
				!internal.ValidateQueueCapacity(g.queueCapacity) {
				g.setErrMessage("invalid parameters")
				return
			}
			g.setErrMessage("")
			_, err := g.env.Update(g.workers, g.queueCapacity, g.requestInterval, g.servingDurationLeft,
				g.servingDurationRight, g.profitRangeLeft, g.profitRangeRight, g.modelingStep, g.lunchDuration, g.distribution)
			if err != nil {
				panic(err)
			}
			g.setDayOff(false)
		},
	}
}
