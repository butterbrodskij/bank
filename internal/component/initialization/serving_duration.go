package initialization

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"strconv"
)

func ServingDurationInitialization(servingDurationLeft, servingDurationRight int,
	setServingDurationLeft, setServingDurationRight func(next int), dayOff bool) []spot.Component {
	return []spot.Component{
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
	}
}
