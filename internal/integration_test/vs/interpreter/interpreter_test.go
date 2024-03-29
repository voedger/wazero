package interpreter

import (
	"testing"

	"github.com/voedger/wazero/internal/integration_test/vs"
)

var runtime = vs.NewWazeroInterpreterRuntime

func TestAllocation(t *testing.T) {
	vs.RunTestAllocation(t, runtime)
}

func TestAllocationEx(t *testing.T) {
	vs.RunTestAllocationEx(t, runtime)
}

func TestFactorial(t *testing.T) {
	vs.RunTestFactorial(t, runtime)
}
