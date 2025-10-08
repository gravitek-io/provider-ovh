# OVH Resources with Schema Issues

This document lists resources from the OVH Terraform provider that have schema issues and could not be automatically generated with Upjet.

## Generation Status

- **Total resources in OVH Terraform provider**: 139
- **Successfully generated resources**: 127 (91%)
- **Resources with schema issues**: 12 (9%)

## Problematic Resources

The following resources have been temporarily excluded from generation and require custom configuration:

### 1. `ovh_cloud_project_alerting`
- **Problematic field**: `formatted_monthly_threshold`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 14

### 2. `ovh_cloud_project_loadbalancer`
- **Problematic field**: `floating_ip`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 50

### 3. `ovh_cloud_project_rancher`
- **Problematic field**: `current_state`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 54

### 4. `ovh_cloud_project_region`
- **Problematic field**: `services`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 55

### 5. `ovh_cloud_project_region_network`
- **Problematic field**: `subnet`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 57

### 6. `ovh_cloud_project_storage`
- **Problematic field**: `encryption`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 60

### 7. `ovh_cloud_project_volume`
- **Problematic field**: `sub_operations`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 64

### 8. `ovh_dedicated_server`
- **Problematic field**: `customizations`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 83
- **Note**: Important resource for dedicated server management

### 9. `ovh_domain_name`
- **Problematic field**: `current_state`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 91
- **Note**: Important resource for domain management

### 10. `ovh_okms`
- **Problematic field**: `iam`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 146

### 11. `ovh_okms_service_key_jwk`
- **Problematic field**: `iam`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 149

### 12. `ovh_vps`
- **Problematic field**: `iam`
- **Error type**: TypeInvalid
- **Line in external_name.go**: 165
- **Note**: Important resource for VPS management

## Problem Analysis

### Identified Issue Types

1. **`iam` fields** (3 resources): `ovh_okms`, `ovh_okms_service_key_jwk`, `ovh_vps`
   - Related to new IAM fields added by OVH
   - Invalid schema type for Upjet

2. **`current_state` fields** (2 resources): `ovh_cloud_project_rancher`, `ovh_domain_name`
   - Complex state fields

3. **Complex configuration fields** (7 resources)
   - Complex or nested data types that Upjet cannot automatically infer

## Possible Solutions

### Option 1: Custom Configuration per Resource

Create custom configurations in `config/<group>/config.go` for each problematic resource:

```go
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("ovh_dedicated_server", func(r *config.Resource) {
        // Ignore or transform the problematic field
        r.Schema.IgnoreFields("customizations")
        // OR
        r.Schema.SetEmbeddedObject("customizations")
    })
}
```

### Option 2: Terraform Schema Patch

Contact OVH to report schema issues in their Terraform provider.

### Option 3: Upjet Custom Types

Define custom types for complex fields.

### Option 4: Keep Commented Temporarily

Keep these resources commented and activate them progressively based on user needs.

## Next Steps

1. ✅ Document all problematic resources
2. ⏳ Create tickets for each resource requiring special attention
3. ⏳ Implement custom configurations for priority resources:
   - `ovh_dedicated_server` (high priority)
   - `ovh_domain_name` (high priority)
   - `ovh_vps` (medium priority)
4. ⏳ Test generation with custom configurations
5. ⏳ Document workarounds in user documentation

## References

- [Upjet Documentation - Custom Resource Configuration](https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md)
- [OVH Terraform Provider - Issues](https://github.com/ovh/terraform-provider-ovh/issues)
- Generation logs: `/tmp/gen_test_*.log`

## Important Notes

- The 12 excluded resources represent 9% of the total
- 127 resources work perfectly (91%)
- The provider is usable right now with the majority of OVH resources
- Problematic resources can be added progressively with custom configurations
