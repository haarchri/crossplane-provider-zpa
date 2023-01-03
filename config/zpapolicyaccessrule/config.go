package zpapolicyaccessrule

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_policy_access_rule", func(r *config.Resource) {
		r.ShortGroup = "zpaaccessrule"
	})
}
