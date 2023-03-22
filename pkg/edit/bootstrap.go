package edit

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"kusionstack.io/kclvm-go"
	"sigs.k8s.io/kustomize/kyaml/errors"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

//go:embed _code.tmpl
var codeTemplateString string

var codeTemplate = template.Must(template.New("code.tmpl").Funcs(sprig.TxtFuncMap()).Parse(codeTemplateString))

const resourceListOptionName = "resource_list"

// runKCL runs the KCL script in the kpt environment
func runKCL(name, source string, resourceList *yaml.RNode) (string, error) {
	resourceListOptionKCLValue, err := toKCLValueString(resourceList)
	if err != nil {
		return "", errors.Wrap(err)
	}
	buffer := new(bytes.Buffer)
	codeTemplate.Execute(buffer, &struct{ Source string }{source})
	r, err := kclvm.RunFiles([]string{name}, kclvm.WithCode(buffer.String()), kclvm.WithOptions(fmt.Sprintf("%s=%s", resourceListOptionName, resourceListOptionKCLValue)))
	if err != nil {
		return "", errors.Wrap(err)
	}
	return r.GetRawYamlResult(), nil
}

// toKCLValueString converts YAML value to KCL value.
func toKCLValueString(resourceList *yaml.RNode) (string, error) {
	kclCode, err := resourceList.MarshalJSON()
	if err != nil {
		return "", errors.Wrap(err)
	}
	// In KCL, `true`, `false` and `null` are denoted by `True`, `False` and `None`.
	result := strings.Replace(string(kclCode), ": true", ": True", -1)
	result = strings.Replace(result, ": false", ": False", -1)
	result = strings.Replace(result, ": null", ": None", -1)
	return result, nil
}
