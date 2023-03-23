package runner

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"kusionstack.io/kpt-kcl-sdk/pkg/process"
)

/// Run evaluates the ResourceList from STDIN to STDOUT
func Run() error {
	return fn.AsMain(fn.ResourceListProcessorFunc(process.Process))
}
