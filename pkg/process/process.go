package process

import (
	"os"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"kcl-lang.io/krm-kcl/pkg/config"
	"kcl-lang.io/krm-kcl/pkg/kube"
)

func init() {
	// Set the fast eval mode for KCL
	os.Setenv("KCL_FAST_EVAL", "1")
}

// Process is a function that takes a pointer to a ResourceList and processes
// it using the KCL function. It returns a boolean indicating whether the
// processing was successful, and an error (if any).
func Process(resourceList *fn.ResourceList) (bool, error) {
	resourceListYAML, err := resourceList.ToYAML()
	if err != nil {
		return false, err
	}
	res, err := kube.ParseResourceList(resourceListYAML)
	if err != nil {
		return false, err
	}
	err = func() error {
		r := &config.KCLRun{}
		if err := r.Config(res.FunctionConfig); err != nil {
			return err
		}
		return r.TransformResourceList(res)
	}()
	if err != nil {
		resourceList.Results = []*fn.Result{
			{
				Message:  err.Error(),
				Severity: fn.Error,
			},
		}
		return false, nil
	}
	processedItems, err := fn.ParseKubeObjects([]byte(res.Items.MustString()))
	if err != nil {
		return false, err
	}
	resourceList.Items = processedItems
	return true, nil
}
