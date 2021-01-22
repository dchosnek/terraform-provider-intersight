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

// IamLdapProvider LDAP Provider or LDAP Server for user authentication.
type IamLdapProvider struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// LDAP Server Port for connection establishment.
	Port *int64 `json:"Port,omitempty"`
	// LDAP Server Address, can be IP address or hostname.
	Server               *string                    `json:"Server,omitempty"`
	LdapPolicy           *IamLdapPolicyRelationship `json:"LdapPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamLdapProvider IamLdapProvider

// NewIamLdapProvider instantiates a new IamLdapProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLdapProvider(classId string, objectType string) *IamLdapProvider {
	this := IamLdapProvider{}
	this.ClassId = classId
	this.ObjectType = objectType
	var port int64 = 389
	this.Port = &port
	return &this
}

// NewIamLdapProviderWithDefaults instantiates a new IamLdapProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLdapProviderWithDefaults() *IamLdapProvider {
	this := IamLdapProvider{}
	var classId string = "iam.LdapProvider"
	this.ClassId = classId
	var objectType string = "iam.LdapProvider"
	this.ObjectType = objectType
	var port int64 = 389
	this.Port = &port
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamLdapProvider) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamLdapProvider) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamLdapProvider) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamLdapProvider) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamLdapProvider) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamLdapProvider) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *IamLdapProvider) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapProvider) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *IamLdapProvider) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *IamLdapProvider) SetPort(v int64) {
	o.Port = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *IamLdapProvider) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapProvider) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *IamLdapProvider) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *IamLdapProvider) SetServer(v string) {
	o.Server = &v
}

// GetLdapPolicy returns the LdapPolicy field value if set, zero value otherwise.
func (o *IamLdapProvider) GetLdapPolicy() IamLdapPolicyRelationship {
	if o == nil || o.LdapPolicy == nil {
		var ret IamLdapPolicyRelationship
		return ret
	}
	return *o.LdapPolicy
}

// GetLdapPolicyOk returns a tuple with the LdapPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapProvider) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool) {
	if o == nil || o.LdapPolicy == nil {
		return nil, false
	}
	return o.LdapPolicy, true
}

// HasLdapPolicy returns a boolean if a field has been set.
func (o *IamLdapProvider) HasLdapPolicy() bool {
	if o != nil && o.LdapPolicy != nil {
		return true
	}

	return false
}

// SetLdapPolicy gets a reference to the given IamLdapPolicyRelationship and assigns it to the LdapPolicy field.
func (o *IamLdapProvider) SetLdapPolicy(v IamLdapPolicyRelationship) {
	o.LdapPolicy = &v
}

func (o IamLdapProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.LdapPolicy != nil {
		toSerialize["LdapPolicy"] = o.LdapPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamLdapProvider) UnmarshalJSON(bytes []byte) (err error) {
	type IamLdapProviderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// LDAP Server Port for connection establishment.
		Port *int64 `json:"Port,omitempty"`
		// LDAP Server Address, can be IP address or hostname.
		Server     *string                    `json:"Server,omitempty"`
		LdapPolicy *IamLdapPolicyRelationship `json:"LdapPolicy,omitempty"`
	}

	varIamLdapProviderWithoutEmbeddedStruct := IamLdapProviderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamLdapProviderWithoutEmbeddedStruct)
	if err == nil {
		varIamLdapProvider := _IamLdapProvider{}
		varIamLdapProvider.ClassId = varIamLdapProviderWithoutEmbeddedStruct.ClassId
		varIamLdapProvider.ObjectType = varIamLdapProviderWithoutEmbeddedStruct.ObjectType
		varIamLdapProvider.Port = varIamLdapProviderWithoutEmbeddedStruct.Port
		varIamLdapProvider.Server = varIamLdapProviderWithoutEmbeddedStruct.Server
		varIamLdapProvider.LdapPolicy = varIamLdapProviderWithoutEmbeddedStruct.LdapPolicy
		*o = IamLdapProvider(varIamLdapProvider)
	} else {
		return err
	}

	varIamLdapProvider := _IamLdapProvider{}

	err = json.Unmarshal(bytes, &varIamLdapProvider)
	if err == nil {
		o.MoBaseMo = varIamLdapProvider.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "LdapPolicy")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamLdapProvider struct {
	value *IamLdapProvider
	isSet bool
}

func (v NullableIamLdapProvider) Get() *IamLdapProvider {
	return v.value
}

func (v *NullableIamLdapProvider) Set(val *IamLdapProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLdapProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLdapProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLdapProvider(val *IamLdapProvider) *NullableIamLdapProvider {
	return &NullableIamLdapProvider{value: val, isSet: true}
}

func (v NullableIamLdapProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLdapProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
