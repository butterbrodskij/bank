package main

import (
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/butterbrodskij/bank_branch/internal/component"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func main() {
	stats := entities.NewStatistics()
	queue := entities.NewQueue(stats, internal.MinQueueCapacity)
	sch := entities.NewSchedule(internal.MinLunchDuration)
	bank := entities.NewBankBranch(internal.MinWorkers, queue, stats)
	generator := entities.NewRequestGenerator(bank, queue, internal.MaxRequestInterval, internal.MinServingDuration,
		internal.MaxServingDuration, internal.MinProfitRange, internal.MaxProfitRange, internal.NormalDistribution)
	env := entities.NewEnvironment(bank, generator, sch)

	graphics := component.NewGraphics(env, stats)
	ui.Init()

	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
		graphics.UpdateState(ctx)
		return graphics.Window()
	})

	ui.Run()
}
