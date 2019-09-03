package heapmemory

import "errors"

type HeapMemory struct {
	allocator *allocatorBuddy
}

func NewHeapMemory(totalSize uint) *HeapMemory {
	return &HeapMemory{ allocator: newAllocatorBuddy(totalSize) }
}

func (hm *HeapMemory) Malloc(size uint) (int, error) {
	if hm.allocator == nil {
		return -1, errors.New("HeapMemory's allocator nil")
	}

	return hm.allocator.alloc(size)
}

func (hm *HeapMemory) Free(offset int) error {
	if hm.allocator == nil {
		return errors.New("HeapMemory's allocator nil")
	}

	return hm.allocator.free(offset)
}

func (hm *HeapMemory) GrowMemory(size uint) error {
	if hm.allocator == nil {
		return errors.New("HeapMemory's allocator nil")
	}

	return hm.allocator.growTotalSize(size)
}
