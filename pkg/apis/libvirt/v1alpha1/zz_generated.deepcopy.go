// +build !ignore_autogenerated

/*
Copyright 2019 GOJEK TECH.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtMachineProviderSpec) DeepCopyInto(out *LibvirtMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtMachineProviderSpec.
func (in *LibvirtMachineProviderSpec) DeepCopy() *LibvirtMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(LibvirtMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibvirtMachineProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtMachineProviderSpecList) DeepCopyInto(out *LibvirtMachineProviderSpecList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LibvirtMachineProviderSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtMachineProviderSpecList.
func (in *LibvirtMachineProviderSpecList) DeepCopy() *LibvirtMachineProviderSpecList {
	if in == nil {
		return nil
	}
	out := new(LibvirtMachineProviderSpecList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LibvirtMachineProviderSpecList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtMachineProviderSpecSpec) DeepCopyInto(out *LibvirtMachineProviderSpecSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtMachineProviderSpecSpec.
func (in *LibvirtMachineProviderSpecSpec) DeepCopy() *LibvirtMachineProviderSpecSpec {
	if in == nil {
		return nil
	}
	out := new(LibvirtMachineProviderSpecSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LibvirtMachineProviderSpecStatus) DeepCopyInto(out *LibvirtMachineProviderSpecStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LibvirtMachineProviderSpecStatus.
func (in *LibvirtMachineProviderSpecStatus) DeepCopy() *LibvirtMachineProviderSpecStatus {
	if in == nil {
		return nil
	}
	out := new(LibvirtMachineProviderSpecStatus)
	in.DeepCopyInto(out)
	return out
}
