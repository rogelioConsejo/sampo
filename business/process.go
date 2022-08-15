package business

import "time"

type Process interface {
	StartingDate() time.Time
	EndingDate() time.Time
	Clients() []Client
}

type process struct {
}

func (p process) Clients() []Client {
	return nil
}

func (p process) StartingDate() time.Time {
	return time.Now()
}

func (p process) EndingDate() time.Time {
	return time.Now()
}
