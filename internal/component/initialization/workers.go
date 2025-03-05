package initialization

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"strconv"
)

func WorkersInitialization(workers int, setWorkers func(next int), dayOff bool) []spot.Component {
	return []spot.Component{
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
	}
}
