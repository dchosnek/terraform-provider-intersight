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

// KubernetesAddonPolicyAllOf Definition of the list of properties defined in 'kubernetes.AddonPolicy', excluding properties defined in parent classes.
type KubernetesAddonPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// To determine if Addon Policy is automatically managed by the system.
	SystemManaged *bool `json:"SystemManaged,omitempty"`
	// An array of relationships to kubernetesAddon resources.
	Addons []KubernetesAddonRelationship `json:"Addons,omitempty"`
	// An array of relationships to kubernetesClusterProfile resources.
	ClusterProfiles      []KubernetesClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddonPolicyAllOf KubernetesAddonPolicyAllOf

// NewKubernetesAddonPolicyAllOf instantiates a new KubernetesAddonPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddonPolicyAllOf(classId string, objectType string) *KubernetesAddonPolicyAllOf {
	this := KubernetesAddonPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAddonPolicyAllOfWithDefaults instantiates a new KubernetesAddonPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonPolicyAllOfWithDefaults() *KubernetesAddonPolicyAllOf {
	this := KubernetesAddonPolicyAllOf{}
	var classId string = "kubernetes.AddonPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.AddonPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddonPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddonPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddonPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddonPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSystemManaged returns the SystemManaged field value if set, zero value otherwise.
func (o *KubernetesAddonPolicyAllOf) GetSystemManaged() bool {
	if o == nil || o.SystemManaged == nil {
		var ret bool
		return ret
	}
	return *o.SystemManaged
}

// GetSystemManagedOk returns a tuple with the SystemManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicyAllOf) GetSystemManagedOk() (*bool, bool) {
	if o == nil || o.SystemManaged == nil {
		return nil, false
	}
	return o.SystemManaged, true
}

// HasSystemManaged returns a boolean if a field has been set.
func (o *KubernetesAddonPolicyAllOf) HasSystemManaged() bool {
	if o != nil && o.SystemManaged != nil {
		return true
	}

	return false
}

// SetSystemManaged gets a reference to the given bool and assigns it to the SystemManaged field.
func (o *KubernetesAddonPolicyAllOf) SetSystemManaged(v bool) {
	o.SystemManaged = &v
}

// GetAddons returns the Addons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonPolicyAllOf) GetAddons() []KubernetesAddonRelationship {
	if o == nil {
		var ret []KubernetesAddonRelationship
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonPolicyAllOf) GetAddonsOk() (*[]KubernetesAddonRelationship, bool) {
	if o == nil || o.Addons == nil {
		return nil, false
	}
	return &o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *KubernetesAddonPolicyAllOf) HasAddons() bool {
	if o != nil && o.Addons != nil {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []KubernetesAddonRelationship and assigns it to the Addons field.
func (o *KubernetesAddonPolicyAllOf) SetAddons(v []KubernetesAddonRelationship) {
	o.Addons = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonPolicyAllOf) GetClusterProfiles() []KubernetesClusterProfileRelationship {
	if o == nil {
		var ret []KubernetesClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonPolicyAllOf) GetClusterProfilesOk() (*[]KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *KubernetesAddonPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []KubernetesClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *KubernetesAddonPolicyAllOf) SetClusterProfiles(v []KubernetesClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesAddonPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesAddonPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesAddonPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesAddonPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SystemManaged != nil {
		toSerialize["SystemManaged"] = o.SystemManaged
	}
	if o.Addons != nil {
		toSerialize["Addons"] = o.Addons
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAddonPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesAddonPolicyAllOf := _KubernetesAddonPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesAddonPolicyAllOf); err == nil {
		*o = KubernetesAddonPolicyAllOf(varKubernetesAddonPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SystemManaged")
		delete(additionalProperties, "Addons")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesAddonPolicyAllOf struct {
	value *KubernetesAddonPolicyAllOf
	isSet bool
}

func (v NullableKubernetesAddonPolicyAllOf) Get() *KubernetesAddonPolicyAllOf {
	return v.value
}

func (v *NullableKubernetesAddonPolicyAllOf) Set(val *KubernetesAddonPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonPolicyAllOf(val *KubernetesAddonPolicyAllOf) *NullableKubernetesAddonPolicyAllOf {
	return &NullableKubernetesAddonPolicyAllOf{value: val, isSet: true}
}

func (v NullableKubernetesAddonPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
