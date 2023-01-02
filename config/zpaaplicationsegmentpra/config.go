package zpaapplicationsegmenpra

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_application_segment_pra", func(r *config.Resource) {
		r.ShortGroup = "applicationsegmentpra"
	})
}
