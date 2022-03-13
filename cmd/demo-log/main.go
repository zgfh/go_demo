package main

import (
	"errors"
	"flag"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

/**
log 组件
1. klog: k8s.io/klog  https://github.com/kubernetes/klog 设置: k8s.io/component-base/logs
2. kit log: github.com/go-kit/kit/log
3. https://github.com/coreos/pkg/tree/master/capnslog
4. https://pkg.go.dev/go.uber.org/zap

*/
/**
规范： https://github.com/kubernetes/community/blob/master/contributors/devel/sig-instrumentation/logging.md


klog.Warningf() - Something unexpected, but probably not an error

klog.Infof() has multiple levels:
*/
func demoklog() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Set("alsologtostderr", "true")
	flag.Set("v", "5")
	//flag.Lookup("v").Value.Set("5")

	klog.Errorf("demo err : %s", errors.New("fake error"))     // Always an error
	klog.Warningf("Warning  : %s", errors.New("fake Warning")) //Something unexpected, but probably not an error
	klog.Info("nice to meet you, I'm klog")

	klog.V(0).Info(" v0 (默认级别) 通常对此有用，始终对操作员可见: Programmer errors/Logging extra info about a panic/CLI argument handling")
	klog.V(1).Info(" v1      Information about config (listening on X, watching Y)/Errors that repeat frequently that relate to conditions that can be corrected (pod detected as unhealthy)")
	klog.V(2).Info(" v2 info: Logging HTTP requests and their exit code/System state changing (killing pod)/Controller state change events (starting pods)/Scheduler log messages")
	klog.V(3).Info(" v3       More info about system state changes")
	klog.V(4).Info(" v4 debug: ")
	klog.V(5).Info(" v5 Trace: ")
	klog.Flush()
}

func main() {
	demoklog()
}
