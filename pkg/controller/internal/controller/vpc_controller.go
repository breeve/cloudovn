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
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	ovnclient "github.com/kubeovn/kube-ovn/pkg/ovs"

	controllerv1 "github.com/breeve/cloudovn/pkg/controller/api/v1"
	"github.com/breeve/cloudovn/pkg/controller/internal/utils"
)

// VPCReconciler reconciles a VPC object
type VPCReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	ovn    ovnclient.OVNNbClient

	log logr.Logger
}

// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=vpcs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=vpcs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=vpcs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *VPCReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	result := ctrl.Result{}

	r.log.Info("reconile vpc", "name", req.NamespacedName.String())

	vpc := &controllerv1.VPC{}
	if err := r.Client.Get(ctx, req.NamespacedName, vpc); err != nil {
		if errors.IsNotFound(err) {
			return result, nil
		}
		r.log.Error(err, "failed to get vpc", "name", req.NamespacedName.String())
		return result, err
	}

	if vpc.DeletionTimestamp.IsZero() {
		// 1. finalizer
		if !controllerutil.ContainsFinalizer(vpc, VPCFinalizer) {
			r.log.Info("vpc add finalizer", "name", req.NamespacedName.String(), "finalizer", VPCFinalizer)
			if err := utils.AddFinalizer(ctx, r.log, r.Client, func() client.Object { return vpc.DeepCopy() }, VPCFinalizer); err != nil {
				r.log.Error(err, "failed to add finalizer", "name", req.NamespacedName.String())
				return result, err
			}
			return result, nil
		}

		// 2. update
		if err := r.update(ctx, vpc); err != nil {
			return result, err
		}

		return result, nil
	}

	// 3. delete
	if err := r.delete(ctx, vpc); err != nil {
		return result, err
	}

	return result, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VPCReconciler) SetupWithManager(mgr ctrl.Manager) error {
	log := logf.FromContext(context.Background())
	r.log = log

	return ctrl.NewControllerManagedBy(mgr).
		For(&controllerv1.VPC{}).
		Named("vpc").
		Complete(r)
}

const (
	VPCFinalizer = "cloudovn.io/vpc"
)

func (r *VPCReconciler) update(ctx context.Context, vpc *controllerv1.VPC) error {
	lrName := vpcLogicRouterName(vpc.Name)
	lr, err := r.ovn.GetLogicalRouter(lrName, true)
	if err != nil {
		r.log.Error(err, "failed to update vpc, get logic router error", "name", fmt.Sprintf("%s/%s", vpc.Namespace, vpc.Name), "lr-name", lrName)
		return err
	}
	if lr == nil {
		r.log.Info("update vpc, create logic router skip, exist already", "name", fmt.Sprintf("%s/%s", vpc.Namespace, vpc.Name), "lr-name", lrName)
		if err := r.ovn.CreateLogicalRouter(lrName); err != nil {
			r.log.Error(err, "failed to update vpc, create logic router error", "name", fmt.Sprintf("%s/%s", vpc.Namespace, vpc.Name), "lr-name", lrName)
			return err
		}
	}

	labels := map[string]string{
		controllerv1.LabelKeyLogicRouterName: lrName,
	}
	if err := utils.PatchLabels(ctx, r.log, r.Client, func() client.Object { return vpc.DeepCopy() }, labels); err != nil {
		r.log.Error(err, "failed to update vpc, patch vpc logic router name error", "name", fmt.Sprintf("%s/%s", vpc.Namespace, vpc.Name), "lr-name", lrName)
		return err
	}
	r.log.Info("update vpc, vpc logic router name label set success", "name", fmt.Sprintf("%s/%s", vpc.Namespace, vpc.Name), "lr-name", lrName)

	err = utils.PatchStatus(ctx, r.log, r.Client, func() client.Object { return vpc.DeepCopy() }, func(o client.Object) {
		object := o.(*controllerv1.VPC)
		object.Status.Message = ""
		object.Status.State = controllerv1.ACTIVE
		object.Status.LogicRouterName = lrName
	})
	if err != nil {
		r.log.Error(err, "failed to update vpc status", "name", fmt.Sprintf("%s/%s", vpc.Namespace, vpc.Name), "lr-name", lrName, "state", controllerv1.ACTIVE)
		return err
	}
	r.log.Info("update vpc status success", "name", fmt.Sprintf("%s/%s", vpc.Namespace, vpc.Name), "lr-name", lrName, "state", controllerv1.ACTIVE)

	return nil
}

func (r *VPCReconciler) delete(ctx context.Context, vpc *controllerv1.VPC) error {
	// TODO(flynn): Impl it.
	return nil
}

func vpcLogicRouterName(vpcName string) string {
	return vpcName
}
