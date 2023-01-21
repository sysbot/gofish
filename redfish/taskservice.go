//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// OverWritePolicy is
type OverWritePolicy string

const (
	// ManualOverWritePolicy Completed tasks are not automatically overwritten.
	ManualOverWritePolicy OverWritePolicy = "Manual"
	// OldestOverWritePolicy Oldest completed tasks are overwritten.
	OldestOverWritePolicy OverWritePolicy = "Oldest"
)

// OemActions shall contain the available OEM-specific actions for this resource.
type OemActions struct {
}

// UnmarshalJSON unmarshals a OemActions object from the raw JSON.
func (oemactions *OemActions) UnmarshalJSON(b []byte) error {
	type temp OemActions
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*oemactions = OemActions(t.temp)

	// Extract the links to other entities for later

	return nil
}

// TaskService This resource contains a task service for a Redfish implementation.
type TaskService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// CompletedTaskOverWritePolicy shall contain the overwrite policy for completed tasks. This property shall
	// indicate if the task service overwrites completed task information.
	CompletedTaskOverWritePolicy string
	// DateTime shall contain the current date and time for the task service, with UTC offset.
	DateTime string
	// Description provides a description of this resource.
	Description string
	// LifeCycleEventOnTaskStateChange shall indicate whether a task state change sends an event. Services should send
	// an event containing a message defined in the Task Event Message Registry when the state of a task changes.
	LifeCycleEventOnTaskStateChange string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TaskAutoDeleteTimeoutMinutes shall contain the number of minutes after which a completed task, where TaskState
	// contains the value 'Completed', 'Killed', 'Cancelled', or 'Exception', is deleted by the service.
	TaskAutoDeleteTimeoutMinutes string
	// Tasks shall contain a link to a resource collection of type TaskCollection.
	Tasks string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a TaskService object from the raw JSON.
func (taskservice *TaskService) UnmarshalJSON(b []byte) error {
	type temp TaskService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*taskservice = TaskService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	taskservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (taskservice *TaskService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(TaskService)
	original.UnmarshalJSON(taskservice.rawData)

	readWriteFields := []string{
		"ServiceEnabled",
		"TaskAutoDeleteTimeoutMinutes",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(taskservice).Elem()

	return taskservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetTaskService will get a TaskService instance from the service.
func GetTaskService(c common.Client, uri string) (*TaskService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var taskservice TaskService
	err = json.NewDecoder(resp.Body).Decode(&taskservice)
	if err != nil {
		return nil, err
	}

	taskservice.SetClient(c)
	return &taskservice, nil
}

// ListReferencedTaskServices gets the collection of TaskService from
// a provided reference.
func ListReferencedTaskServices(c common.Client, link string) ([]*TaskService, error) {
	var result []*TaskService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, taskserviceLink := range links.ItemLinks {
		taskservice, err := GetTaskService(c, taskserviceLink)
		if err != nil {
			collectionError.Failures[taskserviceLink] = err
		} else {
			result = append(result, taskservice)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
