/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-27T10:08:12Z.
 *
 * API version: 1.0.9-4136
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// MemoryUnitAllOf Definition of the list of properties defined in 'memory.Unit', excluding properties defined in parent classes.
type MemoryUnitAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This represents the ID of a regular DIMM on a server.
	MemoryId             *int64                               `json:"MemoryId,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	MemoryArray          *MemoryArrayRelationship             `json:"MemoryArray,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryUnitAllOf MemoryUnitAllOf

// NewMemoryUnitAllOf instantiates a new MemoryUnitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryUnitAllOf(classId string, objectType string) *MemoryUnitAllOf {
	this := MemoryUnitAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryUnitAllOfWithDefaults instantiates a new MemoryUnitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryUnitAllOfWithDefaults() *MemoryUnitAllOf {
	this := MemoryUnitAllOf{}
	var classId string = "memory.Unit"
	this.ClassId = classId
	var objectType string = "memory.Unit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryUnitAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryUnitAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryUnitAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryUnitAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryUnitAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryUnitAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMemoryId returns the MemoryId field value if set, zero value otherwise.
func (o *MemoryUnitAllOf) GetMemoryId() int64 {
	if o == nil || o.MemoryId == nil {
		var ret int64
		return ret
	}
	return *o.MemoryId
}

// GetMemoryIdOk returns a tuple with the MemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryUnitAllOf) GetMemoryIdOk() (*int64, bool) {
	if o == nil || o.MemoryId == nil {
		return nil, false
	}
	return o.MemoryId, true
}

// HasMemoryId returns a boolean if a field has been set.
func (o *MemoryUnitAllOf) HasMemoryId() bool {
	if o != nil && o.MemoryId != nil {
		return true
	}

	return false
}

// SetMemoryId gets a reference to the given int64 and assigns it to the MemoryId field.
func (o *MemoryUnitAllOf) SetMemoryId(v int64) {
	o.MemoryId = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *MemoryUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryUnitAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMemoryArray returns the MemoryArray field value if set, zero value otherwise.
func (o *MemoryUnitAllOf) GetMemoryArray() MemoryArrayRelationship {
	if o == nil || o.MemoryArray == nil {
		var ret MemoryArrayRelationship
		return ret
	}
	return *o.MemoryArray
}

// GetMemoryArrayOk returns a tuple with the MemoryArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryUnitAllOf) GetMemoryArrayOk() (*MemoryArrayRelationship, bool) {
	if o == nil || o.MemoryArray == nil {
		return nil, false
	}
	return o.MemoryArray, true
}

// HasMemoryArray returns a boolean if a field has been set.
func (o *MemoryUnitAllOf) HasMemoryArray() bool {
	if o != nil && o.MemoryArray != nil {
		return true
	}

	return false
}

// SetMemoryArray gets a reference to the given MemoryArrayRelationship and assigns it to the MemoryArray field.
func (o *MemoryUnitAllOf) SetMemoryArray(v MemoryArrayRelationship) {
	o.MemoryArray = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *MemoryUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryUnitAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o MemoryUnitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MemoryId != nil {
		toSerialize["MemoryId"] = o.MemoryId
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.MemoryArray != nil {
		toSerialize["MemoryArray"] = o.MemoryArray
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryUnitAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMemoryUnitAllOf := _MemoryUnitAllOf{}

	if err = json.Unmarshal(bytes, &varMemoryUnitAllOf); err == nil {
		*o = MemoryUnitAllOf(varMemoryUnitAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MemoryId")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryArray")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryUnitAllOf struct {
	value *MemoryUnitAllOf
	isSet bool
}

func (v NullableMemoryUnitAllOf) Get() *MemoryUnitAllOf {
	return v.value
}

func (v *NullableMemoryUnitAllOf) Set(val *MemoryUnitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryUnitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryUnitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryUnitAllOf(val *MemoryUnitAllOf) *NullableMemoryUnitAllOf {
	return &NullableMemoryUnitAllOf{value: val, isSet: true}
}

func (v NullableMemoryUnitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryUnitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
