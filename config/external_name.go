/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Cloud Project resources
	"ovh_cloud_project":                                              config.IdentifierFromProvider,
	"ovh_cloud_project_alerting":                                     config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry":                            config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_iam":                        config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_ip_restrictions_management": config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_ip_restrictions_registry":   config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_oidc":                       config.IdentifierFromProvider,
	"ovh_cloud_project_containerregistry_user":                       config.IdentifierFromProvider,
	"ovh_cloud_project_database":                                     config.IdentifierFromProvider,
	"ovh_cloud_project_database_database":                            config.IdentifierFromProvider,
	"ovh_cloud_project_database_integration":                         config.IdentifierFromProvider,
	"ovh_cloud_project_database_ip_restriction":                      config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_acl":                           config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_schemaregistryacl":             config.IdentifierFromProvider,
	"ovh_cloud_project_database_kafka_topic":                         config.IdentifierFromProvider,
	"ovh_cloud_project_database_log_subscription":                    config.IdentifierFromProvider,
	"ovh_cloud_project_database_m3db_namespace":                      config.IdentifierFromProvider,
	"ovh_cloud_project_database_m3db_user":                           config.IdentifierFromProvider,
	"ovh_cloud_project_database_mongodb_prometheus":                  config.IdentifierFromProvider,
	"ovh_cloud_project_database_mongodb_user":                        config.IdentifierFromProvider,
	"ovh_cloud_project_database_opensearch_pattern":                  config.IdentifierFromProvider,
	"ovh_cloud_project_database_opensearch_user":                     config.IdentifierFromProvider,
	"ovh_cloud_project_database_postgresql_connection_pool":          config.IdentifierFromProvider,
	"ovh_cloud_project_database_postgresql_user":                     config.IdentifierFromProvider,
	"ovh_cloud_project_database_prometheus":                          config.IdentifierFromProvider,
	"ovh_cloud_project_database_redis_user":                          config.IdentifierFromProvider,
	"ovh_cloud_project_database_user":                                config.IdentifierFromProvider,
	"ovh_cloud_project_database_valkey_user":                         config.IdentifierFromProvider,
	"ovh_cloud_project_failover_ip_attach":                           config.IdentifierFromProvider,
	"ovh_cloud_project_gateway":                                      config.IdentifierFromProvider,
	"ovh_cloud_project_gateway_interface":                            config.IdentifierFromProvider,
	"ovh_cloud_project_instance":                                     config.IdentifierFromProvider,
	"ovh_cloud_project_instance_snapshot":                            config.IdentifierFromProvider,
	"ovh_cloud_project_kube":                                         config.IdentifierFromProvider,
	"ovh_cloud_project_kube_iprestrictions":                          config.IdentifierFromProvider,
	"ovh_cloud_project_kube_nodepool":                                config.IdentifierFromProvider,
	"ovh_cloud_project_kube_oidc":                                    config.IdentifierFromProvider,
	"ovh_cloud_project_loadbalancer":                                 config.IdentifierFromProvider,
	"ovh_cloud_project_network_private":                              config.IdentifierFromProvider,
	"ovh_cloud_project_network_private_subnet":                       config.IdentifierFromProvider,
	"ovh_cloud_project_network_private_subnet_v2":                    config.IdentifierFromProvider,
	"ovh_cloud_project_rancher":                                      config.IdentifierFromProvider,
	"ovh_cloud_project_region":                                       config.IdentifierFromProvider,
	"ovh_cloud_project_region_loadbalancer_log_subscription":         config.IdentifierFromProvider,
	"ovh_cloud_project_region_network":                               config.IdentifierFromProvider,
	"ovh_cloud_project_region_storage_presign":                       config.IdentifierFromProvider,
	"ovh_cloud_project_ssh_key":                                      config.IdentifierFromProvider,
	"ovh_cloud_project_storage":                                      config.IdentifierFromProvider,
	"ovh_cloud_project_user":                                         config.IdentifierFromProvider,
	"ovh_cloud_project_user_s3_credential":                           config.IdentifierFromProvider,
	"ovh_cloud_project_user_s3_policy":                               config.IdentifierFromProvider,
	"ovh_cloud_project_volume":                                       config.IdentifierFromProvider,
	"ovh_cloud_project_volume_backup":                                config.IdentifierFromProvider,
	"ovh_cloud_project_workflow_backup":                              config.IdentifierFromProvider,

	// DBaaS Logs resources
	"ovh_dbaas_logs_cluster":                   config.IdentifierFromProvider,
	"ovh_dbaas_logs_input":                     config.IdentifierFromProvider,
	"ovh_dbaas_logs_output_graylog_stream":     config.IdentifierFromProvider,
	"ovh_dbaas_logs_output_opensearch_alias":   config.IdentifierFromProvider,
	"ovh_dbaas_logs_output_opensearch_index":   config.IdentifierFromProvider,
	"ovh_dbaas_logs_role":                      config.IdentifierFromProvider,
	"ovh_dbaas_logs_role_permission_stream":    config.IdentifierFromProvider,
	"ovh_dbaas_logs_token":                     config.IdentifierFromProvider,

	// Dedicated Server resources
	"ovh_dedicated_ceph_acl":              config.IdentifierFromProvider,
	"ovh_dedicated_nasha_partition":       config.IdentifierFromProvider,
	"ovh_dedicated_nasha_partition_access": config.IdentifierFromProvider,
	"ovh_dedicated_nasha_partition_snapshot": config.IdentifierFromProvider,
	"ovh_dedicated_server":                config.IdentifierFromProvider,
	"ovh_dedicated_server_networking":     config.IdentifierFromProvider,
	"ovh_dedicated_server_reboot_task":    config.IdentifierFromProvider,
	"ovh_dedicated_server_reinstall_task": config.IdentifierFromProvider,
	"ovh_dedicated_server_update":         config.IdentifierFromProvider,

	// Domain resources
	"ovh_domain_ds_records":         config.IdentifierFromProvider,
	"ovh_domain_name":               config.IdentifierFromProvider,
	"ovh_domain_name_servers":       config.IdentifierFromProvider,
	"ovh_domain_zone":               config.IdentifierFromProvider,
	"ovh_domain_zone_dnssec":        config.IdentifierFromProvider,
	"ovh_domain_zone_dynhost_login": config.IdentifierFromProvider,
	"ovh_domain_zone_dynhost_record": config.IdentifierFromProvider,
	"ovh_domain_zone_import":        config.IdentifierFromProvider,
	"ovh_domain_zone_record":        config.IdentifierFromProvider,
	"ovh_domain_zone_redirection":   config.IdentifierFromProvider,

	// Hosting resources
	"ovh_hosting_privatedatabase":            config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase_database":   config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase_user":       config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase_user_grant": config.IdentifierFromProvider,
	"ovh_hosting_privatedatabase_whitelist":  config.IdentifierFromProvider,

	// IAM resources
	"ovh_iam_permissions_group": config.IdentifierFromProvider,
	"ovh_iam_policy":            config.IdentifierFromProvider,
	"ovh_iam_resource_group":    config.IdentifierFromProvider,

	// IP resources
	"ovh_ip_firewall":      config.IdentifierFromProvider,
	"ovh_ip_firewall_rule": config.IdentifierFromProvider,
	"ovh_ip_mitigation":    config.IdentifierFromProvider,
	"ovh_ip_move":          config.IdentifierFromProvider,
	"ovh_ip_reverse":       config.IdentifierFromProvider,
	"ovh_ip_service":       config.IdentifierFromProvider,

	// IP Load Balancing resources
	"ovh_iploadbalancing":                   config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_farm":         config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_farm_server":  config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_frontend":     config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_route":        config.IdentifierFromProvider,
	"ovh_iploadbalancing_http_route_rule":   config.IdentifierFromProvider,
	"ovh_iploadbalancing_refresh":           config.IdentifierFromProvider,
	"ovh_iploadbalancing_ssl":               config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_farm":          config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_farm_server":   config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_frontend":      config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_route":         config.IdentifierFromProvider,
	"ovh_iploadbalancing_tcp_route_rule":    config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_farm":          config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_farm_server":   config.IdentifierFromProvider,
	"ovh_iploadbalancing_udp_frontend":      config.IdentifierFromProvider,
	"ovh_iploadbalancing_vrack_network":     config.IdentifierFromProvider,

	// Me/Account resources
	"ovh_me_api_oauth2_client": config.IdentifierFromProvider,
	"ovh_me_identity_group":    config.IdentifierFromProvider,
	"ovh_me_identity_user":     config.IdentifierFromProvider,

	// OKMS (OVHcloud Key Management Service) resources
	"ovh_okms":                 config.IdentifierFromProvider,
	"ovh_okms_credential":      config.IdentifierFromProvider,
	"ovh_okms_service_key":     config.IdentifierFromProvider,
	"ovh_okms_service_key_jwk": config.IdentifierFromProvider,

	// OVHcloud Connect resources
	"ovh_ovhcloud_connect_pop_config":                  config.IdentifierFromProvider,
	"ovh_ovhcloud_connect_pop_datacenter_config":       config.IdentifierFromProvider,
	"ovh_ovhcloud_connect_pop_datacenter_extra_config": config.IdentifierFromProvider,

	// Savings Plan resources
	"ovh_savings_plan": config.IdentifierFromProvider,

	// Storage EFS resources
	"ovh_storage_efs_share":          config.IdentifierFromProvider,
	"ovh_storage_efs_share_acl":      config.IdentifierFromProvider,
	"ovh_storage_efs_share_snapshot": config.IdentifierFromProvider,

	// VPS resources
	"ovh_vps": config.IdentifierFromProvider,

	// vRack resources
	"ovh_vrack":                              config.IdentifierFromProvider,
	"ovh_vrack_cloudproject":                 config.IdentifierFromProvider,
	"ovh_vrack_dedicated_cloud":              config.IdentifierFromProvider,
	"ovh_vrack_dedicated_cloud_datacenter":   config.IdentifierFromProvider,
	"ovh_vrack_dedicated_server":             config.IdentifierFromProvider,
	"ovh_vrack_dedicated_server_interface":   config.IdentifierFromProvider,
	"ovh_vrack_ip":                           config.IdentifierFromProvider,
	"ovh_vrack_iploadbalancing":              config.IdentifierFromProvider,
	"ovh_vrack_ipv6":                         config.IdentifierFromProvider,
	"ovh_vrack_ipv6_routed_subrange":         config.IdentifierFromProvider,
	"ovh_vrack_ovhcloudconnect":              config.IdentifierFromProvider,
	"ovh_vrack_vrackservices":                config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
