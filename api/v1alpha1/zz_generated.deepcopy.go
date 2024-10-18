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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/external-dns/endpoint"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalHeader) DeepCopyInto(out *AdditionalHeader) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalHeader.
func (in *AdditionalHeader) DeepCopy() *AdditionalHeader {
	if in == nil {
		return nil
	}
	out := new(AdditionalHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AdditionalHeaders) DeepCopyInto(out *AdditionalHeaders) {
	{
		in := &in
		*out = make(AdditionalHeaders, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalHeaders.
func (in AdditionalHeaders) DeepCopy() AdditionalHeaders {
	if in == nil {
		return nil
	}
	out := new(AdditionalHeaders)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalHeadersRef) DeepCopyInto(out *AdditionalHeadersRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalHeadersRef.
func (in *AdditionalHeadersRef) DeepCopy() *AdditionalHeadersRef {
	if in == nil {
		return nil
	}
	out := new(AdditionalHeadersRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSHealthCheckProbe) DeepCopyInto(out *DNSHealthCheckProbe) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSHealthCheckProbe.
func (in *DNSHealthCheckProbe) DeepCopy() *DNSHealthCheckProbe {
	if in == nil {
		return nil
	}
	out := new(DNSHealthCheckProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSHealthCheckProbe) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSHealthCheckProbeList) DeepCopyInto(out *DNSHealthCheckProbeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSHealthCheckProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSHealthCheckProbeList.
func (in *DNSHealthCheckProbeList) DeepCopy() *DNSHealthCheckProbeList {
	if in == nil {
		return nil
	}
	out := new(DNSHealthCheckProbeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSHealthCheckProbeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSHealthCheckProbeSpec) DeepCopyInto(out *DNSHealthCheckProbeSpec) {
	*out = *in
	out.Interval = in.Interval
	if in.AdditionalHeadersRef != nil {
		in, out := &in.AdditionalHeadersRef, &out.AdditionalHeadersRef
		*out = new(AdditionalHeadersRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSHealthCheckProbeSpec.
func (in *DNSHealthCheckProbeSpec) DeepCopy() *DNSHealthCheckProbeSpec {
	if in == nil {
		return nil
	}
	out := new(DNSHealthCheckProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSHealthCheckProbeStatus) DeepCopyInto(out *DNSHealthCheckProbeStatus) {
	*out = *in
	in.LastCheckedAt.DeepCopyInto(&out.LastCheckedAt)
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSHealthCheckProbeStatus.
func (in *DNSHealthCheckProbeStatus) DeepCopy() *DNSHealthCheckProbeStatus {
	if in == nil {
		return nil
	}
	out := new(DNSHealthCheckProbeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecord) DeepCopyInto(out *DNSRecord) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecord.
func (in *DNSRecord) DeepCopy() *DNSRecord {
	if in == nil {
		return nil
	}
	out := new(DNSRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSRecord) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordList) DeepCopyInto(out *DNSRecordList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DNSRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordList.
func (in *DNSRecordList) DeepCopy() *DNSRecordList {
	if in == nil {
		return nil
	}
	out := new(DNSRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DNSRecordList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordSpec) DeepCopyInto(out *DNSRecordSpec) {
	*out = *in
	out.ProviderRef = in.ProviderRef
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]*endpoint.Endpoint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(endpoint.Endpoint)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordSpec.
func (in *DNSRecordSpec) DeepCopy() *DNSRecordSpec {
	if in == nil {
		return nil
	}
	out := new(DNSRecordSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordStatus) DeepCopyInto(out *DNSRecordStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.QueuedAt.DeepCopyInto(&out.QueuedAt)
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]*endpoint.Endpoint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(endpoint.Endpoint)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ZoneEndpoints != nil {
		in, out := &in.ZoneEndpoints, &out.ZoneEndpoints
		*out = make([]*endpoint.Endpoint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(endpoint.Endpoint)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.DomainOwners != nil {
		in, out := &in.DomainOwners, &out.DomainOwners
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordStatus.
func (in *DNSRecordStatus) DeepCopy() *DNSRecordStatus {
	if in == nil {
		return nil
	}
	out := new(DNSRecordStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckSpec) DeepCopyInto(out *HealthCheckSpec) {
	*out = *in
	out.Interval = in.Interval
	if in.AdditionalHeadersRef != nil {
		in, out := &in.AdditionalHeadersRef, &out.AdditionalHeadersRef
		*out = new(AdditionalHeadersRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckSpec.
func (in *HealthCheckSpec) DeepCopy() *HealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(HealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckStatus) DeepCopyInto(out *HealthCheckStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = make([]HealthCheckStatusProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckStatus.
func (in *HealthCheckStatus) DeepCopy() *HealthCheckStatus {
	if in == nil {
		return nil
	}
	out := new(HealthCheckStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckStatusProbe) DeepCopyInto(out *HealthCheckStatusProbe) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckStatusProbe.
func (in *HealthCheckStatusProbe) DeepCopy() *HealthCheckStatusProbe {
	if in == nil {
		return nil
	}
	out := new(HealthCheckStatusProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderRef) DeepCopyInto(out *ProviderRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderRef.
func (in *ProviderRef) DeepCopy() *ProviderRef {
	if in == nil {
		return nil
	}
	out := new(ProviderRef)
	in.DeepCopyInto(out)
	return out
}
