package zpainspectionprofile

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_inspection_profile", func(r *config.Resource) {
		r.ShortGroup = "zpainspectionprofile"
	})
}
