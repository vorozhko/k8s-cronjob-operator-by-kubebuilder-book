/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConcurrencyPolicy describes how the job will be handled.
// +kubebuilder:validation:Enum=Allow;Forbid;Replace
type ConcurrencyPolicy string

const (
	// AllowConcurrent allows CronJobs to run concurrently.
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent forbids concurrent runs, skipping next run if previous
	// hasn't finished yet.
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent cancels currently running job and replaces it with a new one.
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

// CronjobSpec defines the desired state of Cronjob.
// +kubebuilder:pruning:PreserveUnknownFields
type CronjobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule string `json:"schedule"`

	// +kubebuilder:validation:Minimum=0

	// Optional deadline in seconds for starting the job
	// +optional
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	// Specifies how to treat concurrent executions of a Job.
	// Valid values are:
	// - "Allow" (default): allows CronJobs to run concurrently;
	// - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet;
	// - "Replace": cancels currently running job and replaces it with a new one
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

	// This flag tells the controller to suspend subsequent executions
	// +optional
	Suspend *bool `json:"suspend,omitempty"`

	// Specifies the job that will be created when executing a CronJob.
	// +kubebuilder:pruning:PreserveUnknownFields
	JobTemplate batchv1.JobTemplateSpec `json:"jobTemplate"`

	// +kubebuilder:validation:Minimum=0

	// The number of successful finished jobs to retain.
	// +optional
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// +kubebuilder:validation:Minimum=0

	// The number of failed finished jobs to retain.
	// +optional
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}

// CronjobStatus defines the observed state of Cronjob.
type CronjobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A list of pointers to currently running jobs.
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`

	// Information when was the last time the job was successfully scheduled.
	// +optional
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Cronjob is the Schema for the cronjobs API.
type Cronjob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronjobSpec   `json:"spec,omitempty"`
	Status CronjobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronjobList contains a list of Cronjob.
type CronjobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cronjob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cronjob{}, &CronjobList{})
}
