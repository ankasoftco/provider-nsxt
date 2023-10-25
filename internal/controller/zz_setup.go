/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	policydhcprelay "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcprelay"
	policydhcpserver "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpserver"
	policydhcpv4staticbinding "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpv4staticbinding"
	policydhcpv6staticbinding "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpv6staticbinding"
	policydnsforwarderzone "github.com/ankasoftco/provider-nsxt/internal/controller/dns/policydnsforwarderzone"
	policygatewaydnsforwarder "github.com/ankasoftco/provider-nsxt/internal/controller/dns/policygatewaydnsforwarder"
	policyevpnconfig "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpnconfig"
	policyevpntenant "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpntenant"
	policyevpntunnelendpoint "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpntunnelendpoint"
	policyvnipool "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyvnipool"
	policycontextprofile "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policycontextprofile"
	policycontextprofilecustomattribute "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policycontextprofilecustomattribute"
	policygatewaypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policygatewaypolicy"
	policyinstrusionservicepolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyinstrusionservicepolicy"
	policyinstrusionserviceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyinstrusionserviceprofile"
	policypredefinedgatewaypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policypredefinedgatewaypolicy"
	policypredefinedsecuritypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policypredefinedsecuritypolicy"
	policysecuritypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policysecuritypolicy"
	policyservice "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyservice"
	policybgpconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policybgpconfig"
	policybgpmeighbor "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policybgpmeighbor"
	policygatewaycommunitylist "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewaycommunitylist"
	policygatewayprefixlist "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayprefixlist"
	policygatewayqosprofile "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayqosprofile"
	policygatewayredistributionconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayredistributionconfig"
	policygatewayroutemap "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayroutemap"
	policynatrule "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policynatrule"
	policyospfconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policyospfconfig"
	policystaticroute "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policystaticroute"
	policystaticroutebfdpeer "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policystaticroutebfdpeer"
	policytier0gateway "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gateway"
	policytier0gatewayhavipconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gatewayhavipconfig"
	policytier0gatewayinterface "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gatewayinterface"
	policytier1gateway "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier1gateway"
	policytier1gatewayinterface "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier1gatewayinterface"
	policydomain "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policydomain"
	policygroup "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policygroup"
	policyvmtags "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policyvmtags"
	policyipaddressallocation "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyipaddressallocation"
	policyipblock "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyipblock"
	policyippool "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippool"
	policyippoolblocksubnet "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippoolblocksubnet"
	policyippoolstaticsubnet "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippoolstaticsubnet"
	policylbpool "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbpool"
	policylbservice "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbservice"
	policylbvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbvirtualserver"
	policyproject "github.com/ankasoftco/provider-nsxt/internal/controller/multitenancy/policyproject"
	algorithmtypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_algorithm_type_ns_service/algorithmtypensservice"
	clustervirtualip "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_cluster_virtual_ip/clustervirtualip"
	computemanager "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_compute_manager/computemanager"
	dhcprelayprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_relay_profile/dhcprelayprofile"
	dhcprelayservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_relay_service/dhcprelayservice"
	dhcpserverippool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_server_ip_pool/dhcpserverippool"
	dhcpserverprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_server_profile/dhcpserverprofile"
	edgecluster "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_edge_cluster/edgecluster"
	ethertypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ether_type_ns_service/ethertypensservice"
	failuredomain "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_failure_domain/failuredomain"
	firewallsection "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_firewall_section/firewallsection"
	icmptypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_icmp_type_ns_service/icmptypensservice"
	igmptypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_igmp_type_ns_service/igmptypensservice"
	ipblock "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_block/ipblock"
	ipblocksubnet "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_block_subnet/ipblocksubnet"
	ipdiscoveryswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_discovery_switching_profile/ipdiscoveryswitchingprofile"
	ippool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_pool/ippool"
	ippoolallocationipaddress "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_pool_allocation_ip_address/ippoolallocationipaddress"
	ipprotocolnsservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_protocol_ns_service/ipprotocolnsservice"
	ipset "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_set/ipset"
	l4portsetnsservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_l4_port_set_ns_service/l4portsetnsservice"
	lbclientsslprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_client_ssl_profile/lbclientsslprofile"
	lbcookiepersistenceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_cookie_persistence_profile/lbcookiepersistenceprofile"
	lbfasttcpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_fast_tcp_application_profile/lbfasttcpapplicationprofile"
	lbfastudpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_fast_udp_application_profile/lbfastudpapplicationprofile"
	lbhttpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_application_profile/lbhttpapplicationprofile"
	lbhttpforwardingrule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_forwarding_rule/lbhttpforwardingrule"
	lbhttpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_monitor/lbhttpmonitor"
	lbhttprequestrewriterule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_request_rewrite_rule/lbhttprequestrewriterule"
	lbhttpresponserewriterule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_response_rewrite_rule/lbhttpresponserewriterule"
	lbhttpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_virtual_server/lbhttpvirtualserver"
	lbhttpsmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_https_monitor/lbhttpsmonitor"
	lbicmpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_icmp_monitor/lbicmpmonitor"
	lbpassivemonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_passive_monitor/lbpassivemonitor"
	lbpool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_pool/lbpool"
	lbserversslprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_server_ssl_profile/lbserversslprofile"
	lbservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_service/lbservice"
	lbsourceippersistenceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_source_ip_persistence_profile/lbsourceippersistenceprofile"
	lbtcpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_tcp_monitor/lbtcpmonitor"
	lbtcpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_tcp_virtual_server/lbtcpvirtualserver"
	lbudpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_udp_monitor/lbudpmonitor"
	lbudpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_udp_virtual_server/lbudpvirtualserver"
	logicaldhcpport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_dhcp_port/logicaldhcpport"
	logicaldhcpserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_dhcp_server/logicaldhcpserver"
	logicalport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_port/logicalport"
	logicalroutercentralizedserviceport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_centralized_service_port/logicalroutercentralizedserviceport"
	logicalrouterdownlinkport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_downlink_port/logicalrouterdownlinkport"
	logicalrouterlinkportontier0 "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_link_port_on_tier0/logicalrouterlinkportontier0"
	logicalrouterlinkportontier1 "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_link_port_on_tier1/logicalrouterlinkportontier1"
	logicalswitch "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_switch/logicalswitch"
	logicaltier0router "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_tier0_router/logicaltier0router"
	logicaltier1router "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_tier1_router/logicaltier1router"
	macmanagementswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_mac_management_switching_profile/macmanagementswitchingprofile"
	managercluster "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_manager_cluster/managercluster"
	natrule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_nat_rule/natrule"
	nsgroup "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ns_group/nsgroup"
	nsservicegroup "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ns_service_group/nsservicegroup"
	policyhosttransportnodeprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_host_transport_node_profile/policyhosttransportnodeprofile"
	policyipsecvpndpdprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_dpd_profile/policyipsecvpndpdprofile"
	policyipsecvpnikeprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_ike_profile/policyipsecvpnikeprofile"
	policyipsecvpnlocalendpoint "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_local_endpoint/policyipsecvpnlocalendpoint"
	policyipsecvpnservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_service/policyipsecvpnservice"
	policyipsecvpnsession "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_session/policyipsecvpnsession"
	policyl2vpnsession "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_l2_vpn_session/policyl2vpnsession"
	policytransportzone "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_transport_zone/policytransportzone"
	qosswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_qos_switching_profile/qosswitchingprofile"
	spoofguardswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_spoofguard_switching_profile/spoofguardswitchingprofile"
	staticroute "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_static_route/staticroute"
	switchsecurityswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_switch_security_switching_profile/switchsecurityswitchingprofile"
	transportnode "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_transport_node/transportnode"
	uplinkhostswitchprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_uplink_host_switch_profile/uplinkhostswitchprofile"
	vlanlogicalswitch "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_vlan_logical_switch/vlanlogicalswitch"
	vmtags "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_vm_tags/vmtags"
	providerconfig "github.com/ankasoftco/provider-nsxt/internal/controller/providerconfig"
	policyfixedsegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyfixedsegment"
	policyipdiscoveryprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyipdiscoveryprofile"
	policymacdiscoveryprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policymacdiscoveryprofile"
	policyqosprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyqosprofile"
	policysegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policysegment"
	policysegmentsecurityprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policysegmentsecurityprofile"
	policyspoofguardprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyspoofguardprofile"
	policyvlansegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyvlansegment"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policydhcprelay.Setup,
		policydhcpserver.Setup,
		policydhcpv4staticbinding.Setup,
		policydhcpv6staticbinding.Setup,
		policydnsforwarderzone.Setup,
		policygatewaydnsforwarder.Setup,
		policyevpnconfig.Setup,
		policyevpntenant.Setup,
		policyevpntunnelendpoint.Setup,
		policyvnipool.Setup,
		policycontextprofile.Setup,
		policycontextprofilecustomattribute.Setup,
		policygatewaypolicy.Setup,
		policyinstrusionservicepolicy.Setup,
		policyinstrusionserviceprofile.Setup,
		policypredefinedgatewaypolicy.Setup,
		policypredefinedsecuritypolicy.Setup,
		policysecuritypolicy.Setup,
		policyservice.Setup,
		policybgpconfig.Setup,
		policybgpmeighbor.Setup,
		policygatewaycommunitylist.Setup,
		policygatewayprefixlist.Setup,
		policygatewayqosprofile.Setup,
		policygatewayredistributionconfig.Setup,
		policygatewayroutemap.Setup,
		policynatrule.Setup,
		policyospfconfig.Setup,
		policystaticroute.Setup,
		policystaticroutebfdpeer.Setup,
		policytier0gateway.Setup,
		policytier0gatewayhavipconfig.Setup,
		policytier0gatewayinterface.Setup,
		policytier1gateway.Setup,
		policytier1gatewayinterface.Setup,
		policydomain.Setup,
		policygroup.Setup,
		policyvmtags.Setup,
		policyipaddressallocation.Setup,
		policyipblock.Setup,
		policyippool.Setup,
		policyippoolblocksubnet.Setup,
		policyippoolstaticsubnet.Setup,
		policylbpool.Setup,
		policylbservice.Setup,
		policylbvirtualserver.Setup,
		policyproject.Setup,
		algorithmtypensservice.Setup,
		clustervirtualip.Setup,
		computemanager.Setup,
		dhcprelayprofile.Setup,
		dhcprelayservice.Setup,
		dhcpserverippool.Setup,
		dhcpserverprofile.Setup,
		edgecluster.Setup,
		ethertypensservice.Setup,
		failuredomain.Setup,
		firewallsection.Setup,
		icmptypensservice.Setup,
		igmptypensservice.Setup,
		ipblock.Setup,
		ipblocksubnet.Setup,
		ipdiscoveryswitchingprofile.Setup,
		ippool.Setup,
		ippoolallocationipaddress.Setup,
		ipprotocolnsservice.Setup,
		ipset.Setup,
		l4portsetnsservice.Setup,
		lbclientsslprofile.Setup,
		lbcookiepersistenceprofile.Setup,
		lbfasttcpapplicationprofile.Setup,
		lbfastudpapplicationprofile.Setup,
		lbhttpapplicationprofile.Setup,
		lbhttpforwardingrule.Setup,
		lbhttpmonitor.Setup,
		lbhttprequestrewriterule.Setup,
		lbhttpresponserewriterule.Setup,
		lbhttpvirtualserver.Setup,
		lbhttpsmonitor.Setup,
		lbicmpmonitor.Setup,
		lbpassivemonitor.Setup,
		lbpool.Setup,
		lbserversslprofile.Setup,
		lbservice.Setup,
		lbsourceippersistenceprofile.Setup,
		lbtcpmonitor.Setup,
		lbtcpvirtualserver.Setup,
		lbudpmonitor.Setup,
		lbudpvirtualserver.Setup,
		logicaldhcpport.Setup,
		logicaldhcpserver.Setup,
		logicalport.Setup,
		logicalroutercentralizedserviceport.Setup,
		logicalrouterdownlinkport.Setup,
		logicalrouterlinkportontier0.Setup,
		logicalrouterlinkportontier1.Setup,
		logicalswitch.Setup,
		logicaltier0router.Setup,
		logicaltier1router.Setup,
		macmanagementswitchingprofile.Setup,
		managercluster.Setup,
		natrule.Setup,
		nsgroup.Setup,
		nsservicegroup.Setup,
		policyhosttransportnodeprofile.Setup,
		policyipsecvpndpdprofile.Setup,
		policyipsecvpnikeprofile.Setup,
		policyipsecvpnlocalendpoint.Setup,
		policyipsecvpnservice.Setup,
		policyipsecvpnsession.Setup,
		policyl2vpnsession.Setup,
		policytransportzone.Setup,
		qosswitchingprofile.Setup,
		spoofguardswitchingprofile.Setup,
		staticroute.Setup,
		switchsecurityswitchingprofile.Setup,
		transportnode.Setup,
		uplinkhostswitchprofile.Setup,
		vlanlogicalswitch.Setup,
		vmtags.Setup,
		providerconfig.Setup,
		policyfixedsegment.Setup,
		policyipdiscoveryprofile.Setup,
		policymacdiscoveryprofile.Setup,
		policyqosprofile.Setup,
		policysegment.Setup,
		policysegmentsecurityprofile.Setup,
		policyspoofguardprofile.Setup,
		policyvlansegment.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
