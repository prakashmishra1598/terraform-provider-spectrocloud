---
page_title: "spectrocloud_application Resource - terraform-provider-spectrocloud"
subcategory: ""
description: |-
  
---

# spectrocloud_application (Resource)

  

## Example Usage


### Application deployment into Cluster Group.

```hcl
resource "spectrocloud_application" "application" {
  name                    = "app-beru-whitesun-lars"
  application_profile_uid = data.spectrocloud_application_profile.id

  config {
    cluster_name      = "sandbox-scorpius"
    cluster_group_uid = "6358d799fad5aa39fa26a8c2" # or data.spectrocloud_cluster_group.id
    limits {
      cpu     = 2
      memory  = 4096
      storage = 10
    }
  }
}
   
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_profile_uid` (String)
- `name` (String)

### Optional

- `config` (Block List, Max: 1) (see [below for nested schema](#nestedblock--config))
- `tags` (Set of String)
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--config"></a>
### Nested Schema for `config`

Optional:

- `cluster_group_uid` (String)
- `cluster_name` (String)
- `cluster_uid` (String)
- `limits` (Block List) (see [below for nested schema](#nestedblock--config--limits))

<a id="nestedblock--config--limits"></a>
### Nested Schema for `config.limits`

Optional:

- `cpu` (Number)
- `memory` (Number)
- `storage` (Number)



<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `update` (String)