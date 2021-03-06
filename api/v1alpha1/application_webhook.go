/*
Copyright 2021.

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
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var Applicationlog = logf.Log.WithName("Application-resource")

func (r *Application) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-cache-example-com-v1alpha1-Application,mutating=true,failurePolicy=fail,sideEffects=None,groups=cache.example.com,resources=Applications,verbs=create;update,versions=v1alpha1,name=mApplication.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &Application{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Application) Default() {
	Applicationlog.Info("default", "name", r.Name)

	if r.Spec.Size < r.Spec.MinSize {
		r.Spec.Size = r.Spec.MinSize
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-cache-example-com-v1alpha1-Application,mutating=false,failurePolicy=fail,sideEffects=None,groups=cache.example.com,resources=Applications,verbs=create;update,versions=v1alpha1,name=vApplication.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &Application{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Application) ValidateCreate() error {
	Applicationlog.Info("validate create", "name", r.Name)

	return validateInRange(r.Spec.Size, r.Spec.MinSize, r.Spec.MaxSize)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Application) ValidateUpdate(old runtime.Object) error {
	Applicationlog.Info("validate update", "name", r.Name)
	return validateInRange(r.Spec.Size, r.Spec.MinSize, r.Spec.MaxSize)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Application) ValidateDelete() error {
	Applicationlog.Info("validate delete", "name", r.Name)

	return nil
}
func validateInRange(requestedAmount int32, minRange int32, maxRange int32) error {
	if maxRange < requestedAmount || requestedAmount < minRange {
		return errors.New("cluster size must be within the range specified by the resource request for this operator")
	}

	return nil
}
