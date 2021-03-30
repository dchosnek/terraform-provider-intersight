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

// ManagementController A specialized service processor that monitors the physical state of a server, using sensors and communicating with the system administrator through an independent connection.
type ManagementController struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Model of the endpoint that houses the management controller.
	Model                       *string                                  `json:"Model,omitempty"`
	AdapterUnit                 *AdapterUnitRelationship                 `json:"AdapterUnit,omitempty"`
	ComputeBlade                *ComputeBladeRelationship                `json:"ComputeBlade,omitempty"`
	ComputeRackUnit             *ComputeRackUnitRelationship             `json:"ComputeRackUnit,omitempty"`
	EquipmentIoCardBase         *EquipmentIoCardBaseRelationship         `json:"EquipmentIoCardBase,omitempty"`
	EquipmentSharedIoModule     *EquipmentSharedIoModuleRelationship     `json:"EquipmentSharedIoModule,omitempty"`
	EquipmentSystemIoController *EquipmentSystemIoControllerRelationship `json:"EquipmentSystemIoController,omitempty"`
	InventoryDeviceInfo         *InventoryDeviceInfoRelationship         `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to managementInterface resources.
	ManagementInterfaces []ManagementInterfaceRelationship    `json:"ManagementInterfaces,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware      []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	StorageSasExpander   *StorageSasExpanderRelationship       `json:"StorageSasExpander,omitempty"`
	TopSystem            *TopSystemRelationship                `json:"TopSystem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagementController ManagementController

// NewManagementController instantiates a new ManagementController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementController(classId string, objectType string) *ManagementController {
	this := ManagementController{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewManagementControllerWithDefaults instantiates a new ManagementController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementControllerWithDefaults() *ManagementController {
	this := ManagementController{}
	var classId string = "management.Controller"
	this.ClassId = classId
	var objectType string = "management.Controller"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ManagementController) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ManagementController) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ManagementController) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ManagementController) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ManagementController) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ManagementController) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ManagementController) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ManagementController) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ManagementController) SetModel(v string) {
	o.Model = &v
}

// GetAdapterUnit returns the AdapterUnit field value if set, zero value otherwise.
func (o *ManagementController) GetAdapterUnit() AdapterUnitRelationship {
	if o == nil || o.AdapterUnit == nil {
		var ret AdapterUnitRelationship
		return ret
	}
	return *o.AdapterUnit
}

// GetAdapterUnitOk returns a tuple with the AdapterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetAdapterUnitOk() (*AdapterUnitRelationship, bool) {
	if o == nil || o.AdapterUnit == nil {
		return nil, false
	}
	return o.AdapterUnit, true
}

// HasAdapterUnit returns a boolean if a field has been set.
func (o *ManagementController) HasAdapterUnit() bool {
	if o != nil && o.AdapterUnit != nil {
		return true
	}

	return false
}

// SetAdapterUnit gets a reference to the given AdapterUnitRelationship and assigns it to the AdapterUnit field.
func (o *ManagementController) SetAdapterUnit(v AdapterUnitRelationship) {
	o.AdapterUnit = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *ManagementController) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *ManagementController) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *ManagementController) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *ManagementController) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *ManagementController) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *ManagementController) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetEquipmentIoCardBase returns the EquipmentIoCardBase field value if set, zero value otherwise.
func (o *ManagementController) GetEquipmentIoCardBase() EquipmentIoCardBaseRelationship {
	if o == nil || o.EquipmentIoCardBase == nil {
		var ret EquipmentIoCardBaseRelationship
		return ret
	}
	return *o.EquipmentIoCardBase
}

// GetEquipmentIoCardBaseOk returns a tuple with the EquipmentIoCardBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetEquipmentIoCardBaseOk() (*EquipmentIoCardBaseRelationship, bool) {
	if o == nil || o.EquipmentIoCardBase == nil {
		return nil, false
	}
	return o.EquipmentIoCardBase, true
}

// HasEquipmentIoCardBase returns a boolean if a field has been set.
func (o *ManagementController) HasEquipmentIoCardBase() bool {
	if o != nil && o.EquipmentIoCardBase != nil {
		return true
	}

	return false
}

// SetEquipmentIoCardBase gets a reference to the given EquipmentIoCardBaseRelationship and assigns it to the EquipmentIoCardBase field.
func (o *ManagementController) SetEquipmentIoCardBase(v EquipmentIoCardBaseRelationship) {
	o.EquipmentIoCardBase = &v
}

// GetEquipmentSharedIoModule returns the EquipmentSharedIoModule field value if set, zero value otherwise.
func (o *ManagementController) GetEquipmentSharedIoModule() EquipmentSharedIoModuleRelationship {
	if o == nil || o.EquipmentSharedIoModule == nil {
		var ret EquipmentSharedIoModuleRelationship
		return ret
	}
	return *o.EquipmentSharedIoModule
}

// GetEquipmentSharedIoModuleOk returns a tuple with the EquipmentSharedIoModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetEquipmentSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool) {
	if o == nil || o.EquipmentSharedIoModule == nil {
		return nil, false
	}
	return o.EquipmentSharedIoModule, true
}

// HasEquipmentSharedIoModule returns a boolean if a field has been set.
func (o *ManagementController) HasEquipmentSharedIoModule() bool {
	if o != nil && o.EquipmentSharedIoModule != nil {
		return true
	}

	return false
}

// SetEquipmentSharedIoModule gets a reference to the given EquipmentSharedIoModuleRelationship and assigns it to the EquipmentSharedIoModule field.
func (o *ManagementController) SetEquipmentSharedIoModule(v EquipmentSharedIoModuleRelationship) {
	o.EquipmentSharedIoModule = &v
}

// GetEquipmentSystemIoController returns the EquipmentSystemIoController field value if set, zero value otherwise.
func (o *ManagementController) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship {
	if o == nil || o.EquipmentSystemIoController == nil {
		var ret EquipmentSystemIoControllerRelationship
		return ret
	}
	return *o.EquipmentSystemIoController
}

// GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool) {
	if o == nil || o.EquipmentSystemIoController == nil {
		return nil, false
	}
	return o.EquipmentSystemIoController, true
}

// HasEquipmentSystemIoController returns a boolean if a field has been set.
func (o *ManagementController) HasEquipmentSystemIoController() bool {
	if o != nil && o.EquipmentSystemIoController != nil {
		return true
	}

	return false
}

// SetEquipmentSystemIoController gets a reference to the given EquipmentSystemIoControllerRelationship and assigns it to the EquipmentSystemIoController field.
func (o *ManagementController) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship) {
	o.EquipmentSystemIoController = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ManagementController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ManagementController) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ManagementController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetManagementInterfaces returns the ManagementInterfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagementController) GetManagementInterfaces() []ManagementInterfaceRelationship {
	if o == nil {
		var ret []ManagementInterfaceRelationship
		return ret
	}
	return o.ManagementInterfaces
}

// GetManagementInterfacesOk returns a tuple with the ManagementInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagementController) GetManagementInterfacesOk() (*[]ManagementInterfaceRelationship, bool) {
	if o == nil || o.ManagementInterfaces == nil {
		return nil, false
	}
	return &o.ManagementInterfaces, true
}

// HasManagementInterfaces returns a boolean if a field has been set.
func (o *ManagementController) HasManagementInterfaces() bool {
	if o != nil && o.ManagementInterfaces != nil {
		return true
	}

	return false
}

// SetManagementInterfaces gets a reference to the given []ManagementInterfaceRelationship and assigns it to the ManagementInterfaces field.
func (o *ManagementController) SetManagementInterfaces(v []ManagementInterfaceRelationship) {
	o.ManagementInterfaces = v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *ManagementController) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *ManagementController) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *ManagementController) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ManagementController) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ManagementController) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ManagementController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagementController) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagementController) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.RunningFirmware == nil {
		return nil, false
	}
	return &o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *ManagementController) HasRunningFirmware() bool {
	if o != nil && o.RunningFirmware != nil {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *ManagementController) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

// GetStorageSasExpander returns the StorageSasExpander field value if set, zero value otherwise.
func (o *ManagementController) GetStorageSasExpander() StorageSasExpanderRelationship {
	if o == nil || o.StorageSasExpander == nil {
		var ret StorageSasExpanderRelationship
		return ret
	}
	return *o.StorageSasExpander
}

// GetStorageSasExpanderOk returns a tuple with the StorageSasExpander field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetStorageSasExpanderOk() (*StorageSasExpanderRelationship, bool) {
	if o == nil || o.StorageSasExpander == nil {
		return nil, false
	}
	return o.StorageSasExpander, true
}

// HasStorageSasExpander returns a boolean if a field has been set.
func (o *ManagementController) HasStorageSasExpander() bool {
	if o != nil && o.StorageSasExpander != nil {
		return true
	}

	return false
}

// SetStorageSasExpander gets a reference to the given StorageSasExpanderRelationship and assigns it to the StorageSasExpander field.
func (o *ManagementController) SetStorageSasExpander(v StorageSasExpanderRelationship) {
	o.StorageSasExpander = &v
}

// GetTopSystem returns the TopSystem field value if set, zero value otherwise.
func (o *ManagementController) GetTopSystem() TopSystemRelationship {
	if o == nil || o.TopSystem == nil {
		var ret TopSystemRelationship
		return ret
	}
	return *o.TopSystem
}

// GetTopSystemOk returns a tuple with the TopSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementController) GetTopSystemOk() (*TopSystemRelationship, bool) {
	if o == nil || o.TopSystem == nil {
		return nil, false
	}
	return o.TopSystem, true
}

// HasTopSystem returns a boolean if a field has been set.
func (o *ManagementController) HasTopSystem() bool {
	if o != nil && o.TopSystem != nil {
		return true
	}

	return false
}

// SetTopSystem gets a reference to the given TopSystemRelationship and assigns it to the TopSystem field.
func (o *ManagementController) SetTopSystem(v TopSystemRelationship) {
	o.TopSystem = &v
}

func (o ManagementController) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.AdapterUnit != nil {
		toSerialize["AdapterUnit"] = o.AdapterUnit
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.EquipmentIoCardBase != nil {
		toSerialize["EquipmentIoCardBase"] = o.EquipmentIoCardBase
	}
	if o.EquipmentSharedIoModule != nil {
		toSerialize["EquipmentSharedIoModule"] = o.EquipmentSharedIoModule
	}
	if o.EquipmentSystemIoController != nil {
		toSerialize["EquipmentSystemIoController"] = o.EquipmentSystemIoController
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.ManagementInterfaces != nil {
		toSerialize["ManagementInterfaces"] = o.ManagementInterfaces
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}
	if o.StorageSasExpander != nil {
		toSerialize["StorageSasExpander"] = o.StorageSasExpander
	}
	if o.TopSystem != nil {
		toSerialize["TopSystem"] = o.TopSystem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ManagementController) UnmarshalJSON(bytes []byte) (err error) {
	type ManagementControllerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Model of the endpoint that houses the management controller.
		Model                       *string                                  `json:"Model,omitempty"`
		AdapterUnit                 *AdapterUnitRelationship                 `json:"AdapterUnit,omitempty"`
		ComputeBlade                *ComputeBladeRelationship                `json:"ComputeBlade,omitempty"`
		ComputeRackUnit             *ComputeRackUnitRelationship             `json:"ComputeRackUnit,omitempty"`
		EquipmentIoCardBase         *EquipmentIoCardBaseRelationship         `json:"EquipmentIoCardBase,omitempty"`
		EquipmentSharedIoModule     *EquipmentSharedIoModuleRelationship     `json:"EquipmentSharedIoModule,omitempty"`
		EquipmentSystemIoController *EquipmentSystemIoControllerRelationship `json:"EquipmentSystemIoController,omitempty"`
		InventoryDeviceInfo         *InventoryDeviceInfoRelationship         `json:"InventoryDeviceInfo,omitempty"`
		// An array of relationships to managementInterface resources.
		ManagementInterfaces []ManagementInterfaceRelationship    `json:"ManagementInterfaces,omitempty"`
		NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to firmwareRunningFirmware resources.
		RunningFirmware    []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
		StorageSasExpander *StorageSasExpanderRelationship       `json:"StorageSasExpander,omitempty"`
		TopSystem          *TopSystemRelationship                `json:"TopSystem,omitempty"`
	}

	varManagementControllerWithoutEmbeddedStruct := ManagementControllerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varManagementControllerWithoutEmbeddedStruct)
	if err == nil {
		varManagementController := _ManagementController{}
		varManagementController.ClassId = varManagementControllerWithoutEmbeddedStruct.ClassId
		varManagementController.ObjectType = varManagementControllerWithoutEmbeddedStruct.ObjectType
		varManagementController.Model = varManagementControllerWithoutEmbeddedStruct.Model
		varManagementController.AdapterUnit = varManagementControllerWithoutEmbeddedStruct.AdapterUnit
		varManagementController.ComputeBlade = varManagementControllerWithoutEmbeddedStruct.ComputeBlade
		varManagementController.ComputeRackUnit = varManagementControllerWithoutEmbeddedStruct.ComputeRackUnit
		varManagementController.EquipmentIoCardBase = varManagementControllerWithoutEmbeddedStruct.EquipmentIoCardBase
		varManagementController.EquipmentSharedIoModule = varManagementControllerWithoutEmbeddedStruct.EquipmentSharedIoModule
		varManagementController.EquipmentSystemIoController = varManagementControllerWithoutEmbeddedStruct.EquipmentSystemIoController
		varManagementController.InventoryDeviceInfo = varManagementControllerWithoutEmbeddedStruct.InventoryDeviceInfo
		varManagementController.ManagementInterfaces = varManagementControllerWithoutEmbeddedStruct.ManagementInterfaces
		varManagementController.NetworkElement = varManagementControllerWithoutEmbeddedStruct.NetworkElement
		varManagementController.RegisteredDevice = varManagementControllerWithoutEmbeddedStruct.RegisteredDevice
		varManagementController.RunningFirmware = varManagementControllerWithoutEmbeddedStruct.RunningFirmware
		varManagementController.StorageSasExpander = varManagementControllerWithoutEmbeddedStruct.StorageSasExpander
		varManagementController.TopSystem = varManagementControllerWithoutEmbeddedStruct.TopSystem
		*o = ManagementController(varManagementController)
	} else {
		return err
	}

	varManagementController := _ManagementController{}

	err = json.Unmarshal(bytes, &varManagementController)
	if err == nil {
		o.InventoryBase = varManagementController.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "AdapterUnit")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "EquipmentIoCardBase")
		delete(additionalProperties, "EquipmentSharedIoModule")
		delete(additionalProperties, "EquipmentSystemIoController")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "ManagementInterfaces")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningFirmware")
		delete(additionalProperties, "StorageSasExpander")
		delete(additionalProperties, "TopSystem")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableManagementController struct {
	value *ManagementController
	isSet bool
}

func (v NullableManagementController) Get() *ManagementController {
	return v.value
}

func (v *NullableManagementController) Set(val *ManagementController) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementController) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementController(val *ManagementController) *NullableManagementController {
	return &NullableManagementController{value: val, isSet: true}
}

func (v NullableManagementController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
