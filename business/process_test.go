package business

import (
	"testing"
	"time"
)

func TestBiz_Processes(t *testing.T) {
	var testProcess Process = process{}
	var _ time.Time = testProcess.StartingDate()
	var _ time.Time = testProcess.EndingDate()
	var _ []Client = testProcess.Clients()
}
