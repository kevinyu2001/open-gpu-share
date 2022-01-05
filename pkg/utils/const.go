package utils

const (
	ResourceName = "alibabacloud.com/gpu-mem"
	CountName    = "alibabacloud.com/gpu-count"

	EnvNVGPU              = "NVIDIA_VISIBLE_DEVICES"
	EnvResourceIndex      = "ALIYUN_COM_GPU_MEM_IDX"
	EnvResourceByPod      = "ALIYUN_COM_GPU_MEM_POD"
	EnvResourceByDev      = "ALIYUN_COM_GPU_MEM_DEV"
	EnvAssignedFlag       = "ALIYUN_COM_GPU_MEM_ASSIGNED"
	EnvResourceAssumeTime = "ALIYUN_COM_GPU_MEM_ASSUME_TIME"
)
