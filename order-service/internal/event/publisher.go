package event

type Publisher interface {
	Publish(event any) error
}
