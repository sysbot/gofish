//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// MetricReport shall represent a metric report in a Redfish implementation. When a metric report is deleted, the
// historic metric data used to generate the report shall be deleted as well unless other metric reports are
// consuming the data.
type MetricReport struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Context shall contain a client supplied context for the event destination to which this event is being sent.
	// This property shall only be present when sent as a payload in an event.
	Context string
	// Description provides a description of this resource.
	Description string
	// MetricReportDefinition shall contain a link to a resource of type MetricReportDefinition.
	MetricReportDefinition string
	// MetricValues shall be metric values for this metric report.
	MetricValues []MetricValue
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Timestamp shall contain the time when the metric report was generated.
	Timestamp string
}

// UnmarshalJSON unmarshals a MetricReport object from the raw JSON.
func (metricreport *MetricReport) UnmarshalJSON(b []byte) error {
	type temp MetricReport
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*metricreport = MetricReport(t.temp)

	// Extract the links to other entities for later

	return nil
}

// GetMetricReport will get a MetricReport instance from the service.
func GetMetricReport(c common.Client, uri string) (*MetricReport, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var metricreport MetricReport
	err = json.NewDecoder(resp.Body).Decode(&metricreport)
	if err != nil {
		return nil, err
	}

	metricreport.SetClient(c)
	return &metricreport, nil
}

// ListReferencedMetricReports gets the collection of MetricReport from
// a provided reference.
func ListReferencedMetricReports(c common.Client, link string) ([]*MetricReport, error) {
	var result []*MetricReport
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	collectionError := common.NewCollectionError()
	for _, metricreportLink := range links.ItemLinks {
		metricreport, err := GetMetricReport(c, metricreportLink)
		if err != nil {
			collectionError.Failures[metricreportLink] = err
		} else {
			result = append(result, metricreport)
		}
	}

	if collectionError.Empty() {
		return result, nil
	} else {
		return result, collectionError
	}
}

// MetricValue shall contain properties that capture a metric value and other associated information.
type MetricValue struct {
	// MetricId shall contain the value of the Id property of the MetricDefinition resource that contains additional
	// information for the source metric.
	MetricId string
	// MetricProperty shall contain a URI following RFC6901-specified JSON pointer notation to the property from which
	// this metric is derived. The value of MetricValue may contain additional calculations performed on the property
	// based upon the configuration of the MetricReportDefinition.
	MetricProperty string
	// MetricValue shall contain the metric value, as a string.
	MetricValue string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Timestamp shall time when the metric value was obtained. Note that this value may be different from the time
	// when this instance is created.
	Timestamp string
}

// UnmarshalJSON unmarshals a MetricValue object from the raw JSON.
func (metricvalue *MetricValue) UnmarshalJSON(b []byte) error {
	type temp MetricValue
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*metricvalue = MetricValue(t.temp)

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
