package groupingandtagging

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_domain", func(r *config.Resource) {
		r.ShortGroup = "groupingandtagging"
		r.Kind = "PolicyDomain"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_group", func(r *config.Resource) {
		r.ShortGroup = "groupingandtagging"
		r.Kind = "PolicyGroup"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_vm_tags", func(r *config.Resource) {
		r.ShortGroup = "groupingandtagging"
		r.Kind = "PolicyVmTags"
		r.Version = "v1alpha1"
	})
}
