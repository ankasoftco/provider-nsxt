/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda

	"nsxt_policy_dhcp_relay":                                  config.IdentifierFromProvider,
	"nsxt_policy_dhcp_server":                                 config.IdentifierFromProvider,
	"nsxt_policy_dhcp_v4_static_binding":                      config.IdentifierFromProvider,
	"nsxt_policy_dhcp_v6_static_binding":                      config.IdentifierFromProvider,
	"nsxt_policy_dns_forwarder_zone":                      	   config.IdentifierFromProvider,
	"nsxt_policy_gateway_dns_forwarder":                       config.IdentifierFromProvider,
	"nsxt_policy_evpn_config":                                 config.IdentifierFromProvider,
	"nsxt_policy_evpn_tenant":                                 config.IdentifierFromProvider,
	"nsxt_policy_evpn_tunnel_endpoint":                        config.IdentifierFromProvider,
	"nsxt_policy_vni_pool":                                    config.IdentifierFromProvider,
	"nsxt_policy_context_profile":                             config.IdentifierFromProvider,
	"nsxt_policy_context_profile_custom_attribute":            config.IdentifierFromProvider,
	"nsxt_policy_gateway_policy":                              config.IdentifierFromProvider,
	"nsxt_policy_intrusion_service_policy":                    config.IdentifierFromProvider,
	"nsxt_policy_intrusion_service_profile":                   config.IdentifierFromProvider,
	"nsxt_policy_predefined_gateway_policy":                   config.IdentifierFromProvider,
	"nsxt_policy_predefined_security_policy":                  config.IdentifierFromProvider,
	"nsxt_policy_security_policy":                             config.IdentifierFromProvider,
	"nsxt_policy_service":                                     config.IdentifierFromProvider,
	"nsxt_policy_bgp_config":                                  config.IdentifierFromProvider,
	"nsxt_policy_bgp_neighbor":                                config.IdentifierFromProvider,
	"nsxt_policy_gateway_community_list":                      config.IdentifierFromProvider,
	"nsxt_policy_gateway_prefix_list":                         config.IdentifierFromProvider,
	"nsxt_policy_gateway_qos_profile":                         config.IdentifierFromProvider,
	"nsxt_policy_gateway_redistribution_config":               config.IdentifierFromProvider,
	"nsxt_policy_gateway_route_map":                           config.IdentifierFromProvider,
	"nsxt_policy_static_route_bfd_peer":                       config.IdentifierFromProvider,
	"nsxt_policy_nat_rule":                                    config.IdentifierFromProvider,
	"nsxt_policy_ospf_config":                                 config.IdentifierFromProvider,
	"nsxt_policy_static_route":                                config.IdentifierFromProvider,
	"nsxt_policy_tier0_gateway":                               config.IdentifierFromProvider,
	"nsxt_policy_tier0_gateway_ha_vip_config":                 config.IdentifierFromProvider,
	"nsxt_policy_tier0_gateway_interface":                     config.IdentifierFromProvider,
	"nsxt_policy_tier1_gateway":                               config.IdentifierFromProvider,
	"nsxt_policy_tier1_gateway_interface":                     config.IdentifierFromProvider,
	"nsxt_policy_domain":                                      config.IdentifierFromProvider,
	"nsxt_policy_group":                                       config.IdentifierFromProvider,
	"nsxt_policy_vm_tags":                                     config.IdentifierFromProvider,
	"nsxt_policy_ip_address_allocation":                       config.IdentifierFromProvider,
	"nsxt_policy_ip_block":                                    config.IdentifierFromProvider,
	"nsxt_policy_ip_pool":                                     config.IdentifierFromProvider,
	"nsxt_policy_ip_pool_block_subnet":                        config.IdentifierFromProvider,
	"nsxt_policy_ip_pool_static_subnet":                       config.IdentifierFromProvider,
	"nsxt_policy_lb_pool":                                     config.IdentifierFromProvider,
	"nsxt_policy_lb_service":                                  config.IdentifierFromProvider,
	"nsxt_policy_lb_virtual_server":                           config.IdentifierFromProvider,
	"nsxt_policy_project":                                     config.IdentifierFromProvider,
	"nsxt_policy_fixed_segment":                               config.IdentifierFromProvider,
	"nsxt_policy_ip_discovery_profile":                        config.IdentifierFromProvider,
	"nsxt_policy_mac_discovery_profile":                       config.IdentifierFromProvider,
	"nsxt_policy_qos_profile":                                 config.IdentifierFromProvider,
	"nsxt_policy_segment":                                     config.IdentifierFromProvider,
	"nsxt_policy_segment_security_profile":                    config.IdentifierFromProvider,
	"nsxt_policy_spoof_guard_profile":                         config.IdentifierFromProvider,
	"nsxt_policy_vlan_segment":                                config.IdentifierFromProvider,

}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
