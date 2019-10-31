package gas

import (
  "math/big"

  "github.com/Ankr-network/wagon/exec/internal/compile"
  ops "github.com/Ankr-network/wagon/wasm/operators"
)

const (
  GasQuickStep    uint64 = 2
  GasFastestStep  uint64 = 3
  GasFastStep     uint64 = 5
  GasMidStep      uint64 = 8
  GasSlowStep     uint64 = 10
  GasExtStep      uint64 = 20

  GasReturn       uint64 = 0
  GasStop         uint64 = 0
  GasContractByte uint64 = 200
)

type GasMetric interface {
  SpendGas(gas *big.Int) bool
}

var OpsGasTable = map[byte]uint64{
  ops.I32Clz : GasQuickStep,
  ops.I32Ctz : GasQuickStep,
  ops.I32Popcnt : GasQuickStep,
  ops.I32Add : GasFastestStep,
  ops.I32Sub : GasFastestStep,
  ops.I32Mul : GasFastestStep,
  ops.I32DivS : GasFastestStep,
  ops.I32DivU : GasFastestStep,
  ops.I32RemS : GasFastestStep,
  ops.I32RemU : GasFastestStep,
  ops.I32And : GasFastestStep,
  ops.I32Or : GasFastestStep,
  ops.I32Xor : GasFastestStep,
  ops.I32Shl : GasFastestStep,
  ops.I32ShrS : GasFastestStep,
  ops.I32ShrU : GasFastestStep,
  ops.I32Rotl : GasFastestStep,
  ops.I32Rotr : GasFastestStep,
  ops.I64Clz : GasQuickStep,
  ops.I64Ctz : GasQuickStep,
  ops.I64Popcnt : GasQuickStep,
  ops.I64Add : GasFastestStep,
  ops.I64Sub : GasFastestStep,
  ops.I64Mul : GasFastestStep,
  ops.I64DivS : GasFastestStep,
  ops.I64DivU : GasFastestStep,
  ops.I64RemS : GasFastestStep,
  ops.I64RemU : GasFastestStep,
  ops.I64And : GasFastestStep,
  ops.I64Or : GasFastestStep,
  ops.I64Xor : GasFastestStep,
  ops.I64Shl : GasFastestStep,
  ops.I64ShrS : GasFastestStep,
  ops.I64ShrU : GasFastestStep,
  ops.I64Rotl : GasFastestStep,
  ops.I64Rotr : GasFastestStep,
  ops.I32Eqz : GasFastestStep,
  ops.I32Eq : GasFastestStep,
  ops.I32Ne : GasFastestStep,
  ops.I32LtS : GasFastestStep,
  ops.I32LtU : GasFastestStep,
  ops.I32GtS : GasFastestStep,
  ops.I32GtU : GasFastestStep,
  ops.I32LeS : GasFastestStep,
  ops.I32LeU : GasFastestStep,
  ops.I32GeS : GasFastestStep,
  ops.I32GeU : GasFastestStep,
  ops.I64Eqz : GasFastestStep,
  ops.I64Eq : GasFastestStep,
  ops.I64Ne : GasFastestStep,
  ops.I64LtS : GasFastestStep,
  ops.I64LtU : GasFastestStep,
  ops.I64GtS : GasFastestStep,
  ops.I64GtU : GasFastestStep,
  ops.I64LeS : GasFastestStep,
  ops.I64LeU : GasFastestStep,
  ops.I64GeS : GasFastestStep,
  ops.I64GeU : GasFastestStep,
  ops.I32Const : GasFastestStep,
  ops.I64Const : GasFastestStep,
  ops.I32WrapI64 : GasFastestStep,
  ops.I64ExtendSI32 : GasFastestStep,
  ops.I64ExtendUI32 : GasFastestStep,
  ops.I32Load : GasFastestStep,
  ops.I64Load : GasFastestStep,
  ops.I32Load8s : GasFastestStep,
  ops.I32Load8u : GasFastestStep,
  ops.I32Load16s : GasFastestStep,
  ops.I32Load16u : GasFastestStep,
  ops.I64Load8s : GasFastestStep,
  ops.I64Load8u : GasFastestStep,
  ops.I64Load16s : GasFastestStep,
  ops.I64Load16u : GasFastestStep,
  ops.I64Load32s : GasFastestStep,
  ops.I64Load32u : GasFastestStep,
  ops.I32Store : GasFastestStep,
  ops.I64Store : GasFastestStep,
  ops.I32Store8 : GasFastestStep,
  ops.I32Store16 : GasFastestStep,
  ops.I64Store8 : GasFastestStep,
  ops.I64Store16 : GasFastestStep,
  ops.I64Store32 : GasFastestStep,
  ops.CurrentMemory : GasFastestStep,
  ops.GrowMemory : GasFastestStep,
  ops.Drop : GasFastestStep,
  ops.Select : GasFastestStep,
  ops.GetLocal : GasFastestStep,
  ops.SetLocal : GasFastestStep,
  ops.TeeLocal : GasFastestStep,
  ops.GetGlobal : GasFastestStep,
  ops.SetGlobal : GasFastestStep,
  ops.Unreachable : GasFastestStep,
  ops.Nop : GasFastestStep,
  ops.Block : GasFastestStep,
  ops.Loop : GasFastestStep,
  ops.If : GasFastestStep,
  ops.Else : GasFastestStep,
  ops.End : GasFastestStep,
  ops.Br : GasFastestStep,
  ops.BrIf : GasFastestStep,
  ops.BrTable : GasFastestStep,
  ops.Return : GasFastestStep,
  ops.Call : GasFastestStep,
  ops.CallIndirect : GasFastestStep,
  ops.WagonNativeExec : GasFastestStep,
}

var CompileGasTable = map[byte]uint64{
  compile.OpJmp : GasQuickStep,
  compile.OpJmpZ : GasQuickStep,
  compile.OpJmpNz : GasQuickStep,
  compile.OpDiscard : GasQuickStep,
  compile.OpDiscardPreserveTop : GasQuickStep,
}
