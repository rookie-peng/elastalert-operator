// +build !ignore_autogenerated

/*
Copyright 2021.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elastalert) DeepCopyInto(out *Elastalert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elastalert.
func (in *Elastalert) DeepCopy() *Elastalert {
	if in == nil {
		return nil
	}
	out := new(Elastalert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Elastalert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastalertList) DeepCopyInto(out *ElastalertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Elastalert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastalertList.
func (in *ElastalertList) DeepCopy() *ElastalertList {
	if in == nil {
		return nil
	}
	out := new(ElastalertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElastalertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastalertSpec) DeepCopyInto(out *ElastalertSpec) {
	*out = *in
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = make([]FreeForm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PodTemplateSpec.DeepCopyInto(&out.PodTemplateSpec)
	in.ConfigSetting.DeepCopyInto(&out.ConfigSetting)
	in.Alert.DeepCopyInto(&out.Alert)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastalertSpec.
func (in *ElastalertSpec) DeepCopy() *ElastalertSpec {
	if in == nil {
		return nil
	}
	out := new(ElastalertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElastalertStatus) DeepCopyInto(out *ElastalertStatus) {
	*out = *in
	if in.Condictions != nil {
		in, out := &in.Condictions, &out.Condictions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElastalertStatus.
func (in *ElastalertStatus) DeepCopy() *ElastalertStatus {
	if in == nil {
		return nil
	}
	out := new(ElastalertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FreeForm) DeepCopyInto(out *FreeForm) {
	*out = *in
	if in.FF != nil {
		in, out := &in.FF, &out.FF
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.json != nil {
		in, out := &in.json, &out.json
		*out = new([]byte)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FreeForm.
func (in *FreeForm) DeepCopy() *FreeForm {
	if in == nil {
		return nil
	}
	out := new(FreeForm)
	in.DeepCopyInto(out)
	return out
}
