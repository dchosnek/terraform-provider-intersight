---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_disk_group_policy"
description: |-
  A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
---

# Resource: intersight_storage_disk_group_policy
A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
## Usage Example
### Resource Creation

```hcl
resource "intersight_storage_disk_group_policy" "storage_disk_group3" {
  name        = "storage_disk_group2"
  description = "Disk Group Test policy"
  raid_level  = "Raid1"
  use_jbods   = true
  span_groups {
    disks {
      slot_number = 2
    }
    disks {
      slot_number = 4
    }
  }
  dedicated_hot_spares {
    slot_number = 5
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(Computed) The Account ID for this managed object. 
* `ancestors`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(Computed) The time when this managed object was created. 
* `dedicated_hot_spares`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `slot_number`:(int) The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string)(Computed) The DomainGroup ID for this managed object. 
* `mod_time`:(string)(Computed) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(Computed)
* `parent`:(HashMap) -(Computed) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `raid_level`:(string) The supported RAID level for the disk group.* `Raid0` - RAID 0 Stripe Raid Level.* `Raid1` - RAID 1 Mirror Raid Level.* `Raid5` - RAID 5 Mirror Raid Level.* `Raid6` - RAID 6 Mirror Raid Level.* `Raid10` - RAID 10 Mirror Raid Level.* `Raid50` - RAID 50 Mirror Raid Level.* `Raid60` - RAID 60 Mirror Raid Level. 
* `shared_scope`:(string)(Computed) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `span_groups`:(Array)
This complex property has following sub-properties:
  + `disks`:(Array)
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `slot_number`:(int) The slot number of the disk to be referenced. As this is a policy, this slot number may or may not be valid depending on the number of disks in the associated server. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `storage_policies`:(Array) An array of relationships to storageStoragePolicy resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `use_jbods`:(bool) Selected disks in the disk group in JBOD state will be set to Unconfigured Good state before they are used in virtual drive creation. 
* `version_context`:(HashMap) -(Computed) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(Computed) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(Computed) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(Computed) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(Computed) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_storage_disk_group_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_storage_disk_group_policy.example 1234567890987654321abcde
``` 