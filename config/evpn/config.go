package evpn

import "github.com/upbound/upjet/pkg/config"

const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_evpn_config", func(r *config.Resource) {
		r.ShortGroup = "evpn"
		r.Kind = "PolicyEvpnConfig"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_evpn_tenant", func(r *config.Resource) {
		r.ShortGroup = "evpn"
		r.Kind = "PolicyEvpnTenant"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_evpn_tunnel_endpoint", func(r *config.Resource) {
		r.ShortGroup = "evpn"
		r.Kind = "PolicyEvpnTunnelEndPoint"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_vni_pool", func(r *config.Resource) {
		r.ShortGroup = "evpn"
		r.Kind = "PolicyVniPool"
		r.Version = version
	})
}
