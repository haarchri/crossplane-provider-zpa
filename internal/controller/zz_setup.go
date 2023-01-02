/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accessrule "github.com/zscaler/provider-zpa/internal/controller/accessrule/accessrule"
	connectorgroup "github.com/zscaler/provider-zpa/internal/controller/app/connectorgroup"
	segmentbrowseraccess "github.com/zscaler/provider-zpa/internal/controller/application/segmentbrowseraccess"
	segmentinspection "github.com/zscaler/provider-zpa/internal/controller/application/segmentinspection"
	segmentpra "github.com/zscaler/provider-zpa/internal/controller/application/segmentpra"
	segment "github.com/zscaler/provider-zpa/internal/controller/applicationsegment/segment"
	server "github.com/zscaler/provider-zpa/internal/controller/applicationserver/server"
	customcontrols "github.com/zscaler/provider-zpa/internal/controller/inspection/customcontrols"
	profile "github.com/zscaler/provider-zpa/internal/controller/inspection/profile"
	configcontroller "github.com/zscaler/provider-zpa/internal/controller/lss/configcontroller"
	forwardingrule "github.com/zscaler/provider-zpa/internal/controller/policy/forwardingrule"
	inspectionrule "github.com/zscaler/provider-zpa/internal/controller/policy/inspectionrule"
	timeoutrule "github.com/zscaler/provider-zpa/internal/controller/policy/timeoutrule"
	providerconfig "github.com/zscaler/provider-zpa/internal/controller/providerconfig"
	key "github.com/zscaler/provider-zpa/internal/controller/provisioning/key"
	group "github.com/zscaler/provider-zpa/internal/controller/segmentgroup/group"
	groupservergroup "github.com/zscaler/provider-zpa/internal/controller/servergroup/group"
	edgegroup "github.com/zscaler/provider-zpa/internal/controller/serviceedgegroup/edgegroup"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accessrule.Setup,
		connectorgroup.Setup,
		segmentbrowseraccess.Setup,
		segmentinspection.Setup,
		segmentpra.Setup,
		segment.Setup,
		server.Setup,
		customcontrols.Setup,
		profile.Setup,
		configcontroller.Setup,
		forwardingrule.Setup,
		inspectionrule.Setup,
		timeoutrule.Setup,
		providerconfig.Setup,
		key.Setup,
		group.Setup,
		groupservergroup.Setup,
		edgegroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
