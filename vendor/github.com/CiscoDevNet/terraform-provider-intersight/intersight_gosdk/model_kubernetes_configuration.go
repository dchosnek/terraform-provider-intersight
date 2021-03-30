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

// KubernetesConfiguration Configuration to access a Kubernetes cluster.
type KubernetesConfiguration struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Kubernetes configuration to connect to the cluster as an system administrator.
	KubeConfig           *string `json:"KubeConfig,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesConfiguration KubernetesConfiguration

// NewKubernetesConfiguration instantiates a new KubernetesConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesConfiguration(classId string, objectType string) *KubernetesConfiguration {
	this := KubernetesConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesConfigurationWithDefaults instantiates a new KubernetesConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesConfigurationWithDefaults() *KubernetesConfiguration {
	this := KubernetesConfiguration{}
	var classId string = "kubernetes.Configuration"
	this.ClassId = classId
	var objectType string = "kubernetes.Configuration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKubeConfig returns the KubeConfig field value if set, zero value otherwise.
func (o *KubernetesConfiguration) GetKubeConfig() string {
	if o == nil || o.KubeConfig == nil {
		var ret string
		return ret
	}
	return *o.KubeConfig
}

// GetKubeConfigOk returns a tuple with the KubeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesConfiguration) GetKubeConfigOk() (*string, bool) {
	if o == nil || o.KubeConfig == nil {
		return nil, false
	}
	return o.KubeConfig, true
}

// HasKubeConfig returns a boolean if a field has been set.
func (o *KubernetesConfiguration) HasKubeConfig() bool {
	if o != nil && o.KubeConfig != nil {
		return true
	}

	return false
}

// SetKubeConfig gets a reference to the given string and assigns it to the KubeConfig field.
func (o *KubernetesConfiguration) SetKubeConfig(v string) {
	o.KubeConfig = &v
}

func (o KubernetesConfiguration) MarshalJSON() ([]byte, error) {
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
	if o.KubeConfig != nil {
		toSerialize["KubeConfig"] = o.KubeConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Kubernetes configuration to connect to the cluster as an system administrator.
		KubeConfig *string `json:"KubeConfig,omitempty"`
	}

	varKubernetesConfigurationWithoutEmbeddedStruct := KubernetesConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesConfiguration := _KubernetesConfiguration{}
		varKubernetesConfiguration.ClassId = varKubernetesConfigurationWithoutEmbeddedStruct.ClassId
		varKubernetesConfiguration.ObjectType = varKubernetesConfigurationWithoutEmbeddedStruct.ObjectType
		varKubernetesConfiguration.KubeConfig = varKubernetesConfigurationWithoutEmbeddedStruct.KubeConfig
		*o = KubernetesConfiguration(varKubernetesConfiguration)
	} else {
		return err
	}

	varKubernetesConfiguration := _KubernetesConfiguration{}

	err = json.Unmarshal(bytes, &varKubernetesConfiguration)
	if err == nil {
		o.MoBaseComplexType = varKubernetesConfiguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KubeConfig")

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

type NullableKubernetesConfiguration struct {
	value *KubernetesConfiguration
	isSet bool
}

func (v NullableKubernetesConfiguration) Get() *KubernetesConfiguration {
	return v.value
}

func (v *NullableKubernetesConfiguration) Set(val *KubernetesConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesConfiguration(val *KubernetesConfiguration) *NullableKubernetesConfiguration {
	return &NullableKubernetesConfiguration{value: val, isSet: true}
}

func (v NullableKubernetesConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
