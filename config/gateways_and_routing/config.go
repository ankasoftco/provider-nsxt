package gatewaysandrouting

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_bgp_config", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyBgpConfig"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_bgp_neighbor", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyBgpMeighbor"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_gateway_community_list", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyGatewayCommunityList"
		r.Version = "v1alpha1"
	})
	
	p.AddResourceConfigurator("nsxt_policy_gateway_prefix_list", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyGatewayPrefixList"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_gateway_qos_profile", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyGatewayQosProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_gateway_redistribution_config", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyGatewayRedistributionConfig"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_gateway_route_map", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyGatewayRouteMap"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_static_route_bfd_peer", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyStaticRouteBfdPeer"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_nat_rule", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyNatRule"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_ospf_config", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyOspfConfig"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_static_route", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyStaticRoute"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_tier0_gateway", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyTier0Gateway"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_tier0_gateway_ha_vip_config", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyTier0GatewayHaVipConfig"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_tier0_gateway_interface", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyTier0GatewayInterface"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_tier1_gateway", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyTier1Gateway"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_tier1_gateway_interface", func(r *config.Resource) {
		r.ShortGroup = "gatewaysandrouting"
		r.Kind = "PolicyTier1GatewayInterface"
		r.Version = "v1alpha1"
	})
}
