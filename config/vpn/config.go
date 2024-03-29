package vpn

import "github.com/upbound/upjet/pkg/config"

const version string = "v1alpha1"

// Configure configures the VPN settings based on the provided configuration provider.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_dpd_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_policy_ipsec_vpn_dpd_profile"
		r.Kind = "PolicyIpsecVpnDpdProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_local_endpoint", func(r *config.Resource) {
		r.ShortGroup = "nsxt_policy_ipsec_vpn_local_endpoint"
		r.Kind = "PolicyIpsecVpnLocalEndpoint"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_ike_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_policy_ipsec_vpn_ike_profile"
		r.Kind = "PolicyIpsecVpnIkeProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_session", func(r *config.Resource) {
		r.ShortGroup = "nsxt_policy_ipsec_vpn_session"
		r.Kind = "PolicyIpsecVpnSession"
		r.Version = version

	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_policy_ipsec_vpn_service"
		r.Kind = "PolicyIpsecVpnService"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_tunnel_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_policy_ipsec_vpn_tunnel_profile"
		r.Kind = "PolicyIpsecVpnTunnelProfile"
		r.Version = version

	})

	p.AddResourceConfigurator("nsxt_policy_l2_vpn_server", func(r *config.Resource) {
		r.ShortGroup = "nsxt_policy_l2_vpn_server"
		r.Kind = "PolicyL2VpnServer"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_l2_vpn_session", func(r *config.Resource) {
		r.ShortGroup = "nsxt_policy_l2_vpn_session"
		r.Kind = "PolicyL2VpnSession"
		r.Version = version
	})

}
