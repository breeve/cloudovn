package utils

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/equality"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func AddFinalizer(ctx context.Context, log logr.Logger, c client.Client, deepCopyObject func() client.Object, finalizer string) error {
	base := deepCopyObject()
	if controllerutil.ContainsFinalizer(base, finalizer) {
		return nil
	}

	patched := deepCopyObject()
	if !controllerutil.AddFinalizer(patched, finalizer) {
		return nil
	}

	patch := client.MergeFrom(base)
	if err := c.Patch(ctx, patched, patch); client.IgnoreNotFound(err) != nil {
		log.Error(err, "failed to patch finalizer", "name", base.GetName())
		return err
	}

	return nil
}

func PatchLabels(ctx context.Context, log logr.Logger, c client.Client, deepCopyObject func() client.Object, labels map[string]string) error {
	if len(labels) == 0 {
		return nil
	}

	base := deepCopyObject()
	if ContainsMaps(base.GetLabels(), labels) {
		return nil
	}

	patched := deepCopyObject()
	patched.SetLabels(MergeMaps(base.GetLabels(), labels))

	patch := client.MergeFrom(base)
	if err := c.Patch(ctx, patched, patch); client.IgnoreNotFound(err) != nil {
		log.Error(err, "failed to patch labels", "name", base.GetName())
		return err
	}

	return nil
}

func PatchStatus(ctx context.Context, log logr.Logger, c client.Client, deepCopyObject func() client.Object, updateStatus func(client.Object)) error {
	base := deepCopyObject()

	patched := deepCopyObject()
	updateStatus(patched)

	if equality.Semantic.DeepEqual(base, patched) {
		return nil
	}

	patch := client.MergeFrom(base)
	if err := c.Status().Patch(ctx, patched, patch, client.FieldOwner("cloudovn-controller")); client.IgnoreNotFound(err) != nil {
		log.Error(err, "failed to patch status", "name", base.GetName())
		return err
	}

	return nil
}
