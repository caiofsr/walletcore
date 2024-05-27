package event

import "time"

type TransferCreated struct {
	Name    string
	Payload interface{}
}

func NewTransferCreated() *TransferCreated {
	return &TransferCreated{
		Name: "TransferCreated",
	}
}

func (e *TransferCreated) GetName() string {
	return e.Name
}

func (e *TransferCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *TransferCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *TransferCreated) GetDateTime() time.Time {
	return time.Now()
}
