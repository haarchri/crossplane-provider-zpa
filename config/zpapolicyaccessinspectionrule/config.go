package zpapolicyaccessinspectionrule

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zpa_policy_inspection_rule", func(r *config.Resource) {
		r.ShortGroup = "zpainspectionrule"
	})
}
