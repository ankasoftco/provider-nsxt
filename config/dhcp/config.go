package dhcp

import "github.com/upbound/upjet/pkg/config"

const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_dhcp_relay", func(r *config.Resource) {
		r.ShortGroup = "dhcp"
		r.Kind = "PolicyDhcpRelay"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_dhcp_server", func(r *config.Resource) {
		r.ShortGroup = "dhcp"
		r.Kind = "PolicyDhcpServer"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_dhcp_v4_static_binding", func(r *config.Resource) {
		r.ShortGroup = "dhcp"
		r.Kind = "PolicyDhcpV4StaticBinding"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_dhcp_v6_static_binding", func(r *config.Resource) {
		r.ShortGroup = "dhcp"
		r.Kind = "PolicyDhcpV6StaticBinding"
		r.Version = version
	})
}
