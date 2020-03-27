// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vulnerability) DeepCopyInto(out *Vulnerability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vulnerability.
func (in *Vulnerability) DeepCopy() *Vulnerability {
	if in == nil {
		return nil
	}
	out := new(Vulnerability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vulnerability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityItem) DeepCopyInto(out *VulnerabilityItem) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityItem.
func (in *VulnerabilityItem) DeepCopy() *VulnerabilityItem {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityList) DeepCopyInto(out *VulnerabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vulnerability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityList.
func (in *VulnerabilityList) DeepCopy() *VulnerabilityList {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityReport) DeepCopyInto(out *VulnerabilityReport) {
	*out = *in
	out.Summary = in.Summary
	if in.Vulnerabilities != nil {
		in, out := &in.Vulnerabilities, &out.Vulnerabilities
		*out = make([]VulnerabilityItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityReport.
func (in *VulnerabilityReport) DeepCopy() *VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilitySummary) DeepCopyInto(out *VulnerabilitySummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilitySummary.
func (in *VulnerabilitySummary) DeepCopy() *VulnerabilitySummary {
	if in == nil {
		return nil
	}
	out := new(VulnerabilitySummary)
	in.DeepCopyInto(out)
	return out
}