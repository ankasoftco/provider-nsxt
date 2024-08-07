package segments

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "segments"
const version = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_fixed_segment", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyFixedSegment"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_ip_discovery_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyIpDiscoveryProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_mac_discovery_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyMacDiscoveryProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_qos_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyQosProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_segment", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicySegment"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_segment_security_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicySegmentSecurityProfile"
		r.Version = version
	})
	p.AddResourceConfigurator("nsxt_policy_spoof_guard_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicySpoofGuardProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("nsxt_policy_vlan_segment", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PolicyVlanSegment"
		r.Version = "v1alpha1"
	})

}
