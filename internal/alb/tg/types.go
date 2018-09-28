package tg

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/alb/tags"
	"github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/aws/albelbv2"
	"github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/ingress/controller/store"
	"github.com/kubernetes-sigs/aws-alb-ingress-controller/pkg/util/log"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// TargetGroups is a slice of TargetGroups
type TargetGroups []*TargetGroup

// TargetGroup contains the current & desired configuration
type TargetGroup struct {
	ID         string
	SvcName    string
	SvcPort    intstr.IntOrString
	TargetType string

	tg         tg
	attributes *Attributes
	tags       *tags.Tags
	targets    targets

	deleted bool
	logger  *log.Logger
}

type tg struct {
	current *elbv2.TargetGroup
	desired *elbv2.TargetGroup
}

type targets struct {
	current albelbv2.TargetDescriptions
	desired albelbv2.TargetDescriptions
}

type ReconcileOptions struct {
	Store                  store.Storer
	Eventf                 func(string, string, string, ...interface{})
	VpcID                  *string
	IgnoreDeletes          bool
	TgAttributesController AttributesController
	TagsController         tags.TagsController
}

type tgChange uint

const (
	paramsModified tgChange = 1 << iota
	targetsModified
)

// CopyCurrentToDesired is used for testing other packages against tg
func CopyCurrentToDesired(a *TargetGroup) {
	if a != nil {
		a.tg.desired = a.tg.current
		a.tags = a.tags
		a.targets.desired = a.targets.current
	}
}
