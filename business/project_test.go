package business

import "testing"

func TestBiz_Projects(t *testing.T) {
	var testProject Project = project{}
	var _ []Process = testProject.Processes()
	var _ []Client = testProject.Clients()
	var _ []Objective = testProject.Objectives()
	var _ []KeyResult = testProject.KeyResults()
}
