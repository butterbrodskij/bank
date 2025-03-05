package component

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal/component/initialization"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"os"
	"slices"
)

func (g *Graphics) Window() spot.Component {
	return &ui.Window{
		Title:    "bank branch",
		Width:    1000,
		Height:   1000,
		Children: g.Components(),
	}
}

func (g *Graphics) Components() []spot.Component {
	if g == nil {
		return nil
	}
	components := []spot.Component{
		&ui.Label{
			X:        10,
			Y:        10,
			Width:    300,
			Height:   20,
			Value:    fmt.Sprintf("simulation initialization"),
			FontSize: 16,
			Align:    ui.LabelAlignmentLeft,
		}, &ui.Button{
			X:      410,
			Y:      950,
			Width:  180,
			Height: 25,
			Title:  "Quit",
			OnClick: func() {
				os.Exit(0)
			},
		},
		g.StartButton(),
		g.StepButton(),
		g.SkipDayButton(),
		g.SkipMonthButton(),
		g.Schedule(),
		g.Queue(),
		g.Table(),
		g.Flow(),
		g.Updates(),
		g.Stats(),
		g.BreakButton(),
	}
	components = slices.Concat(components, initialization.WorkersInitialization(g.workers, g.setWorkers, g.dayOff),
		initialization.QueueCapacityInitialization(g.queueCapacity, g.setQueueCapacity, g.dayOff),
		initialization.RequestIntervalInitialization(g.requestInterval, g.setRequestInterval, g.dayOff),
		initialization.ServingDurationInitialization(g.servingDurationLeft, g.servingDurationRight, g.setServingDurationLeft, g.setServingDurationRight, g.dayOff),
		initialization.ProfitRangeInitialization(g.profitRangeLeft, g.profitRangeRight, g.setProfitRangeLeft, g.setProfitRangeRight, g.dayOff),
		initialization.ModelingStepInitialization(g.modelingStep, g.setModelingStep, g.dayOff),
		initialization.LunchDurationInitialization(g.lunchDuration, g.setLunchDuration, g.dayOff),
		initialization.DistributionInitialization(g.distributions, g.distribution, g.setDistribution, g.dayOff))
	return components
}
