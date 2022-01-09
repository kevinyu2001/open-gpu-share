package scheduler

import (
	"github.com/alibaba/open-gpu-share/pkg/cache"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type Result struct {
	Nodes []*Node `json:"nodes"`
	Error string  `json:"error,omitempty"`
}

type Node struct {
	Name     string    `json:"name"`
	TotalGPU uint      `json:"totalGPU"`
	UsedGPU  uint      `json:"usedGPU"`
	Devices  []*Device `json:"devs"`
}

type Device struct {
	ID       int    `json:"id"`
	TotalGPU uint   `json:"totalGPU"`
	UsedGPU  uint   `json:"usedGPU"`
	Pods     []*Pod `json:"pods"`
}

type Pod struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	UsedGPU   int    `json:"usedGPU"`
}

type Inspect struct {
	Name  string
	cache *cache.SchedulerCache
}

// Bind is responsible for binding node and pod
type Bind struct {
	Name  string
	Func  func(podName string, podNamespace string, podUID types.UID, node string, cache *cache.SchedulerCache) error
	cache *cache.SchedulerCache
}

type Predicate struct {
	Name  string
	Func  func(pod *v1.Pod, nodeName string, c *cache.SchedulerCache) (bool, error)
	cache *cache.SchedulerCache
}
