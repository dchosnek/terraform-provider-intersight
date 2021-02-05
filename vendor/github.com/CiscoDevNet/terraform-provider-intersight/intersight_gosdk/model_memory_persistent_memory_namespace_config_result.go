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
	"reflect"
	"strings"
)

// MemoryPersistentMemoryNamespaceConfigResult Result of a previously configured Persistent Memory Namespace on a server.
type MemoryPersistentMemoryNamespaceConfigResult struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the Persistent Memory Namespace needed to be configured.
	ConfigStatus *string `json:"ConfigStatus,omitempty"`
	// Name of a Persistent Memory Namespace that needed to be configured.
	Name *string `json:"Name,omitempty"`
	// Socket ID in which the Persistent Memory Namespace needed to be configured.
	SocketId *string `json:"SocketId,omitempty"`
	// Socket Memory ID in which the Persistent Memory Namespace needed to be configured.
	SocketMemoryId                     *string                                         `json:"SocketMemoryId,omitempty"`
	InventoryDeviceInfo                *InventoryDeviceInfoRelationship                `json:"InventoryDeviceInfo,omitempty"`
	MemoryPersistentMemoryConfigResult *MemoryPersistentMemoryConfigResultRelationship `json:"MemoryPersistentMemoryConfigResult,omitempty"`
	RegisteredDevice                   *AssetDeviceRegistrationRelationship            `json:"RegisteredDevice,omitempty"`
	AdditionalProperties               map[string]interface{}
}

type _MemoryPersistentMemoryNamespaceConfigResult MemoryPersistentMemoryNamespaceConfigResult

// NewMemoryPersistentMemoryNamespaceConfigResult instantiates a new MemoryPersistentMemoryNamespaceConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryNamespaceConfigResult(classId string, objectType string) *MemoryPersistentMemoryNamespaceConfigResult {
	this := MemoryPersistentMemoryNamespaceConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryNamespaceConfigResultWithDefaults instantiates a new MemoryPersistentMemoryNamespaceConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryNamespaceConfigResultWithDefaults() *MemoryPersistentMemoryNamespaceConfigResult {
	this := MemoryPersistentMemoryNamespaceConfigResult{}
	var classId string = "memory.PersistentMemoryNamespaceConfigResult"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryNamespaceConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigStatus returns the ConfigStatus field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetConfigStatus() string {
	if o == nil || o.ConfigStatus == nil {
		var ret string
		return ret
	}
	return *o.ConfigStatus
}

// GetConfigStatusOk returns a tuple with the ConfigStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetConfigStatusOk() (*string, bool) {
	if o == nil || o.ConfigStatus == nil {
		return nil, false
	}
	return o.ConfigStatus, true
}

// HasConfigStatus returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasConfigStatus() bool {
	if o != nil && o.ConfigStatus != nil {
		return true
	}

	return false
}

// SetConfigStatus gets a reference to the given string and assigns it to the ConfigStatus field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetConfigStatus(v string) {
	o.ConfigStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetName(v string) {
	o.Name = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketId() string {
	if o == nil || o.SocketId == nil {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketIdOk() (*string, bool) {
	if o == nil || o.SocketId == nil {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasSocketId() bool {
	if o != nil && o.SocketId != nil {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetSocketId(v string) {
	o.SocketId = &v
}

// GetSocketMemoryId returns the SocketMemoryId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketMemoryId() string {
	if o == nil || o.SocketMemoryId == nil {
		var ret string
		return ret
	}
	return *o.SocketMemoryId
}

// GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketMemoryIdOk() (*string, bool) {
	if o == nil || o.SocketMemoryId == nil {
		return nil, false
	}
	return o.SocketMemoryId, true
}

// HasSocketMemoryId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasSocketMemoryId() bool {
	if o != nil && o.SocketMemoryId != nil {
		return true
	}

	return false
}

// SetSocketMemoryId gets a reference to the given string and assigns it to the SocketMemoryId field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetSocketMemoryId(v string) {
	o.SocketMemoryId = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMemoryPersistentMemoryConfigResult returns the MemoryPersistentMemoryConfigResult field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetMemoryPersistentMemoryConfigResult() MemoryPersistentMemoryConfigResultRelationship {
	if o == nil || o.MemoryPersistentMemoryConfigResult == nil {
		var ret MemoryPersistentMemoryConfigResultRelationship
		return ret
	}
	return *o.MemoryPersistentMemoryConfigResult
}

// GetMemoryPersistentMemoryConfigResultOk returns a tuple with the MemoryPersistentMemoryConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetMemoryPersistentMemoryConfigResultOk() (*MemoryPersistentMemoryConfigResultRelationship, bool) {
	if o == nil || o.MemoryPersistentMemoryConfigResult == nil {
		return nil, false
	}
	return o.MemoryPersistentMemoryConfigResult, true
}

// HasMemoryPersistentMemoryConfigResult returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasMemoryPersistentMemoryConfigResult() bool {
	if o != nil && o.MemoryPersistentMemoryConfigResult != nil {
		return true
	}

	return false
}

// SetMemoryPersistentMemoryConfigResult gets a reference to the given MemoryPersistentMemoryConfigResultRelationship and assigns it to the MemoryPersistentMemoryConfigResult field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetMemoryPersistentMemoryConfigResult(v MemoryPersistentMemoryConfigResultRelationship) {
	o.MemoryPersistentMemoryConfigResult = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o MemoryPersistentMemoryNamespaceConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigStatus != nil {
		toSerialize["ConfigStatus"] = o.ConfigStatus
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SocketId != nil {
		toSerialize["SocketId"] = o.SocketId
	}
	if o.SocketMemoryId != nil {
		toSerialize["SocketMemoryId"] = o.SocketMemoryId
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.MemoryPersistentMemoryConfigResult != nil {
		toSerialize["MemoryPersistentMemoryConfigResult"] = o.MemoryPersistentMemoryConfigResult
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryPersistentMemoryNamespaceConfigResult) UnmarshalJSON(bytes []byte) (err error) {
	type MemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of the Persistent Memory Namespace needed to be configured.
		ConfigStatus *string `json:"ConfigStatus,omitempty"`
		// Name of a Persistent Memory Namespace that needed to be configured.
		Name *string `json:"Name,omitempty"`
		// Socket ID in which the Persistent Memory Namespace needed to be configured.
		SocketId *string `json:"SocketId,omitempty"`
		// Socket Memory ID in which the Persistent Memory Namespace needed to be configured.
		SocketMemoryId                     *string                                         `json:"SocketMemoryId,omitempty"`
		InventoryDeviceInfo                *InventoryDeviceInfoRelationship                `json:"InventoryDeviceInfo,omitempty"`
		MemoryPersistentMemoryConfigResult *MemoryPersistentMemoryConfigResultRelationship `json:"MemoryPersistentMemoryConfigResult,omitempty"`
		RegisteredDevice                   *AssetDeviceRegistrationRelationship            `json:"RegisteredDevice,omitempty"`
	}

	varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct := MemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryNamespaceConfigResult := _MemoryPersistentMemoryNamespaceConfigResult{}
		varMemoryPersistentMemoryNamespaceConfigResult.ClassId = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryNamespaceConfigResult.ObjectType = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryNamespaceConfigResult.ConfigStatus = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.ConfigStatus
		varMemoryPersistentMemoryNamespaceConfigResult.Name = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.Name
		varMemoryPersistentMemoryNamespaceConfigResult.SocketId = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.SocketId
		varMemoryPersistentMemoryNamespaceConfigResult.SocketMemoryId = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.SocketMemoryId
		varMemoryPersistentMemoryNamespaceConfigResult.InventoryDeviceInfo = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.InventoryDeviceInfo
		varMemoryPersistentMemoryNamespaceConfigResult.MemoryPersistentMemoryConfigResult = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.MemoryPersistentMemoryConfigResult
		varMemoryPersistentMemoryNamespaceConfigResult.RegisteredDevice = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.RegisteredDevice
		*o = MemoryPersistentMemoryNamespaceConfigResult(varMemoryPersistentMemoryNamespaceConfigResult)
	} else {
		return err
	}

	varMemoryPersistentMemoryNamespaceConfigResult := _MemoryPersistentMemoryNamespaceConfigResult{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryNamespaceConfigResult)
	if err == nil {
		o.InventoryBase = varMemoryPersistentMemoryNamespaceConfigResult.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigStatus")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SocketId")
		delete(additionalProperties, "SocketMemoryId")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryPersistentMemoryConfigResult")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryPersistentMemoryNamespaceConfigResult struct {
	value *MemoryPersistentMemoryNamespaceConfigResult
	isSet bool
}

func (v NullableMemoryPersistentMemoryNamespaceConfigResult) Get() *MemoryPersistentMemoryNamespaceConfigResult {
	return v.value
}

func (v *NullableMemoryPersistentMemoryNamespaceConfigResult) Set(val *MemoryPersistentMemoryNamespaceConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryNamespaceConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryNamespaceConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryNamespaceConfigResult(val *MemoryPersistentMemoryNamespaceConfigResult) *NullableMemoryPersistentMemoryNamespaceConfigResult {
	return &NullableMemoryPersistentMemoryNamespaceConfigResult{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryNamespaceConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryNamespaceConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}