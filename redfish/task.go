//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// TaskState is
type TaskState string

const (
	// NewTaskState shall represent that the task is newly created, but has not started.
	NewTaskState TaskState = "New"
	// StartingTaskState shall represent that the task is starting.
	StartingTaskState TaskState = "Starting"
	// RunningTaskState shall represent that the task is executing.
	RunningTaskState TaskState = "Running"
	// SuspendedTaskState shall represent that the task has been suspended but is expected to restart and is therefore
	// not complete.
	SuspendedTaskState TaskState = "Suspended"
	// InterruptedTaskState shall represent that the task has been interrupted but is expected to restart and is
	// therefore not complete.
	InterruptedTaskState TaskState = "Interrupted"
	// PendingTaskState shall represent that the task is pending some condition and has not yet begun to execute.
	PendingTaskState TaskState = "Pending"
	// StoppingTaskState shall represent that the task is stopping but is not yet complete.
	StoppingTaskState TaskState = "Stopping"
	// CompletedTaskState shall represent that the task completed successfully or with warnings.
	CompletedTaskState TaskState = "Completed"
	// KilledTaskState shall represent that the task is complete because an operator killed it.
	KilledTaskState TaskState = "Killed"
	// ExceptionTaskState shall represent that the task completed with errors.
	ExceptionTaskState TaskState = "Exception"
	// ServiceTaskState shall represent that the task is now running as a service and expected to continue operation
	// until stopped or killed.
	ServiceTaskState TaskState = "Service"
	// CancellingTaskState shall represent that the task is in the process of being cancelled.
	CancellingTaskState TaskState = "Cancelling"
	// CancelledTaskState shall represent that either a DELETE operation on a task monitor or Task resource or by an
	// internal process cancelled the task.
	CancelledTaskState TaskState = "Cancelled"
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

// Payload shall contain information detailing the HTTP and JSON payload information for executing this task.
type Payload struct {
	// HttpHeaders shall contain an array of HTTP headers that this task includes.
	HttpHeaders []string
	// HttpOperation shall contain the HTTP operation to execute for this task.
	HttpOperation string
	// JsonBody shall contain JSON formatted payload used for this task.
	JsonBody string
	// TargetUri shall contain a link to the location to use as the target for an HTTP operation.
	TargetUri string
}

// UnmarshalJSON unmarshals a Payload object from the raw JSON.
func (payload *Payload) UnmarshalJSON(b []byte) error {
	type temp Payload
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*payload = Payload(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Task This resource contains a task for a Redfish implementation.
type Task struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// EndTime shall indicate the date and time when the task was completed. This property shall not appear if the task
	// is running or otherwise has not been completed. This property shall appear only if the TaskState is Completed,
	// Killed, Cancelled, or Exception.
	EndTime string
	// EstimatedDuration shall indicate the estimated total time needed to complete the task. The value is not expected
	// to change while the task is in progress, but the service may update the value if it obtains new information that
	// significantly changes the expected duration. Services should be conservative in the reported estimate and
	// clients should treat this value as an estimate.
	EstimatedDuration string
	// HidePayload shall indicate whether the contents of the payload should be hidden from view after the task has
	// been created. If 'true', responses shall not return the Payload property. If 'false', responses shall return the
	// Payload property. If this property is not present when the task is created, the default is 'false'. This
	// property shall be supported if the Payload property is supported.
	HidePayload string
	// Messages shall contain an array of messages associated with the task.
	Messages []Message
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Payload shall contain information detailing the HTTP and JSON payload information for executing this task. This
	// property shall not be included in the response if the HidePayload property is 'true'.
	Payload string
	// PercentComplete shall indicate the completion progress of the task, reported in percent of completion. If the
	// task has not been started, the value shall be zero.
	PercentComplete int
	// StartTime shall indicate the date and time when the task was started.
	StartTime string
	// SubTasks shall contain a link to a resource collection of type TaskCollection. This property shall not be
	// present if this resource represents a sub-task for a task.
	SubTasks string
	// TaskMonitor shall contain a URI to task monitor as defined in the Redfish Specification.
	TaskMonitor string
	// TaskState shall indicate the state of the task.
	TaskState string
	// TaskStatus shall contain the completion status of the task and shall not be set until the task completes. This
	// property should contain 'Critical' if one or more messages in the Messages array contains the severity
	// 'Critical'. This property should contain 'Warning' if one or more messages in the Messages array contains the
	// severity 'Warning' and no messages contain the severity 'Critical'. This property should contain 'OK' if all
	// messages in the Messages array contain the severity 'OK' or the array is empty.
	TaskStatus string
}

// UnmarshalJSON unmarshals a Task object from the raw JSON.
func (task *Task) UnmarshalJSON(b []byte) error {
	type temp Task
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*task = Task(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetTask will get a Task instance from the service.
func GetTask(c common.Client, uri string) (*Task, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var task Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		return nil, err
	}

	task.SetClient(c)
	return &task, nil
}

// ListReferencedTasks gets the collection of Task from
// a provided reference.
func ListReferencedTasks(c common.Client, link string) ([]*Task, error) {
	var result []*Task
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, taskLink := range links.ItemLinks {
		task, err := GetTask(c, taskLink)
		if err != nil {
			collectionError.Failures[taskLink] = err
		} else {
			result = append(result, task)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}
