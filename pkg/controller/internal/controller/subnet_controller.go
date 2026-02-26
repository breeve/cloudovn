/*
Copyright 2026.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	controllerv1 "github.com/breeve/cloudovn/pkg/controller/api/v1"
	"github.com/breeve/cloudovn/pkg/controller/internal/utils"
	"github.com/go-logr/logr"
)

// SubnetReconciler reconciles a Subnet object
type SubnetReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	log logr.Logger
}

// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=subnets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=subnets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=subnets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Subnet object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
func (r *SubnetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	result := ctrl.Result{}

	r.log.Info("reconile subnet", "name", req.NamespacedName.String())

	subnet := &controllerv1.Subnet{}
	if err := r.Client.Get(ctx, req.NamespacedName, subnet); err != nil {
		if errors.IsNotFound(err) {
			return result, nil
		}
		r.log.Error(err, "failed to get subnet", "name", req.NamespacedName.String())
		return result, err
	}

	if subnet.DeletionTimestamp.IsZero() {
		// 1. finalizer
		if !controllerutil.ContainsFinalizer(subnet, SubnetFinalizer) {
			r.log.Info("subnet add finalizer", "name", req.NamespacedName.String(), "finalizer", SubnetFinalizer)
			if err := utils.AddFinalizer(ctx, r.log, r.Client, func() client.Object { return subnet.DeepCopy() }, SubnetFinalizer); err != nil {
				r.log.Error(err, "failed to add finalizer", "name", req.NamespacedName.String())
				return result, err
			}
			return result, nil
		}

		// 2. update
		if err := r.update(ctx, subnet); err != nil {
			return result, err
		}

		return result, nil
	}

	// 3. delete
	if err := r.delete(ctx, subnet); err != nil {
		return result, err
	}

	return result, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SubnetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	log := logf.FromContext(context.Background())
	r.log = log

	return ctrl.NewControllerManagedBy(mgr).
		For(&controllerv1.Subnet{}).
		Named("subnet").
		Complete(r)
}

const (
	SubnetFinalizer = "cloudovn.io/subnet"
)

func (r *SubnetReconciler) update(ctx context.Context, subnet *controllerv1.Subnet) error {
	// TODO(flynn): Impl it.
	return nil
}

func (r *SubnetReconciler) delete(ctx context.Context, subnet *controllerv1.Subnet) error {
	// TODO(flynn): Impl it.
	return nil
}
