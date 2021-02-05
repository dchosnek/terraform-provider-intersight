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

// StoragePureHostUtilization Storage space utilization of Pure FlashArray host entity.
type StoragePureHostUtilization struct {
	StorageStorageUtilization
	AdditionalProperties map[string]interface{}
}

type _StoragePureHostUtilization StoragePureHostUtilization

// NewStoragePureHostUtilization instantiates a new StoragePureHostUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureHostUtilization() *StoragePureHostUtilization {
	this := StoragePureHostUtilization{}
	return &this
}

// NewStoragePureHostUtilizationWithDefaults instantiates a new StoragePureHostUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureHostUtilizationWithDefaults() *StoragePureHostUtilization {
	this := StoragePureHostUtilization{}
	return &this
}

func (o StoragePureHostUtilization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageStorageUtilization, errStorageStorageUtilization := json.Marshal(o.StorageStorageUtilization)
	if errStorageStorageUtilization != nil {
		return []byte{}, errStorageStorageUtilization
	}
	errStorageStorageUtilization = json.Unmarshal([]byte(serializedStorageStorageUtilization), &toSerialize)
	if errStorageStorageUtilization != nil {
		return []byte{}, errStorageStorageUtilization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureHostUtilization) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureHostUtilizationWithoutEmbeddedStruct struct {
	}

	varStoragePureHostUtilizationWithoutEmbeddedStruct := StoragePureHostUtilizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureHostUtilizationWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureHostUtilization := _StoragePureHostUtilization{}
		*o = StoragePureHostUtilization(varStoragePureHostUtilization)
	} else {
		return err
	}

	varStoragePureHostUtilization := _StoragePureHostUtilization{}

	err = json.Unmarshal(bytes, &varStoragePureHostUtilization)
	if err == nil {
		o.StorageStorageUtilization = varStoragePureHostUtilization.StorageStorageUtilization
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectStorageStorageUtilization := reflect.ValueOf(o.StorageStorageUtilization)
		for i := 0; i < reflectStorageStorageUtilization.Type().NumField(); i++ {
			t := reflectStorageStorageUtilization.Type().Field(i)

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

type NullableStoragePureHostUtilization struct {
	value *StoragePureHostUtilization
	isSet bool
}

func (v NullableStoragePureHostUtilization) Get() *StoragePureHostUtilization {
	return v.value
}

func (v *NullableStoragePureHostUtilization) Set(val *StoragePureHostUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureHostUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureHostUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureHostUtilization(val *StoragePureHostUtilization) *NullableStoragePureHostUtilization {
	return &NullableStoragePureHostUtilization{value: val, isSet: true}
}

func (v NullableStoragePureHostUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureHostUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}