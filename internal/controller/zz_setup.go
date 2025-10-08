// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	project "github.com/gravitek-io/provider-ovh/internal/controller/cloud/project"
	projectcontainerregistry "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectcontainerregistry"
	projectcontainerregistryiam "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectcontainerregistryiam"
	projectcontainerregistryiprestrictionsmanagement "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectcontainerregistryiprestrictionsmanagement"
	projectcontainerregistryiprestrictionsregistry "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectcontainerregistryiprestrictionsregistry"
	projectcontainerregistryoidc "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectcontainerregistryoidc"
	projectcontainerregistryuser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectcontainerregistryuser"
	projectdatabase "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabase"
	projectdatabasedatabase "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasedatabase"
	projectdatabaseintegration "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabaseintegration"
	projectdatabaseiprestriction "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabaseiprestriction"
	projectdatabasekafkaacl "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasekafkaacl"
	projectdatabasekafkaschemaregistryacl "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasekafkaschemaregistryacl"
	projectdatabasekafkatopic "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasekafkatopic"
	projectdatabaselogsubscription "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabaselogsubscription"
	projectdatabasem3dbnamespace "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasem3dbnamespace"
	projectdatabasem3dbuser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasem3dbuser"
	projectdatabasemongodbprometheus "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasemongodbprometheus"
	projectdatabasemongodbuser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasemongodbuser"
	projectdatabaseopensearchpattern "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabaseopensearchpattern"
	projectdatabaseopensearchuser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabaseopensearchuser"
	projectdatabasepostgresqlconnectionpool "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasepostgresqlconnectionpool"
	projectdatabasepostgresqluser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasepostgresqluser"
	projectdatabaseprometheus "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabaseprometheus"
	projectdatabaseredisuser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabaseredisuser"
	projectdatabaseuser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabaseuser"
	projectdatabasevalkeyuser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectdatabasevalkeyuser"
	projectfailoveripattach "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectfailoveripattach"
	projectgateway "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectgateway"
	projectgatewayinterface "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectgatewayinterface"
	projectinstance "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectinstance"
	projectinstancesnapshot "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectinstancesnapshot"
	projectkube "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectkube"
	projectkubeiprestrictions "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectkubeiprestrictions"
	projectkubenodepool "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectkubenodepool"
	projectkubeoidc "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectkubeoidc"
	projectnetworkprivate "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectnetworkprivate"
	projectnetworkprivatesubnet "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectnetworkprivatesubnet"
	projectnetworkprivatesubnetv2 "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectnetworkprivatesubnetv2"
	projectregionloadbalancerlogsubscription "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectregionloadbalancerlogsubscription"
	projectregionstoragepresign "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectregionstoragepresign"
	projectsshkey "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectsshkey"
	projectuser "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectuser"
	projectusers3credential "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectusers3credential"
	projectusers3policy "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectusers3policy"
	projectvolumebackup "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectvolumebackup"
	projectworkflowbackup "github.com/gravitek-io/provider-ovh/internal/controller/cloud/projectworkflowbackup"
	logscluster "github.com/gravitek-io/provider-ovh/internal/controller/dbaas/logscluster"
	logsinput "github.com/gravitek-io/provider-ovh/internal/controller/dbaas/logsinput"
	logsoutputgraylogstream "github.com/gravitek-io/provider-ovh/internal/controller/dbaas/logsoutputgraylogstream"
	logsoutputopensearchalias "github.com/gravitek-io/provider-ovh/internal/controller/dbaas/logsoutputopensearchalias"
	logsoutputopensearchindex "github.com/gravitek-io/provider-ovh/internal/controller/dbaas/logsoutputopensearchindex"
	logsrole "github.com/gravitek-io/provider-ovh/internal/controller/dbaas/logsrole"
	logsrolepermissionstream "github.com/gravitek-io/provider-ovh/internal/controller/dbaas/logsrolepermissionstream"
	logstoken "github.com/gravitek-io/provider-ovh/internal/controller/dbaas/logstoken"
	cephacl "github.com/gravitek-io/provider-ovh/internal/controller/dedicated/cephacl"
	nashapartition "github.com/gravitek-io/provider-ovh/internal/controller/dedicated/nashapartition"
	nashapartitionaccess "github.com/gravitek-io/provider-ovh/internal/controller/dedicated/nashapartitionaccess"
	nashapartitionsnapshot "github.com/gravitek-io/provider-ovh/internal/controller/dedicated/nashapartitionsnapshot"
	servernetworking "github.com/gravitek-io/provider-ovh/internal/controller/dedicated/servernetworking"
	serverreboottask "github.com/gravitek-io/provider-ovh/internal/controller/dedicated/serverreboottask"
	serverreinstalltask "github.com/gravitek-io/provider-ovh/internal/controller/dedicated/serverreinstalltask"
	serverupdate "github.com/gravitek-io/provider-ovh/internal/controller/dedicated/serverupdate"
	dsrecords "github.com/gravitek-io/provider-ovh/internal/controller/domain/dsrecords"
	nameservers "github.com/gravitek-io/provider-ovh/internal/controller/domain/nameservers"
	zone "github.com/gravitek-io/provider-ovh/internal/controller/domain/zone"
	zonednssec "github.com/gravitek-io/provider-ovh/internal/controller/domain/zonednssec"
	zonedynhostlogin "github.com/gravitek-io/provider-ovh/internal/controller/domain/zonedynhostlogin"
	zonedynhostrecord "github.com/gravitek-io/provider-ovh/internal/controller/domain/zonedynhostrecord"
	zoneimport "github.com/gravitek-io/provider-ovh/internal/controller/domain/zoneimport"
	zonerecord "github.com/gravitek-io/provider-ovh/internal/controller/domain/zonerecord"
	zoneredirection "github.com/gravitek-io/provider-ovh/internal/controller/domain/zoneredirection"
	privatedatabase "github.com/gravitek-io/provider-ovh/internal/controller/hosting/privatedatabase"
	privatedatabasedatabase "github.com/gravitek-io/provider-ovh/internal/controller/hosting/privatedatabasedatabase"
	privatedatabaseuser "github.com/gravitek-io/provider-ovh/internal/controller/hosting/privatedatabaseuser"
	privatedatabaseusergrant "github.com/gravitek-io/provider-ovh/internal/controller/hosting/privatedatabaseusergrant"
	privatedatabasewhitelist "github.com/gravitek-io/provider-ovh/internal/controller/hosting/privatedatabasewhitelist"
	permissionsgroup "github.com/gravitek-io/provider-ovh/internal/controller/iam/permissionsgroup"
	policy "github.com/gravitek-io/provider-ovh/internal/controller/iam/policy"
	resourcegroup "github.com/gravitek-io/provider-ovh/internal/controller/iam/resourcegroup"
	firewall "github.com/gravitek-io/provider-ovh/internal/controller/ip/firewall"
	firewallrule "github.com/gravitek-io/provider-ovh/internal/controller/ip/firewallrule"
	mitigation "github.com/gravitek-io/provider-ovh/internal/controller/ip/mitigation"
	move "github.com/gravitek-io/provider-ovh/internal/controller/ip/move"
	reverse "github.com/gravitek-io/provider-ovh/internal/controller/ip/reverse"
	service "github.com/gravitek-io/provider-ovh/internal/controller/ip/service"
	httpfarm "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/httpfarm"
	httpfarmserver "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/httpfarmserver"
	httpfrontend "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/httpfrontend"
	httproute "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/httproute"
	httprouterule "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/httprouterule"
	refresh "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/refresh"
	ssl "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/ssl"
	tcpfarm "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/tcpfarm"
	tcpfarmserver "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/tcpfarmserver"
	tcpfrontend "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/tcpfrontend"
	tcproute "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/tcproute"
	tcprouterule "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/tcprouterule"
	udpfarm "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/udpfarm"
	udpfarmserver "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/udpfarmserver"
	udpfrontend "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/udpfrontend"
	vracknetwork "github.com/gravitek-io/provider-ovh/internal/controller/iploadbalancing/vracknetwork"
	apioauth2client "github.com/gravitek-io/provider-ovh/internal/controller/me/apioauth2client"
	identitygroup "github.com/gravitek-io/provider-ovh/internal/controller/me/identitygroup"
	identityuser "github.com/gravitek-io/provider-ovh/internal/controller/me/identityuser"
	credential "github.com/gravitek-io/provider-ovh/internal/controller/okms/credential"
	servicekey "github.com/gravitek-io/provider-ovh/internal/controller/okms/servicekey"
	iploadbalancing "github.com/gravitek-io/provider-ovh/internal/controller/ovh/iploadbalancing"
	vrack "github.com/gravitek-io/provider-ovh/internal/controller/ovh/vrack"
	connectpopconfig "github.com/gravitek-io/provider-ovh/internal/controller/ovhcloud/connectpopconfig"
	connectpopdatacenterconfig "github.com/gravitek-io/provider-ovh/internal/controller/ovhcloud/connectpopdatacenterconfig"
	connectpopdatacenterextraconfig "github.com/gravitek-io/provider-ovh/internal/controller/ovhcloud/connectpopdatacenterextraconfig"
	providerconfig "github.com/gravitek-io/provider-ovh/internal/controller/providerconfig"
	plan "github.com/gravitek-io/provider-ovh/internal/controller/savings/plan"
	efsshare "github.com/gravitek-io/provider-ovh/internal/controller/storage/efsshare"
	efsshareacl "github.com/gravitek-io/provider-ovh/internal/controller/storage/efsshareacl"
	efssharesnapshot "github.com/gravitek-io/provider-ovh/internal/controller/storage/efssharesnapshot"
	cloudproject "github.com/gravitek-io/provider-ovh/internal/controller/vrack/cloudproject"
	dedicatedcloud "github.com/gravitek-io/provider-ovh/internal/controller/vrack/dedicatedcloud"
	dedicatedclouddatacenter "github.com/gravitek-io/provider-ovh/internal/controller/vrack/dedicatedclouddatacenter"
	dedicatedserver "github.com/gravitek-io/provider-ovh/internal/controller/vrack/dedicatedserver"
	dedicatedserverinterface "github.com/gravitek-io/provider-ovh/internal/controller/vrack/dedicatedserverinterface"
	ip "github.com/gravitek-io/provider-ovh/internal/controller/vrack/ip"
	iploadbalancingvrack "github.com/gravitek-io/provider-ovh/internal/controller/vrack/iploadbalancing"
	ipv6 "github.com/gravitek-io/provider-ovh/internal/controller/vrack/ipv6"
	ipv6routedsubrange "github.com/gravitek-io/provider-ovh/internal/controller/vrack/ipv6routedsubrange"
	ovhcloudconnect "github.com/gravitek-io/provider-ovh/internal/controller/vrack/ovhcloudconnect"
	vrackservices "github.com/gravitek-io/provider-ovh/internal/controller/vrack/vrackservices"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		projectcontainerregistry.Setup,
		projectcontainerregistryiam.Setup,
		projectcontainerregistryiprestrictionsmanagement.Setup,
		projectcontainerregistryiprestrictionsregistry.Setup,
		projectcontainerregistryoidc.Setup,
		projectcontainerregistryuser.Setup,
		projectdatabase.Setup,
		projectdatabasedatabase.Setup,
		projectdatabaseintegration.Setup,
		projectdatabaseiprestriction.Setup,
		projectdatabasekafkaacl.Setup,
		projectdatabasekafkaschemaregistryacl.Setup,
		projectdatabasekafkatopic.Setup,
		projectdatabaselogsubscription.Setup,
		projectdatabasem3dbnamespace.Setup,
		projectdatabasem3dbuser.Setup,
		projectdatabasemongodbprometheus.Setup,
		projectdatabasemongodbuser.Setup,
		projectdatabaseopensearchpattern.Setup,
		projectdatabaseopensearchuser.Setup,
		projectdatabasepostgresqlconnectionpool.Setup,
		projectdatabasepostgresqluser.Setup,
		projectdatabaseprometheus.Setup,
		projectdatabaseredisuser.Setup,
		projectdatabaseuser.Setup,
		projectdatabasevalkeyuser.Setup,
		projectfailoveripattach.Setup,
		projectgateway.Setup,
		projectgatewayinterface.Setup,
		projectinstance.Setup,
		projectinstancesnapshot.Setup,
		projectkube.Setup,
		projectkubeiprestrictions.Setup,
		projectkubenodepool.Setup,
		projectkubeoidc.Setup,
		projectnetworkprivate.Setup,
		projectnetworkprivatesubnet.Setup,
		projectnetworkprivatesubnetv2.Setup,
		projectregionloadbalancerlogsubscription.Setup,
		projectregionstoragepresign.Setup,
		projectsshkey.Setup,
		projectuser.Setup,
		projectusers3credential.Setup,
		projectusers3policy.Setup,
		projectvolumebackup.Setup,
		projectworkflowbackup.Setup,
		logscluster.Setup,
		logsinput.Setup,
		logsoutputgraylogstream.Setup,
		logsoutputopensearchalias.Setup,
		logsoutputopensearchindex.Setup,
		logsrole.Setup,
		logsrolepermissionstream.Setup,
		logstoken.Setup,
		cephacl.Setup,
		nashapartition.Setup,
		nashapartitionaccess.Setup,
		nashapartitionsnapshot.Setup,
		servernetworking.Setup,
		serverreboottask.Setup,
		serverreinstalltask.Setup,
		serverupdate.Setup,
		dsrecords.Setup,
		nameservers.Setup,
		zone.Setup,
		zonednssec.Setup,
		zonedynhostlogin.Setup,
		zonedynhostrecord.Setup,
		zoneimport.Setup,
		zonerecord.Setup,
		zoneredirection.Setup,
		privatedatabase.Setup,
		privatedatabasedatabase.Setup,
		privatedatabaseuser.Setup,
		privatedatabaseusergrant.Setup,
		privatedatabasewhitelist.Setup,
		permissionsgroup.Setup,
		policy.Setup,
		resourcegroup.Setup,
		firewall.Setup,
		firewallrule.Setup,
		mitigation.Setup,
		move.Setup,
		reverse.Setup,
		service.Setup,
		httpfarm.Setup,
		httpfarmserver.Setup,
		httpfrontend.Setup,
		httproute.Setup,
		httprouterule.Setup,
		refresh.Setup,
		ssl.Setup,
		tcpfarm.Setup,
		tcpfarmserver.Setup,
		tcpfrontend.Setup,
		tcproute.Setup,
		tcprouterule.Setup,
		udpfarm.Setup,
		udpfarmserver.Setup,
		udpfrontend.Setup,
		vracknetwork.Setup,
		apioauth2client.Setup,
		identitygroup.Setup,
		identityuser.Setup,
		credential.Setup,
		servicekey.Setup,
		iploadbalancing.Setup,
		vrack.Setup,
		connectpopconfig.Setup,
		connectpopdatacenterconfig.Setup,
		connectpopdatacenterextraconfig.Setup,
		providerconfig.Setup,
		plan.Setup,
		efsshare.Setup,
		efsshareacl.Setup,
		efssharesnapshot.Setup,
		cloudproject.Setup,
		dedicatedcloud.Setup,
		dedicatedclouddatacenter.Setup,
		dedicatedserver.Setup,
		dedicatedserverinterface.Setup,
		ip.Setup,
		iploadbalancingvrack.Setup,
		ipv6.Setup,
		ipv6routedsubrange.Setup,
		ovhcloudconnect.Setup,
		vrackservices.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
