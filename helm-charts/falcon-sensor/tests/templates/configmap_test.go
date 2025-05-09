package template_test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	corev1 "k8s.io/api/core/v1"
)

func TestConfigMapTemplateRender(t *testing.T) {
	releaseName := "test-release"
	helmChartPath := "../../../falcon-sensor"
	options := &helm.Options{
		SetValues: map[string]string{"falcon.cid": cid},
	}

	output := helm.RenderTemplate(
		t, options, helmChartPath, releaseName,
		[]string{"templates/configmap.yaml"})

	var configMap corev1.ConfigMap
	helm.UnmarshalK8SYaml(t, output, &configMap)
	t.Logf("configMap: %v", configMap)
	// expectedContainerImage := "nginx:1.15.8"
	// podContainers := configMap.Spec.Containers
	// if podContainers[0].Image != expectedContainerImage {
	// 	t.Fatalf(
	// 		"Rendered container image (%s) is not expected (%s)",
	// 		podContainers[0].Image,
	// 		expectedContainerImage,
	// 	)
	// }
}
