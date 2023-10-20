package ospf

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_ospf_area", func(r *config.Resource) {
		r.ShortGroup = "ospf"
		r.Kind = "PolicyOspfArea"
		r.Version = "v1alpha1"
	})
	
}
