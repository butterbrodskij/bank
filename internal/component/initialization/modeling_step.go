package initialization

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"strconv"
)

func ModelingStepInitialization(modelingStep int, setModelingStep func(next int), dayOff bool) []spot.Component {
	return []spot.Component{
		&ui.TextEditor{
			X:        10,
			Y:        190,
			Width:    30,
			Height:   20,
			Text:     internal.StringValue(modelingStep),
			FontSize: 16,
			OnChange: func(content string) {
				newModelingStep, err := strconv.Atoi(content)
				if dayOff {
					if err == nil && internal.ValidateModelingStep(newModelingStep) {
						setModelingStep(newModelingStep)
					} else {
						setModelingStep(internal.InvalidValue)
					}
				}
			},
		},
		&ui.Label{
			X:        50,
			Y:        190,
			Width:    400,
			Height:   20,
			Value:    fmt.Sprintf("modeling step (10..60), current value: %s", internal.StringValue(modelingStep)),
			FontSize: 16,
			Align:    ui.LabelAlignmentLeft,
		},
	}
}
