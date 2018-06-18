/*
Jobs may be created by:
* a request handler that doesn't want to keep the response open until completion
* a background process that needs somewhere to store the result

TODO: Do they execute themselves, or are they only used as a communication method?
*/
package models

import (
	"encoding/json"
	"strings"

	"github.com/hjkelly/zbbapi/common"
)

type Job struct {
	ID     string                 `json:"id"`
	Status string                 `json:"status"`
	Result map[string]interface{} `json:"result"`
	Timestamped
}

func NewJob(id string) *Job {
	job := &Job{
		ID:     id,
		Status: JobStatusPending,
	}
	job.SetCreationTimestamp()
	return job
}

func (job Job) GetValidated() (Job, error) {
	var errs []error

	// status ----------
	validStatus := false
	for _, status := range JobStatuses {
		if job.Status == status {
			validStatus = true
			break
		}
	}
	if !validStatus {
		errs = append(errs, common.NewValidationError("status", common.BadEnumChoiceCode, "Job status must be one of: %s", strings.Join(JobStatuses, ", ")))
	}

	// id ----------
	if len(job.ID) == 0 {
		errs = append(errs, common.NewValidationError(common.MissingCode, "id", "You must provide this field."))
	}

	// combining and wrapping up ----------
	err := common.CombineErrors(errs...)
	if err != nil {
		return Job{}, err
	}
	return job, nil
}

func (job Job) MarshalText() (text []byte, err error) {
	originalJSON, err := json.Marshal(job)
	if err != nil {
		return []byte{}, err
	}
	data := map[string]interface{}{}
	err = json.Unmarshal(originalJSON, data)
	if err != nil {
		return []byte{}, err
	}
	data["url"] = "https://example.com/v1/jobs/" + job.ID
	return json.Marshal(data)
}

const (
	JobStatusPending  = "pending"
	JobStatusFinished = "finished"
	JobStatusFailed   = "failed"
)

var JobStatuses = []string{
	JobStatusPending,
	JobStatusFinished,
	JobStatusFailed,
}
