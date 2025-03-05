package entities

import (
	"errors"
	"github.com/butterbrodskij/bank_branch/internal"
	"math"
)

type BankBranch struct {
	workers []*Worker
	*Queue
	*Table
	stats *Statistics
}

func NewBankBranch(workersCount int, q *Queue, stats *Statistics) *BankBranch {
	workers := make([]*Worker, workersCount)
	return &BankBranch{
		workers: workers,
		Queue:   q,
		Table:   NewTable(),
		stats:   stats,
	}
}

func (b *BankBranch) Update(workersCount, capacity int) (*BankBranch, error) {
	workers := make([]*Worker, workersCount)
	for i, _ := range workers {
		workers[i] = NewWorker(i + 1)
	}
	b.workers = workers
	if _, err := b.Queue.Update(capacity); err != nil {
		return nil, err
	}
	b.Table.ClearFlow()
	return b, nil
}

func (b *BankBranch) HasFreeWorker() bool {
	for _, worker := range b.workers {
		if worker.IsFree() {
			return true
		}
	}
	return false
}

func (b *BankBranch) BeginServingClient(c *Client) error {
	for _, worker := range b.workers {
		if worker.IsFree() {
			worker.AcceptClient(c)
			b.Table.WorkerIsServing(worker)
			return nil
		}
	}
	return errors.New("cannot begin serving client")
}

func (b *BankBranch) GetMinTimeToServe() int {
	res := math.MaxInt
	for _, worker := range b.workers {
		if !worker.IsFree() {
			res = min(res, worker.client.toServe)
		}
	}
	return res
}

func (b *BankBranch) ServeClients(min int) {
	b.Queue.AddWaitingTime(min)
	for _, worker := range b.workers {
		servedClient := worker.ServeClient(min)
		if servedClient != nil {
			b.Table.ClientServed(worker, servedClient.id)
			b.stats.AddProfit(servedClient.profit)
			b.stats.AddServingStat(servedClient.difficulty)
			b.stats.AddWaitingStat(servedClient.waiting)
		}
	}
}

func (b *BankBranch) ServeAll() {
	for _, worker := range b.workers {
		servedClient := worker.ServeClient(math.MaxInt)
		if servedClient != nil {
			b.Table.ClientServed(worker, servedClient.id)
			b.stats.AddProfit(servedClient.profit)
			b.stats.AddServingStat(servedClient.difficulty)
			b.stats.AddWaitingStat(servedClient.waiting)
		}
	}
}

func (b *BankBranch) CloseShifts() {
	b.ServeAll()
	b.ClientLost(b.Queue.EmptyQueue())
	b.stats.AddLostClients(b.Queue.EmptyQueue())
	b.stats.AddProfit(-len(b.workers) * internal.WorkerSalary)
}

func (b *BankBranch) GetInfo() string {
	changes := b.Table.StringInfo()
	b.Table.ClearInfo()
	return changes
}

func (b *BankBranch) GetFlow() string {
	changes := b.Table.StringFlow()
	b.Table.ClearFlow()
	return changes
}

func (b *BankBranch) GetUpdates() string {
	updates := b.Table.StringUpdates()
	b.Table.ClearUpdates()
	return updates
}

func (b *BankBranch) NotifyClientUpdated(id int) {
	if id == 0 {
		b.Table.ClientLost(1)
		b.stats.AddLostClients(1)
	} else {
		b.Table.ClientCreated(id)
	}
}
