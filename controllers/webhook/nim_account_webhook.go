package webhook

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:webhook:admissionReviewVersions=v1,path=/validate-nim-opendatahub-io-v1-account,mutating=false,failurePolicy=fail,groups=nim.opendatahub.io,resources=accounts,verbs=create;update,versions=v1,name=validating.nim.account.odh-model-controller.opendatahub.io,sideEffects=None

type accountValidator struct {
	client client.Client
}

func NewAccountValidator(client client.Client) admission.CustomValidator {
	return &accountValidator{client: client}
}

func (w *accountValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	// TODO Only the ODH NIM Operator is allowed to Create or Delete Account
	//return w.verifyOnlyOneInNamespace(ctx, obj)
	return nil, nil
}

func (w *accountValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (warnings admission.Warnings, err error) {
	return nil, nil
}

func (w *accountValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (warnings admission.Warnings, err error) {
	return nil, nil
}
