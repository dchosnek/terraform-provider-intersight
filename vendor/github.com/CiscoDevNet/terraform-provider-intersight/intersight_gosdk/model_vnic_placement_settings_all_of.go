/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VnicPlacementSettingsAllOf Definition of the list of properties defined in 'vnic.PlacementSettings', excluding properties defined in parent classes.
type VnicPlacementSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// PCIe Slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
	Id *string `json:"Id,omitempty"`
	// The PCI Link used as transport for the virtual interface. All VIC adapters have a single PCI link except VIC 1385 which has two.
	PciLink *int64 `json:"PciLink,omitempty"`
	// The fabric port to which the vNICs will be associated. * `None` - Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used. * `A` - Fabric A of the FI cluster. * `B` - Fabric B of the FI cluster.
	SwitchId *string `json:"SwitchId,omitempty"`
	// Adapter port on which the virtual interface will be created.
	Uplink               *int64 `json:"Uplink,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicPlacementSettingsAllOf VnicPlacementSettingsAllOf

// NewVnicPlacementSettingsAllOf instantiates a new VnicPlacementSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicPlacementSettingsAllOf(classId string, objectType string) *VnicPlacementSettingsAllOf {
	this := VnicPlacementSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var pciLink int64 = 0
	this.PciLink = &pciLink
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// NewVnicPlacementSettingsAllOfWithDefaults instantiates a new VnicPlacementSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicPlacementSettingsAllOfWithDefaults() *VnicPlacementSettingsAllOf {
	this := VnicPlacementSettingsAllOf{}
	var classId string = "vnic.PlacementSettings"
	this.ClassId = classId
	var objectType string = "vnic.PlacementSettings"
	this.ObjectType = objectType
	var pciLink int64 = 0
	this.PciLink = &pciLink
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicPlacementSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicPlacementSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicPlacementSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicPlacementSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VnicPlacementSettingsAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettingsAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VnicPlacementSettingsAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VnicPlacementSettingsAllOf) SetId(v string) {
	o.Id = &v
}

// GetPciLink returns the PciLink field value if set, zero value otherwise.
func (o *VnicPlacementSettingsAllOf) GetPciLink() int64 {
	if o == nil || o.PciLink == nil {
		var ret int64
		return ret
	}
	return *o.PciLink
}

// GetPciLinkOk returns a tuple with the PciLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettingsAllOf) GetPciLinkOk() (*int64, bool) {
	if o == nil || o.PciLink == nil {
		return nil, false
	}
	return o.PciLink, true
}

// HasPciLink returns a boolean if a field has been set.
func (o *VnicPlacementSettingsAllOf) HasPciLink() bool {
	if o != nil && o.PciLink != nil {
		return true
	}

	return false
}

// SetPciLink gets a reference to the given int64 and assigns it to the PciLink field.
func (o *VnicPlacementSettingsAllOf) SetPciLink(v int64) {
	o.PciLink = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *VnicPlacementSettingsAllOf) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettingsAllOf) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *VnicPlacementSettingsAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *VnicPlacementSettingsAllOf) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *VnicPlacementSettingsAllOf) GetUplink() int64 {
	if o == nil || o.Uplink == nil {
		var ret int64
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicPlacementSettingsAllOf) GetUplinkOk() (*int64, bool) {
	if o == nil || o.Uplink == nil {
		return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *VnicPlacementSettingsAllOf) HasUplink() bool {
	if o != nil && o.Uplink != nil {
		return true
	}

	return false
}

// SetUplink gets a reference to the given int64 and assigns it to the Uplink field.
func (o *VnicPlacementSettingsAllOf) SetUplink(v int64) {
	o.Uplink = &v
}

func (o VnicPlacementSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.PciLink != nil {
		toSerialize["PciLink"] = o.PciLink
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Uplink != nil {
		toSerialize["Uplink"] = o.Uplink
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicPlacementSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicPlacementSettingsAllOf := _VnicPlacementSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicPlacementSettingsAllOf); err == nil {
		*o = VnicPlacementSettingsAllOf(varVnicPlacementSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "PciLink")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Uplink")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicPlacementSettingsAllOf struct {
	value *VnicPlacementSettingsAllOf
	isSet bool
}

func (v NullableVnicPlacementSettingsAllOf) Get() *VnicPlacementSettingsAllOf {
	return v.value
}

func (v *NullableVnicPlacementSettingsAllOf) Set(val *VnicPlacementSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicPlacementSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicPlacementSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicPlacementSettingsAllOf(val *VnicPlacementSettingsAllOf) *NullableVnicPlacementSettingsAllOf {
	return &NullableVnicPlacementSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicPlacementSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicPlacementSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}