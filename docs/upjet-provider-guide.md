# Generating a Crossplane Provider with Upjet

**Source:** https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md

This document is a reference guide for creating and configuring a Crossplane provider using the Upjet framework.

## Overview

Upjet is a code generation framework that allows creating Crossplane providers by wrapping existing Terraform providers. This guide demonstrates the process using examples.

## Preparation Steps

### 1. Create Repository
- Use the [upjet-provider-template](https://github.com/upbound/upjet-provider-template) repository as a starting point
- Naming convention: `provider-<name>` (e.g., `provider-github`, `provider-ovhcloud`)

### 2. Initial Setup

1. Clone the repository
2. Fetch upbound/build submodule:
```bash
make submodules
```

3. Prepare the code:
```bash
./hack/prepare.sh
```

### 3. Configure Terraform Provider Variables

Update the Makefile with these key variables:

| Variable | Description | Example |
|----------|-------------|---------|
| `TERRAFORM_PROVIDER_SOURCE` | Provider source from Terraform registry | `ovh/ovh` |
| `TERRAFORM_PROVIDER_REPO` | Provider's code repository URL | `https://github.com/ovh/terraform-provider-ovh` |
| `TERRAFORM_PROVIDER_VERSION` | Provider version | `2.8.0` |
| `TERRAFORM_PROVIDER_DOWNLOAD_NAME` | Provider name in Terraform registry | `terraform-provider-ovh` |
| `TERRAFORM_NATIVE_PROVIDER_BINARY` | Binary name pattern | `terraform-provider-ovh_2.8.0` |
| `TERRAFORM_DOCS_PATH` | Path to provider resource documentation | `docs/resources` |

### 4. Provider Configuration

1. **Add ProviderConfig logic** in `internal/clients/<provider>.go`:
   - Implement credential extraction
   - Define authentication keys
   - Configure Terraform provider setup function

2. **Key implementation points:**
   - Use `resource.CommonCredentialExtractor` for credential retrieval
   - Support multiple authentication methods if applicable
   - Return `terraform.Setup` with proper configuration

### 5. External Name Configuration

Configure external names in `config/external_name.go`:

**Common patterns:**
- `config.NameAsIdentifier` - Use Kubernetes resource name as Terraform identifier
- `config.IdentifierFromProvider` - Let Terraform generate the identifier
- `config.TemplatedStringAsIdentifier` - Use a template for complex identifiers

**Example:**
```go
var ExternalNameConfigs = map[string]config.ExternalName{
    "ovh_cloud_project": config.NameAsIdentifier,
    "ovh_cloud_project_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.service_name }}/{{ .external_name }}"),
}
```

### 6. Custom Resource Configurations

For advanced resource configuration:

1. **Create configuration directories:**
```bash
mkdir config/<resource-group>
```

2. **Create configuration file** `config/<resource-group>/config.go`:
```go
package resourcegroup

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("terraform_resource_name", func(r *config.Resource) {
        // Custom configuration here
        r.ShortGroup = "group"
        r.Kind = "ResourceKind"
    })
}
```

3. **Register configurations** in `config/provider.go`:
```go
for _, configure := range []func(provider *ujconfig.Provider){
    resourcegroup.Configure,
} {
    configure(pc)
}
```

### 7. Generate Provider

1. **Install goimports:**
```bash
go install golang.org/x/tools/cmd/goimports@latest
```

2. **Generate provider code:**
```bash
make generate
```

This will:
- Fetch Terraform provider schema
- Generate API types
- Generate controllers
- Generate CRDs
- Generate example manifests

## Testing Resources

### Local Testing

1. **Build and run locally:**
```bash
make run
```

2. **Create ProviderConfig:**
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: provider-secret
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "endpoint": "ovh-eu",
      "application_key": "...",
      "application_secret": "...",
      "consumer_key": "..."
    }
---
apiVersion: <provider>.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
```

3. **Create test resources** using examples from `examples/` directory

### End-to-End Testing

```bash
# Set up environment
export UPTEST_EXAMPLE_LIST="examples/resource1/resource.yaml,examples/resource2/resource.yaml"

# Run tests
make e2e
```

## Important Concepts

### Generated Files
Files prefixed with `zz_` are auto-generated and should **never** be manually edited:
- `zz_generated.deepcopy.go`
- `zz_generated.managed.go`
- `zz_controller.go`
- etc.

### External Name
The external name configuration determines:
- How Kubernetes resource names map to Terraform resource identifiers
- Import behavior for existing resources
- Cross-resource references

### Resource Configuration
Advanced configuration options include:
- Custom field mappings
- Reference resolution
- Sensitive field handling
- Resource-specific behavior

## Best Practices

1. **Start with simple resources** - Begin with resources that have minimal dependencies
2. **Test incrementally** - Generate and test one resource group at a time
3. **Document authentication** - Clearly document credential format and authentication methods
4. **Use meaningful external names** - Choose external name patterns that make sense for the provider
5. **Follow conventions** - Use consistent naming and organization patterns

## Troubleshooting

### Generation Issues
- Ensure `make submodules` has been run
- Verify Terraform provider variables are correct
- Check that external name configurations don't conflict

### Runtime Issues
- Verify ProviderConfig credentials are correct
- Check controller logs: `kubectl logs -n upbound-system <pod-name>`
- Ensure Terraform provider version compatibility

### Testing Issues
- Confirm kubeconfig is properly configured
- Verify all required secrets are created
- Check resource examples for correct API versions

## Additional Resources

- [Upjet Documentation](https://github.com/crossplane/upjet)
- [Crossplane Documentation](https://docs.crossplane.io)
- [Provider Template Repository](https://github.com/upbound/upjet-provider-template)
