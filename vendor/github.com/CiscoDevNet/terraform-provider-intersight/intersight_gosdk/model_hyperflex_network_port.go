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

// HyperflexNetworkPort Network port connected to a vswitch.
type HyperflexNetworkPort struct {
	VirtualizationBaseNetworkPort
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                     `json:"ObjectType"`
	BondState     NullableHyperflexBondState `json:"BondState,omitempty"`
	NetInterfaces []string                   `json:"NetInterfaces,omitempty"`
	// The type of the network port. * `unknown` - This port is of an unknown port type. * `hypervisor` - This port is connected to the hypervisor. * `vm` - This port is connected to a VM. * `uplink` - This port is an uplink port.
	PortType *string `json:"PortType,omitempty"`
	// The vlan id associated with this port.
	Vlans                *string `json:"Vlans,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexNetworkPort HyperflexNetworkPort

// NewHyperflexNetworkPort instantiates a new HyperflexNetworkPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNetworkPort(classId string, objectType string) *HyperflexNetworkPort {
	this := HyperflexNetworkPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	var portType string = "unknown"
	this.PortType = &portType
	return &this
}

// NewHyperflexNetworkPortWithDefaults instantiates a new HyperflexNetworkPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNetworkPortWithDefaults() *HyperflexNetworkPort {
	this := HyperflexNetworkPort{}
	var classId string = "hyperflex.NetworkPort"
	this.ClassId = classId
	var objectType string = "hyperflex.NetworkPort"
	this.ObjectType = objectType
	var portType string = "unknown"
	this.PortType = &portType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexNetworkPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexNetworkPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexNetworkPort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexNetworkPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexNetworkPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexNetworkPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBondState returns the BondState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNetworkPort) GetBondState() HyperflexBondState {
	if o == nil || o.BondState.Get() == nil {
		var ret HyperflexBondState
		return ret
	}
	return *o.BondState.Get()
}

// GetBondStateOk returns a tuple with the BondState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNetworkPort) GetBondStateOk() (*HyperflexBondState, bool) {
	if o == nil {
		return nil, false
	}
	return o.BondState.Get(), o.BondState.IsSet()
}

// HasBondState returns a boolean if a field has been set.
func (o *HyperflexNetworkPort) HasBondState() bool {
	if o != nil && o.BondState.IsSet() {
		return true
	}

	return false
}

// SetBondState gets a reference to the given NullableHyperflexBondState and assigns it to the BondState field.
func (o *HyperflexNetworkPort) SetBondState(v HyperflexBondState) {
	o.BondState.Set(&v)
}

// SetBondStateNil sets the value for BondState to be an explicit nil
func (o *HyperflexNetworkPort) SetBondStateNil() {
	o.BondState.Set(nil)
}

// UnsetBondState ensures that no value is present for BondState, not even an explicit nil
func (o *HyperflexNetworkPort) UnsetBondState() {
	o.BondState.Unset()
}

// GetNetInterfaces returns the NetInterfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexNetworkPort) GetNetInterfaces() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NetInterfaces
}

// GetNetInterfacesOk returns a tuple with the NetInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexNetworkPort) GetNetInterfacesOk() (*[]string, bool) {
	if o == nil || o.NetInterfaces == nil {
		return nil, false
	}
	return &o.NetInterfaces, true
}

// HasNetInterfaces returns a boolean if a field has been set.
func (o *HyperflexNetworkPort) HasNetInterfaces() bool {
	if o != nil && o.NetInterfaces != nil {
		return true
	}

	return false
}

// SetNetInterfaces gets a reference to the given []string and assigns it to the NetInterfaces field.
func (o *HyperflexNetworkPort) SetNetInterfaces(v []string) {
	o.NetInterfaces = v
}

// GetPortType returns the PortType field value if set, zero value otherwise.
func (o *HyperflexNetworkPort) GetPortType() string {
	if o == nil || o.PortType == nil {
		var ret string
		return ret
	}
	return *o.PortType
}

// GetPortTypeOk returns a tuple with the PortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNetworkPort) GetPortTypeOk() (*string, bool) {
	if o == nil || o.PortType == nil {
		return nil, false
	}
	return o.PortType, true
}

// HasPortType returns a boolean if a field has been set.
func (o *HyperflexNetworkPort) HasPortType() bool {
	if o != nil && o.PortType != nil {
		return true
	}

	return false
}

// SetPortType gets a reference to the given string and assigns it to the PortType field.
func (o *HyperflexNetworkPort) SetPortType(v string) {
	o.PortType = &v
}

// GetVlans returns the Vlans field value if set, zero value otherwise.
func (o *HyperflexNetworkPort) GetVlans() string {
	if o == nil || o.Vlans == nil {
		var ret string
		return ret
	}
	return *o.Vlans
}

// GetVlansOk returns a tuple with the Vlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNetworkPort) GetVlansOk() (*string, bool) {
	if o == nil || o.Vlans == nil {
		return nil, false
	}
	return o.Vlans, true
}

// HasVlans returns a boolean if a field has been set.
func (o *HyperflexNetworkPort) HasVlans() bool {
	if o != nil && o.Vlans != nil {
		return true
	}

	return false
}

// SetVlans gets a reference to the given string and assigns it to the Vlans field.
func (o *HyperflexNetworkPort) SetVlans(v string) {
	o.Vlans = &v
}

func (o HyperflexNetworkPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseNetworkPort, errVirtualizationBaseNetworkPort := json.Marshal(o.VirtualizationBaseNetworkPort)
	if errVirtualizationBaseNetworkPort != nil {
		return []byte{}, errVirtualizationBaseNetworkPort
	}
	errVirtualizationBaseNetworkPort = json.Unmarshal([]byte(serializedVirtualizationBaseNetworkPort), &toSerialize)
	if errVirtualizationBaseNetworkPort != nil {
		return []byte{}, errVirtualizationBaseNetworkPort
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BondState.IsSet() {
		toSerialize["BondState"] = o.BondState.Get()
	}
	if o.NetInterfaces != nil {
		toSerialize["NetInterfaces"] = o.NetInterfaces
	}
	if o.PortType != nil {
		toSerialize["PortType"] = o.PortType
	}
	if o.Vlans != nil {
		toSerialize["Vlans"] = o.Vlans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexNetworkPort) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexNetworkPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                     `json:"ObjectType"`
		BondState     NullableHyperflexBondState `json:"BondState,omitempty"`
		NetInterfaces []string                   `json:"NetInterfaces,omitempty"`
		// The type of the network port. * `unknown` - This port is of an unknown port type. * `hypervisor` - This port is connected to the hypervisor. * `vm` - This port is connected to a VM. * `uplink` - This port is an uplink port.
		PortType *string `json:"PortType,omitempty"`
		// The vlan id associated with this port.
		Vlans *string `json:"Vlans,omitempty"`
	}

	varHyperflexNetworkPortWithoutEmbeddedStruct := HyperflexNetworkPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexNetworkPortWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexNetworkPort := _HyperflexNetworkPort{}
		varHyperflexNetworkPort.ClassId = varHyperflexNetworkPortWithoutEmbeddedStruct.ClassId
		varHyperflexNetworkPort.ObjectType = varHyperflexNetworkPortWithoutEmbeddedStruct.ObjectType
		varHyperflexNetworkPort.BondState = varHyperflexNetworkPortWithoutEmbeddedStruct.BondState
		varHyperflexNetworkPort.NetInterfaces = varHyperflexNetworkPortWithoutEmbeddedStruct.NetInterfaces
		varHyperflexNetworkPort.PortType = varHyperflexNetworkPortWithoutEmbeddedStruct.PortType
		varHyperflexNetworkPort.Vlans = varHyperflexNetworkPortWithoutEmbeddedStruct.Vlans
		*o = HyperflexNetworkPort(varHyperflexNetworkPort)
	} else {
		return err
	}

	varHyperflexNetworkPort := _HyperflexNetworkPort{}

	err = json.Unmarshal(bytes, &varHyperflexNetworkPort)
	if err == nil {
		o.VirtualizationBaseNetworkPort = varHyperflexNetworkPort.VirtualizationBaseNetworkPort
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BondState")
		delete(additionalProperties, "NetInterfaces")
		delete(additionalProperties, "PortType")
		delete(additionalProperties, "Vlans")

		// remove fields from embedded structs
		reflectVirtualizationBaseNetworkPort := reflect.ValueOf(o.VirtualizationBaseNetworkPort)
		for i := 0; i < reflectVirtualizationBaseNetworkPort.Type().NumField(); i++ {
			t := reflectVirtualizationBaseNetworkPort.Type().Field(i)

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

type NullableHyperflexNetworkPort struct {
	value *HyperflexNetworkPort
	isSet bool
}

func (v NullableHyperflexNetworkPort) Get() *HyperflexNetworkPort {
	return v.value
}

func (v *NullableHyperflexNetworkPort) Set(val *HyperflexNetworkPort) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNetworkPort) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNetworkPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNetworkPort(val *HyperflexNetworkPort) *NullableHyperflexNetworkPort {
	return &NullableHyperflexNetworkPort{value: val, isSet: true}
}

func (v NullableHyperflexNetworkPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNetworkPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
