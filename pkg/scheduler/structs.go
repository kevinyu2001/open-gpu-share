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
	TotalGpu int64     `json:"totalGpu"`
	UsedGpu  int64     `json:"usedGpu"`
	Devices  []*Device `json:"devs"`
}

type Device struct {
	ID       int    `json:"id"`
	TotalGpu int64  `json:"totalGpu"`
	UsedGpu  int64  `json:"usedGpu"`
	Pods     []*Pod `json:"pods"`
}

type Pod struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	UsedGpu   int    `json:"usedGpu"`
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
