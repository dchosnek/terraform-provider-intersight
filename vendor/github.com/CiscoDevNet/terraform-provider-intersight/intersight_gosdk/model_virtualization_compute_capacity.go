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

// VirtualizationComputeCapacity Details of available CPU power across all cores and the free capacity still available.
type VirtualizationComputeCapacity struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total capacity of the entity in MHz.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Free CPU capacity in MHz, as a point-in-time snapshot. The available CPU capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used CPU capacity is also reported.
	Free *int64 `json:"Free,omitempty"`
	// Used CPU capacity of the entity in MHz, as a point-in-time snapshot.
	Used                 *int64 `json:"Used,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationComputeCapacity VirtualizationComputeCapacity

// NewVirtualizationComputeCapacity instantiates a new VirtualizationComputeCapacity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationComputeCapacity(classId string, objectType string) *VirtualizationComputeCapacity {
	this := VirtualizationComputeCapacity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationComputeCapacityWithDefaults instantiates a new VirtualizationComputeCapacity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationComputeCapacityWithDefaults() *VirtualizationComputeCapacity {
	this := VirtualizationComputeCapacity{}
	var classId string = "virtualization.ComputeCapacity"
	this.ClassId = classId
	var objectType string = "virtualization.ComputeCapacity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationComputeCapacity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationComputeCapacity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationComputeCapacity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationComputeCapacity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationComputeCapacity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationComputeCapacity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationComputeCapacity) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationComputeCapacity) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationComputeCapacity) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *VirtualizationComputeCapacity) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *VirtualizationComputeCapacity) GetFree() int64 {
	if o == nil || o.Free == nil {
		var ret int64
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationComputeCapacity) GetFreeOk() (*int64, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *VirtualizationComputeCapacity) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given int64 and assigns it to the Free field.
func (o *VirtualizationComputeCapacity) SetFree(v int64) {
	o.Free = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *VirtualizationComputeCapacity) GetUsed() int64 {
	if o == nil || o.Used == nil {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationComputeCapacity) GetUsedOk() (*int64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *VirtualizationComputeCapacity) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *VirtualizationComputeCapacity) SetUsed(v int64) {
	o.Used = &v
}

func (o VirtualizationComputeCapacity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Free != nil {
		toSerialize["Free"] = o.Free
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationComputeCapacity) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationComputeCapacityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Total capacity of the entity in MHz.
		Capacity *int64 `json:"Capacity,omitempty"`
		// Free CPU capacity in MHz, as a point-in-time snapshot. The available CPU capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used CPU capacity is also reported.
		Free *int64 `json:"Free,omitempty"`
		// Used CPU capacity of the entity in MHz, as a point-in-time snapshot.
		Used *int64 `json:"Used,omitempty"`
	}

	varVirtualizationComputeCapacityWithoutEmbeddedStruct := VirtualizationComputeCapacityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationComputeCapacityWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationComputeCapacity := _VirtualizationComputeCapacity{}
		varVirtualizationComputeCapacity.ClassId = varVirtualizationComputeCapacityWithoutEmbeddedStruct.ClassId
		varVirtualizationComputeCapacity.ObjectType = varVirtualizationComputeCapacityWithoutEmbeddedStruct.ObjectType
		varVirtualizationComputeCapacity.Capacity = varVirtualizationComputeCapacityWithoutEmbeddedStruct.Capacity
		varVirtualizationComputeCapacity.Free = varVirtualizationComputeCapacityWithoutEmbeddedStruct.Free
		varVirtualizationComputeCapacity.Used = varVirtualizationComputeCapacityWithoutEmbeddedStruct.Used
		*o = VirtualizationComputeCapacity(varVirtualizationComputeCapacity)
	} else {
		return err
	}

	varVirtualizationComputeCapacity := _VirtualizationComputeCapacity{}

	err = json.Unmarshal(bytes, &varVirtualizationComputeCapacity)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationComputeCapacity.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Free")
		delete(additionalProperties, "Used")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVirtualizationComputeCapacity struct {
	value *VirtualizationComputeCapacity
	isSet bool
}

func (v NullableVirtualizationComputeCapacity) Get() *VirtualizationComputeCapacity {
	return v.value
}

func (v *NullableVirtualizationComputeCapacity) Set(val *VirtualizationComputeCapacity) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationComputeCapacity) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationComputeCapacity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationComputeCapacity(val *VirtualizationComputeCapacity) *NullableVirtualizationComputeCapacity {
	return &NullableVirtualizationComputeCapacity{value: val, isSet: true}
}

func (v NullableVirtualizationComputeCapacity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationComputeCapacity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
