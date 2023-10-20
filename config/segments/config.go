package segments

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nsxt_policy_fixed_segment", func(r *config.Resource) {
		r.ShortGroup = "segments"
		r.Kind = "PolicyFixedSegment"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_ip_discovery_profile", func(r *config.Resource) {
		r.ShortGroup = "segments"
		r.Kind = "PolicyIpDiscoveryProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_mac_discovery_profile", func(r *config.Resource) {
		r.ShortGroup = "segments"
		r.Kind = "PolicyMacDiscoveryProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_qos_profile", func(r *config.Resource) {
		r.ShortGroup = "segments"
		r.Kind = "PolicyQosProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_segment", func(r *config.Resource) {
		r.ShortGroup = "segments"
		r.Kind = "PolicySegment"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_segment_security_profile", func(r *config.Resource) {
		r.ShortGroup = "segments"
		r.Kind = "PolicySegmentSecurityProfile"
		r.Version = "v1alpha1"
	})
	p.AddResourceConfigurator("nsxt_policy_spoof_guard_profile", func(r *config.Resource) {
		r.ShortGroup = "segments"
		r.Kind = "PolicySpoofGuardProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("nsxt_policy_vlan_segment", func(r *config.Resource) {
		r.ShortGroup = "segments"
		r.Kind = "PolicyVlanSegment"
		r.Version = "v1alpha1"
	})

	
}
