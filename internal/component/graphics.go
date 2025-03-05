package component

import (
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
)

type Graphics struct {
	env   *entities.Environment
	stats *entities.Statistics

	dayOff                  bool
	setDayOff               func(next bool)
	workers                 int
	setWorkers              func(next int)
	queueCapacity           int
	setQueueCapacity        func(next int)
	requestInterval         int
	setRequestInterval      func(next int)
	servingDurationLeft     int
	setServingDurationLeft  func(next int)
	servingDurationRight    int
	setServingDurationRight func(next int)
	profitRangeLeft         int
	setProfitRangeLeft      func(next int)
	profitRangeRight        int
	setProfitRangeRight     func(next int)
	modelingStep            int
	setModelingStep         func(next int)
	lunchDuration           int
	setLunchDuration        func(next int)
	distribution            string
	setDistribution         func(next string)
	distributions           []string
	updated                 func(struct{})
}

func NewGraphics(env *entities.Environment, stats *entities.Statistics) *Graphics {
	return &Graphics{
		env:   env,
		stats: stats,
	}
}

func (g *Graphics) UpdateState(ctx *spot.RenderContext) *Graphics {
	if g == nil {
		return nil
	}
	g.dayOff, g.setDayOff = spot.UseState[bool](ctx, true)
	g.workers, g.setWorkers = spot.UseState[int](ctx, internal.MinWorkers)
	g.queueCapacity, g.setQueueCapacity = spot.UseState[int](ctx, internal.MinQueueCapacity)
	g.requestInterval, g.setRequestInterval = spot.UseState[int](ctx, internal.MaxRequestInterval)
	g.servingDurationLeft, g.setServingDurationLeft = spot.UseState[int](ctx, internal.MinServingDuration)
	g.servingDurationRight, g.setServingDurationRight = spot.UseState[int](ctx, internal.MaxServingDuration)
	g.profitRangeLeft, g.setProfitRangeLeft = spot.UseState[int](ctx, internal.MinProfitRange)
	g.profitRangeRight, g.setProfitRangeRight = spot.UseState[int](ctx, internal.MaxProfitRange)
	g.modelingStep, g.setModelingStep = spot.UseState[int](ctx, internal.MaxModelingStep)
	g.lunchDuration, g.setLunchDuration = spot.UseState[int](ctx, internal.MinLunchDuration)
	g.distribution, g.setDistribution = spot.UseState[string](ctx, internal.UniformDistribution)
	if !g.dayOff {
		g.distributions = []string{g.distribution}
	} else {
		g.distributions = internal.GetDistributions()
	}
	_, g.updated = spot.UseState[struct{}](ctx, struct{}{})
	return g
}
