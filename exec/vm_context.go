package exec

import (
	"encoding/json"
	"fmt"

	"github.com/Ankr-network/wagon/exec/gas"
)

const (
	maxVMNest = 64
)

type VMContext struct{
	runningVM *VM
	callVM    []*VM
	vmIndex   int
	gasMetric gas.GasMetric
	JsonObjectCache []map[string]json.RawMessage
}

func NewVMContext() *VMContext {
	return &VMContext{callVM: make([]*VM, maxVMNest)}
}

func (vmc *VMContext) SetRunningVM(vm *VM) {
	vmc.runningVM = vm
}

func (vmc *VMContext) RunningVM() *VM {
	return vmc.runningVM
}

func (vmc *VMContext) PushVM(vm *VM) (int, error) {
	if vmc.vmIndex > maxVMNest-1 {
		return -1, fmt.Errorf("beyond maxVMNest: vmIndex=%d, maxVMNest=%d", vmc.vmIndex, maxVMNest)
	}

	vmc.callVM[vmc.vmIndex] = vm
	vmc.vmIndex++

	return vmc.vmIndex, nil
}

func (vmc *VMContext) PopVM() (*VM, error) {
	if vmc.vmIndex < 0 {
		return  nil, fmt.Errorf("blank vms: vmIndex=%d", vmc.vmIndex)
	}

	vm := vmc.callVM[vmc.vmIndex]
	vmc.callVM[vmc.vmIndex] = nil
	vmc.vmIndex--

	return vm, nil
}

func (vmc *VMContext) TopVM() (*VM, error) {
	if vmc.vmIndex < 0 {
		return nil, fmt.Errorf("blank vms: vmIndex=%d", vmc.vmIndex)
	}

	return vmc.callVM[vmc.vmIndex], nil
}

func (vmc *VMContext) SetGasMetric(metric gas.GasMetric) {
	vmc.gasMetric = metric
}

func (vmc *VMContext) GasMetric() gas.GasMetric{
	return vmc.gasMetric
}
