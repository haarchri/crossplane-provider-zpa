package zpapolicyaccessforwardingrule

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_policy_forwarding_rule", func(r *config.Resource) {
		r.ShortGroup = "zpaforwardingrule"
	})
}
