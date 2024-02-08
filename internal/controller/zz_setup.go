/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	policytier0gateway "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gateway"
dhcprelayservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_relay_service/dhcprelayservice"
lbpool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_pool/lbpool"
natrule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_nat_rule/natrule"
logicaltier1router "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_tier1_router/logicaltier1router"
vlanlogicalswitch "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_vlan_logical_switch/vlanlogicalswitch"
policynatrule "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policynatrule"
policystaticroute "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policystaticroute"
policystaticroutebfdpeer "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policystaticroutebfdpeer"
clustervirtualip "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_cluster_virtual_ip/clustervirtualip"
policyippool "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippool"
lbclientsslprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_client_ssl_profile/lbclientsslprofile"
logicalport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_port/logicalport"
logicaltier0router "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_tier0_router/logicaltier0router"
logicalrouterdownlinkport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_downlink_port/logicalrouterdownlinkport"
policygatewayqosprofile "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayqosprofile"
edgecluster "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_edge_cluster/edgecluster"
lbsourceippersistenceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_source_ip_persistence_profile/lbsourceippersistenceprofile"
lbtcpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_tcp_monitor/lbtcpmonitor"
policyqosprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyqosprofile"
policyippoolblocksubnet "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippoolblocksubnet"
ipblocksubnet "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_block_subnet/ipblocksubnet"
lbtcpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_tcp_virtual_server/lbtcpvirtualserver"
policyipsecvpndpdprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_dpd_profile/policyipsecvpndpdprofile"
policyvlansegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyvlansegment"
policydhcpv6staticbinding "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpv6staticbinding"
ipprotocolnsservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_protocol_ns_service/ipprotocolnsservice"
logicalswitch "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_switch/logicalswitch"
transportnode "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_transport_node/transportnode"
policylbpool "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbpool"
lbservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_service/lbservice"
spoofguardswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_spoofguard_switching_profile/spoofguardswitchingprofile"
policysegmentsecurityprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policysegmentsecurityprofile"
staticroute "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_static_route/staticroute"
policyspoofguardprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyspoofguardprofile"
policycontextprofile "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policycontextprofile"
policytier1gatewayinterface "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier1gatewayinterface"
policylbservice "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbservice"
l4portsetnsservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_l4_port_set_ns_service/l4portsetnsservice"
ipset "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_set/ipset"
policyipsecvpnsession "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_session/policyipsecvpnsession"
policymacdiscoveryprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policymacdiscoveryprofile"
policyinstrusionserviceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyinstrusionserviceprofile"
policytier0gatewayinterface "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gatewayinterface"
policydomain "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policydomain"
policyippoolstaticsubnet "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyippoolstaticsubnet"
policyevpnconfig "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpnconfig"
policypredefinedsecuritypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policypredefinedsecuritypolicy"
lbhttpsmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_https_monitor/lbhttpsmonitor"
qosswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_qos_switching_profile/qosswitchingprofile"
policyservice "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyservice"
ippool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_pool/ippool"
lbhttprequestrewriterule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_request_rewrite_rule/lbhttprequestrewriterule"
policyfixedsegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyfixedsegment"
ethertypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ether_type_ns_service/ethertypensservice"
lbpassivemonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_passive_monitor/lbpassivemonitor"
switchsecurityswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_switch_security_switching_profile/switchsecurityswitchingprofile"
lbhttpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_virtual_server/lbhttpvirtualserver"
policyevpntenant "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpntenant"
policysecuritypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policysecuritypolicy"
policyvmtags "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policyvmtags"
ippoolallocationipaddress "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_pool_allocation_ip_address/ippoolallocationipaddress"
nsservicegroup "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ns_service_group/nsservicegroup"
policygatewayroutemap "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayroutemap"
policyproject "github.com/ankasoftco/provider-nsxt/internal/controller/multitenancy/policyproject"
ipblock "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_block/ipblock"
lbhttpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_application_profile/lbhttpapplicationprofile"
policydhcpv4staticbinding "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpv4staticbinding"
lbudpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_udp_monitor/lbudpmonitor"
policytier1gateway "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier1gateway"
dhcpserverprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_server_profile/dhcpserverprofile"
lbfastudpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_fast_udp_application_profile/lbfastudpapplicationprofile"
policyipdiscoveryprofile "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policyipdiscoveryprofile"
policyospfconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policyospfconfig"
logicalrouterlinkportontier1 "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_link_port_on_tier1/logicalrouterlinkportontier1"
managercluster "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_manager_cluster/managercluster"
policytransportzone "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_transport_zone/policytransportzone"
policyinstrusionservicepolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policyinstrusionservicepolicy"
policygatewayredistributionconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayredistributionconfig"
vmtags "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_vm_tags/vmtags"
lbhttpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_monitor/lbhttpmonitor"
logicaldhcpport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_dhcp_port/logicaldhcpport"
nsgroup "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ns_group/nsgroup"
policytier0gatewayhavipconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policytier0gatewayhavipconfig"
policyipblock "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyipblock"
computemanager "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_compute_manager/computemanager"
lbfasttcpapplicationprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_fast_tcp_application_profile/lbfasttcpapplicationprofile"
lbserversslprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_server_ssl_profile/lbserversslprofile"
policyhosttransportnodeprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_host_transport_node_profile/policyhosttransportnodeprofile"
policydnsforwarderzone "github.com/ankasoftco/provider-nsxt/internal/controller/dns/policydnsforwarderzone"
policygatewayprefixlist "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewayprefixlist"
failuredomain "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_failure_domain/failuredomain"
ipdiscoveryswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_ip_discovery_switching_profile/ipdiscoveryswitchingprofile"
policydhcpserver "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcpserver"
policybgpconfig "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policybgpconfig"
uplinkhostswitchprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_uplink_host_switch_profile/uplinkhostswitchprofile"
logicalroutercentralizedserviceport "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_centralized_service_port/logicalroutercentralizedserviceport"
policysegment "github.com/ankasoftco/provider-nsxt/internal/controller/segments/policysegment"
policyevpntunnelendpoint "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyevpntunnelendpoint"
policyvnipool "github.com/ankasoftco/provider-nsxt/internal/controller/evpn/policyvnipool"
policygatewaypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policygatewaypolicy"
icmptypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_icmp_type_ns_service/icmptypensservice"
policygatewaycommunitylist "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policygatewaycommunitylist"
policyipaddressallocation "github.com/ankasoftco/provider-nsxt/internal/controller/ipam/policyipaddressallocation"
policypredefinedgatewaypolicy "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policypredefinedgatewaypolicy"
policylbvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/loadbalancer/policylbvirtualserver"
policyipsecvpnlocalendpoint "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_local_endpoint/policyipsecvpnlocalendpoint"
policyipsecvpnservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_service/policyipsecvpnservice"
lbicmpmonitor "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_icmp_monitor/lbicmpmonitor"
policybgpmeighbor "github.com/ankasoftco/provider-nsxt/internal/controller/gatewaysandrouting/policybgpmeighbor"
lbhttpforwardingrule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_forwarding_rule/lbhttpforwardingrule"
logicalrouterlinkportontier0 "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_router_link_port_on_tier0/logicalrouterlinkportontier0"
policygatewaydnsforwarder "github.com/ankasoftco/provider-nsxt/internal/controller/dns/policygatewaydnsforwarder"
lbudpvirtualserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_udp_virtual_server/lbudpvirtualserver"
policycontextprofilecustomattribute "github.com/ankasoftco/provider-nsxt/internal/controller/firewall/policycontextprofilecustomattribute"
policygroup "github.com/ankasoftco/provider-nsxt/internal/controller/groupingandtagging/policygroup"
dhcprelayprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_relay_profile/dhcprelayprofile"
igmptypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_igmp_type_ns_service/igmptypensservice"
logicaldhcpserver "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_logical_dhcp_server/logicaldhcpserver"
providerconfig "github.com/ankasoftco/provider-nsxt/internal/controller/providerconfig"
policyl2vpnsession "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_l2_vpn_session/policyl2vpnsession"
algorithmtypensservice "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_algorithm_type_ns_service/algorithmtypensservice"
dhcpserverippool "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_dhcp_server_ip_pool/dhcpserverippool"
lbhttpresponserewriterule "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_http_response_rewrite_rule/lbhttpresponserewriterule"
policyipsecvpnikeprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_policy_ipsec_vpn_ike_profile/policyipsecvpnikeprofile"
firewallsection "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_firewall_section/firewallsection"
lbcookiepersistenceprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_lb_cookie_persistence_profile/lbcookiepersistenceprofile"
policydhcprelay "github.com/ankasoftco/provider-nsxt/internal/controller/dhcp/policydhcprelay"
macmanagementswitchingprofile "github.com/ankasoftco/provider-nsxt/internal/controller/nsxt_mac_management_switching_profile/macmanagementswitchingprofile"

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
