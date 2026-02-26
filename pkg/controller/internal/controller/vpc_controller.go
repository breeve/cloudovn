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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	ovnclient "github.com/kubeovn/kube-ovn/pkg/ovs"

	controllerv1 "github.com/breeve/cloudovn/pkg/controller/api/v1"
)

// VPCReconciler reconciles a VPC object
type VPCReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	ovn    ovnclient.OVNNbClient
}

// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=vpcs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=vpcs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=controller.cloudovn.io,resources=vpcs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *VPCReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	r.ovn.CreateLogicalRouter("")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *VPCReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&controllerv1.VPC{}).
		Named("vpc").
		Complete(r)
}
