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

// EquipmentFanModule This represents Fan module housing multiple fans for chassis/server.
type EquipmentFanModule struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field is to provide description for the fan module.
	Description *string `json:"Description,omitempty"`
	// This field acts as the identifier for this particular Module, within the Fabric Interconnect.
	ModuleId *int64 `json:"ModuleId,omitempty"`
	// This field is used to indicate this fan module's operational state.
	OperState *string `json:"OperState,omitempty"`
	// This field identifies the Part Number for this Fan Module.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for the fan module.
	Pid *string `json:"Pid,omitempty"`
	// This field is used to indicate this fan module's presence.
	Presence *string `json:"Presence,omitempty"`
	// This field identifies the Stockkeeping Unit for this Fan Module.
	Sku *string `json:"Sku,omitempty"`
	// Tray identifier for the fan module.
	TrayId *int64 `json:"TrayId,omitempty"`
	// This field identifies the Vendor ID for this Fan Module.
	Vid                    *string                             `json:"Vid,omitempty"`
	ComputeRackUnit        *ComputeRackUnitRelationship        `json:"ComputeRackUnit,omitempty"`
	EquipmentChassis       *EquipmentChassisRelationship       `json:"EquipmentChassis,omitempty"`
	EquipmentRackEnclosure *EquipmentRackEnclosureRelationship `json:"EquipmentRackEnclosure,omitempty"`
	// An array of relationships to equipmentFan resources.
	Fans                 []EquipmentFanRelationship           `json:"Fans,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentFanModule EquipmentFanModule

// NewEquipmentFanModule instantiates a new EquipmentFanModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentFanModule(classId string, objectType string) *EquipmentFanModule {
	this := EquipmentFanModule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentFanModuleWithDefaults instantiates a new EquipmentFanModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentFanModuleWithDefaults() *EquipmentFanModule {
	this := EquipmentFanModule{}
	var classId string = "equipment.FanModule"
	this.ClassId = classId
	var objectType string = "equipment.FanModule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentFanModule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentFanModule) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentFanModule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentFanModule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentFanModule) SetDescription(v string) {
	o.Description = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentFanModule) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentFanModule) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentFanModule) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentFanModule) SetPid(v string) {
	o.Pid = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *EquipmentFanModule) SetPresence(v string) {
	o.Presence = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentFanModule) SetSku(v string) {
	o.Sku = &v
}

// GetTrayId returns the TrayId field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetTrayId() int64 {
	if o == nil || o.TrayId == nil {
		var ret int64
		return ret
	}
	return *o.TrayId
}

// GetTrayIdOk returns a tuple with the TrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetTrayIdOk() (*int64, bool) {
	if o == nil || o.TrayId == nil {
		return nil, false
	}
	return o.TrayId, true
}

// HasTrayId returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasTrayId() bool {
	if o != nil && o.TrayId != nil {
		return true
	}

	return false
}

// SetTrayId gets a reference to the given int64 and assigns it to the TrayId field.
func (o *EquipmentFanModule) SetTrayId(v int64) {
	o.TrayId = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentFanModule) SetVid(v string) {
	o.Vid = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *EquipmentFanModule) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentFanModule) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship {
	if o == nil || o.EquipmentRackEnclosure == nil {
		var ret EquipmentRackEnclosureRelationship
		return ret
	}
	return *o.EquipmentRackEnclosure
}

// GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool) {
	if o == nil || o.EquipmentRackEnclosure == nil {
		return nil, false
	}
	return o.EquipmentRackEnclosure, true
}

// HasEquipmentRackEnclosure returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasEquipmentRackEnclosure() bool {
	if o != nil && o.EquipmentRackEnclosure != nil {
		return true
	}

	return false
}

// SetEquipmentRackEnclosure gets a reference to the given EquipmentRackEnclosureRelationship and assigns it to the EquipmentRackEnclosure field.
func (o *EquipmentFanModule) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship) {
	o.EquipmentRackEnclosure = &v
}

// GetFans returns the Fans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentFanModule) GetFans() []EquipmentFanRelationship {
	if o == nil {
		var ret []EquipmentFanRelationship
		return ret
	}
	return o.Fans
}

// GetFansOk returns a tuple with the Fans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentFanModule) GetFansOk() (*[]EquipmentFanRelationship, bool) {
	if o == nil || o.Fans == nil {
		return nil, false
	}
	return &o.Fans, true
}

// HasFans returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasFans() bool {
	if o != nil && o.Fans != nil {
		return true
	}

	return false
}

// SetFans gets a reference to the given []EquipmentFanRelationship and assigns it to the Fans field.
func (o *EquipmentFanModule) SetFans(v []EquipmentFanRelationship) {
	o.Fans = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentFanModule) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentFanModule) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentFanModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFanModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentFanModule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentFanModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentFanModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ModuleId != nil {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.TrayId != nil {
		toSerialize["TrayId"] = o.TrayId
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.EquipmentRackEnclosure != nil {
		toSerialize["EquipmentRackEnclosure"] = o.EquipmentRackEnclosure
	}
	if o.Fans != nil {
		toSerialize["Fans"] = o.Fans
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentFanModule) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentFanModuleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This field is to provide description for the fan module.
		Description *string `json:"Description,omitempty"`
		// This field acts as the identifier for this particular Module, within the Fabric Interconnect.
		ModuleId *int64 `json:"ModuleId,omitempty"`
		// This field is used to indicate this fan module's operational state.
		OperState *string `json:"OperState,omitempty"`
		// This field identifies the Part Number for this Fan Module.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the Product ID for the fan module.
		Pid *string `json:"Pid,omitempty"`
		// This field is used to indicate this fan module's presence.
		Presence *string `json:"Presence,omitempty"`
		// This field identifies the Stockkeeping Unit for this Fan Module.
		Sku *string `json:"Sku,omitempty"`
		// Tray identifier for the fan module.
		TrayId *int64 `json:"TrayId,omitempty"`
		// This field identifies the Vendor ID for this Fan Module.
		Vid                    *string                             `json:"Vid,omitempty"`
		ComputeRackUnit        *ComputeRackUnitRelationship        `json:"ComputeRackUnit,omitempty"`
		EquipmentChassis       *EquipmentChassisRelationship       `json:"EquipmentChassis,omitempty"`
		EquipmentRackEnclosure *EquipmentRackEnclosureRelationship `json:"EquipmentRackEnclosure,omitempty"`
		// An array of relationships to equipmentFan resources.
		Fans                []EquipmentFanRelationship           `json:"Fans,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement      *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentFanModuleWithoutEmbeddedStruct := EquipmentFanModuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentFanModuleWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentFanModule := _EquipmentFanModule{}
		varEquipmentFanModule.ClassId = varEquipmentFanModuleWithoutEmbeddedStruct.ClassId
		varEquipmentFanModule.ObjectType = varEquipmentFanModuleWithoutEmbeddedStruct.ObjectType
		varEquipmentFanModule.Description = varEquipmentFanModuleWithoutEmbeddedStruct.Description
		varEquipmentFanModule.ModuleId = varEquipmentFanModuleWithoutEmbeddedStruct.ModuleId
		varEquipmentFanModule.OperState = varEquipmentFanModuleWithoutEmbeddedStruct.OperState
		varEquipmentFanModule.PartNumber = varEquipmentFanModuleWithoutEmbeddedStruct.PartNumber
		varEquipmentFanModule.Pid = varEquipmentFanModuleWithoutEmbeddedStruct.Pid
		varEquipmentFanModule.Presence = varEquipmentFanModuleWithoutEmbeddedStruct.Presence
		varEquipmentFanModule.Sku = varEquipmentFanModuleWithoutEmbeddedStruct.Sku
		varEquipmentFanModule.TrayId = varEquipmentFanModuleWithoutEmbeddedStruct.TrayId
		varEquipmentFanModule.Vid = varEquipmentFanModuleWithoutEmbeddedStruct.Vid
		varEquipmentFanModule.ComputeRackUnit = varEquipmentFanModuleWithoutEmbeddedStruct.ComputeRackUnit
		varEquipmentFanModule.EquipmentChassis = varEquipmentFanModuleWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentFanModule.EquipmentRackEnclosure = varEquipmentFanModuleWithoutEmbeddedStruct.EquipmentRackEnclosure
		varEquipmentFanModule.Fans = varEquipmentFanModuleWithoutEmbeddedStruct.Fans
		varEquipmentFanModule.InventoryDeviceInfo = varEquipmentFanModuleWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentFanModule.NetworkElement = varEquipmentFanModuleWithoutEmbeddedStruct.NetworkElement
		varEquipmentFanModule.RegisteredDevice = varEquipmentFanModuleWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentFanModule(varEquipmentFanModule)
	} else {
		return err
	}

	varEquipmentFanModule := _EquipmentFanModule{}

	err = json.Unmarshal(bytes, &varEquipmentFanModule)
	if err == nil {
		o.EquipmentBase = varEquipmentFanModule.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "TrayId")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "EquipmentRackEnclosure")
		delete(additionalProperties, "Fans")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableEquipmentFanModule struct {
	value *EquipmentFanModule
	isSet bool
}

func (v NullableEquipmentFanModule) Get() *EquipmentFanModule {
	return v.value
}

func (v *NullableEquipmentFanModule) Set(val *EquipmentFanModule) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFanModule) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFanModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFanModule(val *EquipmentFanModule) *NullableEquipmentFanModule {
	return &NullableEquipmentFanModule{value: val, isSet: true}
}

func (v NullableEquipmentFanModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFanModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}