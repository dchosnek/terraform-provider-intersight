/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-05-12T14:10:48Z.
 *
 * API version: 1.0.9-4289
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexHxapDvswitch A HyperFlex Application Platform cluster wise distributed vSwitch entity.
type HyperflexHxapDvswitch struct {
	VirtualizationBaseDvswitch
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the dvUplink referenced by this dvswitch.
	DvUplink *string `json:"DvUplink,omitempty"`
	// The last host that update this object.
	LastHostname *string                           `json:"LastHostname,omitempty"`
	Cluster      *HyperflexHxapClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to hyperflexHxapHost resources.
	MemberHosts []HyperflexHxapHostRelationship `json:"MemberHosts,omitempty"`
	// An array of relationships to hyperflexHxapHostVswitch resources.
	MemberVswitches      []HyperflexHxapHostVswitchRelationship `json:"MemberVswitches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxapDvswitch HyperflexHxapDvswitch

// NewHyperflexHxapDvswitch instantiates a new HyperflexHxapDvswitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxapDvswitch(classId string, objectType string) *HyperflexHxapDvswitch {
	this := HyperflexHxapDvswitch{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxapDvswitchWithDefaults instantiates a new HyperflexHxapDvswitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxapDvswitchWithDefaults() *HyperflexHxapDvswitch {
	this := HyperflexHxapDvswitch{}
	var classId string = "hyperflex.HxapDvswitch"
	this.ClassId = classId
	var objectType string = "hyperflex.HxapDvswitch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxapDvswitch) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitch) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxapDvswitch) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxapDvswitch) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitch) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxapDvswitch) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDvUplink returns the DvUplink field value if set, zero value otherwise.
func (o *HyperflexHxapDvswitch) GetDvUplink() string {
	if o == nil || o.DvUplink == nil {
		var ret string
		return ret
	}
	return *o.DvUplink
}

// GetDvUplinkOk returns a tuple with the DvUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitch) GetDvUplinkOk() (*string, bool) {
	if o == nil || o.DvUplink == nil {
		return nil, false
	}
	return o.DvUplink, true
}

// HasDvUplink returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitch) HasDvUplink() bool {
	if o != nil && o.DvUplink != nil {
		return true
	}

	return false
}

// SetDvUplink gets a reference to the given string and assigns it to the DvUplink field.
func (o *HyperflexHxapDvswitch) SetDvUplink(v string) {
	o.DvUplink = &v
}

// GetLastHostname returns the LastHostname field value if set, zero value otherwise.
func (o *HyperflexHxapDvswitch) GetLastHostname() string {
	if o == nil || o.LastHostname == nil {
		var ret string
		return ret
	}
	return *o.LastHostname
}

// GetLastHostnameOk returns a tuple with the LastHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitch) GetLastHostnameOk() (*string, bool) {
	if o == nil || o.LastHostname == nil {
		return nil, false
	}
	return o.LastHostname, true
}

// HasLastHostname returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitch) HasLastHostname() bool {
	if o != nil && o.LastHostname != nil {
		return true
	}

	return false
}

// SetLastHostname gets a reference to the given string and assigns it to the LastHostname field.
func (o *HyperflexHxapDvswitch) SetLastHostname(v string) {
	o.LastHostname = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHxapDvswitch) GetCluster() HyperflexHxapClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexHxapClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxapDvswitch) GetClusterOk() (*HyperflexHxapClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitch) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexHxapClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHxapDvswitch) SetCluster(v HyperflexHxapClusterRelationship) {
	o.Cluster = &v
}

// GetMemberHosts returns the MemberHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapDvswitch) GetMemberHosts() []HyperflexHxapHostRelationship {
	if o == nil {
		var ret []HyperflexHxapHostRelationship
		return ret
	}
	return o.MemberHosts
}

// GetMemberHostsOk returns a tuple with the MemberHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapDvswitch) GetMemberHostsOk() (*[]HyperflexHxapHostRelationship, bool) {
	if o == nil || o.MemberHosts == nil {
		return nil, false
	}
	return &o.MemberHosts, true
}

// HasMemberHosts returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitch) HasMemberHosts() bool {
	if o != nil && o.MemberHosts != nil {
		return true
	}

	return false
}

// SetMemberHosts gets a reference to the given []HyperflexHxapHostRelationship and assigns it to the MemberHosts field.
func (o *HyperflexHxapDvswitch) SetMemberHosts(v []HyperflexHxapHostRelationship) {
	o.MemberHosts = v
}

// GetMemberVswitches returns the MemberVswitches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxapDvswitch) GetMemberVswitches() []HyperflexHxapHostVswitchRelationship {
	if o == nil {
		var ret []HyperflexHxapHostVswitchRelationship
		return ret
	}
	return o.MemberVswitches
}

// GetMemberVswitchesOk returns a tuple with the MemberVswitches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxapDvswitch) GetMemberVswitchesOk() (*[]HyperflexHxapHostVswitchRelationship, bool) {
	if o == nil || o.MemberVswitches == nil {
		return nil, false
	}
	return &o.MemberVswitches, true
}

// HasMemberVswitches returns a boolean if a field has been set.
func (o *HyperflexHxapDvswitch) HasMemberVswitches() bool {
	if o != nil && o.MemberVswitches != nil {
		return true
	}

	return false
}

// SetMemberVswitches gets a reference to the given []HyperflexHxapHostVswitchRelationship and assigns it to the MemberVswitches field.
func (o *HyperflexHxapDvswitch) SetMemberVswitches(v []HyperflexHxapHostVswitchRelationship) {
	o.MemberVswitches = v
}

func (o HyperflexHxapDvswitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDvswitch, errVirtualizationBaseDvswitch := json.Marshal(o.VirtualizationBaseDvswitch)
	if errVirtualizationBaseDvswitch != nil {
		return []byte{}, errVirtualizationBaseDvswitch
	}
	errVirtualizationBaseDvswitch = json.Unmarshal([]byte(serializedVirtualizationBaseDvswitch), &toSerialize)
	if errVirtualizationBaseDvswitch != nil {
		return []byte{}, errVirtualizationBaseDvswitch
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DvUplink != nil {
		toSerialize["DvUplink"] = o.DvUplink
	}
	if o.LastHostname != nil {
		toSerialize["LastHostname"] = o.LastHostname
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.MemberHosts != nil {
		toSerialize["MemberHosts"] = o.MemberHosts
	}
	if o.MemberVswitches != nil {
		toSerialize["MemberVswitches"] = o.MemberVswitches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxapDvswitch) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHxapDvswitchWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the dvUplink referenced by this dvswitch.
		DvUplink *string `json:"DvUplink,omitempty"`
		// The last host that update this object.
		LastHostname *string                           `json:"LastHostname,omitempty"`
		Cluster      *HyperflexHxapClusterRelationship `json:"Cluster,omitempty"`
		// An array of relationships to hyperflexHxapHost resources.
		MemberHosts []HyperflexHxapHostRelationship `json:"MemberHosts,omitempty"`
		// An array of relationships to hyperflexHxapHostVswitch resources.
		MemberVswitches []HyperflexHxapHostVswitchRelationship `json:"MemberVswitches,omitempty"`
	}

	varHyperflexHxapDvswitchWithoutEmbeddedStruct := HyperflexHxapDvswitchWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHxapDvswitchWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxapDvswitch := _HyperflexHxapDvswitch{}
		varHyperflexHxapDvswitch.ClassId = varHyperflexHxapDvswitchWithoutEmbeddedStruct.ClassId
		varHyperflexHxapDvswitch.ObjectType = varHyperflexHxapDvswitchWithoutEmbeddedStruct.ObjectType
		varHyperflexHxapDvswitch.DvUplink = varHyperflexHxapDvswitchWithoutEmbeddedStruct.DvUplink
		varHyperflexHxapDvswitch.LastHostname = varHyperflexHxapDvswitchWithoutEmbeddedStruct.LastHostname
		varHyperflexHxapDvswitch.Cluster = varHyperflexHxapDvswitchWithoutEmbeddedStruct.Cluster
		varHyperflexHxapDvswitch.MemberHosts = varHyperflexHxapDvswitchWithoutEmbeddedStruct.MemberHosts
		varHyperflexHxapDvswitch.MemberVswitches = varHyperflexHxapDvswitchWithoutEmbeddedStruct.MemberVswitches
		*o = HyperflexHxapDvswitch(varHyperflexHxapDvswitch)
	} else {
		return err
	}

	varHyperflexHxapDvswitch := _HyperflexHxapDvswitch{}

	err = json.Unmarshal(bytes, &varHyperflexHxapDvswitch)
	if err == nil {
		o.VirtualizationBaseDvswitch = varHyperflexHxapDvswitch.VirtualizationBaseDvswitch
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DvUplink")
		delete(additionalProperties, "LastHostname")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "MemberHosts")
		delete(additionalProperties, "MemberVswitches")

		// remove fields from embedded structs
		reflectVirtualizationBaseDvswitch := reflect.ValueOf(o.VirtualizationBaseDvswitch)
		for i := 0; i < reflectVirtualizationBaseDvswitch.Type().NumField(); i++ {
			t := reflectVirtualizationBaseDvswitch.Type().Field(i)

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

type NullableHyperflexHxapDvswitch struct {
	value *HyperflexHxapDvswitch
	isSet bool
}

func (v NullableHyperflexHxapDvswitch) Get() *HyperflexHxapDvswitch {
	return v.value
}

func (v *NullableHyperflexHxapDvswitch) Set(val *HyperflexHxapDvswitch) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxapDvswitch) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxapDvswitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxapDvswitch(val *HyperflexHxapDvswitch) *NullableHyperflexHxapDvswitch {
	return &NullableHyperflexHxapDvswitch{value: val, isSet: true}
}

func (v NullableHyperflexHxapDvswitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxapDvswitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
