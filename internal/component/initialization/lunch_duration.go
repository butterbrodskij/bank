package initialization

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"strconv"
)

func LunchDurationInitialization(lunchDuration int, setLunchDuration func(next int), dayOff bool) []spot.Component {
	return []spot.Component{
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
	}
}
