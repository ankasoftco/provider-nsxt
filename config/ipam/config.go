package ipam

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "ipam"
const version = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_ip_address_allocation", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpAddressAllocation"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ip_block", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpBlock"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ip_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpPool"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_ip_pool_block_subnet", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpPoolBlockSubnet"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ip_pool_static_subnet", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpPoolStaticSubnet"
		r.Version = version
	})
}
