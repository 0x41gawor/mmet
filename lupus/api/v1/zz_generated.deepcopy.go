//go:build !ignore_autogenerated

/*
Copyright 2024.

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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Decide) DeepCopyInto(out *Decide) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Decide.
func (in *Decide) DeepCopy() *Decide {
	if in == nil {
		return nil
	}
	out := new(Decide)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Decide) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecideList) DeepCopyInto(out *DecideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Decide, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecideList.
func (in *DecideList) DeepCopy() *DecideList {
	if in == nil {
		return nil
	}
	out := new(DecideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DecideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecideSpec) DeepCopyInto(out *DecideSpec) {
	*out = *in
	out.Url = in.Url
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecideSpec.
func (in *DecideSpec) DeepCopy() *DecideSpec {
	if in == nil {
		return nil
	}
	out := new(DecideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecideStatus) DeepCopyInto(out *DecideStatus) {
	*out = *in
	in.Input.DeepCopyInto(&out.Input)
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecideStatus.
func (in *DecideStatus) DeepCopy() *DecideStatus {
	if in == nil {
		return nil
	}
	out := new(DecideStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Element) DeepCopyInto(out *Element) {
	*out = *in
	out.Url = in.Url
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Element.
func (in *Element) DeepCopy() *Element {
	if in == nil {
		return nil
	}
	out := new(Element)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Execute) DeepCopyInto(out *Execute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Execute.
func (in *Execute) DeepCopy() *Execute {
	if in == nil {
		return nil
	}
	out := new(Execute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Execute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecuteList) DeepCopyInto(out *ExecuteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Execute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecuteList.
func (in *ExecuteList) DeepCopy() *ExecuteList {
	if in == nil {
		return nil
	}
	out := new(ExecuteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExecuteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecuteSpec) DeepCopyInto(out *ExecuteSpec) {
	*out = *in
	out.Url = in.Url
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecuteSpec.
func (in *ExecuteSpec) DeepCopy() *ExecuteSpec {
	if in == nil {
		return nil
	}
	out := new(ExecuteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecuteStatus) DeepCopyInto(out *ExecuteStatus) {
	*out = *in
	in.Input.DeepCopyInto(&out.Input)
	in.LastUpdated.DeepCopyInto(&out.LastUpdated)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecuteStatus.
func (in *ExecuteStatus) DeepCopy() *ExecuteStatus {
	if in == nil {
		return nil
	}
	out := new(ExecuteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Loop) DeepCopyInto(out *Loop) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Loop.
func (in *Loop) DeepCopy() *Loop {
	if in == nil {
		return nil
	}
	out := new(Loop)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Loop) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopList) DeepCopyInto(out *LoopList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Loop, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopList.
func (in *LoopList) DeepCopy() *LoopList {
	if in == nil {
		return nil
	}
	out := new(LoopList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoopList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopSpec) DeepCopyInto(out *LoopSpec) {
	*out = *in
	if in.Elements != nil {
		in, out := &in.Elements, &out.Elements
		*out = make([]Element, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopSpec.
func (in *LoopSpec) DeepCopy() *LoopSpec {
	if in == nil {
		return nil
	}
	out := new(LoopSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoopStatus) DeepCopyInto(out *LoopStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoopStatus.
func (in *LoopStatus) DeepCopy() *LoopStatus {
	if in == nil {
		return nil
	}
	out := new(LoopStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observe) DeepCopyInto(out *Observe) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observe.
func (in *Observe) DeepCopy() *Observe {
	if in == nil {
		return nil
	}
	out := new(Observe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Observe) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObserveList) DeepCopyInto(out *ObserveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Observe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObserveList.
func (in *ObserveList) DeepCopy() *ObserveList {
	if in == nil {
		return nil
	}
	out := new(ObserveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObserveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObserveSpec) DeepCopyInto(out *ObserveSpec) {
	*out = *in
	out.Url = in.Url
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObserveSpec.
func (in *ObserveSpec) DeepCopy() *ObserveSpec {
	if in == nil {
		return nil
	}
	out := new(ObserveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObserveStatus) DeepCopyInto(out *ObserveStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObserveStatus.
func (in *ObserveStatus) DeepCopy() *ObserveStatus {
	if in == nil {
		return nil
	}
	out := new(ObserveStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *URL) DeepCopyInto(out *URL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new URL.
func (in *URL) DeepCopy() *URL {
	if in == nil {
		return nil
	}
	out := new(URL)
	in.DeepCopyInto(out)
	return out
}
