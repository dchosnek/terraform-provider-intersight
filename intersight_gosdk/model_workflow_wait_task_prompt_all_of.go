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

// WorkflowWaitTaskPromptAllOf Definition of the list of properties defined in 'workflow.WaitTaskPrompt', excluding properties defined in parent classes.
type WorkflowWaitTaskPromptAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description that give more details about what it means to pick this prompt option for the wait task.
	Description *string `json:"Description,omitempty"`
	// User friendly label for the prompt. This label will be shown to the user as one of available options for the wait task.
	Label *string `json:"Label,omitempty"`
	// Name for the wait prompt.
	Name *string `json:"Name,omitempty"`
	// Task status for the wait task when this prompt option is selected. * `Scheduled` - The enum represents the status when task is in scheduled state. * `InProgress` - The enum represents the status when task is in-progress state. * `NoOp` - The enum represents the status when task is a noop. * `Timeout` - The enum represents the status when task has timed out. * `Completed` - The enum represents the status when task has completed. * `Failed` - The enum represents the status when task has failed.
	TaskStatus           *string `json:"TaskStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWaitTaskPromptAllOf WorkflowWaitTaskPromptAllOf

// NewWorkflowWaitTaskPromptAllOf instantiates a new WorkflowWaitTaskPromptAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWaitTaskPromptAllOf(classId string, objectType string) *WorkflowWaitTaskPromptAllOf {
	this := WorkflowWaitTaskPromptAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var taskStatus string = "Scheduled"
	this.TaskStatus = &taskStatus
	return &this
}

// NewWorkflowWaitTaskPromptAllOfWithDefaults instantiates a new WorkflowWaitTaskPromptAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWaitTaskPromptAllOfWithDefaults() *WorkflowWaitTaskPromptAllOf {
	this := WorkflowWaitTaskPromptAllOf{}
	var classId string = "workflow.WaitTaskPrompt"
	this.ClassId = classId
	var objectType string = "workflow.WaitTaskPrompt"
	this.ObjectType = objectType
	var taskStatus string = "Scheduled"
	this.TaskStatus = &taskStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWaitTaskPromptAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWaitTaskPromptAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWaitTaskPromptAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWaitTaskPromptAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPromptAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPromptAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowWaitTaskPromptAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPromptAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPromptAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowWaitTaskPromptAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPromptAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPromptAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowWaitTaskPromptAllOf) SetName(v string) {
	o.Name = &v
}

// GetTaskStatus returns the TaskStatus field value if set, zero value otherwise.
func (o *WorkflowWaitTaskPromptAllOf) GetTaskStatus() string {
	if o == nil || o.TaskStatus == nil {
		var ret string
		return ret
	}
	return *o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskPromptAllOf) GetTaskStatusOk() (*string, bool) {
	if o == nil || o.TaskStatus == nil {
		return nil, false
	}
	return o.TaskStatus, true
}

// HasTaskStatus returns a boolean if a field has been set.
func (o *WorkflowWaitTaskPromptAllOf) HasTaskStatus() bool {
	if o != nil && o.TaskStatus != nil {
		return true
	}

	return false
}

// SetTaskStatus gets a reference to the given string and assigns it to the TaskStatus field.
func (o *WorkflowWaitTaskPromptAllOf) SetTaskStatus(v string) {
	o.TaskStatus = &v
}

func (o WorkflowWaitTaskPromptAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.TaskStatus != nil {
		toSerialize["TaskStatus"] = o.TaskStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWaitTaskPromptAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowWaitTaskPromptAllOf := _WorkflowWaitTaskPromptAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowWaitTaskPromptAllOf); err == nil {
		*o = WorkflowWaitTaskPromptAllOf(varWorkflowWaitTaskPromptAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "TaskStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowWaitTaskPromptAllOf struct {
	value *WorkflowWaitTaskPromptAllOf
	isSet bool
}

func (v NullableWorkflowWaitTaskPromptAllOf) Get() *WorkflowWaitTaskPromptAllOf {
	return v.value
}

func (v *NullableWorkflowWaitTaskPromptAllOf) Set(val *WorkflowWaitTaskPromptAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWaitTaskPromptAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWaitTaskPromptAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWaitTaskPromptAllOf(val *WorkflowWaitTaskPromptAllOf) *NullableWorkflowWaitTaskPromptAllOf {
	return &NullableWorkflowWaitTaskPromptAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWaitTaskPromptAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWaitTaskPromptAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
