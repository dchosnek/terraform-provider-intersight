---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_server_firmware_version_entry"
description: |-
  An entry specifying supported server firmware version in regex format.
---

# Data Source: intersight_hyperflex_server_firmware_version_entry
An entry specifying supported server firmware version in regex format.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_server_firmware_version_entry.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_platform`:(string) The server platform type that is applicable for the server firmware bundle version.* `M5` - M5 generation of UCS server.* `M3` - M3 generation of UCS server.* `M4` - M4 generation of UCS server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) The server firmware bundle version. 
 