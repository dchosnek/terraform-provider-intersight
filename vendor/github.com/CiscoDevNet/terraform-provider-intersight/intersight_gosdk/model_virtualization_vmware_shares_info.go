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
	"reflect"
	"strings"
)

// VirtualizationVmwareSharesInfo Shares are used to determine relative allocation between resource consumers. In general, a consumer with more shares gets proportionally more of the resource, subject to certain other constraints.
type VirtualizationVmwareSharesInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. If the shares value does not map to a predefined size, then the level is set as custom.
	Level *string `json:"Level,omitempty"`
	// The number of shares allocated. It is used to determine resource allocation in case of resource contention. Set if level is custom. If level is not set to custom, this value is ignored. Therefore, only shares with custom values can be compared.
	Shares               *int64 `json:"Shares,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareSharesInfo VirtualizationVmwareSharesInfo

// NewVirtualizationVmwareSharesInfo instantiates a new VirtualizationVmwareSharesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareSharesInfo(classId string, objectType string) *VirtualizationVmwareSharesInfo {
	this := VirtualizationVmwareSharesInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareSharesInfoWithDefaults instantiates a new VirtualizationVmwareSharesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareSharesInfoWithDefaults() *VirtualizationVmwareSharesInfo {
	this := VirtualizationVmwareSharesInfo{}
	var classId string = "virtualization.VmwareSharesInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareSharesInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareSharesInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareSharesInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareSharesInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareSharesInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareSharesInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareSharesInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *VirtualizationVmwareSharesInfo) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareSharesInfo) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *VirtualizationVmwareSharesInfo) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *VirtualizationVmwareSharesInfo) SetLevel(v string) {
	o.Level = &v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *VirtualizationVmwareSharesInfo) GetShares() int64 {
	if o == nil || o.Shares == nil {
		var ret int64
		return ret
	}
	return *o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareSharesInfo) GetSharesOk() (*int64, bool) {
	if o == nil || o.Shares == nil {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *VirtualizationVmwareSharesInfo) HasShares() bool {
	if o != nil && o.Shares != nil {
		return true
	}

	return false
}

// SetShares gets a reference to the given int64 and assigns it to the Shares field.
func (o *VirtualizationVmwareSharesInfo) SetShares(v int64) {
	o.Shares = &v
}

func (o VirtualizationVmwareSharesInfo) MarshalJSON() ([]byte, error) {
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
	if o.Level != nil {
		toSerialize["Level"] = o.Level
	}
	if o.Shares != nil {
		toSerialize["Shares"] = o.Shares
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareSharesInfo) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareSharesInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. If the shares value does not map to a predefined size, then the level is set as custom.
		Level *string `json:"Level,omitempty"`
		// The number of shares allocated. It is used to determine resource allocation in case of resource contention. Set if level is custom. If level is not set to custom, this value is ignored. Therefore, only shares with custom values can be compared.
		Shares *int64 `json:"Shares,omitempty"`
	}

	varVirtualizationVmwareSharesInfoWithoutEmbeddedStruct := VirtualizationVmwareSharesInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareSharesInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareSharesInfo := _VirtualizationVmwareSharesInfo{}
		varVirtualizationVmwareSharesInfo.ClassId = varVirtualizationVmwareSharesInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareSharesInfo.ObjectType = varVirtualizationVmwareSharesInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareSharesInfo.Level = varVirtualizationVmwareSharesInfoWithoutEmbeddedStruct.Level
		varVirtualizationVmwareSharesInfo.Shares = varVirtualizationVmwareSharesInfoWithoutEmbeddedStruct.Shares
		*o = VirtualizationVmwareSharesInfo(varVirtualizationVmwareSharesInfo)
	} else {
		return err
	}

	varVirtualizationVmwareSharesInfo := _VirtualizationVmwareSharesInfo{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareSharesInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmwareSharesInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Level")
		delete(additionalProperties, "Shares")

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

type NullableVirtualizationVmwareSharesInfo struct {
	value *VirtualizationVmwareSharesInfo
	isSet bool
}

func (v NullableVirtualizationVmwareSharesInfo) Get() *VirtualizationVmwareSharesInfo {
	return v.value
}

func (v *NullableVirtualizationVmwareSharesInfo) Set(val *VirtualizationVmwareSharesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareSharesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareSharesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareSharesInfo(val *VirtualizationVmwareSharesInfo) *NullableVirtualizationVmwareSharesInfo {
	return &NullableVirtualizationVmwareSharesInfo{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareSharesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareSharesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
