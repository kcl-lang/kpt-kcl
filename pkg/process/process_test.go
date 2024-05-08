package process

import (
	"testing"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func TestProcess(t *testing.T) {
	p := fn.ResourceListProcessorFunc(Process)
	input := `kind: ResourceList
items:
- apiVersion: apps/v1
  kind: Deployment
  spec:
    replicas: '2'
- kind: Service
functionConfig:
  apiVersion: krm.kcl.dev/v1alpha1
  kind: KCLRun
  metadata:
    name: conditionally-add-annotations
  spec:
    params:
      replicas: "5"
    source: |
      params = option("params")
      replicas = params.replicas
      setReplicas = lambda items, replicas {
         [item | {if item.kind == "Deployment": spec.replicas = replicas} for item in items]
      }
      items = setReplicas(option("items"), replicas)
`
	expected := `apiVersion: config.kubernetes.io/v1
kind: ResourceList
items:
- kind: Service
- apiVersion: apps/v1
  kind: Deployment
  spec:
    replicas: '5'
functionConfig:
  apiVersion: krm.kcl.dev/v1alpha1
  kind: KCLRun
  metadata:
    name: conditionally-add-annotations
  spec:
    params:
      replicas: "5"
    source: |
      params = option("params")
      replicas = params.replicas
      setReplicas = lambda items, replicas {
         [item | {if item.kind == "Deployment": spec.replicas = replicas} for item in items]
      }
      items = setReplicas(option("items"), replicas)
`
	output, err := fn.Run(p, []byte(input))
	if err != nil {
		t.Fatal(err)
	}
	expectedYaml, err := yaml.Parse(expected)
	if err != nil {
		t.Fatal(err)
	}
	outputYaml, err := yaml.Parse(string(output))
	if err != nil {
		t.Fatal(err)
	}
	expectedYamlString := expectedYaml.MustString()
	outputYamlString := outputYaml.MustString()
	if expectedYamlString != outputYamlString {
		t.Errorf("test failed, expected %s got %s", expectedYamlString, outputYamlString)
	}
}
