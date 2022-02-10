package main

import (
	"flag"

	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()
	flag.Parse()

	klog.Info("nice to meet you, I'm klog")
	klog.V(2).Info("testing level 2")
	klog.V(3).Info("testing level 3")

}
