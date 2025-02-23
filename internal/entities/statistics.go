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
}

func NewStatistics() *Statistics {
	return &Statistics{}
}

func (s *Statistics) AddProfit(profit int) {
	s.profit += profit
}

func (s *Statistics) AddLostClients(lostClients int) {
	s.lostClients += lostClients
}

func (s *Statistics) AddQueueStat(queue int, delta int) {
	s.totalQueue += int64(queue)
	s.queueCount += delta
}

func (s *Statistics) AddWaitingStat(waiting int) {
	s.totalWaiting += int64(waiting)
	s.waitingCount++
}

func (s *Statistics) AddServingStat(serving int) {
	s.totalServing += int64(serving)
	s.servingCount++
}

func (s *Statistics) String() string {
	return fmt.Sprintf(`
Statistics
profit: %d
served clients: %d
lost clients: %d
avg queue length: %.2f
avg waiting time: %.2f
avg serving time: %.2f
		`, s.profit, s.servingCount, s.lostClients, float64(s.totalQueue)/float64(s.queueCount),
		float64(s.totalWaiting)/float64(s.waitingCount), float64(s.servingCount)/float64(s.servingCount))
}
