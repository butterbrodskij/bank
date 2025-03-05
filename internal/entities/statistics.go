package entities

import "fmt"

type Statistics struct {
	profit       int
	lostClients  int
	totalQueue   int64
	totalWaiting int64
	totalServing int64
	queueCount   int
	waitingCount int
	servingCount int
	maxWaiting   int
}

func NewStatistics() *Statistics {
	return &Statistics{}
}

func (s *Statistics) Empty() {
	if s == nil {
		return
	}
	*s = Statistics{}
}

func (s *Statistics) AddProfit(profit int) {
	s.profit += profit
}

func (s *Statistics) AddLostClients(lostClients int) {
	s.lostClients += lostClients
}

func (s *Statistics) AddQueueStat(queue int, delta int) {
	s.totalQueue += int64(queue * delta)
	s.queueCount += delta
}

func (s *Statistics) AddWaitingStat(waiting int) {
	s.maxWaiting = max(waiting, s.maxWaiting)
	s.totalWaiting += int64(waiting)
	s.waitingCount++
}

func (s *Statistics) AddServingStat(serving int) {
	s.totalServing += int64(serving)
	s.servingCount++
}

func (s *Statistics) String() string {
	avgQueue := float64(0)
	if s.queueCount > 0 {
		avgQueue = float64(s.totalQueue) / float64(s.queueCount)
	}
	avgWaiting := float64(0)
	if s.waitingCount > 0 {
		avgWaiting = float64(s.totalWaiting) / float64(s.waitingCount)
	}
	avgServing := float64(0)
	if s.servingCount > 0 {
		avgServing = float64(s.totalServing) / float64(s.servingCount)
	}
	return fmt.Sprintf(`
Statistics
profit: %d
served clients: %d
lost clients: %d
avg queue length: %.2f
avg waiting time: %.2f
max waiting time: %d
avg serving time: %.2f
		`, s.profit, s.servingCount, s.lostClients, avgQueue, avgWaiting, s.maxWaiting, avgServing)
}
