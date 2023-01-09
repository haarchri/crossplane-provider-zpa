package zpaapplicationsegment

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_application_segment", func(r *config.Resource) {
		r.ShortGroup = "zpaapplicationsegment"
		r.References = config.References{
			"server_groups.id": {
				Type: "github.com/zscaler/crossplane-provider-zpa/apis/zpaservergroup/v1alpha1.Group",
			},
			"segment_group_id": {
				Type: "github.com/zscaler/crossplane-provider-zpa/apis/zpasegmentgroup/v1alpha1.Group",
			},
		}
	})
}
