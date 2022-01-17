# Open GPU Share

## Overview
`open-gpu-share` mainly defines the data structure of a scheduler cache. It is used to extend the scheduler, under [Kubernetes scheduling framework](https://github.com/kubernetes/enhancements/tree/master/keps/sig-scheduling/624-scheduling-framework#proposal), to allocate pods that can share a GPU, i.e., apply for a portion of full GPU memory.

## Acknowledgments

This project heavily relies on [AliyunContainerService/gpushare-scheduler-extender](https://github.com/AliyunContainerService/gpushare-scheduler-extender), referring to [Nvidia Docker2](https://github.com/NVIDIA/nvidia-docker) and their [GPU sharing design](https://docs.google.com/document/d/1ZgKH_K4SEfdiE_OfxQ836s4yQWxZfSjS288Tq9YIWCA/edit#heading=h.r88v2xgacqr).

## License

[Apache 2.0 License](LICENSE)