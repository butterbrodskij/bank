package entities

type Client struct {
	*Application
	id      int
	waiting int
	toServe int
}
