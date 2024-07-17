package gatewaysandrouting

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "gatewaysandrouting"
const version = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_bgp_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyBgpConfig"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_bgp_neighbor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyBgpMeighbor"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_gateway_community_list", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGatewayCommunityList"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_gateway_prefix_list", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGatewayPrefixList"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_gateway_qos_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGatewayQosProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_gateway_redistribution_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGatewayRedistributionConfig"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_gateway_route_map", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyGatewayRouteMap"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_static_route_bfd_peer", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyStaticRouteBfdPeer"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_nat_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyNatRule"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_ospf_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyOspfConfig"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_static_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyStaticRoute"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_tier0_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyTier0Gateway"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_tier0_gateway_ha_vip_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyTier0GatewayHaVipConfig"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_tier0_gateway_interface", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyTier0GatewayInterface"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_tier1_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyTier1Gateway"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_tier1_gateway_interface", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyTier1GatewayInterface"
		r.Version = version
	})
}
