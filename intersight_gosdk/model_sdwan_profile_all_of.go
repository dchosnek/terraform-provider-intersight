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

// SdwanProfileAllOf Definition of the list of properties defined in 'sdwan.Profile', excluding properties defined in parent classes.
type SdwanProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                                `json:"ObjectType"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to sdwanRouterNode resources.
	RouterNodes          []SdwanRouterNodeRelationship          `json:"RouterNodes,omitempty"`
	RouterPolicy         *SdwanRouterPolicyRelationship         `json:"RouterPolicy,omitempty"`
	VmanageAccount       *SdwanVmanageAccountPolicyRelationship `json:"VmanageAccount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdwanProfileAllOf SdwanProfileAllOf

// NewSdwanProfileAllOf instantiates a new SdwanProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdwanProfileAllOf(classId string, objectType string) *SdwanProfileAllOf {
	this := SdwanProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSdwanProfileAllOfWithDefaults instantiates a new SdwanProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdwanProfileAllOfWithDefaults() *SdwanProfileAllOf {
	this := SdwanProfileAllOf{}
	var classId string = "sdwan.Profile"
	this.ClassId = classId
	var objectType string = "sdwan.Profile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdwanProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdwanProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdwanProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdwanProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdwanProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdwanProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SdwanProfileAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanProfileAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SdwanProfileAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SdwanProfileAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRouterNodes returns the RouterNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdwanProfileAllOf) GetRouterNodes() []SdwanRouterNodeRelationship {
	if o == nil {
		var ret []SdwanRouterNodeRelationship
		return ret
	}
	return o.RouterNodes
}

// GetRouterNodesOk returns a tuple with the RouterNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdwanProfileAllOf) GetRouterNodesOk() (*[]SdwanRouterNodeRelationship, bool) {
	if o == nil || o.RouterNodes == nil {
		return nil, false
	}
	return &o.RouterNodes, true
}

// HasRouterNodes returns a boolean if a field has been set.
func (o *SdwanProfileAllOf) HasRouterNodes() bool {
	if o != nil && o.RouterNodes != nil {
		return true
	}

	return false
}

// SetRouterNodes gets a reference to the given []SdwanRouterNodeRelationship and assigns it to the RouterNodes field.
func (o *SdwanProfileAllOf) SetRouterNodes(v []SdwanRouterNodeRelationship) {
	o.RouterNodes = v
}

// GetRouterPolicy returns the RouterPolicy field value if set, zero value otherwise.
func (o *SdwanProfileAllOf) GetRouterPolicy() SdwanRouterPolicyRelationship {
	if o == nil || o.RouterPolicy == nil {
		var ret SdwanRouterPolicyRelationship
		return ret
	}
	return *o.RouterPolicy
}

// GetRouterPolicyOk returns a tuple with the RouterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanProfileAllOf) GetRouterPolicyOk() (*SdwanRouterPolicyRelationship, bool) {
	if o == nil || o.RouterPolicy == nil {
		return nil, false
	}
	return o.RouterPolicy, true
}

// HasRouterPolicy returns a boolean if a field has been set.
func (o *SdwanProfileAllOf) HasRouterPolicy() bool {
	if o != nil && o.RouterPolicy != nil {
		return true
	}

	return false
}

// SetRouterPolicy gets a reference to the given SdwanRouterPolicyRelationship and assigns it to the RouterPolicy field.
func (o *SdwanProfileAllOf) SetRouterPolicy(v SdwanRouterPolicyRelationship) {
	o.RouterPolicy = &v
}

// GetVmanageAccount returns the VmanageAccount field value if set, zero value otherwise.
func (o *SdwanProfileAllOf) GetVmanageAccount() SdwanVmanageAccountPolicyRelationship {
	if o == nil || o.VmanageAccount == nil {
		var ret SdwanVmanageAccountPolicyRelationship
		return ret
	}
	return *o.VmanageAccount
}

// GetVmanageAccountOk returns a tuple with the VmanageAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanProfileAllOf) GetVmanageAccountOk() (*SdwanVmanageAccountPolicyRelationship, bool) {
	if o == nil || o.VmanageAccount == nil {
		return nil, false
	}
	return o.VmanageAccount, true
}

// HasVmanageAccount returns a boolean if a field has been set.
func (o *SdwanProfileAllOf) HasVmanageAccount() bool {
	if o != nil && o.VmanageAccount != nil {
		return true
	}

	return false
}

// SetVmanageAccount gets a reference to the given SdwanVmanageAccountPolicyRelationship and assigns it to the VmanageAccount field.
func (o *SdwanProfileAllOf) SetVmanageAccount(v SdwanVmanageAccountPolicyRelationship) {
	o.VmanageAccount = &v
}

func (o SdwanProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RouterNodes != nil {
		toSerialize["RouterNodes"] = o.RouterNodes
	}
	if o.RouterPolicy != nil {
		toSerialize["RouterPolicy"] = o.RouterPolicy
	}
	if o.VmanageAccount != nil {
		toSerialize["VmanageAccount"] = o.VmanageAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdwanProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSdwanProfileAllOf := _SdwanProfileAllOf{}

	if err = json.Unmarshal(bytes, &varSdwanProfileAllOf); err == nil {
		*o = SdwanProfileAllOf(varSdwanProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RouterNodes")
		delete(additionalProperties, "RouterPolicy")
		delete(additionalProperties, "VmanageAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdwanProfileAllOf struct {
	value *SdwanProfileAllOf
	isSet bool
}

func (v NullableSdwanProfileAllOf) Get() *SdwanProfileAllOf {
	return v.value
}

func (v *NullableSdwanProfileAllOf) Set(val *SdwanProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSdwanProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSdwanProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdwanProfileAllOf(val *SdwanProfileAllOf) *NullableSdwanProfileAllOf {
	return &NullableSdwanProfileAllOf{value: val, isSet: true}
}

func (v NullableSdwanProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdwanProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}