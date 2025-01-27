package workspace

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/prometheusservice"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	svcapitypes "github.com/crossplane/provider-aws/apis/prometheusservice/v1alpha1"
	awsclients "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errNotWorkspace     = "managed resource is not an Workspace custom resource"
	errKubeUpdateFailed = "cannot update Workspace custom resource"
)

// SetupWorkspace adds a controller that reconciles Workspace for PrometheusService.
func SetupWorkspace(mgr ctrl.Manager, l logging.Logger, rl workqueue.RateLimiter, poll time.Duration) error {
	name := managed.ControllerName(svcapitypes.WorkspaceGroupKind)
	opts := []option{
		func(e *external) {
			e.postObserve = postObserve
			e.preObserve = preObserve
			e.postCreate = postCreate
			e.preDelete = preDelete
			e.postDelete = postDelete
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(controller.Options{
			RateLimiter: ratelimiter.NewController(rl),
		}).
		For(&svcapitypes.Workspace{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(svcapitypes.WorkspaceGroupVersionKind),
			managed.WithExternalConnecter(&connector{kube: mgr.GetClient(), opts: opts}),
			managed.WithInitializers(managed.NewDefaultProviderConfig(mgr.GetClient()), &tagger{kube: mgr.GetClient()}),
			managed.WithPollInterval(poll),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

func preObserve(_ context.Context, cr *svcapitypes.Workspace, obj *svcsdk.DescribeWorkspaceInput) error {
	obj.WorkspaceId = aws.String(meta.GetExternalName(cr))
	return nil
}

func postObserve(_ context.Context, cr *svcapitypes.Workspace, resp *svcsdk.DescribeWorkspaceOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	switch aws.StringValue(resp.Workspace.Status.StatusCode) {
	case string(svcapitypes.WorkspaceStatusCode_ACTIVE):
		cr.SetConditions(xpv1.Available())
	case string(svcapitypes.WorkspaceStatusCode_CREATING):
		cr.SetConditions(xpv1.Creating())
	case string(svcapitypes.WorkspaceStatusCode_CREATION_FAILED):
		cr.SetConditions(xpv1.Unavailable())
	}

	cr.Status.AtProvider.ARN = resp.Workspace.Arn
	cr.Status.AtProvider.PrometheusEndpoint = resp.Workspace.PrometheusEndpoint
	cr.Status.AtProvider.Status.StatusCode = resp.Workspace.Status.StatusCode

	obs.ConnectionDetails = managed.ConnectionDetails{
		"arn":                []byte(awsclients.StringValue(resp.Workspace.Arn)),
		"prometheusEndpoint": []byte(awsclients.StringValue(resp.Workspace.PrometheusEndpoint)),
		"workspaceId":        []byte(awsclients.StringValue(resp.Workspace.WorkspaceId)),
	}

	return obs, nil
}

func postCreate(_ context.Context, cr *svcapitypes.Workspace, resp *svcsdk.CreateWorkspaceOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	if err != nil {
		return managed.ExternalCreation{}, err
	}
	meta.SetExternalName(cr, aws.StringValue(resp.WorkspaceId))
	return cre, nil
}

func preDelete(_ context.Context, cr *svcapitypes.Workspace, obj *svcsdk.DeleteWorkspaceInput) (bool, error) {
	obj.WorkspaceId = aws.String(meta.GetExternalName(cr))
	return false, nil
}

func postDelete(_ context.Context, cr *svcapitypes.Workspace, obj *svcsdk.DeleteWorkspaceOutput, err error) error {
	if err != nil {
		if strings.Contains(err.Error(), svcsdk.ErrCodeConflictException) {
			// skip: Can't delete workspace in non-ACTIVE state. Current status is DELETING
			return nil
		}
		return err
	}
	return err
}

type tagger struct {
	kube client.Client
}

func (t *tagger) Initialize(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*svcapitypes.Workspace)
	if !ok {
		return errors.New(errNotWorkspace)
	}
	if cr.Spec.ForProvider.Tags == nil {
		cr.Spec.ForProvider.Tags = map[string]*string{}
	}
	for k, v := range resource.GetExternalTags(mg) {
		cr.Spec.ForProvider.Tags[k] = awsclients.String(v)
	}
	return errors.Wrap(t.kube.Update(ctx, cr), errKubeUpdateFailed)
}
