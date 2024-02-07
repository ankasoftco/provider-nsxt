/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	dhcprelayservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_relay_service/dhcprelayservice"
nsgroup "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ns_group/nsgroup"
policytier1gateway "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier1gateway"
policytier1gatewayinterface "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier1gatewayinterface"
policyfixedsegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyfixedsegment"
policyqosprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyqosprofile"
edgecluster "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_edge_cluster/edgecluster"
lbserversslprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_server_ssl_profile/lbserversslprofile"
policytier0gateway "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gateway"
policylbpool "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbpool"
ipdiscoveryswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_discovery_switching_profile/ipdiscoveryswitchingprofile"
logicalrouterlinkportontier1 "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_link_port_on_tier1/logicalrouterlinkportontier1"
policydhcprelay "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcprelay"
policytier0gatewayhavipconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gatewayhavipconfig"
policygroup "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policygroup"
lbhttpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_monitor/lbhttpmonitor"
spoofguardswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_spoofguard_switching_profile/spoofguardswitchingprofile"
vlanlogicalswitch "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_vlan_logical_switch/vlanlogicalswitch"
policybgpmeighbor "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policybgpmeighbor"
policyvmtags "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policyvmtags"
lbicmpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_icmp_monitor/lbicmpmonitor"
lbudpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_udp_virtual_server/lbudpvirtualserver"
policyhosttransportnodeprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_host_transport_node_profile/policyhosttransportnodeprofile"
policyipsecvpndpdprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_dpd_profile/policyipsecvpndpdprofile"
qosswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_qos_switching_profile/qosswitchingprofile"
policysegmentsecurityprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policysegmentsecurityprofile"
policyevpntenant "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpntenant"
lbclientsslprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_client_ssl_profile/lbclientsslprofile"
policyproject "github.com/ankasoftco/provider-nsxt/internal/controller/multitenancy/policyproject"
ipblocksubnet "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_block_subnet/ipblocksubnet"
lbfastudpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_fast_udp_application_profile/lbfastudpapplicationprofile"
logicaltier1router "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_tier1_router/logicaltier1router"
policyvnipool "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyvnipool"
policyospfconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policyospfconfig"
lbsourceippersistenceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_source_ip_persistence_profile/lbsourceippersistenceprofile"
policyipsecvpnlocalendpoint "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_local_endpoint/policyipsecvpnlocalendpoint"
transportnode "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_transport_node/transportnode"
policyippool "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippool"
ethertypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ether_type_ns_service/ethertypensservice"
policyipsecvpnservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_service/policyipsecvpnservice"
policypredefinedgatewaypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policypredefinedgatewaypolicy"
policytier0gatewayinterface "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gatewayinterface"
ipprotocolnsservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_protocol_ns_service/ipprotocolnsservice"
logicaltier0router "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_tier0_router/logicaltier0router"
natrule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_nat_rule/natrule"
policyipsecvpnikeprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_ike_profile/policyipsecvpnikeprofile"
policysecuritypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policysecuritypolicy"
dhcpserverprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_server_profile/dhcpserverprofile"
ipblock "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_block/ipblock"
lbhttpsmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_https_monitor/lbhttpsmonitor"
lbtcpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_tcp_monitor/lbtcpmonitor"
failuredomain "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_failure_domain/failuredomain"
firewallsection "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_firewall_section/firewallsection"
icmptypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_icmp_type_ns_service/icmptypensservice"
switchsecurityswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_switch_security_switching_profile/switchsecurityswitchingprofile"
policynatrule "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policynatrule"
algorithmtypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_algorithm_type_ns_service/algorithmtypensservice"
policyservice "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyservice"
policygatewayqosprofile "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayqosprofile"
policygatewayroutemap "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayroutemap"
policydhcpv4staticbinding "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpv4staticbinding"
policycontextprofilecustomattribute "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policycontextprofilecustomattribute"
ippoolallocationipaddress "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_pool_allocation_ip_address/ippoolallocationipaddress"
logicalrouterdownlinkport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_downlink_port/logicalrouterdownlinkport"
managercluster "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_manager_cluster/managercluster"
policyevpntunnelendpoint "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpntunnelendpoint"
logicalroutercentralizedserviceport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_centralized_service_port/logicalroutercentralizedserviceport"
lbhttprequestrewriterule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_request_rewrite_rule/lbhttprequestrewriterule"
policyvlansegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyvlansegment"
clustervirtualip "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_cluster_virtual_ip/clustervirtualip"
lbhttpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_application_profile/lbhttpapplicationprofile"
policypredefinedsecuritypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policypredefinedsecuritypolicy"
lbhttpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_virtual_server/lbhttpvirtualserver"
policydomain "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policydomain"
computemanager "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_compute_manager/computemanager"
ippool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_pool/ippool"
lbtcpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_tcp_virtual_server/lbtcpvirtualserver"
vmtags "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_vm_tags/vmtags"
policysegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policysegment"
policydhcpv6staticbinding "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpv6staticbinding"
policygatewaydnsforwarder "github.com/ankasoftco/provider-nsxt/internal/controller/dns/policygatewaydnsforwarder"
policystaticroute "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policystaticroute"
policyippoolblocksubnet "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippoolblocksubnet"
lbhttpforwardingrule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_forwarding_rule/lbhttpforwardingrule"
macmanagementswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_mac_management_switching_profile/macmanagementswitchingprofile"
nsservicegroup "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ns_service_group/nsservicegroup"
staticroute "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_static_route/staticroute"
policyevpnconfig "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpnconfig"
policyinstrusionservicepolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyinstrusionservicepolicy"
dhcprelayprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_relay_profile/dhcprelayprofile"
policygatewaypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policygatewaypolicy"
policyippoolstaticsubnet "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippoolstaticsubnet"
dhcpserverippool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_server_ip_pool/dhcpserverippool"
logicaldhcpport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_dhcp_port/logicaldhcpport"
policyspoofguardprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyspoofguardprofile"
policybgpconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policybgpconfig"
policyipblock "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyipblock"
providerconfig "github.com/ankasoftco/provider-nsxt/internal/controller/providerconfig"
policylbservice "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbservice"
policyipsecvpnsession "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_session/policyipsecvpnsession"
policytransportzone "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_transport_zone/policytransportzone"
policystaticroutebfdpeer "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policystaticroutebfdpeer"
logicaldhcpserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_dhcp_server/logicaldhcpserver"
policymacdiscoveryprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policymacdiscoveryprofile"
policydnsforwarderzone "github.com/ankasoftco/provider-nsxt/internal/controller/dns/policydnsforwarderzone"
logicalrouterlinkportontier0 "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_link_port_on_tier0/logicalrouterlinkportontier0"
lbfasttcpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_fast_tcp_application_profile/lbfasttcpapplicationprofile"
lbservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_service/lbservice"
policyl2vpnsession "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_l2_vpn_session/policyl2vpnsession"
uplinkhostswitchprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_uplink_host_switch_profile/uplinkhostswitchprofile"
policyipdiscoveryprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyipdiscoveryprofile"
policygatewaycommunitylist "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewaycommunitylist"
lbcookiepersistenceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_cookie_persistence_profile/lbcookiepersistenceprofile"
igmptypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_igmp_type_ns_service/igmptypensservice"
lbhttpresponserewriterule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_response_rewrite_rule/lbhttpresponserewriterule"
lbudpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_udp_monitor/lbudpmonitor"
logicalswitch "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_switch/logicalswitch"
policydhcpserver "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpserver"
policylbvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbvirtualserver"
lbpassivemonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_passive_monitor/lbpassivemonitor"
lbpool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_pool/lbpool"
logicalport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_port/logicalport"
policyinstrusionserviceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyinstrusionserviceprofile"
policygatewayredistributionconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayredistributionconfig"
policyipaddressallocation "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyipaddressallocation"
ipset "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_set/ipset"
l4portsetnsservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_l4_port_set_ns_service/l4portsetnsservice"
policycontextprofile "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policycontextprofile"
policygatewayprefixlist "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayprefixlist"

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
