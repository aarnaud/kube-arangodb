//go:build !ignore_autogenerated
// +build !ignore_autogenerated

//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	types "k8s.io/apimachinery/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in HashList) DeepCopyInto(out *HashList) {
	{
		in := &in
		*out = make(HashList, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HashList.
func (in HashList) DeepCopy() HashList {
	if in == nil {
		return nil
	}
	out := new(HashList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Object) DeepCopyInto(out *Object) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(types.UID)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Object.
func (in *Object) DeepCopy() *Object {
	if in == nil {
		return nil
	}
	out := new(Object)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
	*out = *in
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(Object)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespaced != nil {
		in, out := &in.Namespaced, &out.Namespaced
		*out = new(ServiceAccountRole)
		(*in).DeepCopyInto(*out)
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(ServiceAccountRole)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccount.
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountRole) DeepCopyInto(out *ServiceAccountRole) {
	*out = *in
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(Object)
		(*in).DeepCopyInto(*out)
	}
	if in.Binding != nil {
		in, out := &in.Binding, &out.Binding
		*out = new(Object)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountRole.
func (in *ServiceAccountRole) DeepCopy() *ServiceAccountRole {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountRole)
	in.DeepCopyInto(out)
	return out
}
