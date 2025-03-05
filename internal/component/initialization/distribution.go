package initialization

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"slices"
)

func DistributionInitialization(distributions []string, distribution string, setDistribution func(next string), dayOff bool) []spot.Component {
	return []spot.Component{
		&ui.Dropdown{
			X:             10,
			Y:             250,
			Width:         200,
			Height:        20,
			Items:         distributions,
			SelectedIndex: slices.Index(distributions, distribution),
			Editable:      false,
			OnSelectionDidChange: func(index int) {
				if dayOff && index < len(distributions) {
					setDistribution(distributions[index])
				}
			},
		},
	}
}
