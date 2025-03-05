package component

import (
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/butterbrodskij/bank_branch/internal/entities"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func Queue(dayOff bool, q *entities.Queue) spot.Component {
	if dayOff {
		return nil
	}
	return &ui.Label{
		X:        50,
		Y:        360,
		Width:    200,
		Height:   internal.MaxQueueCapacity * 40,
		Value:    q.String(),
		FontSize: 16,
		Align:    ui.LabelAlignmentCenter,
	}
}
