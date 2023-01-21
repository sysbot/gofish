//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// JobService shall represent a job service for a Redfish implementation.
type JobService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// DateTime shall contain the current date and time setting for the job service.
	DateTime string
	// Description provides a description of this resource.
	Description string
	// Jobs shall contain a link to a resource collection of type JobCollection.
	Jobs string
	// Log shall contain a link to a resource of type LogService that this job service uses.
	Log string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceCapabilities shall contain properties that describe the capabilities or supported features of this
	// implementation of a job service.
	ServiceCapabilities string
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a JobService object from the raw JSON.
func (jobservice *JobService) UnmarshalJSON(b []byte) error {
	type temp JobService
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*jobservice = JobService(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	jobservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (jobservice *JobService) Update() error {

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(JobService)
	original.UnmarshalJSON(jobservice.rawData)

	readWriteFields := []string{
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(jobservice).Elem()

	return jobservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetJobService will get a JobService instance from the service.
func GetJobService(c common.Client, uri string) (*JobService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jobservice JobService
	err = json.NewDecoder(resp.Body).Decode(&jobservice)
	if err != nil {
		return nil, err
	}

	jobservice.SetClient(c)
	return &jobservice, nil
}

// ListReferencedJobServices gets the collection of JobService from
// a provided reference.
func ListReferencedJobServices(c common.Client, link string) ([]*JobService, error) {
	var result []*JobService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, jobserviceLink := range links.ItemLinks {
		jobservice, err := GetJobService(c, jobserviceLink)
		if err != nil {
			collectionError.Failures[jobserviceLink] = err
		} else {
			result = append(result, jobservice)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// JobServiceCapabilities shall contain properties that describe the capabilities or supported features of this
// implementation of a job service.
type JobServiceCapabilities struct {
	// MaxJobs shall contain the maximum number of jobs supported by the implementation.
	MaxJobs int
	// MaxSteps shall contain the maximum number of steps supported by a single job instance.
	MaxSteps int
	// Scheduling shall indicate whether the Schedule property within the job supports scheduling of jobs.
	Scheduling bool
}

// UnmarshalJSON unmarshals a JobServiceCapabilities object from the raw JSON.
func (jobservicecapabilities *JobServiceCapabilities) UnmarshalJSON(b []byte) error {
	type temp JobServiceCapabilities
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*jobservicecapabilities = JobServiceCapabilities(t.temp)

	// Extract the links to other entities for later

	return nil
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
