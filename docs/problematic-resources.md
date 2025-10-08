# Ressources OVH avec Problèmes de Schéma

Ce document liste les ressources du provider Terraform OVH qui ont des problèmes de schéma et qui n'ont pas pu être générées automatiquement avec Upjet.

## Statut de la Génération

- **Total de ressources dans le provider Terraform OVH** : 139
- **Ressources générées avec succès** : 127 (91%)
- **Ressources avec problèmes de schéma** : 12 (9%)

## Ressources Problématiques

Les ressources suivantes ont été temporairement exclues de la génération et nécessitent une configuration personnalisée :

### 1. `ovh_cloud_project_alerting`
- **Champ problématique** : `formatted_monthly_threshold`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 14

### 2. `ovh_cloud_project_loadbalancer`
- **Champ problématique** : `floating_ip`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 50

### 3. `ovh_cloud_project_rancher`
- **Champ problématique** : `current_state`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 54

### 4. `ovh_cloud_project_region`
- **Champ problématique** : `services`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 55

### 5. `ovh_cloud_project_region_network`
- **Champ problématique** : `subnet`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 57

### 6. `ovh_cloud_project_storage`
- **Champ problématique** : `encryption`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 60

### 7. `ovh_cloud_project_volume`
- **Champ problématique** : `sub_operations`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 64

### 8. `ovh_dedicated_server`
- **Champ problématique** : `customizations`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 83
- **Note** : Ressource importante pour la gestion des serveurs dédiés

### 9. `ovh_domain_name`
- **Champ problématique** : `current_state`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 91
- **Note** : Ressource importante pour la gestion des domaines

### 10. `ovh_okms`
- **Champ problématique** : `iam`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 146

### 11. `ovh_okms_service_key_jwk`
- **Champ problématique** : `iam`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 149

### 12. `ovh_vps`
- **Champ problématique** : `iam`
- **Type d'erreur** : TypeInvalid
- **Ligne dans external_name.go** : 165
- **Note** : Ressource importante pour la gestion des VPS

## Analyse des Problèmes

### Types de problèmes identifiés

1. **Champs `iam`** (3 ressources) : `ovh_okms`, `ovh_okms_service_key_jwk`, `ovh_vps`
   - Problème lié aux nouveaux champs IAM ajoutés par OVH
   - Type de schéma invalide pour Upjet

2. **Champs `current_state`** (2 ressources) : `ovh_cloud_project_rancher`, `ovh_domain_name`
   - Champs d'état complexes

3. **Champs de configuration complexe** (7 ressources)
   - Types de données complexes ou imbriqués que Upjet ne peut pas inférer automatiquement

## Solutions Possibles

### Option 1 : Configuration Personnalisée par Ressource

Créer des configurations personnalisées dans `config/<group>/config.go` pour chaque ressource problématique :

```go
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("ovh_dedicated_server", func(r *config.Resource) {
        // Ignorer ou transformer le champ problématique
        r.Schema.IgnoreFields("customizations")
        // OU
        r.Schema.SetEmbeddedObject("customizations")
    })
}
```

### Option 2 : Patch du Schéma Terraform

Contacte OVH pour signaler les problèmes de schéma dans leur provider Terraform.

### Option 3 : Upjet Custom Types

Définir des types personnalisés pour les champs complexes.

### Option 4 : Laisser Commentées Temporairement

Garder ces ressources commentées et les activer progressivement selon les besoins des utilisateurs.

## Prochaines Étapes

1. ✅ Documenter toutes les ressources problématiques
2. ⏳ Créer des tickets pour chaque ressource nécessitant une attention particulière
3. ⏳ Implémenter des configurations personnalisées pour les ressources prioritaires :
   - `ovh_dedicated_server` (priorité haute)
   - `ovh_domain_name` (priorité haute)
   - `ovh_vps` (priorité moyenne)
4. ⏳ Tester la génération avec les configurations personnalisées
5. ⏳ Documenter les workarounds dans la documentation utilisateur

## Références

- [Upjet Documentation - Custom Resource Configuration](https://github.com/crossplane/upjet/blob/main/docs/configuring-a-resource.md)
- [OVH Terraform Provider - Issues](https://github.com/ovh/terraform-provider-ovh/issues)
- Logs de génération : `/tmp/gen_test_*.log`

## Notes Importantes

- Les 12 ressources exclues représentent 9% du total
- 127 ressources fonctionnent parfaitement (91%)
- Le provider est utilisable dès maintenant avec la majorité des ressources OVH
- Les ressources problématiques peuvent être ajoutées progressivement avec des configurations personnalisées
