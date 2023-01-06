package zpaservergroup

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_server_group", func(r *config.Resource) {
		r.ShortGroup = "zpaservergroup"
		r.References = config.References{
			"app_connector_groups.id": {
				Type: "github.com/zscaler/crossplane-provider-zpa/apis/zpaappconnectorgroup/v1alpha1.ConnectorGroup",
			},
			// (ToDo haarchri): Import Cycles not allowed https://github.com/upbound/upjet/issues/96
			// "servers.id": {
			// 	Type: "github.com/zscaler/crossplane-provider-zpa/apis/zpaapplicationserver/v1alpha1.Server",
			// },
		}
	})
}
