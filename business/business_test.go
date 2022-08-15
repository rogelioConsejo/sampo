package business

import (
	"testing"
)

func TestBusiness(t *testing.T) {
	var testBusiness Business = New()
	var _ []Project = testBusiness.Projects()
	var _ []Process = testBusiness.Processes()
	var _ []Objective = testBusiness.Objectives()
	var _ []Client = testBusiness.Clients()
}

type objective struct {
}

func (o objective) Name() string {
	return ""
}

func (o objective) Description() string {
	return ""
}

func TestObjective(t *testing.T) {
	var testObjective Objective = objective{}
	var _ string = testObjective.Name()
	var _ string = testObjective.Description()
}
