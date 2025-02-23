package entities

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
)

type ApplicationFlow struct {
	bank                *BankBranch
	queue               *Queue
	ApplicationInterval *internal.Range
	ServingDuration     *internal.Range
	ProfitRange         *internal.Range
	Distribution        string
	NextApp             int
}

func NewApplicationFlow(bank *BankBranch, queue *Queue, applicationInterval, servingDurationLeft, servingDurationRight,
	profitRangeLeft, profitRangeRight int, distribution string) *ApplicationFlow {
	applications := internal.NewRange(0, applicationInterval)
	return &ApplicationFlow{
		bank:                bank,
		queue:               queue,
		ApplicationInterval: applications,
		ServingDuration:     internal.NewRange(servingDurationLeft, servingDurationRight),
		ProfitRange:         internal.NewRange(profitRangeLeft, profitRangeRight),
		Distribution:        distribution,
		NextApp:             internal.RandValue(applications, distribution),
	}
}

func (a *ApplicationFlow) Update(applicationInterval, servingDurationLeft, servingDurationRight,
	profitRangeLeft, profitRangeRight int, distribution string) (*ApplicationFlow, error) {
	if a == nil {
		return nil, fmt.Errorf("nil Application Flow")
	}
	a.ApplicationInterval = internal.NewRange(0, applicationInterval)
	a.ServingDuration = internal.NewRange(servingDurationLeft, servingDurationRight)
	a.ProfitRange = internal.NewRange(profitRangeLeft, profitRangeRight)
	a.Distribution = distribution
	a.NextApp = internal.RandValue(a.ApplicationInterval, distribution)
	return a, nil
}

func (a *ApplicationFlow) GenerateApplication(now *timestamp) {
	application := NewApplication(now, internal.RandValue(a.ServingDuration, a.Distribution), internal.RandValue(a.ProfitRange, a.Distribution))
	id := a.queue.AddClient(application)
	a.bank.NotifyClientUpdated(id)
	a.NextApp = internal.RandValue(a.ApplicationInterval, a.Distribution)
}
