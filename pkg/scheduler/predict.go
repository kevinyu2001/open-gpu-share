package scheduler

import (
	"fmt"
	"log"

	"github.com/alibaba/open-gpu-share/pkg/cache"
	"github.com/alibaba/open-gpu-share/pkg/utils"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	kubeschedv1 "k8s.io/kube-scheduler/extender/v1"
)

func NewGPUsharePredicate(clientset *kubernetes.Clientset, c *cache.SchedulerCache) *Predicate {
	return &Predicate{
		Name: "gpusharingfilter",
		Func: func(pod *v1.Pod, nodeName string, c *cache.SchedulerCache) (bool, error) {
			log.Printf("info: check if the pod name %s can be scheduled on node %s", pod.Name, nodeName)
			nodeInfo, err := c.GetGpuNodeInfo(nodeName)
			if err != nil {
				return false, err
			}

			if !utils.IsGPUSharingNode(nodeInfo.GetNode()) {
				return false, fmt.Errorf("The node %s is not for GPU share, need skip", nodeName)
			}

			allocatable := nodeInfo.Assume(pod)
			if !allocatable {
				return false, fmt.Errorf("Insufficient GPU Memory in one device")
			} else {
				log.Printf("info: The pod %s in the namespace %s can be scheduled on %s",
					pod.Name,
					pod.Namespace,
					nodeName)
			}
			return true, nil
		},
		cache: c,
	}
}

func (p Predicate) Handler(args kubeschedv1.ExtenderArgs) *kubeschedv1.ExtenderFilterResult {
	pod := args.Pod
	nodeNames := *args.NodeNames
	canSchedule := make([]string, 0, len(nodeNames))
	canNotSchedule := make(map[string]string)

	for _, nodeName := range nodeNames {
		result, err := p.Func(pod, nodeName, p.cache)
		if err != nil {
			canNotSchedule[nodeName] = err.Error()
		} else {
			if result {
				canSchedule = append(canSchedule, nodeName)
			}
		}
	}

	result := kubeschedv1.ExtenderFilterResult{
		NodeNames:   &canSchedule,
		FailedNodes: canNotSchedule,
		Error:       "",
	}

	return &result
}
