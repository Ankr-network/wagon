// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"errors"
	"fmt"
	"math"

	"github.com/go-interpreter/wagon/exec/common"
	"github.com/go-interpreter/wagon/wasm"
)

// ErrOutOfBoundsMemoryAccess is the error value used while trapping the VM
// when it detects an out of bounds access to the linear memory.
var ErrOutOfBoundsMemoryAccess = errors.New("exec: out of bounds memory access")

const (
	wasmStackSize     = 16384 //16 * 1024
	vmStackStartIndex = 16384 //16 *1024
	MinHeapMemorySize = 64 * 1024 //64k
	MaxHeapMemorySize = 1024 *1024 //1M
)

var (
	InvalidMemIndex = errors.New("invalid memory index")
)

func (vm *VM) fetchBaseAddr() int {
	return int(vm.fetchUint32() + uint32(vm.popInt32()))
}

// inBounds returns true when the next vm.fetchBaseAddr() + offset
// indices are in bounds accesses to the linear memory.
func (vm *VM) inBounds(offset int) bool {
	addr := endianess.Uint32(vm.ctx.code[vm.ctx.pc:]) + uint32(vm.ctx.stack[len(vm.ctx.stack)-1])
	return int(addr)+offset < len(vm.memory)
}

// curMem returns a slice to the memory segment pointed to by
// the current base address on the bytecode stream.
func (vm *VM) curMem() []byte {
	return vm.memory[vm.fetchBaseAddr():]
}

func (vm *VM) i32Load() {
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint32(endianess.Uint32(vm.curMem()))
}

func (vm *VM) i32Load8s() {
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushInt32(int32(int8(vm.memory[vm.fetchBaseAddr()])))
}

func (vm *VM) i32Load8u() {
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint32(uint32(uint8(vm.memory[vm.fetchBaseAddr()])))
}

func (vm *VM) i32Load16s() {
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushInt32(int32(int16(endianess.Uint16(vm.curMem()))))
}

func (vm *VM) i32Load16u() {
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint32(uint32(endianess.Uint16(vm.curMem())))
}

func (vm *VM) i64Load() {
	if !vm.inBounds(7) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint64(endianess.Uint64(vm.curMem()))
}

func (vm *VM) i64Load8s() {
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushInt64(int64(int8(vm.memory[vm.fetchBaseAddr()])))
}

func (vm *VM) i64Load8u() {
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint64(uint64(uint8(vm.memory[vm.fetchBaseAddr()])))
}

func (vm *VM) i64Load16s() {
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushInt64(int64(int16(endianess.Uint16(vm.curMem()))))
}

func (vm *VM) i64Load16u() {
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint64(uint64(endianess.Uint16(vm.curMem())))
}

func (vm *VM) i64Load32s() {
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushInt64(int64(int32(endianess.Uint32(vm.curMem()))))
}

func (vm *VM) i64Load32u() {
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint64(uint64(endianess.Uint32(vm.curMem())))
}

func (vm *VM) f32Store() {
	v := math.Float32bits(vm.popFloat32())
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint32(vm.curMem(), v)
}

func (vm *VM) f32Load() {
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushFloat32(math.Float32frombits(endianess.Uint32(vm.curMem())))
}

func (vm *VM) f64Store() {
	v := math.Float64bits(vm.popFloat64())
	if !vm.inBounds(7) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint64(vm.curMem(), v)
}

func (vm *VM) f64Load() {
	if !vm.inBounds(7) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushFloat64(math.Float64frombits(endianess.Uint64(vm.curMem())))
}

func (vm *VM) i32Store() {
	v := vm.popUint32()
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint32(vm.curMem(), v)
}

func (vm *VM) i32Store8() {
	v := byte(uint8(vm.popUint32()))
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.memory[vm.fetchBaseAddr()] = v
}

func (vm *VM) i32Store16() {
	v := uint16(vm.popUint32())
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint16(vm.curMem(), v)
}

func (vm *VM) i64Store() {
	v := vm.popUint64()
	if !vm.inBounds(7) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint64(vm.curMem(), v)
}

func (vm *VM) i64Store8() {
	v := byte(uint8(vm.popUint64()))
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.memory[vm.fetchBaseAddr()] = v
}

func (vm *VM) i64Store16() {
	v := uint16(vm.popUint64())
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint16(vm.curMem(), v)
}

func (vm *VM) i64Store32() {
	v := uint32(vm.popUint64())
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint32(vm.curMem(), v)
}

func (vm *VM) currentMemory() {
	_ = vm.fetchInt8() // reserved (https://github.com/WebAssembly/design/blob/27ac254c854994103c24834a994be16f74f54186/BinaryEncoding.md#memory-related-operators-described-here)
	vm.pushInt32(int32(len(vm.memory) / wasmPageSize))
}

func (vm *VM) growMemory() {
	_ = vm.fetchInt8() // reserved (https://github.com/WebAssembly/design/blob/27ac254c854994103c24834a994be16f74f54186/BinaryEncoding.md#memory-related-operators-described-here)
	curLen := len(vm.memory) / wasmPageSize
	n := vm.popInt32()
	vm.module.HeapMem.GrowMemory(uint(n*wasmPageSize))
	vm.pushInt32(int32(curLen))
}

func (vm *VM) heapBase() (int32, error) {
	if vm.module == nil {
		return -1, errors.New("vm module nil")
	}

	hbExportEntry, ok := vm.module.Export.Entries["__heap_base"]
	if !ok {
	   return -1, errors.New("there is no __heap_base")
	}

	if hbExportEntry.Kind != wasm.ExternalGlobal {
		return -1, errors.New("invlid __heap_base")
	}

	gEntry := vm.module.GetGlobal(int(hbExportEntry.Index))
	if gEntry == nil {
		return -1, fmt.Errorf("can't find global entry: index=%d", hbExportEntry.Index)
	}

	heapBaseIndex, err := vm.module.ExecInitExpr(gEntry.Init)
	if err != nil {
		return -1, err
	}

	return heapBaseIndex.(int32), nil
}

func (vm *VM) initMemory() error {
	if vm.module  == nil {
		return errors.New("vm module nil")
	}

	initSize := uint(vm.module.Memory.Entries[0].Limits.Initial)*wasmPageSize
	if !common.IsPowOf2(initSize) {
		initSize = common.FixSize(initSize)
	}

	if initSize < MinHeapMemorySize {
		initSize = MinHeapMemorySize
	}

	heapBaseIndex, err := vm.heapBase()
	if err != nil {
		return err
	}
	vm.memory  = make([]byte, uint(heapBaseIndex) + initSize)
	vm.module.HeapMem.Init(initSize)

    if vm.module.LinearMemoryIndexSpace[0] != nil {
    	copy(vm.memory[vmStackStartIndex:], vm.module.LinearMemoryIndexSpace[0][wasmStackSize:])
	}

	return nil
}

func (vm *VM) Strlen(memIndex uint) (int, error) {
	memLen :=len(vm.memory)
	if memIndex >= uint(memLen) {
		return 0, InvalidMemIndex
	}

	l := 0
	s := memIndex
	for vm.memory[s] != byte(0) && l < (memLen-1) {
		l++
		s++
	}

	if memIndex < vmStackStartIndex && (memIndex + uint(l)) >= vmStackStartIndex {
		return 0, InvalidMemIndex
	}

	heapBaseIndex, err := vm.heapBase()
	if err != nil {
		return 0, err
	}
	if memIndex < uint(heapBaseIndex) && (memIndex + uint(l)) >= uint(heapBaseIndex) {
		return 0, InvalidMemIndex
	}

	return l, nil
}

func (vm *VM) Strcmp(memIndex1 uint, memIndex2 uint) (int, error) {
	memLen :=len(vm.memory)
	if memIndex1 >= uint(memLen) || memIndex2 >= uint(memLen) {
		return 0, InvalidMemIndex
	}

	len1, err1 := vm.Strlen(memIndex1)
	len2, err2 := vm.Strlen(memIndex2)
	if err1 != nil || err2 != nil {
		return 0, InvalidMemIndex
	}

	minLen := common.MinI(len1, len2)
	for i:= uint(0); i < uint(minLen); i++ {
		if vm.memory[memIndex1+i] == vm.memory[memIndex2+i] {
			continue
		} else if vm.memory[memIndex1+i] < vm.memory[memIndex2+i] {
			return -1, nil
		} else {
			return 1, nil
		}
	}

	if len1 > len2 {
		return 1, nil
	} else if len1 < len2 {
		return -1, nil
	}else {
		return 0, nil
	}
}

func (vm *VM) SetBytes(bytes []byte) (uint64, error) {
	lenBytes   := len(bytes)
	index, err := vm.module.HeapMem.Malloc(uint(lenBytes + 1))
	if err != nil {
		return 0, err
	}

	heapBaseIndex, _ := vm.heapBase()
	index = index + uint64(heapBaseIndex)

	copy(vm.memory[index:index+uint64(lenBytes)], bytes)
	vm.memory[index+uint64(lenBytes)] = byte(0)

    return index, nil
}
