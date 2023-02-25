/*
Copyright 2023.

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

package v1alpha1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var samplelog = logf.Log.WithName("sample webhook")

func (r *Sample) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-mygroup-example-com-v1alpha1-sample,mutating=true,failurePolicy=fail,sideEffects=None,groups=mygroup.example.com,resources=samples,verbs=create;update,versions=v1alpha1,name=msample.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Sample{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Sample) Default() {
	samplelog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-mygroup-example-com-v1alpha1-sample,mutating=false,failurePolicy=fail,sideEffects=None,groups=mygroup.example.com,resources=samples,verbs=create;update,versions=v1alpha1,name=vsample.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Sample{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Sample) ValidateCreate() error {
	samplelog.Info("validate create", "name", r.Name)

	if r.Spec.Foo == "forbidden_foo" {
		return fmt.Errorf("foo value of 'forbidden_foo' is not allowed")
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Sample) ValidateUpdate(old runtime.Object) error {
	samplelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Sample) ValidateDelete() error {
	samplelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
