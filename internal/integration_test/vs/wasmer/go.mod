module github.com/voedger/wazero/internal/integration_test/vs/wasmer

go 1.20

require (
	github.com/voedger/wazero v0.0.0
	github.com/wasmerio/wasmer-go v1.0.4
)

replace github.com/voedger/wazero => ../../../..
