package exbench

import (
	_ "embed"
	"testing"

	"github.com/voedger/wazero"
	"github.com/voedger/wazero/api"
	"github.com/voedger/wazero/internal/testing/require"
)

func Benchmark_hwazero_Arg_JustCall(b *testing.B) {
	hcallbaсkCount = 0
	var err error
	rtc := wazero.NewRuntimeConfigInterpreter().WithMemoryLimitPages(2)
	rtm := wazero.NewRuntimeWithConfig(rtc)
	require.NotNil(b, rtm)

	host, err := rtm.NewModuleBuilder("env").
		ExportFunction("callbackp", hcallbackSP).
		ExportFunction("callbackp1", hcallbackSP).
		ExportFunction("callback", hcallbackSP).
		Instantiate(testCtx)

	require.Nil(b, err)
	defer host.Close(testCtx)

	module, err := rtm.InstantiateModuleFromCode(callCtx, calls)
	require.NoError(b, err)
	defer module.Close(testCtx)

	justCall := module.ExportedFunction("justCall")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = justCall.CallExArg(testCtx, nil, nil, nil)
		if nil != err {
			break
		}
	}
	require.Nil(b, err)
}

func Benchmark_hwazero_Arg_CallBackNoParam(b *testing.B) {

	hcallbaсkCount = 0
	var err error
	rtm := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigInterpreter())
	require.NotNil(b, rtm)

	host, err := rtm.NewModuleBuilder("env").
		ExportFunction("callbackp", hcallbackSP).
		ExportFunction("callbackp1", hcallbackSP).
		ExportFunction("callback", hcallbackSP).
		Instantiate(testCtx)
	defer host.Close(testCtx)

	module, err := rtm.InstantiateModuleFromCode(callCtx, calls)
	require.NoError(b, err)
	defer module.Close(testCtx)

	var ce api.ICallEngine = module.NewCallEngine()
	callback := module.ExportedFunction("doCallback")
	cnp := api.CallEngineParams{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = callback.CallExArg(testCtx, ce, &cnp, nil)
		if nil != err {
			break
		}
	}
	require.Nil(b, err)
}

func Benchmark_hwazero_Arg_CallBack1Param(b *testing.B) {

	hcallbaсkCount = 0
	var err error
	rtm := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigInterpreter())
	require.NotNil(b, rtm)

	host, err := rtm.NewModuleBuilder("env").
		ExportFunction("callbackp", hcallbackSP).
		ExportFunction("callbackp1", hcallbackSP).
		ExportFunction("callback", hcallbackSP).
		Instantiate(testCtx)

	require.Nil(b, err)
	defer host.Close(testCtx)

	module, err := rtm.InstantiateModuleFromCode(callCtx, calls)
	require.NoError(b, err)
	defer module.Close(testCtx)

	var ce api.ICallEngine = module.NewCallEngine()
	callbackp := module.ExportedFunction("doCallbackp1")
	cnp := api.CallEngineParams{}

	par := []uint64{2, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = callbackp.CallExArg(testCtx, ce, &cnp, par)
		if nil != err {
			break
		}
	}
	require.Nil(b, err)
}
func Benchmark_hwazero_Arg_CallBack3Param(b *testing.B) {

	hcallbaсkCount = 0
	var err error
	rtm := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigInterpreter())
	require.NotNil(b, rtm)

	host, err := rtm.NewModuleBuilder("env").
		ExportFunction("callbackp", hcallbackSP).
		ExportFunction("callbackp1", hcallbackSP).
		ExportFunction("callback", hcallbackSP).
		Instantiate(testCtx)

	require.Nil(b, err)
	defer host.Close(testCtx)

	module, err := rtm.InstantiateModuleFromCode(callCtx, calls)
	require.NoError(b, err)
	defer module.Close(testCtx)

	var ce api.ICallEngine = module.NewCallEngine()
	callbackp3 := module.ExportedFunction("doCallbackp")
	cnp := api.CallEngineParams{}

	b.ResetTimer()
	args := []uint64{2, 3}
	for i := 0; i < b.N; i++ {
		_, err = callbackp3.CallExArg(testCtx, ce, &cnp, args)
		if nil != err {
			break
		}
	}
	require.Nil(b, err)
}

func Benchmark_hwazero_Arg_fib20(b *testing.B) {
	rtm := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigInterpreter())

	module, _ := rtm.InstantiateModuleFromCode(testCtx, fib)
	defer module.Close(testCtx)

	var ce api.ICallEngine = module.NewCallEngine()
	fibonacci := module.ExportedFunction("fibonacci")
	cep := api.CallEngineParams{}

	args := []uint64{20}
	for i := 0; i < b.N; i++ {
		fibonacci.CallExArg(testCtx, ce, &cep, args)
	}
}

func Benchmark_hwazero_Arg_Root(b *testing.B) {
	rtm := wazero.NewRuntimeWithConfig(wazero.NewRuntimeConfigInterpreter())

	module, _ := rtm.InstantiateModuleFromCode(testCtx, root)
	defer module.Close(testCtx)

	var ce api.ICallEngine = module.NewCallEngine()
	root := module.ExportedFunction("root")
	cep := api.CallEngineParams{}

	par := []uint64{1000}
	for i := 0; i < b.N; i++ {
		root.CallExArg(testCtx, ce, &cep, par)
	}
}

func Benchmark_hwazero_Old_Add2Param(b *testing.B) {
	var err error
	rtm := wazero.NewRuntime()
	require.NotNil(b, rtm)

	module, err := rtm.NewModuleBuilder("host/math").
		ExportFunction("add", func(v1, v2 uint32) uint32 {
			return v1 + v2
		}).Instantiate(testCtx)

	require.Nil(b, err)
	defer module.Close(testCtx)

	var ce api.ICallEngine = module.NewCallEngine()
	addEx := module.ExportedFunction("add")
	cnp := api.CallEngineParams{}

	var res []uint64
	b.ResetTimer()
	args := []uint64{12, 14}
	for i := 0; i < b.N; i++ {
		res, err = addEx.CallExArg(testCtx, ce, &cnp, args)

	}
	require.Equal(b, res[0], uint64(26))
	require.Nil(b, err)
}

func Benchmark_hwazero_Arg_AddParam(b *testing.B) {
	var err error
	rtm := wazero.NewRuntime()
	require.NotNil(b, rtm)

	module, err := rtm.NewModuleBuilder("host/math").
		ExportFunction("add", func(pars []uint64) (res []uint64) {
			var r uint64
			for i := 0; i < len(pars); i++ {
				r = r + pars[i]
			}
			return []uint64{r}
		}).Instantiate(testCtx)

	require.Nil(b, err)
	defer module.Close(testCtx)

	var ce api.ICallEngine = module.NewCallEngine()
	add := module.ExportedFunction("add")
	cnp := api.CallEngineParams{}

	var res []uint64
	b.ResetTimer()
	args := []uint64{12, 14}
	for i := 0; i < b.N; i++ {
		res, err = add.CallExArg(testCtx, ce, &cnp, args)
	}
	require.Equal(b, res[0], uint64(26))
	require.Nil(b, err)
}
