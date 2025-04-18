package go_pubsub

// PublishRequest ...
type PublishRequest struct {
	Topic      string
	Data       []byte
	Attributes map[string]string
}

// PublishResponse ...
type PublishResponse struct {
	ID string
}
