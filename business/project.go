package business

type Project interface {
	Processes() []Process
	Clients() []Client
	Objectives() []Objective
	KeyResults() []KeyResult
}

type KeyResult interface {
}

type project struct {
}

func (p project) KeyResults() []KeyResult {
	return nil
}

func (p project) Objectives() []Objective {
	return nil
}

func (p project) Clients() []Client {
	return nil
}

func (p project) Processes() []Process {
	return nil
}
