package beta

import (
	"github.com/upbound/upjet/pkg/config"
)

// ExternalNameConfigured returns a list of all external name resources
// configured for this provider.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_cluster_virtual_ip", func(r *config.Resource) {
		r.ShortGroup = "nsxtclustervirtualip"
		r.Kind = "ClusterVirtualIp"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_compute_manager", func(r *config.Resource) {
		r.ShortGroup = "nsxtcomputemanager"
		r.Kind = "ComputeManager"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_edge_cluster", func(r *config.Resource) {
		r.ShortGroup = "nsxtedgecluster"
		r.Kind = "EdgeCluster"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_failure_domain", func(r *config.Resource) {
		r.ShortGroup = "nsxtfailuredomain"
		r.Kind = "FailureDomain"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_manager_cluster", func(r *config.Resource) {
		r.ShortGroup = "nsxtmanagercluster"
		r.Kind = "ManagerCluster"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_host_transport_node_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicyhosttransportnodeprofile"
		r.Kind = "PolicyHostTransportNodeProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_transport_zone", func(r *config.Resource) {
		r.ShortGroup = "nsxtpolicytransportzone"
		r.Kind = "PolicyTransportZone"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_transport_node", func(r *config.Resource) {
		r.ShortGroup = "nsxttransportnode"
		r.Kind = "TransportNode"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_uplink_host_switch_profile", func(r *config.Resource) {
		r.ShortGroup = "nsxtuplinkhostswitchprofile"
		r.Kind = "UplinkHostSwitchProfile"
		r.Version = "v1alpha1"
	})

}
