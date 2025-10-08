# ProviderConfig Examples

This directory contains examples for configuring authentication with the OVH provider.

## Authentication Methods

The OVH provider supports two authentication methods:

### 1. Application-based Authentication (Recommended)

This method uses OVH API credentials (Application Key, Application Secret, Consumer Key).

**Create your API credentials:**
1. Go to https://eu.api.ovh.com/createToken/ (or your region's equivalent)
2. Set the required permissions for your application
3. Generate your credentials

**Create the secret:**
```bash
kubectl create secret generic ovh-credentials \
  --namespace crossplane-system \
  --from-literal=credentials='{
    "endpoint": "ovh-eu",
    "application_key": "YOUR_APPLICATION_KEY",
    "application_secret": "YOUR_APPLICATION_SECRET",
    "consumer_key": "YOUR_CONSUMER_KEY"
  }'
```

Or use the template:
```bash
# Edit the template with your credentials
cp secret.yaml.tmpl secret.yaml
vi secret.yaml

# Apply the secret
kubectl apply -f secret.yaml
```

**Available endpoints:**
- `ovh-eu` - OVH Europe
- `ovh-us` - OVH US
- `ovh-ca` - OVH Canada
- `soyoustart-eu` - So you Start Europe
- `soyoustart-ca` - So you Start Canada
- `kimsufi-eu` - Kimsufi Europe
- `kimsufi-ca` - Kimsufi Canada

### 2. OAuth2 Authentication

This method uses OAuth2 client credentials.

**Create your OAuth2 credentials in the OVH Control Panel**

**Create the secret:**
```bash
kubectl create secret generic ovh-credentials-oauth2 \
  --namespace crossplane-system \
  --from-literal=credentials='{
    "endpoint": "ovh-eu",
    "client_id": "YOUR_CLIENT_ID",
    "client_secret": "YOUR_CLIENT_SECRET"
  }'
```

Or use the template:
```bash
# Edit the template with your credentials
cp secret-oauth2.yaml.tmpl secret-oauth2.yaml
vi secret-oauth2.yaml

# Apply the secret
kubectl apply -f secret-oauth2.yaml
```

## Create ProviderConfig

After creating the secret, create a ProviderConfig:

```bash
kubectl apply -f providerconfig.yaml
```

For OAuth2, update the secret reference in `providerconfig.yaml`:
```yaml
spec:
  credentials:
    source: Secret
    secretRef:
      name: ovh-credentials-oauth2  # Change this
      namespace: crossplane-system
      key: credentials
```

## Verify

Check that the ProviderConfig is ready:
```bash
kubectl get providerconfig
kubectl describe providerconfig default
```

## Security Best Practices

⚠️ **Never commit actual credentials to version control!**

- Always use `.tmpl` files for examples
- Store actual secrets in Kubernetes secrets or external secret managers
- Consider using [External Secrets Operator](https://external-secrets.io/) to sync from Vault, AWS Secrets Manager, etc.
- Use StoreConfig for external secret stores (see `../storeconfig/vault.yaml`)
