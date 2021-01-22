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

// WorkflowLoopTask A LoopTask is a control task that runs one or more task multiple times based on a counter. The tasks can be run in serial or parallel.
type WorkflowLoopTask struct {
	WorkflowControlTask
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Count value for the loop, this can be a constant or an expression which will evaluate to an integer value. Example, Use the length of the input A which is an array. Count must be less than or equal to 500.
	Count *string `json:"Count,omitempty"`
	// Start task where the list of tasks will be executed multiple times based on the count value.
	LoopStartTask *string `json:"LoopStartTask,omitempty"`
	// When tasks are run in parallel and the count is large, the actual number of task run in parallel can be controlled by this property. If count is 100 and numberOfBatches is 5 then 20 tasks are run in parallel 5 times. Parallel batch size must be less than the count. In cases where count is dynamic and depends on input given during workflow execution, if that count is less than batch then empty batches might get created which do not have any tasks under them.
	NumberOfBatches *int64 `json:"NumberOfBatches,omitempty"`
	// This specifies the name of the next task to run if all iterations of the loop task succeeds. The unique name given to the task instance within the workflow must be provided here. In a graph model, denotes an edge to another Task Node.
	OnSuccess *string `json:"OnSuccess,omitempty"`
	// When set to true the loop will run in parallel else it will run in a serial fashion. Only one task is supported inside the loop task when the loop is run in parallel. Subworkflow can be used inside the single loop task to build complex conditions.
	Parallel             *bool `json:"Parallel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowLoopTask WorkflowLoopTask

// NewWorkflowLoopTask instantiates a new WorkflowLoopTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowLoopTask(classId string, objectType string) *WorkflowLoopTask {
	this := WorkflowLoopTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	var numberOfBatches int64 = 1
	this.NumberOfBatches = &numberOfBatches
	var parallel bool = true
	this.Parallel = &parallel
	return &this
}

// NewWorkflowLoopTaskWithDefaults instantiates a new WorkflowLoopTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowLoopTaskWithDefaults() *WorkflowLoopTask {
	this := WorkflowLoopTask{}
	var classId string = "workflow.LoopTask"
	this.ClassId = classId
	var objectType string = "workflow.LoopTask"
	this.ObjectType = objectType
	var numberOfBatches int64 = 1
	this.NumberOfBatches = &numberOfBatches
	var parallel bool = true
	this.Parallel = &parallel
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowLoopTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowLoopTask) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowLoopTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowLoopTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WorkflowLoopTask) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTask) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WorkflowLoopTask) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *WorkflowLoopTask) SetCount(v string) {
	o.Count = &v
}

// GetLoopStartTask returns the LoopStartTask field value if set, zero value otherwise.
func (o *WorkflowLoopTask) GetLoopStartTask() string {
	if o == nil || o.LoopStartTask == nil {
		var ret string
		return ret
	}
	return *o.LoopStartTask
}

// GetLoopStartTaskOk returns a tuple with the LoopStartTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTask) GetLoopStartTaskOk() (*string, bool) {
	if o == nil || o.LoopStartTask == nil {
		return nil, false
	}
	return o.LoopStartTask, true
}

// HasLoopStartTask returns a boolean if a field has been set.
func (o *WorkflowLoopTask) HasLoopStartTask() bool {
	if o != nil && o.LoopStartTask != nil {
		return true
	}

	return false
}

// SetLoopStartTask gets a reference to the given string and assigns it to the LoopStartTask field.
func (o *WorkflowLoopTask) SetLoopStartTask(v string) {
	o.LoopStartTask = &v
}

// GetNumberOfBatches returns the NumberOfBatches field value if set, zero value otherwise.
func (o *WorkflowLoopTask) GetNumberOfBatches() int64 {
	if o == nil || o.NumberOfBatches == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfBatches
}

// GetNumberOfBatchesOk returns a tuple with the NumberOfBatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTask) GetNumberOfBatchesOk() (*int64, bool) {
	if o == nil || o.NumberOfBatches == nil {
		return nil, false
	}
	return o.NumberOfBatches, true
}

// HasNumberOfBatches returns a boolean if a field has been set.
func (o *WorkflowLoopTask) HasNumberOfBatches() bool {
	if o != nil && o.NumberOfBatches != nil {
		return true
	}

	return false
}

// SetNumberOfBatches gets a reference to the given int64 and assigns it to the NumberOfBatches field.
func (o *WorkflowLoopTask) SetNumberOfBatches(v int64) {
	o.NumberOfBatches = &v
}

// GetOnSuccess returns the OnSuccess field value if set, zero value otherwise.
func (o *WorkflowLoopTask) GetOnSuccess() string {
	if o == nil || o.OnSuccess == nil {
		var ret string
		return ret
	}
	return *o.OnSuccess
}

// GetOnSuccessOk returns a tuple with the OnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTask) GetOnSuccessOk() (*string, bool) {
	if o == nil || o.OnSuccess == nil {
		return nil, false
	}
	return o.OnSuccess, true
}

// HasOnSuccess returns a boolean if a field has been set.
func (o *WorkflowLoopTask) HasOnSuccess() bool {
	if o != nil && o.OnSuccess != nil {
		return true
	}

	return false
}

// SetOnSuccess gets a reference to the given string and assigns it to the OnSuccess field.
func (o *WorkflowLoopTask) SetOnSuccess(v string) {
	o.OnSuccess = &v
}

// GetParallel returns the Parallel field value if set, zero value otherwise.
func (o *WorkflowLoopTask) GetParallel() bool {
	if o == nil || o.Parallel == nil {
		var ret bool
		return ret
	}
	return *o.Parallel
}

// GetParallelOk returns a tuple with the Parallel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTask) GetParallelOk() (*bool, bool) {
	if o == nil || o.Parallel == nil {
		return nil, false
	}
	return o.Parallel, true
}

// HasParallel returns a boolean if a field has been set.
func (o *WorkflowLoopTask) HasParallel() bool {
	if o != nil && o.Parallel != nil {
		return true
	}

	return false
}

// SetParallel gets a reference to the given bool and assigns it to the Parallel field.
func (o *WorkflowLoopTask) SetParallel(v bool) {
	o.Parallel = &v
}

func (o WorkflowLoopTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowControlTask, errWorkflowControlTask := json.Marshal(o.WorkflowControlTask)
	if errWorkflowControlTask != nil {
		return []byte{}, errWorkflowControlTask
	}
	errWorkflowControlTask = json.Unmarshal([]byte(serializedWorkflowControlTask), &toSerialize)
	if errWorkflowControlTask != nil {
		return []byte{}, errWorkflowControlTask
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.LoopStartTask != nil {
		toSerialize["LoopStartTask"] = o.LoopStartTask
	}
	if o.NumberOfBatches != nil {
		toSerialize["NumberOfBatches"] = o.NumberOfBatches
	}
	if o.OnSuccess != nil {
		toSerialize["OnSuccess"] = o.OnSuccess
	}
	if o.Parallel != nil {
		toSerialize["Parallel"] = o.Parallel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowLoopTask) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowLoopTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Count value for the loop, this can be a constant or an expression which will evaluate to an integer value. Example, Use the length of the input A which is an array. Count must be less than or equal to 500.
		Count *string `json:"Count,omitempty"`
		// Start task where the list of tasks will be executed multiple times based on the count value.
		LoopStartTask *string `json:"LoopStartTask,omitempty"`
		// When tasks are run in parallel and the count is large, the actual number of task run in parallel can be controlled by this property. If count is 100 and numberOfBatches is 5 then 20 tasks are run in parallel 5 times. Parallel batch size must be less than the count. In cases where count is dynamic and depends on input given during workflow execution, if that count is less than batch then empty batches might get created which do not have any tasks under them.
		NumberOfBatches *int64 `json:"NumberOfBatches,omitempty"`
		// This specifies the name of the next task to run if all iterations of the loop task succeeds. The unique name given to the task instance within the workflow must be provided here. In a graph model, denotes an edge to another Task Node.
		OnSuccess *string `json:"OnSuccess,omitempty"`
		// When set to true the loop will run in parallel else it will run in a serial fashion. Only one task is supported inside the loop task when the loop is run in parallel. Subworkflow can be used inside the single loop task to build complex conditions.
		Parallel *bool `json:"Parallel,omitempty"`
	}

	varWorkflowLoopTaskWithoutEmbeddedStruct := WorkflowLoopTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowLoopTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowLoopTask := _WorkflowLoopTask{}
		varWorkflowLoopTask.ClassId = varWorkflowLoopTaskWithoutEmbeddedStruct.ClassId
		varWorkflowLoopTask.ObjectType = varWorkflowLoopTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowLoopTask.Count = varWorkflowLoopTaskWithoutEmbeddedStruct.Count
		varWorkflowLoopTask.LoopStartTask = varWorkflowLoopTaskWithoutEmbeddedStruct.LoopStartTask
		varWorkflowLoopTask.NumberOfBatches = varWorkflowLoopTaskWithoutEmbeddedStruct.NumberOfBatches
		varWorkflowLoopTask.OnSuccess = varWorkflowLoopTaskWithoutEmbeddedStruct.OnSuccess
		varWorkflowLoopTask.Parallel = varWorkflowLoopTaskWithoutEmbeddedStruct.Parallel
		*o = WorkflowLoopTask(varWorkflowLoopTask)
	} else {
		return err
	}

	varWorkflowLoopTask := _WorkflowLoopTask{}

	err = json.Unmarshal(bytes, &varWorkflowLoopTask)
	if err == nil {
		o.WorkflowControlTask = varWorkflowLoopTask.WorkflowControlTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "LoopStartTask")
		delete(additionalProperties, "NumberOfBatches")
		delete(additionalProperties, "OnSuccess")
		delete(additionalProperties, "Parallel")

		// remove fields from embedded structs
		reflectWorkflowControlTask := reflect.ValueOf(o.WorkflowControlTask)
		for i := 0; i < reflectWorkflowControlTask.Type().NumField(); i++ {
			t := reflectWorkflowControlTask.Type().Field(i)

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

type NullableWorkflowLoopTask struct {
	value *WorkflowLoopTask
	isSet bool
}

func (v NullableWorkflowLoopTask) Get() *WorkflowLoopTask {
	return v.value
}

func (v *NullableWorkflowLoopTask) Set(val *WorkflowLoopTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLoopTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLoopTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLoopTask(val *WorkflowLoopTask) *NullableWorkflowLoopTask {
	return &NullableWorkflowLoopTask{value: val, isSet: true}
}

func (v NullableWorkflowLoopTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLoopTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
