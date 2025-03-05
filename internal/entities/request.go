package entities

type Request struct {
	created    *timestamp
	difficulty int
	profit     int
}

func NewRequest(t *timestamp, difficulty int, profit int) *Request {
	return &Request{
		created:    t,
		difficulty: difficulty,
		profit:     profit,
	}
}
