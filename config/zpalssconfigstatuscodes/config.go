package zpalssconfigstatuscodes

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_lss_config_status_codes", func(r *config.Resource) {
		r.ShortGroup = "zpalssconfigstatuscodes"
	})
}
