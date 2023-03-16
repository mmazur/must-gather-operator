//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/must-gather-operator/api/v1alpha1.ProxySpec": schema_openshift_must_gather_operator_api_v1alpha1_ProxySpec(ref),
	}
}

func schema_openshift_must_gather_operator_api_v1alpha1_ProxySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"httpProxy": {
						SchemaProps: spec.SchemaProps{
							Description: "httpProxy is the URL of the proxy for HTTP requests.  Empty means unset and will not result in an env var.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"httpsProxy": {
						SchemaProps: spec.SchemaProps{
							Description: "httpsProxy is the URL of the proxy for HTTPS requests.  Empty means unset and will not result in an env var.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"noProxy": {
						SchemaProps: spec.SchemaProps{
							Description: "noProxy is the list of domains for which the proxy should not be used.  Empty means unset and will not result in an env var.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}