package log

import (
	"k8s.io/klog/v2/klogr"
)

// Logger is the root logger which should be used by all
// other packages in the codebase.
var Logger = klogr.New()
