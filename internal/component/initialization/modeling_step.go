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
			Text:     fmt.Sprint(modelingStep),
			FontSize: 16,
			OnChange: func(content string) {
				newModelingStep, err := strconv.Atoi(content)
				if err == nil && internal.ValidateModelingStep(newModelingStep) && dayOff {
					setModelingStep(newModelingStep)
				}
			},
		},
		&ui.Label{
			X:        50,
			Y:        190,
			Width:    400,
			Height:   20,
			Value:    fmt.Sprintf("modeling step (10..60), current value: %d", modelingStep),
			FontSize: 16,
			Align:    ui.LabelAlignmentLeft,
		},
	}
}
