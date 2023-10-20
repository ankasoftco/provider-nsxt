package ipam

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_ip_address_allocation", func(r *config.Resource) {
		r.ShortGroup = "ipam"
		r.Kind = "PolicyIpAddressAllocation"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_ip_block", func(r *config.Resource) {
		r.ShortGroup = "ipam"
		r.Kind = "PolicyIpBlock"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_ip_pool", func(r *config.Resource) {
		r.ShortGroup = "ipam"
		r.Kind = "PolicyIpPool"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_ip_pool_block_subnet", func(r *config.Resource) {
		r.ShortGroup = "ipam"
		r.Kind = "PolicyIpPoolBlockSubnet"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_ip_pool_static_subnet", func(r *config.Resource) {
		r.ShortGroup = "ipam"
		r.Kind = "PolicyIpPoolStaticSubnet"
		r.Version = "v1alpha1"
	})
}
