package entities

type Application struct {
	created    *timestamp
	difficulty int
	profit     int
}

func NewApplication(t *timestamp, difficulty int, profit int) *Application {
	return &Application{
		created:    t,
		difficulty: difficulty,
		profit:     profit,
	}
}
