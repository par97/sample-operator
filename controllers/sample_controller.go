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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	mygroupv1alpha1 "github.com/example/sample-operator/api/v1alpha1"
)

// SampleReconciler reconciles a Sample object
type SampleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=mygroup.example.com,resources=samples,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=mygroup.example.com,resources=samples/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=mygroup.example.com,resources=samples/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Sample object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *SampleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("enter sample reconcile")

	sample := &mygroupv1alpha1.Sample{}
	err := r.Get(ctx, req.NamespacedName, sample)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Info("sample CR was not found")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "error to get sample CR")
		return ctrl.Result{}, err
	}
	logger.Info("got sample CR, foo=" + sample.Spec.Foo)

	sample.Status.State = sample.Spec.Foo

	return ctrl.Result{}, r.Status().Update(ctx, sample)
}

// SetupWithManager sets up the controller with the Manager.
func (r *SampleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mygroupv1alpha1.Sample{}).
		Complete(r)
}
