---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "spectrocloud_user Data Source - terraform-provider-spectrocloud"
subcategory: ""
description: |-
  
---

# spectrocloud_user (Data Source)



## Example Usage

```terraform
data "spectrocloud_user" "user1" {
  name = "Foo Bar"

  # (alternatively)
  # id =  "5fd0ca727c411c71b55a359c"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `name` (String)

### Read-Only

- `id` (String) The ID of this resource.


