package business

type Business interface {
	Projects() []Project
	Processes() []Process
	Clients() []Client
	Objectives() []Objective
}

func New() Business {
	return biz{}
}

type Client interface {
}

type Objective interface {
	Name() string
	Description() string
}

type biz struct {
}

func (b biz) Objectives() []Objective {
	return nil
}

func (b biz) Clients() []Client {
	return nil
}

func (b biz) Projects() []Project {
	return nil
}

func (b biz) Processes() []Process {
	return nil
}
