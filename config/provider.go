/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	beta "github.com/ankasoftco/provider-nsxt/config/beta"
	deprecated "github.com/ankasoftco/provider-nsxt/config/deprecated"
	dhcp "github.com/ankasoftco/provider-nsxt/config/dhcp"
	dns "github.com/ankasoftco/provider-nsxt/config/dns"
	evpn "github.com/ankasoftco/provider-nsxt/config/evpn"
	firewall "github.com/ankasoftco/provider-nsxt/config/firewall"
	gatewaysandrouting "github.com/ankasoftco/provider-nsxt/config/gateways_and_routing"
	groupingandtagging "github.com/ankasoftco/provider-nsxt/config/grouping_and_tagging"
	ipam "github.com/ankasoftco/provider-nsxt/config/ipam"
	loadbalancer "github.com/ankasoftco/provider-nsxt/config/load_balancer"
	multitenancy "github.com/ankasoftco/provider-nsxt/config/multitenancy"
	ospf "github.com/ankasoftco/provider-nsxt/config/ospf"
	segments "github.com/ankasoftco/provider-nsxt/config/segments"
	vpn "github.com/ankasoftco/provider-nsxt/config/vpn"
)

const (
	resourcePrefix = "nsxt"
	modulePath     = "github.com/ankasoftco/provider-nsxt"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		dhcp.Configure,
		dns.Configure,
		evpn.Configure,
		firewall.Configure,
		gatewaysandrouting.Configure,
		groupingandtagging.Configure,
		ipam.Configure,
		loadbalancer.Configure,
		multitenancy.Configure,
		ospf.Configure,
		segments.Configure,
		vpn.Configure,
		beta.Configure,
		deprecated.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
