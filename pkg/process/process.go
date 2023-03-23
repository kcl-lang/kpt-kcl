package process

import (
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"kusionstack.io/kpt-kcl-sdk/pkg/config"
)

func Process(resourceList *fn.ResourceList) (bool, error) {
	err := func() error {
		r := &config.KCLRun{}
		if err := r.Config(resourceList.FunctionConfig); err != nil {
			return err
		}
		return r.Transform(resourceList)
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
	return true, nil
}
