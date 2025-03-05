package initialization

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"strconv"
)

func QueueCapacityInitialization(queueCapacity int, setQueueCapacity func(next int), dayOff bool) []spot.Component {
	return []spot.Component{
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
	}
}
