package entities

type Client struct {
	*Request
	id      int
	waiting int
	toServe int
}
