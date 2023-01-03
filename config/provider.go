/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"
	"github.com/zscaler/provider-zpa/config/zpaappconnectorgroup"
	"github.com/zscaler/provider-zpa/config/zpaapplicationsegment"
	"github.com/zscaler/provider-zpa/config/zpaapplicationsegmentbrowseraccess"
	"github.com/zscaler/provider-zpa/config/zpaapplicationsegmentinspection"
	"github.com/zscaler/provider-zpa/config/zpaapplicationsegmentpra"
	"github.com/zscaler/provider-zpa/config/zpaapplicationserver"
	"github.com/zscaler/provider-zpa/config/zpalssconfigclienttypes"
	"github.com/zscaler/provider-zpa/config/zpalssconfigcontroller"
	"github.com/zscaler/provider-zpa/config/zpalssconfiglogtypeformats"
	"github.com/zscaler/provider-zpa/config/zpalssconfigstatuscodes"
	"github.com/zscaler/provider-zpa/config/zpapolicyaccessforwardingrule"
	"github.com/zscaler/provider-zpa/config/zpapolicyaccessinspectionrule"
	"github.com/zscaler/provider-zpa/config/zpapolicyaccessrule"
	"github.com/zscaler/provider-zpa/config/zpapolicyaccesstimeoutrule"
	"github.com/zscaler/provider-zpa/config/zpaprovisioningkey"
	"github.com/zscaler/provider-zpa/config/zpasegmentgroup"
	"github.com/zscaler/provider-zpa/config/zpaservergroup"
	"github.com/zscaler/provider-zpa/config/zpaserviceedgegroup"
)

const (
	resourcePrefix = "zpa"
	modulePath     = "github.com/zscaler/provider-zpa"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithShortName("zpa"),
		ujconfig.WithRootGroup("zpa.crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		zpaapplicationsegment.Configure,
		zpaapplicationsegmentpra.Configure,
		zpaapplicationsegmentinspection.Configure,
		zpaapplicationsegmentbrowseraccess.Configure,
		zpaappconnectorgroup.Configure,
		zpaserviceedgegroup.Configure,
		zpaapplicationserver.Configure,
		zpasegmentgroup.Configure,
		zpaservergroup.Configure,
		zpapolicyaccessrule.Configure,
		zpapolicyaccesstimeoutrule.Configure,
		zpapolicyaccessforwardingrule.Configure,
		zpapolicyaccessinspectionrule.Configure,
		zpaprovisioningkey.Configure,
		zpalssconfigcontroller.Configure,
		zpalssconfigclienttypes.Configure,
		zpalssconfiglogtypeformats.Configure,
		zpalssconfigstatuscodes.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
