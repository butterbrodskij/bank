package entities

import (
	"errors"
	"fmt"
)

type Queue struct {
	clients  []*Client
	capacity int
	counter  int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		clients:  make([]*Client, 0, capacity),
		capacity: capacity,
	}
}

func (q *Queue) Update(capacity int) (*Queue, error) {
	if q == nil {
		return nil, errors.New("nil Queue")
	}
	q.capacity = capacity
	q.clients = make([]*Client, 0, capacity)
	return q, nil
}

func (q *Queue) EmptyQueue() int {
	if q == nil {
		return 0
	}
	l := len(q.clients)
	q.clients = nil
	q.counter = 0
	return l
}

func (q *Queue) AddClient(app *Application) int {
	if q == nil || q.capacity <= len(q.clients) {
		return 0
	}
	q.counter++
	q.clients = append(q.clients, &Client{
		Application: app,
		id:          q.counter,
		waiting:     0,
		toServe:     app.difficulty,
	})
	return q.counter
}

func (q *Queue) PopClient() *Client {
	if q == nil || len(q.clients) == 0 {
		return nil
	}
	cl := q.clients[0]
	q.clients = q.clients[1:]
	return cl
}

func (q *Queue) String() string {
	if q == nil {
		return ""
	}
	result := "Queue"
	if len(q.clients) == 0 {
		result += "\n-"
	}
	for _, cl := range q.clients {
		result += fmt.Sprintf("\n%d", cl.id)
	}
	return result
}

func (q *Queue) Len() int {
	if q == nil {
		return 0
	}
	return len(q.clients)
}
