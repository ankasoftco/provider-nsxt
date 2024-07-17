package vpn

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_dpd_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyipsecvpndpdprofile"
		r.Kind = "PolicyIpsecVpnDpdProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_local_endpoint", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyipsecvpnlocalendpoint"
		r.Kind = "PolicyIpsecVpnLocalEndpoint"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_ike_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyipsecvpnikeprofile"
		r.Kind = "PolicyIpsecVpnIkeProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_session", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyipsecvpnsession"
		r.Kind = "PolicyIpsecVpnSession"
		r.Version = "v1alpha1"
	
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_service", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyipsecvpnservice"
		r.Kind = "PolicyIpsecVpnService"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_ipsec_vpn_tunnel_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyipsecvpntunnelprofile"
		r.Kind = "PolicyIpsecVpnTunnelProfile"
		r.Version = "v1alpha1"

	})

	p.AddResourceConfigurator("nsxt_policy_l2_vpn_server", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyl2vpnserver"
		r.Kind = "PolicyL2VpnServer"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_l2_vpn_session", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyl2vpnsession"
		r.Kind = "PolicyL2VpnSession"
		r.Version = "v1alpha1"
	})

}
