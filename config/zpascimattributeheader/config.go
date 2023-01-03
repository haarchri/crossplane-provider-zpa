package zpascimattributeheader

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_scim_attribute_header", func(r *config.Resource) {
		r.ShortGroup = "zpascimattributeheader"
	})
}
