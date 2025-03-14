package events

const CustomerCreatedEventBirthdateFormat = "2006-01-02"

type CustomerCreatedEvent struct {
	ID, Name, Email, CPF, Birthday string
}
