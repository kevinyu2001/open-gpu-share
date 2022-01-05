package plugins

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sync"
)

type DeviceInfo struct {
	idx    int
	podMap map[types.UID]*v1.Pod
	// usedGPUMem  uint
	totalGPUMem uint
	rwmu        *sync.RWMutex
}

func newDeviceInfo(index int, totalGPUMem uint) *DeviceInfo {
	return &DeviceInfo{
		idx:         index,
		totalGPUMem: totalGPUMem,
		podMap:      map[types.UID]*v1.Pod{},
		rwmu:        new(sync.RWMutex),
	}
}