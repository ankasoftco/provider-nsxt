package evpn

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_evpn_config", func(r *config.Resource) {
		r.ShortGroup = "evpn"
		r.Kind = "PolicyEvpnConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_evpn_tenant", func(r *config.Resource) {
		r.ShortGroup = "evpn"
		r.Kind = "PolicyEvpnTenant"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_evpn_tunnel_endpoint", func(r *config.Resource) {
		r.ShortGroup = "evpn"
		r.Kind = "PolicyEvpnTunnelEndPoint"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_vni_pool", func(r *config.Resource) {
		r.ShortGroup = "evpn"
		r.Kind = "PolicyVniPool"
		r.Version = "v1alpha1"
	})
}
