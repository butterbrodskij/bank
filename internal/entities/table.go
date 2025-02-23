package entities

import (
	"fmt"
)

type flowItem struct {
	ClientId int
	WorkerId int
}

type Updates struct {
	NewClients    []int
	ServedClients []int
	LostClients   int
}

type Table struct {
	flow []flowItem
	*Updates
}

func NewUpdates() *Updates {
	return &Updates{}
}

func NewTable() *Table {
	flow := make([]flowItem, 0)
	return &Table{
		flow:    flow,
		Updates: NewUpdates(),
	}
}

func (t *Table) ClearFlow() {
	t.flow = nil
}

func (t *Table) ClearUpdates() {
	t.Updates = NewUpdates()
}

func (t *Table) WorkerIsServing(worker *Worker) {
	if worker == nil || worker.client == nil {
		return
	}
	t.flow = append(t.flow, flowItem{
		ClientId: worker.client.id,
		WorkerId: worker.id,
	})
}

func (t *Table) ClientCreated(id int) {
	t.Updates.NewClients = append(t.Updates.NewClients, id)
}

func (t *Table) ClientLost(count int) {
	t.Updates.LostClients += count
}

func (t *Table) ClientServed(id int) {
	t.Updates.ServedClients = append(t.Updates.ServedClients, id)
}

func (t *Table) String() string {
	result := "Information table\n(client -> worker)"
	if len(t.flow) == 0 {
		result += "\n-"
		return result
	}
	for i := len(t.flow) - 1; i >= 0; i-- {
		result += fmt.Sprintf("\n%d -> %d", t.flow[i].ClientId, t.flow[i].WorkerId)
	}
	return result
}

func (t *Table) StringUpdates() string {
	result := "Updates\nnew clients:\n"
	if len(t.Updates.NewClients) == 0 {
		result += "-"
	}
	for _, u := range t.Updates.NewClients {
		result += fmt.Sprintf("%d ", u)
	}
	result += "\n\nserved clients:\n"
	if len(t.Updates.ServedClients) == 0 {
		result += "-"
	}
	for _, u := range t.Updates.ServedClients {
		result += fmt.Sprintf("%d ", u)
	}
	result += fmt.Sprintf("\n\nnumber of lost clients: %d", t.Updates.LostClients)
	return result
}
