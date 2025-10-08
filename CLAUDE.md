# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Project Overview

This is a Crossplane provider for OVH, built using
[Upjet](https://github.com/crossplane/upjet) code generation framework. It
exposes XRM-conformant managed resources for the OVH API by wrapping the
[Terraform provider for OVH](https://registry.terraform.io/providers/ovh/ovh/latest/docs).

Here a some reference documentation about upjet:

- [Generating a Crossplane provider](https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md)
- [Configuring a resource](https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md)
- [Testing resources by using Uptest](https://github.com/crossplane/upjet/blob/main/docs/testing-with-uptest.md)

**Key Technologies:**

- Go 1.21+
- Crossplane runtime and Upjet framework
- Terraform provider wrapper (OVH Terraform provider v2.8.0)
- Kubernetes controller-runtime
- Terraform version: 1.5.7 (locked to <1.6 due to BSL license restrictions)

## Common Commands

### Code Generation

```bash
# Run the code generation pipeline (must be run after modifying resource configurations)
go run cmd/generator/main.go "$PWD"
```

### Development

```bash
# Run provider locally out-of-cluster (requires kubeconfig)
make run

# Build the provider binary
make build

# Build everything (images, packages, etc.)
make all
```

### Testing

```bash
# Run Go tests
make go.test

# Run end-to-end tests (requires UPTEST_EXAMPLE_LIST and proper cluster setup)
make e2e

# Local deployment for testing
make local-deploy
```

### Code Quality

```bash
# Run linter
make go.lint

# Format code
make go.fmt

# Check for breaking CRD schema changes
make crddiff
```

### Git Submodules

```bash
# Update build submodules (required on first clone or when build scripts change)
make submodules
```

## Architecture

### Two-Binary Structure

1. **Generator** (`cmd/generator/main.go`): Generates Kubernetes CRDs,
   controllers, and Go types from Terraform provider schema
2. **Provider** (`cmd/provider/main.go`): The runtime controller that reconciles
   managed resources

### Directory Structure

- **`config/`**: Provider configuration and resource external name mappings

  - `provider.go`: Main provider configuration using Upjet
  - `external_name.go`: Maps Terraform resource names to Kubernetes external
    names
  - `schema.json`: Terraform provider schema (auto-generated)
  - `provider-metadata.yaml`: Resource metadata from Terraform docs
  - Resource-specific config packages (e.g., `config/null/`) when added

- **`apis/`**: Kubernetes API definitions

  - `v1alpha1/`: StoreConfig APIs for external secret stores
  - `v1beta1/`: ProviderConfig and ProviderConfigUsage APIs
  - Generated versioned APIs for managed resources (auto-generated)

- **`internal/clients/`**: Terraform provider setup and credential handling

  - `ovhcloud.go`: Implements credential extraction and Terraform provider
    configuration
  - Supports two authentication modes:
    - Application-based: `application_key`, `application_secret`, `consumer_key`
    - OAuth2-based: `client_id`, `client_secret`

- **`internal/controller/`**: Kubernetes controllers

  - Auto-generated controllers for each managed resource
  - `providerconfig/`: ProviderConfig controller
  - `zz_setup.go`: Controller registration (auto-generated)

- **`internal/features/`**: Feature flag definitions

- **`package/`**: Crossplane package artifacts (CRDs, metadata)

- **`examples/`**: Example manifests for users
- **`examples-generated/`**: Auto-generated examples

### Code Generation Flow

1. Terraform provider schema is fetched via `make generate.init` → creates
   `config/schema.json`
2. Resource configurations are defined in `config/` (e.g., adding entries to
   `ExternalNameConfigs`)
3. Running the generator (`go run cmd/generator/main.go "$PWD"`) produces:
   - API types in `apis/`
   - Controllers in `internal/controller/`
   - CRDs in `package/crds/`
   - Example manifests

### Adding New Resources

1. Add the Terraform resource name and external name config to
   `config/external_name.go`
2. (Optional) Create a resource-specific config package under
   `config/<resource-group>/config.go` for advanced configuration
3. Register the config function in `config/provider.go`
4. Run code generation: `go run cmd/generator/main.go "$PWD"`
5. Build and test the provider

### Credential Management

Credentials are stored in Kubernetes secrets and referenced via ProviderConfig.
The `internal/clients/ovhcloud.go` file handles:

- Extracting credentials from secrets using Crossplane's
  CommonCredentialExtractor
- Unmarshaling JSON credentials into the OVHCredentials struct
- Configuring the Terraform provider with the appropriate authentication method

Expected credential JSON format:

```json
{
  "endpoint": "ovh-eu",
  "application_key": "...",
  "application_secret": "...",
  "consumer_key": "..."
}
```

or OAuth2:

```json
{
  "endpoint": "ovh-eu",
  "client_id": "...",
  "client_secret": "..."
}
```

## Important Constraints

- **Terraform version must be < 1.6.0** due to BSL license restrictions
  (enforced in Makefile)
- Files prefixed with `zz_` are auto-generated and should not be manually edited
- The provider uses Upjet's code generation, so most controller logic is
  framework-provided
- External name configurations determine how Kubernetes resource names map to
  Terraform IDs

## Current State

The provider currently includes:

- `null_resource` as a reference/example resource
- ProviderConfig infrastructure for credential management
- Support for management policies (beta feature)
- Optional external secret stores (alpha feature)

The null resource configuration serves as a template for adding real OVHcloud
resources.
