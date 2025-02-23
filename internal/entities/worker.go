package entities

type Worker struct {
	id     int
	client *Client
}

func NewWorker(id int) *Worker {
	return &Worker{
		id: id,
	}
}

func (w *Worker) ServeClient(min int) *Client {
	if w.client == nil {
		return nil
	}
	if w.client.toServe > min {
		w.client.toServe -= min
		return nil
	}
	client := w.client
	w.client.toServe = 0
	w.client = nil
	return client
}

func (w *Worker) IsFree() bool {
	return w.client == nil
}

func (w *Worker) AcceptClient(client *Client) {
	w.client = client
}
