package main

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/butterbrodskij/bank_branch/internal/component"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"os"
	"slices"
	"strconv"
)

var (
	distributions = []string{internal.NormalDistribution, internal.UniformDistribution}
)

func main() {
	stats := entities.NewStatistics()
	queue := entities.NewQueue(stats, internal.MinQueueCapacity)
	sch := entities.NewSchedule(internal.MinLunchDuration)
	bank := entities.NewBankBranch(internal.MinWorkers, queue, stats)
	flow := entities.NewApplicationFlow(bank, queue, internal.MaxApplicationInterval, internal.MinServingDuration,
		internal.MaxServingDuration, internal.MinProfitRange, internal.MaxProfitRange, internal.NormalDistribution)
	env := entities.NewEnvironment(bank, flow, sch)

	ui.Init()

	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
		dayOff, setDayOff := spot.UseState[bool](ctx, true)
		workers, setWorkers := spot.UseState[int](ctx, internal.MinWorkers)
		queueCapacity, setQueueCapacity := spot.UseState[int](ctx, internal.MinQueueCapacity)
		applicationInterval, setApplicationInterval := spot.UseState[int](ctx, internal.MaxApplicationInterval)
		servingDurationLeft, setServingDurationLeft := spot.UseState[int](ctx, internal.MinServingDuration)
		servingDurationRight, setServingDurationRight := spot.UseState[int](ctx, internal.MaxServingDuration)
		profitRangeLeft, setProfitRangeLeft := spot.UseState[int](ctx, internal.MinProfitRange)
		profitRangeRight, setProfitRangeRight := spot.UseState[int](ctx, internal.MaxProfitRange)
		modelingStep, setModelingStep := spot.UseState[int](ctx, internal.MaxModelingStep)
		lunchDuration, setLunchDuration := spot.UseState[int](ctx, internal.MinLunchDuration)
		distribution, setDistribution := spot.UseState[string](ctx, internal.UniformDistribution)
		_, updated := spot.UseState[struct{}](ctx, struct{}{})

		availableDistributions := distributions
		if !dayOff {
			availableDistributions = []string{distribution}
		}

		return &ui.Window{
			Title:  "bank branch",
			Width:  1000,
			Height: 1000,
			Children: []spot.Component{
				&ui.Label{
					X:        10,
					Y:        10,
					Width:    300,
					Height:   20,
					Value:    fmt.Sprintf("simulation initialization"),
					FontSize: 16,
					Align:    ui.LabelAlignmentLeft,
				},
				&ui.TextEditor{
					X:        10,
					Y:        40,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(workers),
					FontSize: 16,
					OnChange: func(content string) {
						newWorkers, err := strconv.Atoi(content)
						if err == nil && internal.ValidateWorkers(newWorkers) && dayOff {
							setWorkers(newWorkers)
						}
					},
				},
				&ui.Label{
					X:        50,
					Y:        40,
					Width:    400,
					Height:   20,
					Value:    fmt.Sprintf("workers (2..7), current value: %d", workers),
					FontSize: 16,
					Align:    ui.LabelAlignmentLeft,
				},
				&ui.TextEditor{
					X:        10,
					Y:        70,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(queueCapacity),
					FontSize: 16,
					OnChange: func(content string) {
						newQueueCapacity, err := strconv.Atoi(content)
						if err == nil && internal.ValidateQueueCapacity(newQueueCapacity) && dayOff {
							setQueueCapacity(newQueueCapacity)
						}
					},
				},
				&ui.Label{
					X:        50,
					Y:        70,
					Width:    400,
					Height:   20,
					Value:    fmt.Sprintf("queue capacity (10..25), current value: %d", queueCapacity),
					FontSize: 16,
					Align:    ui.LabelAlignmentLeft,
				},
				&ui.TextEditor{
					X:        10,
					Y:        100,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(applicationInterval),
					FontSize: 16,
					OnChange: func(content string) {
						newApplicationInterval, err := strconv.Atoi(content)
						if err == nil && internal.ValidateApplicationInterval(newApplicationInterval) && dayOff {
							setApplicationInterval(newApplicationInterval)
						}
					},
				},
				&ui.Label{
					X:        50,
					Y:        100,
					Width:    400,
					Height:   20,
					Value:    fmt.Sprintf("max application interval (1..10), current value: %d", applicationInterval),
					FontSize: 16,
					Align:    ui.LabelAlignmentLeft,
				},
				&ui.TextEditor{
					X:        10,
					Y:        130,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(servingDurationLeft),
					FontSize: 16,
					OnChange: func(content string) {
						newServingDurationLeft, err := strconv.Atoi(content)
						if err == nil && internal.ValidateServingDuration(newServingDurationLeft, servingDurationRight) && dayOff {
							setServingDurationLeft(newServingDurationLeft)
						}
					},
				},
				&ui.Label{
					X:        40,
					Y:        130,
					Width:    10,
					Height:   20,
					Value:    fmt.Sprint("-"),
					FontSize: 16,
					Align:    ui.LabelAlignmentCenter,
				},
				&ui.TextEditor{
					X:        50,
					Y:        130,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(servingDurationRight),
					FontSize: 16,
					OnChange: func(content string) {
						newServingDurationRight, err := strconv.Atoi(content)
						if err == nil && internal.ValidateServingDuration(servingDurationLeft, newServingDurationRight) && dayOff {
							setServingDurationRight(newServingDurationRight)
						}
					},
				},
				&ui.Label{
					X:      90,
					Y:      130,
					Width:  400,
					Height: 20,
					Value: fmt.Sprintf("serving duration (2..30-2..30), current value: %d-%d",
						servingDurationLeft, servingDurationRight),
					FontSize: 16,
					Align:    ui.LabelAlignmentLeft,
				},
				&ui.TextEditor{
					X:        10,
					Y:        160,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(profitRangeLeft),
					FontSize: 16,
					OnChange: func(content string) {
						newProfitRangeLeft, err := strconv.Atoi(content)
						if err == nil && internal.ValidateProfitRange(newProfitRangeLeft, profitRangeRight) && dayOff {
							setProfitRangeLeft(newProfitRangeLeft)
						}
					},
				},
				&ui.Label{
					X:        40,
					Y:        160,
					Width:    10,
					Height:   20,
					Value:    fmt.Sprint("-"),
					FontSize: 16,
					Align:    ui.LabelAlignmentCenter,
				},
				&ui.TextEditor{
					X:        50,
					Y:        160,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(profitRangeRight),
					FontSize: 16,
					OnChange: func(content string) {
						newProfitRangeRight, err := strconv.Atoi(content)
						if err == nil && internal.ValidateProfitRange(profitRangeLeft, newProfitRangeRight) && dayOff {
							setProfitRangeRight(newProfitRangeRight)
						}
					},
				},
				&ui.Label{
					X:      90,
					Y:      160,
					Width:  400,
					Height: 20,
					Value: fmt.Sprintf("profit range (3..50-3..50), current value: %d-%d",
						profitRangeLeft, profitRangeRight),
					FontSize: 16,
					Align:    ui.LabelAlignmentLeft,
				},
				&ui.TextEditor{
					X:        10,
					Y:        190,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(modelingStep),
					FontSize: 16,
					OnChange: func(content string) {
						newModelingStep, err := strconv.Atoi(content)
						if err == nil && internal.ValidateModelingStep(newModelingStep) && dayOff {
							setModelingStep(newModelingStep)
						}
					},
				},
				&ui.Label{
					X:        50,
					Y:        190,
					Width:    400,
					Height:   20,
					Value:    fmt.Sprintf("modeling step (10..60), current value: %d", modelingStep),
					FontSize: 16,
					Align:    ui.LabelAlignmentLeft,
				},
				&ui.TextEditor{
					X:        10,
					Y:        220,
					Width:    30,
					Height:   20,
					Text:     fmt.Sprint(lunchDuration),
					FontSize: 16,
					OnChange: func(content string) {
						newLunchDuration, err := strconv.Atoi(content)
						if err == nil && internal.ValidateLunchDuration(newLunchDuration) && dayOff {
							setLunchDuration(newLunchDuration)
						}
					},
				},
				&ui.Label{
					X:        50,
					Y:        220,
					Width:    400,
					Height:   20,
					Value:    fmt.Sprintf("lunch duration (0..60), current value: %d", lunchDuration),
					FontSize: 16,
					Align:    ui.LabelAlignmentLeft,
				},
				&ui.Dropdown{
					X:             10,
					Y:             250,
					Width:         200,
					Height:        20,
					Items:         availableDistributions,
					SelectedIndex: slices.Index(availableDistributions, distribution),
					Editable:      false,
					OnSelectionDidChange: func(index int) {
						if dayOff && index < len(distributions) {
							setDistribution(distributions[index])
						}
					},
				},
				component.StartButton(dayOff, workers, queueCapacity, applicationInterval, servingDurationLeft,
					servingDurationRight, profitRangeLeft, profitRangeRight, modelingStep, lunchDuration, distribution,
					setDayOff, env),
				component.StepButton(dayOff, env, updated),
				component.SkipDayButton(dayOff, env, updated),
				&ui.Button{
					X:      340,
					Y:      280,
					Width:  100,
					Height: 25,
					Title:  "Skip month",
					OnClick: func() {
						setDayOff(true)
					},
				},
				component.Schedule(dayOff, env),
				component.Queue(dayOff, queue),
				component.Table(dayOff, bank),
				component.Updates(dayOff, bank),
				component.Stats(dayOff, stats),
				&ui.Button{
					X:      410,
					Y:      950,
					Width:  180,
					Height: 25,
					Title:  "Quit",
					OnClick: func() {
						os.Exit(0)
					},
				},
			},
		}
	})

	ui.Run()
}
