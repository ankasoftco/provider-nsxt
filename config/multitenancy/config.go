package multitenancy

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_project", func(r *config.Resource) {
		r.ShortGroup = "multitenancy"
		r.Kind = "PolicyProject"
		r.Version = "v1alpha1"
	})
	
}
