package deprecated

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameConfigured returns a list of all external name resources
// configured for this provider.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("nsxt_algorithm_type_ns_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_algorithm_type_ns_service"
		r.Kind = "AlgorithmTypeNsService"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_dhcp_relay_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_dhcp_relay_profile"
		r.Kind = "DhcpRelayProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_dhcp_relay_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_dhcp_relay_service"
		r.Kind = "DhcpRelayService"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_dhcp_server_ip_pool", func(r *config.Resource) {
		r.ShortGroup = "nsxt_dhcp_server_ip_pool"
		r.Kind = "DhcpServerIpPool"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_dhcp_server_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_dhcp_server_profile"
		r.Kind = "DhcpServerProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ether_type_ns_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ether_type_ns_service"
		r.Kind = "EtherTypeNsService"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_firewall_section", func(r *config.Resource) {
		r.ShortGroup = "nsxt_firewall_section"
		r.Kind = "FirewallSection"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_icmp_type_ns_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_icmp_type_ns_service"
		r.Kind = "IcmpTypeNsService"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_igmp_type_ns_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_igmp_type_ns_service"
		r.Kind = "IgmpTypeNsService"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ip_block", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ip_block"
		r.Kind = "IpBlock"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ip_block_subnet", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ip_block_subnet"
		r.Kind = "IpBlockSubnet"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ip_discovery_switching_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ip_discovery_switching_profile"
		r.Kind = "IpDiscoverySwitchingProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ip_pool", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ip_pool"
		r.Kind = "IpPool"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ip_pool_allocation_ip_address", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ip_pool_allocation_ip_address"
		r.Kind = "IpPoolAllocationIpAddress"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ip_protocol_ns_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ip_protocol_ns_service"
		r.Kind = "IpProtocolNsService"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ip_set", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ip_set"
		r.Kind = "IpSet"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_l4_port_set_ns_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_l4_port_set_ns_service"
		r.Kind = "L4PortSetNsService"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_client_ssl_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_client_ssl_profile"
		r.Kind = "LbClientSslProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_cookie_persistence_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_cookie_persistence_profile"
		r.Kind = "LbCookiePersistenceProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_fast_tcp_application_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_fast_tcp_application_profile"
		r.Kind = "LbFastTcpApplicationProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_fast_udp_application_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_fast_udp_application_profile"
		r.Kind = "LbFastUdpApplicationProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_http_application_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_http_application_profile"
		r.Kind = "LbHttpApplicationProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_http_forwarding_rule", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_http_forwarding_rule"
		r.Kind = "LbHttpForwardingRule"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_http_monitor", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_http_monitor"
		r.Kind = "LbHttpMonitor"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_http_request_rewrite_rule", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_http_request_rewrite_rule"
		r.Kind = "LbHttpRequestRewriteRule"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_http_response_rewrite_rule", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_http_response_rewrite_rule"
		r.Kind = "LbHttpResponseRewriteRule"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_http_virtual_server", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_http_virtual_server"
		r.Kind = "LbHttpVirtualServer"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_https_monitor", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_https_monitor"
		r.Kind = "LbHttpsMonitor"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_icmp_monitor", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_icmp_monitor"
		r.Kind = "LbIcmpMonitor"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_passive_monitor", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_passive_monitor"
		r.Kind = "LbPassiveMonitor"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_pool", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_pool"
		r.Kind = "LbPool"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_server_ssl_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_server_ssl_profile"
		r.Kind = "LbServerSslProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_service", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_service"
		r.Kind = "LbService"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_source_ip_persistence_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_source_ip_persistence_profile"
		r.Kind = "LbSourceIpPersistenceProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_tcp_monitor", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_tcp_monitor"
		r.Kind = "LbTcpMonitor"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_tcp_virtual_server", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_tcp_virtual_server"
		r.Kind = "LbTcpVirtualServer"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_udp_monitor", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_udp_monitor"
		r.Kind = "LbUdpMonitor"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_lb_udp_virtual_server", func(r *config.Resource) {
		r.ShortGroup = "nsxt_lb_udp_virtual_server"
		r.Kind = "LbUdpVirtualServer"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_dhcp_port", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_dhcp_port"
		r.Kind = "LogicalDhcpPort"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_dhcp_server", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_dhcp_server"
		r.Kind = "LogicalDhcpServer"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_port", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_port"
		r.Kind = "LogicalPort"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_router_centralized_service_port", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_router_centralized_service_port"
		r.Kind = "LogicalRouterCentralizedServicePort"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_router_downlink_port", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_router_downlink_port"
		r.Kind = "LogicalRouterDownlinkPort"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_router_link_port_on_tier0", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_router_link_port_on_tier0"
		r.Kind = "LogicalRouterLinkPortOnTier0"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_router_link_port_on_tier1", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_router_link_port_on_tier1"
		r.Kind = "LogicalRouterLinkPortOnTier1"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_switch", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_switch"
		r.Kind = "LogicalSwitch"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_tier0_router", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_tier0_router"
		r.Kind = "LogicalTier0Router"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_logical_tier1_router", func(r *config.Resource) {
		r.ShortGroup = "nsxt_logical_tier1_router"
		r.Kind = "LogicalTier1Router"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_mac_management_switching_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_mac_management_switching_profile"
		r.Kind = "MacManagementSwitchingProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_nat_rule", func(r *config.Resource) {
		r.ShortGroup = "nsxt_nat_rule"
		r.Kind = "NatRule"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ns_group", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ns_group"
		r.Kind = "NsGroup"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_ns_service_group", func(r *config.Resource) {
		r.ShortGroup = "nsxt_ns_service_group"
		r.Kind = "NsServiceGroup"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_qos_switching_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_qos_switching_profile"
		r.Kind = "QosSwitchingProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_spoofguard_switching_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_spoofguard_switching_profile"
		r.Kind = "SpoofguardSwitchingProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_static_route", func(r *config.Resource) {
		r.ShortGroup = "nsxt_static_route"
		r.Kind = "StaticRoute"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_switch_security_switching_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxt_switch_security_switching_profile"
		r.Kind = "SwitchSecuritySwitchingProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_vlan_logical_switch", func(r *config.Resource) {
		r.ShortGroup = "nsxt_vlan_logical_switch"
		r.Kind = "VlanLogicalSwitch"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_vm_tags", func(r *config.Resource) {
		r.ShortGroup = "nsxt_vm_tags"
		r.Kind = "VmTags"
		r.Version = "v1alpha1"
	})

}
