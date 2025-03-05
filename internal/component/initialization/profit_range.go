package initialization

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"strconv"
)

func ProfitRangeInitialization(profitRangeLeft, profitRangeRight int,
	setProfitRangeLeft, setProfitRangeRight func(next int), dayOff bool) []spot.Component {
	return []spot.Component{
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
	}
}
