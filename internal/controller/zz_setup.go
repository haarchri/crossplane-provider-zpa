/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/zscaler/provider-zpa/internal/controller/providerconfig"
	edgegroup "github.com/zscaler/provider-zpa/internal/controller/service/edgegroup"
	accessrule "github.com/zscaler/provider-zpa/internal/controller/zpaaccessrule/accessrule"
	connectorgroup "github.com/zscaler/provider-zpa/internal/controller/zpaappconnectorgroup/connectorgroup"
	segment "github.com/zscaler/provider-zpa/internal/controller/zpaapplicationsegment/segment"
	segmentbrowseraccess "github.com/zscaler/provider-zpa/internal/controller/zpaapplicationsegmentbrowseraccess/segmentbrowseraccess"
	segmentinspection "github.com/zscaler/provider-zpa/internal/controller/zpaapplicationsegmentinspection/segmentinspection"
	segmentpra "github.com/zscaler/provider-zpa/internal/controller/zpaapplicationsegmentpra/segmentpra"
	server "github.com/zscaler/provider-zpa/internal/controller/zpaapplicationserver/server"
	forwardingrule "github.com/zscaler/provider-zpa/internal/controller/zpaforwardingrule/forwardingrule"
	customcontrols "github.com/zscaler/provider-zpa/internal/controller/zpainspectioncustomcontrols/customcontrols"
	profile "github.com/zscaler/provider-zpa/internal/controller/zpainspectionprofile/profile"
	inspectionrule "github.com/zscaler/provider-zpa/internal/controller/zpainspectionrule/inspectionrule"
	configcontroller "github.com/zscaler/provider-zpa/internal/controller/zpalssconfigcontroller/configcontroller"
	key "github.com/zscaler/provider-zpa/internal/controller/zpaprovisioningkey/key"
	group "github.com/zscaler/provider-zpa/internal/controller/zpasegmentgroup/group"
	groupzpaservergroup "github.com/zscaler/provider-zpa/internal/controller/zpaservergroup/group"
	timeoutrule "github.com/zscaler/provider-zpa/internal/controller/zpatimeoutrule/timeoutrule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		edgegroup.Setup,
		accessrule.Setup,
		connectorgroup.Setup,
		segment.Setup,
		segmentbrowseraccess.Setup,
		segmentinspection.Setup,
		segmentpra.Setup,
		server.Setup,
		forwardingrule.Setup,
		customcontrols.Setup,
		profile.Setup,
		inspectionrule.Setup,
		configcontroller.Setup,
		key.Setup,
		group.Setup,
		groupzpaservergroup.Setup,
		edgegroup.Setup,
		timeoutrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
