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

// PoolAbstractBlockLeaseAllOf Definition of the list of properties defined in 'pool.AbstractBlockLease', excluding properties defined in parent classes.
type PoolAbstractBlockLeaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Count of number of leases requested.
	Count                *int64 `json:"Count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PoolAbstractBlockLeaseAllOf PoolAbstractBlockLeaseAllOf

// NewPoolAbstractBlockLeaseAllOf instantiates a new PoolAbstractBlockLeaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractBlockLeaseAllOf(classId string, objectType string) *PoolAbstractBlockLeaseAllOf {
	this := PoolAbstractBlockLeaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPoolAbstractBlockLeaseAllOfWithDefaults instantiates a new PoolAbstractBlockLeaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractBlockLeaseAllOfWithDefaults() *PoolAbstractBlockLeaseAllOf {
	this := PoolAbstractBlockLeaseAllOf{}
	var classId string = "ippool.BlockLease"
	this.ClassId = classId
	var objectType string = "ippool.BlockLease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PoolAbstractBlockLeaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockLeaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PoolAbstractBlockLeaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PoolAbstractBlockLeaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockLeaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PoolAbstractBlockLeaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PoolAbstractBlockLeaseAllOf) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractBlockLeaseAllOf) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PoolAbstractBlockLeaseAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *PoolAbstractBlockLeaseAllOf) SetCount(v int64) {
	o.Count = &v
}

func (o PoolAbstractBlockLeaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PoolAbstractBlockLeaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPoolAbstractBlockLeaseAllOf := _PoolAbstractBlockLeaseAllOf{}

	if err = json.Unmarshal(bytes, &varPoolAbstractBlockLeaseAllOf); err == nil {
		*o = PoolAbstractBlockLeaseAllOf(varPoolAbstractBlockLeaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePoolAbstractBlockLeaseAllOf struct {
	value *PoolAbstractBlockLeaseAllOf
	isSet bool
}

func (v NullablePoolAbstractBlockLeaseAllOf) Get() *PoolAbstractBlockLeaseAllOf {
	return v.value
}

func (v *NullablePoolAbstractBlockLeaseAllOf) Set(val *PoolAbstractBlockLeaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractBlockLeaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractBlockLeaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractBlockLeaseAllOf(val *PoolAbstractBlockLeaseAllOf) *NullablePoolAbstractBlockLeaseAllOf {
	return &NullablePoolAbstractBlockLeaseAllOf{value: val, isSet: true}
}

func (v NullablePoolAbstractBlockLeaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractBlockLeaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
