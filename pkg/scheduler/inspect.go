package scheduler

import (
	"github.com/alibaba/open-gpu-share/pkg/cache"
	"github.com/alibaba/open-gpu-share/pkg/utils"
)

func (in Inspect) Handler(name string) *Result {
	nodes := []*Node{}
	errMsg := ""
	if len(name) == 0 {
		nodeInfos := in.cache.GetGpuNodeinfos()
		for _, info := range nodeInfos {
			nodes = append(nodes, buildNode(info))
		}

	} else {
		node, err := in.cache.GetGpuNodeInfo(name)
		if err != nil {
			errMsg = err.Error()
		}
		// nodeInfos = append(nodeInfos, node)
		nodes = append(nodes, buildNode(node))
	}

	return &Result{
		Nodes: nodes,
		Error: errMsg,
	}
}

func buildNode(info *cache.GpuNodeInfo) *Node {
	devInfos := info.GetDevs()
	devs := []*Device{}
	var usedGPU uint

	for i, devInfo := range devInfos {
		dev := &Device{
			ID:       i,
			TotalGPU: devInfo.GetTotalGPUMemory(),
			UsedGPU:  devInfo.GetUsedGPUMemory(),
		}

		podInfos := devInfo.GetPods()
		pods := []*Pod{}
		for _, podInfo := range podInfos {
			if utils.AssignedNonTerminatedPod(podInfo) {
				pod := &Pod{
					Namespace: podInfo.Namespace,
					Name:      podInfo.Name,
					UsedGPU:   utils.GetGPUMemoryFromPodResource(podInfo),
				}
				pods = append(pods, pod)
			}
		}
		dev.Pods = pods
		devs = append(devs, dev)
		usedGPU += devInfo.GetUsedGPUMemory()
	}

	return &Node{
		Name:     info.GetName(),
		TotalGPU: uint(info.GetTotalGPUMemory()),
		UsedGPU:  usedGPU,
		Devices:  devs,
	}

}

func NewGPUShareInspect(c *cache.SchedulerCache) *Inspect {
	return &Inspect{
		Name:  "gpushareinspect",
		cache: c,
	}
}
