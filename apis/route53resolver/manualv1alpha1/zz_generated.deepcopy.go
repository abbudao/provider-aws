//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package manualv1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociation) DeepCopyInto(out *ResolverRuleAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociation.
func (in *ResolverRuleAssociation) DeepCopy() *ResolverRuleAssociation {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverRuleAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationList) DeepCopyInto(out *ResolverRuleAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResolverRuleAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationList.
func (in *ResolverRuleAssociationList) DeepCopy() *ResolverRuleAssociationList {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverRuleAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationObservation) DeepCopyInto(out *ResolverRuleAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationObservation.
func (in *ResolverRuleAssociationObservation) DeepCopy() *ResolverRuleAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationParameters) DeepCopyInto(out *ResolverRuleAssociationParameters) {
	*out = *in
	if in.ResolverRuleID != nil {
		in, out := &in.ResolverRuleID, &out.ResolverRuleID
		*out = new(string)
		**out = **in
	}
	if in.ResolverRuleIDRef != nil {
		in, out := &in.ResolverRuleIDRef, &out.ResolverRuleIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResolverRuleIDSelector != nil {
		in, out := &in.ResolverRuleIDSelector, &out.ResolverRuleIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCId != nil {
		in, out := &in.VPCId, &out.VPCId
		*out = new(string)
		**out = **in
	}
	if in.VPCIdRef != nil {
		in, out := &in.VPCIdRef, &out.VPCIdRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.VPCIdSelector != nil {
		in, out := &in.VPCIdSelector, &out.VPCIdSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationParameters.
func (in *ResolverRuleAssociationParameters) DeepCopy() *ResolverRuleAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationSpec) DeepCopyInto(out *ResolverRuleAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationSpec.
func (in *ResolverRuleAssociationSpec) DeepCopy() *ResolverRuleAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationStatus) DeepCopyInto(out *ResolverRuleAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationStatus.
func (in *ResolverRuleAssociationStatus) DeepCopy() *ResolverRuleAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationStatus)
	in.DeepCopyInto(out)
	return out
}
