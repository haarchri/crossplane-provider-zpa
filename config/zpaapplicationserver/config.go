package zpaapplicationserver

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_application_server", func(r *config.Resource) {
		r.ShortGroup = "zpaapplicationserver"
		r.References = config.References{
			"app_server_group_ids": {
				Type:              "github.com/zscaler/crossplane-provider-zpa/apis/zpaservergroup/v1alpha1.Group",
				RefFieldName:      "AppServerGroupIDRefs",
				SelectorFieldName: "AppServerGroupIDSelector",
			},
		}
	})
}
