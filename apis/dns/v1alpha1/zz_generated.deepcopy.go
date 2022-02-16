//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSet) DeepCopyInto(out *ResourceRecordSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSet.
func (in *ResourceRecordSet) DeepCopy() *ResourceRecordSet {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceRecordSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetList) DeepCopyInto(out *ResourceRecordSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceRecordSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetList.
func (in *ResourceRecordSetList) DeepCopy() *ResourceRecordSetList {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceRecordSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetObservation) DeepCopyInto(out *ResourceRecordSetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetObservation.
func (in *ResourceRecordSetObservation) DeepCopy() *ResourceRecordSetObservation {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetParameters) DeepCopyInto(out *ResourceRecordSetParameters) {
	*out = *in
	if in.RRDatas != nil {
		in, out := &in.RRDatas, &out.RRDatas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SignatureRRDatas != nil {
		in, out := &in.SignatureRRDatas, &out.SignatureRRDatas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetParameters.
func (in *ResourceRecordSetParameters) DeepCopy() *ResourceRecordSetParameters {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetSpec) DeepCopyInto(out *ResourceRecordSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetSpec.
func (in *ResourceRecordSetSpec) DeepCopy() *ResourceRecordSetSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRecordSetStatus) DeepCopyInto(out *ResourceRecordSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRecordSetStatus.
func (in *ResourceRecordSetStatus) DeepCopy() *ResourceRecordSetStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceRecordSetStatus)
	in.DeepCopyInto(out)
	return out
}
