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
	"context"
	"fmt"

	"github.com/robfig/cron"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	batchv1 "tutorial.kubebuilder.io/k8s-cronjob-operator/api/v1"
)

// nolint:unused
// log is for logging in this package.
var cronjoblog = logf.Log.WithName("cronjob-resource")

// SetupCronjobWebhookWithManager registers the webhook for Cronjob in the manager.
func SetupCronjobWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&batchv1.Cronjob{}).
		WithValidator(&CronjobCustomValidator{}).
		WithDefaulter(&CronjobCustomDefaulter{
			DefaultConcurrencyPolicy:          batchv1.AllowConcurrent,
			DefaultSuspend:                    false,
			DefaultSuccessfulJobsHistoryLimit: 3,
			DefaultFailedJobsHistoryLimit:     1,
		}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-batch-tutorial-kubebuilder-io-v1-cronjob,mutating=true,failurePolicy=fail,sideEffects=None,groups=batch.tutorial.kubebuilder.io,resources=cronjobs,verbs=create;update,versions=v1,name=mcronjob-v1.kb.io,admissionReviewVersions=v1

// CronjobCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Cronjob when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type CronjobCustomDefaulter struct {

	// Default values for various CronJob fields
	DefaultConcurrencyPolicy          batchv1.ConcurrencyPolicy
	DefaultSuspend                    bool
	DefaultSuccessfulJobsHistoryLimit int32
	DefaultFailedJobsHistoryLimit     int32
}

var _ webhook.CustomDefaulter = &CronjobCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Cronjob.
func (d *CronjobCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	cronjob, ok := obj.(*batchv1.Cronjob)

	if !ok {
		return fmt.Errorf("expected an Cronjob object but got %T", obj)
	}
	cronjoblog.Info("Defaulting for Cronjob", "name", cronjob.GetName())

	// Set default values
	d.applyDefaults(cronjob)

	return nil
}

// applyDefaults applies default values to CronJob fields.
func (d *CronjobCustomDefaulter) applyDefaults(cronJob *batchv1.Cronjob) {
	if cronJob.Spec.ConcurrencyPolicy == "" {
		cronJob.Spec.ConcurrencyPolicy = d.DefaultConcurrencyPolicy
	}
	if cronJob.Spec.Suspend == nil {
		cronJob.Spec.Suspend = new(bool)
		*cronJob.Spec.Suspend = d.DefaultSuspend
	}
	if cronJob.Spec.SuccessfulJobsHistoryLimit == nil {
		cronJob.Spec.SuccessfulJobsHistoryLimit = new(int32)
		*cronJob.Spec.SuccessfulJobsHistoryLimit = d.DefaultSuccessfulJobsHistoryLimit
	}
	if cronJob.Spec.FailedJobsHistoryLimit == nil {
		cronJob.Spec.FailedJobsHistoryLimit = new(int32)
		*cronJob.Spec.FailedJobsHistoryLimit = d.DefaultFailedJobsHistoryLimit
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-batch-tutorial-kubebuilder-io-v1-cronjob,mutating=false,failurePolicy=fail,sideEffects=None,groups=batch.tutorial.kubebuilder.io,resources=cronjobs,verbs=create;update,versions=v1,name=vcronjob-v1.kb.io,admissionReviewVersions=v1

// CronjobCustomValidator struct is responsible for validating the Cronjob resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type CronjobCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &CronjobCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Cronjob.
func (v *CronjobCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cronjob, ok := obj.(*batchv1.Cronjob)
	if !ok {
		return nil, fmt.Errorf("expected a Cronjob object but got %T", obj)
	}
	cronjoblog.Info("Validation for Cronjob upon creation", "name", cronjob.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, validateCronJob(cronjob)
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Cronjob.
func (v *CronjobCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	cronjob, ok := newObj.(*batchv1.Cronjob)
	if !ok {
		return nil, fmt.Errorf("expected a Cronjob object for the newObj but got %T", newObj)
	}
	cronjoblog.Info("Validation for Cronjob upon update", "name", cronjob.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, validateCronJob(cronjob)
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Cronjob.
func (v *CronjobCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	cronjob, ok := obj.(*batchv1.Cronjob)
	if !ok {
		return nil, fmt.Errorf("expected a Cronjob object but got %T", obj)
	}
	cronjoblog.Info("Validation for Cronjob upon deletion", "name", cronjob.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}

// validateCronJob validates the fields of a CronJob object.
func validateCronJob(cronjob *batchv1.Cronjob) error {
	var allErrs field.ErrorList
	if err := validateCronJobName(cronjob); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validateCronJobSpec(cronjob); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "batch.tutorial.kubebuilder.io", Kind: "Cronjob"},
		cronjob.Name, allErrs)
}

func validateCronJobSpec(cronjob *batchv1.Cronjob) *field.Error {
	// The field helpers from the kubernetes API machinery help us return nicely
	// structured validation errors.
	return validateScheduleFormat(
		cronjob.Spec.Schedule,
		field.NewPath("spec").Child("schedule"))
}

func validateScheduleFormat(schedule string, fldPath *field.Path) *field.Error {
	if _, err := cron.ParseStandard(schedule); err != nil {
		return field.Invalid(fldPath, schedule, err.Error())
	}
	return nil
}

func validateCronJobName(cronjob *batchv1.Cronjob) *field.Error {
	if len(cronjob.ObjectMeta.Name) > validation.DNS1035LabelMaxLength-11 {
		// The job name length is 63 characters like all Kubernetes objects
		// (which must fit in a DNS subdomain). The cronjob controller appends
		// a 11-character suffix to the cronjob (`-$TIMESTAMP`) when creating
		// a job. The job name length limit is 63 characters. Therefore cronjob
		// names must have length <= 63-11=52. If we don't validate this here,
		// then job creation will fail later.
		return field.Invalid(field.NewPath("metadata").Child("name"), cronjob.ObjectMeta.Name, "must be no more than 52 characters")
	}
	return nil
}
