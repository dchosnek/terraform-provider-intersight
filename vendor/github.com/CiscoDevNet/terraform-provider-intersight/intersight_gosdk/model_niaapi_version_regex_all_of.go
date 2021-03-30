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

// NiaapiVersionRegexAllOf Definition of the list of properties defined in 'niaapi.VersionRegex', excluding properties defined in parent classes.
type NiaapiVersionRegexAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                             `json:"ObjectType"`
	Apic       NullableNiaapiVersionRegexPlatform `json:"Apic,omitempty"`
	Dcnm       NullableNiaapiVersionRegexPlatform `json:"Dcnm,omitempty"`
	// Version number for the Version Regex data, also used as identity.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiVersionRegexAllOf NiaapiVersionRegexAllOf

// NewNiaapiVersionRegexAllOf instantiates a new NiaapiVersionRegexAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiVersionRegexAllOf(classId string, objectType string) *NiaapiVersionRegexAllOf {
	this := NiaapiVersionRegexAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiVersionRegexAllOfWithDefaults instantiates a new NiaapiVersionRegexAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiVersionRegexAllOfWithDefaults() *NiaapiVersionRegexAllOf {
	this := NiaapiVersionRegexAllOf{}
	var classId string = "niaapi.VersionRegex"
	this.ClassId = classId
	var objectType string = "niaapi.VersionRegex"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiVersionRegexAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiVersionRegexAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiVersionRegexAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiVersionRegexAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApic returns the Apic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexAllOf) GetApic() NiaapiVersionRegexPlatform {
	if o == nil || o.Apic.Get() == nil {
		var ret NiaapiVersionRegexPlatform
		return ret
	}
	return *o.Apic.Get()
}

// GetApicOk returns a tuple with the Apic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexAllOf) GetApicOk() (*NiaapiVersionRegexPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Apic.Get(), o.Apic.IsSet()
}

// HasApic returns a boolean if a field has been set.
func (o *NiaapiVersionRegexAllOf) HasApic() bool {
	if o != nil && o.Apic.IsSet() {
		return true
	}

	return false
}

// SetApic gets a reference to the given NullableNiaapiVersionRegexPlatform and assigns it to the Apic field.
func (o *NiaapiVersionRegexAllOf) SetApic(v NiaapiVersionRegexPlatform) {
	o.Apic.Set(&v)
}

// SetApicNil sets the value for Apic to be an explicit nil
func (o *NiaapiVersionRegexAllOf) SetApicNil() {
	o.Apic.Set(nil)
}

// UnsetApic ensures that no value is present for Apic, not even an explicit nil
func (o *NiaapiVersionRegexAllOf) UnsetApic() {
	o.Apic.Unset()
}

// GetDcnm returns the Dcnm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegexAllOf) GetDcnm() NiaapiVersionRegexPlatform {
	if o == nil || o.Dcnm.Get() == nil {
		var ret NiaapiVersionRegexPlatform
		return ret
	}
	return *o.Dcnm.Get()
}

// GetDcnmOk returns a tuple with the Dcnm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegexAllOf) GetDcnmOk() (*NiaapiVersionRegexPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dcnm.Get(), o.Dcnm.IsSet()
}

// HasDcnm returns a boolean if a field has been set.
func (o *NiaapiVersionRegexAllOf) HasDcnm() bool {
	if o != nil && o.Dcnm.IsSet() {
		return true
	}

	return false
}

// SetDcnm gets a reference to the given NullableNiaapiVersionRegexPlatform and assigns it to the Dcnm field.
func (o *NiaapiVersionRegexAllOf) SetDcnm(v NiaapiVersionRegexPlatform) {
	o.Dcnm.Set(&v)
}

// SetDcnmNil sets the value for Dcnm to be an explicit nil
func (o *NiaapiVersionRegexAllOf) SetDcnmNil() {
	o.Dcnm.Set(nil)
}

// UnsetDcnm ensures that no value is present for Dcnm, not even an explicit nil
func (o *NiaapiVersionRegexAllOf) UnsetDcnm() {
	o.Dcnm.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiaapiVersionRegexAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegexAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiaapiVersionRegexAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NiaapiVersionRegexAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o NiaapiVersionRegexAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Apic.IsSet() {
		toSerialize["Apic"] = o.Apic.Get()
	}
	if o.Dcnm.IsSet() {
		toSerialize["Dcnm"] = o.Dcnm.Get()
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiVersionRegexAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiVersionRegexAllOf := _NiaapiVersionRegexAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiVersionRegexAllOf); err == nil {
		*o = NiaapiVersionRegexAllOf(varNiaapiVersionRegexAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Apic")
		delete(additionalProperties, "Dcnm")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiVersionRegexAllOf struct {
	value *NiaapiVersionRegexAllOf
	isSet bool
}

func (v NullableNiaapiVersionRegexAllOf) Get() *NiaapiVersionRegexAllOf {
	return v.value
}

func (v *NullableNiaapiVersionRegexAllOf) Set(val *NiaapiVersionRegexAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiVersionRegexAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiVersionRegexAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiVersionRegexAllOf(val *NiaapiVersionRegexAllOf) *NullableNiaapiVersionRegexAllOf {
	return &NullableNiaapiVersionRegexAllOf{value: val, isSet: true}
}

func (v NullableNiaapiVersionRegexAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiVersionRegexAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
