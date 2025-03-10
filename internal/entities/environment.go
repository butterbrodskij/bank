package entities

import (
	"errors"
	"github.com/butterbrodskij/bank_branch/internal"
)

type Environment struct {
	*BankBranch
	*RequestGenerator
	*Schedule
	WorkersCount        int
	QueueCapacity       int
	ApplicationInterval *internal.Range
	ServingDuration     *internal.Range
	Profit              *internal.Range
	ModelingStep        int
	Distribution        string

	*timestamp
}

func NewEnvironment(bank *BankBranch, flow *RequestGenerator, sch *Schedule) *Environment {
	return &Environment{
		BankBranch:       bank,
		RequestGenerator: flow,
		Schedule:         sch,
		timestamp:        newTimestamp(),
	}
}

func (e *Environment) Update(workers, queue, applicationInterval, servingDurationLeft, servingDurationRight,
	profitRangeLeft, profitRangeRight, modelingStep, lunchDuration int, distribution string) (*Environment, error) {
	if e == nil {
		return nil, errors.New("nil Environment")
	}
	if _, err := e.Schedule.Update(lunchDuration); err != nil {
		return nil, err
	}
	if _, err := e.RequestGenerator.Update(applicationInterval, servingDurationLeft, servingDurationRight,
		profitRangeLeft, profitRangeRight, distribution); err != nil {
		return nil, err
	}
	if _, err := e.BankBranch.Update(workers, queue); err != nil {
		return nil, err
	}
	e.WorkersCount = workers
	e.QueueCapacity = queue
	e.ApplicationInterval = internal.NewRange(0, applicationInterval)
	e.ServingDuration = internal.NewRange(servingDurationLeft, servingDurationRight)
	e.Profit = internal.NewRange(profitRangeLeft, profitRangeRight)
	e.ModelingStep = modelingStep
	e.Distribution = distribution
	e.timestamp = newTimestamp()
	return e, nil
}

func (e *Environment) Step() error {
	if e.timestamp.isTheEndOfSimulation() {
		return ErrEndOfSimulation
	}
	if e.timestamp.isTheEndOfDay() {
		e.BankBranch.CloseShifts()
		e.timestamp = e.timestamp.nextDay()
		return nil
	}
	var difference int
	t := *e.timestamp
	e.timestamp, difference = e.timestamp.addValidMinutes(e.ModelingStep)
	for i := difference; i > 0; {
		for e.BankBranch.HasFreeWorker() && e.Queue.Len() > 0 {
			clientToServe := e.Queue.PopClient()
			if err := e.BankBranch.BeginServingClient(clientToServe); err != nil {
				break
			}
		}

		nextApp := e.RequestGenerator.NextApp
		minToServe := e.BankBranch.GetMinTimeToServe()

		if i < nextApp && i < minToServe {
			e.RequestGenerator.NextApp -= i
			e.BankBranch.ServeClients(i)
			break
		} else if nextApp < minToServe {
			e.RequestGenerator.GenerateApplication(t.addMinutes(nextApp))
			e.BankBranch.ServeClients(nextApp)
		} else if minToServe < nextApp {
			e.BankBranch.ServeClients(minToServe)
			e.RequestGenerator.NextApp -= minToServe
		} else {
			e.BankBranch.ServeClients(minToServe)
			for e.BankBranch.HasFreeWorker() && e.Queue.Len() > 0 {
				clientToServe := e.Queue.PopClient()
				if err := e.BankBranch.BeginServingClient(clientToServe); err != nil {
					break
				}
			}
			e.RequestGenerator.GenerateApplication(t.addMinutes(nextApp))
		}
		for e.BankBranch.HasFreeWorker() && e.Queue.Len() > 0 {
			clientToServe := e.Queue.PopClient()
			if err := e.BankBranch.BeginServingClient(clientToServe); err != nil {
				break
			}
		}
		i -= min(minToServe, nextApp)
	}
	if difference < e.ModelingStep {
		// update stats
	}
	return nil
}

func (e *Environment) SkipDay() error {
	if e.timestamp.isTheEndOfSimulation() {
		return ErrEndOfSimulation
	}
	if err := e.Step(); err != nil {
		return err
	}
	for !e.timestamp.isTheEndOfDay() {
		e.Table.Clear()
		if err := e.Step(); err != nil {
			return err
		}
	}
	return nil
}

func (e *Environment) Break() {
	e.BankBranch.CloseShifts()
	e.Table.Clear()
	e.timestamp = newTimestamp()
}

func (e *Environment) IsOver() bool {
	return e.timestamp.isTheEndOfSimulation()
}

func (e *Environment) GetDay() string {
	return e.timestamp.getDayInfo()
}

func (e *Environment) GetTime() string {
	return e.timestamp.getTime()
}

func (e *Environment) GetWorkTime() string {
	return e.Schedule.GetWorkTime(e.timestamp)
}

func (e *Environment) GetLunchTime() string {
	return e.Schedule.GetLunchTime(e.timestamp)
}
