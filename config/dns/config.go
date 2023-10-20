package dns

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_dns_forwarder_zone", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.Kind = "PolicyDnsForwarderZone"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_gateway_dns_forwarder", func(r *config.Resource) {
		r.ShortGroup = "dns"
		r.Kind = "PolicyGatewayDnsForwarder"
		r.Version = "v1alpha1"
	})
}
