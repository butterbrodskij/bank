package entities

import (
	"fmt"
	"github.com/butterbrodskij/bank_branch/internal"
)

type RequestGenerator struct {
	bank            *BankBranch
	queue           *Queue
	RequestInterval *internal.Range
	ServingDuration *internal.Range
	ProfitRange     *internal.Range
	Distribution    string
	NextApp         int
}

func NewRequestGenerator(bank *BankBranch, queue *Queue, applicationInterval, servingDurationLeft, servingDurationRight,
	profitRangeLeft, profitRangeRight int, distribution string) *RequestGenerator {
	applications := internal.NewRange(0, applicationInterval)
	return &RequestGenerator{
		bank:            bank,
		queue:           queue,
		RequestInterval: applications,
		ServingDuration: internal.NewRange(servingDurationLeft, servingDurationRight),
		ProfitRange:     internal.NewRange(profitRangeLeft, profitRangeRight),
		Distribution:    distribution,
		NextApp:         internal.RandValue(applications, distribution),
	}
}

func (a *RequestGenerator) Update(applicationInterval, servingDurationLeft, servingDurationRight,
	profitRangeLeft, profitRangeRight int, distribution string) (*RequestGenerator, error) {
	if a == nil {
		return nil, fmt.Errorf("nil Request Flow")
	}
	a.RequestInterval = internal.NewRange(0, applicationInterval)
	a.ServingDuration = internal.NewRange(servingDurationLeft, servingDurationRight)
	a.ProfitRange = internal.NewRange(profitRangeLeft, profitRangeRight)
	a.Distribution = distribution
	a.NextApp = internal.RandValue(a.RequestInterval, distribution)
	return a, nil
}

func (a *RequestGenerator) GenerateApplication(now *timestamp) {
	application := NewRequest(now, internal.RandValue(a.ServingDuration, a.Distribution), internal.RandValue(a.ProfitRange, a.Distribution))
	id := a.queue.AddClient(application)
	a.bank.NotifyClientUpdated(id)
	a.NextApp = internal.RandValue(a.RequestInterval, a.Distribution)
}
