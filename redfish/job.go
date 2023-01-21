//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// JobState is
type JobState string

const (
	// NewJobState shall represent that this job is newly created but the operation has not yet started.
	NewJobState JobState = "New"
	// StartingJobState shall represent that the operation is starting.
	StartingJobState JobState = "Starting"
	// RunningJobState shall represent that the operation is executing.
	RunningJobState JobState = "Running"
	// SuspendedJobState shall represent that the operation has been suspended but is expected to restart and is
	// therefore not complete.
	SuspendedJobState JobState = "Suspended"
	// InterruptedJobState shall represent that the operation has been interrupted but is expected to restart and is
	// therefore not complete.
	InterruptedJobState JobState = "Interrupted"
	// PendingJobState shall represent that the operation is pending some condition and has not yet begun to execute.
	PendingJobState JobState = "Pending"
	// StoppingJobState shall represent that the operation is stopping but is not yet complete.
	StoppingJobState JobState = "Stopping"
	// CompletedJobState shall represent that the operation completed successfully or with warnings.
	CompletedJobState JobState = "Completed"
	// CancelledJobState shall represent that the operation completed because the job was cancelled by an operator.
	CancelledJobState JobState = "Cancelled"
	// ExceptionJobState shall represent that the operation completed with errors.
	ExceptionJobState JobState = "Exception"
	// ServiceJobState shall represent that the operation is now running as a service and expected to continue
	// operation until stopped or killed.
	ServiceJobState JobState = "Service"
	// UserInterventionJobState shall represent that the operation is waiting for a user to intervene and needs to be
	// manually continued, stopped, or cancelled.
	UserInterventionJobState JobState = "UserIntervention"
	// ContinueJobState shall represent that the operation has been resumed from a paused condition and should return
	// to a Running state.
	ContinueJobState JobState = "Continue"
)

// Job shall contain a job in a Redfish implementation.
type Job struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// CreatedBy shall contain the user name, software program name, or other identifier indicating the creator of this
	// job.
	CreatedBy string
	// Description provides a description of this resource.
	Description string
	// EndTime shall indicate the date and time when the job was completed. This property shall not appear if the job
	// is running or was not completed. This property shall appear only if the JobState is Completed, Cancelled, or
	// Exception.
	EndTime string
	// EstimatedDuration shall indicate the estimated total time needed to complete the job. The value is not expected
	// to change while the job is in progress, but the service may update the value if it obtains new information that
	// significantly changes the expected duration. Services should be conservative in the reported estimate and
	// clients should treat this value as an estimate.
	EstimatedDuration string
	// HidePayload shall indicate whether the contents of the payload should be hidden from view after the job has been
	// created. If 'true', responses shall not return the Payload property. If 'false', responses shall return the
	// Payload property. If this property is not present when the job is created, the default is 'false'.
	HidePayload string
	// JobState shall indicate the state of the job.
	JobState string
	// JobStatus shall indicate the health status of the job. This property should contain 'Critical' if one or more
	// messages in the Messages array contains the severity 'Critical'. This property should contain 'Warning' if one
	// or more messages in the Messages array contains the severity 'Warning' and no messages contain the severity
	// 'Critical'. This property should contain 'OK' if all messages in the Messages array contain the severity 'OK' or
	// the array is empty.
	JobStatus string
	// MaxExecutionTime shall be an ISO 8601 conformant duration describing the maximum duration the job is allowed to
	// execute before being stopped by the service.
	MaxExecutionTime string
	// Messages shall contain an array of messages associated with the job.
	Messages []Message
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Payload shall contain the HTTP and JSON payload information for executing this job. This property shall not be
	// included in the response if the HidePayload property is 'true'.
	Payload string
	// PercentComplete shall indicate the completion progress of the job, reported in percent of completion. If the job
	// has not been started, the value shall be zero.
	PercentComplete int
	// Schedule shall contain the scheduling details for this job and the recurrence frequency for future instances of
	// this job.
	Schedule string
	// StartTime shall indicate the date and time when the job was last started or is scheduled to start.
	StartTime string
	// StepOrder shall contain an array of IDs for the job steps in the order that they shall be executed. Each step
	// shall be completed prior to the execution of the next step in array order. An incomplete list of steps shall be
	// considered an invalid configuration. If this property is not present or contains an empty array it shall
	// indicate that the step execution order is omitted and may occur in parallel or in series as determined by the
	// service.
	StepOrder []string
	// Steps shall contain the link to a resource collection of type JobCollection. This property shall not be present
	// if this resource represents a step for a job.
	Steps string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Job object from the raw JSON.
func (job *Job) UnmarshalJSON(b []byte) error {
	type temp Job
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*job = Job(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	job.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (job *Job) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Job)
	original.UnmarshalJSON(job.rawData)

	readWriteFields := []string{
		"JobState",
		"MaxExecutionTime",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(job).Elem()

	return job.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetJob will get a Job instance from the service.
func GetJob(c common.Client, uri string) (*Job, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var job Job
	err = json.NewDecoder(resp.Body).Decode(&job)
	if err != nil {
		return nil, err
	}

	job.SetClient(c)
	return &job, nil
}

// ListReferencedJobs gets the collection of Job from
// a provided reference.
func ListReferencedJobs(c common.Client, link string) ([]*Job, error) {
	var result []*Job
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, jobLink := range links.ItemLinks {
		job, err := GetJob(c, jobLink)
		if err != nil {
			collectionError.Failures[jobLink] = err
		} else {
			result = append(result, job)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

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

// Payload shall contain information detailing the HTTP and JSON payload information for executing this job.
type Payload struct {
	// HttpHeaders shall contain an array of HTTP headers in this job.
	HttpHeaders []string
	// HttpOperation shall contain the HTTP operation that executes this job.
	HttpOperation string
	// JsonBody shall contain JSON-formatted payload for this job.
	JsonBody string
	// TargetUri shall contain link to a target location for an HTTP operation.
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
