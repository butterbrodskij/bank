package initialization

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"strconv"
)

func RequestIntervalInitialization(requestInterval int, setRequestInterval func(next int), dayOff bool) []spot.Component {
	return []spot.Component{
		&ui.Label{
			X:        50,
			Y:        100,
			Width:    400,
			Height:   20,
			Value:    fmt.Sprintf("max request interval (1..10), current value: %d", requestInterval),
			FontSize: 16,
			Align:    ui.LabelAlignmentLeft,
		},
		&ui.TextEditor{
			X:        10,
			Y:        100,
			Width:    30,
			Height:   20,
			Text:     fmt.Sprint(requestInterval),
			FontSize: 16,
			OnChange: func(content string) {
				newApplicationInterval, err := strconv.Atoi(content)
				if err == nil && internal.ValidateRequestInterval(newApplicationInterval) && dayOff {
					setRequestInterval(newApplicationInterval)
				}
			},
		},
	}
}
