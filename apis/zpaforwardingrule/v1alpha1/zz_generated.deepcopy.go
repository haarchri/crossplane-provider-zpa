//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsInitParameters) DeepCopyInto(out *ConditionsInitParameters) {
	*out = *in
	if in.Negated != nil {
		in, out := &in.Negated, &out.Negated
		*out = new(bool)
		**out = **in
	}
	if in.Operands != nil {
		in, out := &in.Operands, &out.Operands
		*out = make([]OperandsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsInitParameters.
func (in *ConditionsInitParameters) DeepCopy() *ConditionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsObservation) DeepCopyInto(out *ConditionsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Negated != nil {
		in, out := &in.Negated, &out.Negated
		*out = new(bool)
		**out = **in
	}
	if in.Operands != nil {
		in, out := &in.Operands, &out.Operands
		*out = make([]OperandsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsObservation.
func (in *ConditionsObservation) DeepCopy() *ConditionsObservation {
	if in == nil {
		return nil
	}
	out := new(ConditionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionsParameters) DeepCopyInto(out *ConditionsParameters) {
	*out = *in
	if in.Negated != nil {
		in, out := &in.Negated, &out.Negated
		*out = new(bool)
		**out = **in
	}
	if in.Operands != nil {
		in, out := &in.Operands, &out.Operands
		*out = make([]OperandsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionsParameters.
func (in *ConditionsParameters) DeepCopy() *ConditionsParameters {
	if in == nil {
		return nil
	}
	out := new(ConditionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingRule) DeepCopyInto(out *ForwardingRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingRule.
func (in *ForwardingRule) DeepCopy() *ForwardingRule {
	if in == nil {
		return nil
	}
	out := new(ForwardingRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ForwardingRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingRuleInitParameters) DeepCopyInto(out *ForwardingRuleInitParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.ActionID != nil {
		in, out := &in.ActionID, &out.ActionID
		*out = new(string)
		**out = **in
	}
	if in.BypassDefaultRule != nil {
		in, out := &in.BypassDefaultRule, &out.BypassDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomMsg != nil {
		in, out := &in.CustomMsg, &out.CustomMsg
		*out = new(string)
		**out = **in
	}
	if in.DefaultRule != nil {
		in, out := &in.DefaultRule, &out.DefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LssDefaultRule != nil {
		in, out := &in.LssDefaultRule, &out.LssDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.PolicySetID != nil {
		in, out := &in.PolicySetID, &out.PolicySetID
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.ReauthDefaultRule != nil {
		in, out := &in.ReauthDefaultRule, &out.ReauthDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.ReauthIdleTimeout != nil {
		in, out := &in.ReauthIdleTimeout, &out.ReauthIdleTimeout
		*out = new(string)
		**out = **in
	}
	if in.ReauthTimeout != nil {
		in, out := &in.ReauthTimeout, &out.ReauthTimeout
		*out = new(string)
		**out = **in
	}
	if in.RuleOrder != nil {
		in, out := &in.RuleOrder, &out.RuleOrder
		*out = new(string)
		**out = **in
	}
	if in.ZpnCbiProfileID != nil {
		in, out := &in.ZpnCbiProfileID, &out.ZpnCbiProfileID
		*out = new(string)
		**out = **in
	}
	if in.ZpnInspectionProfileID != nil {
		in, out := &in.ZpnInspectionProfileID, &out.ZpnInspectionProfileID
		*out = new(string)
		**out = **in
	}
	if in.ZpnIsolationProfileID != nil {
		in, out := &in.ZpnIsolationProfileID, &out.ZpnIsolationProfileID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingRuleInitParameters.
func (in *ForwardingRuleInitParameters) DeepCopy() *ForwardingRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(ForwardingRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingRuleList) DeepCopyInto(out *ForwardingRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ForwardingRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingRuleList.
func (in *ForwardingRuleList) DeepCopy() *ForwardingRuleList {
	if in == nil {
		return nil
	}
	out := new(ForwardingRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ForwardingRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingRuleObservation) DeepCopyInto(out *ForwardingRuleObservation) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.ActionID != nil {
		in, out := &in.ActionID, &out.ActionID
		*out = new(string)
		**out = **in
	}
	if in.BypassDefaultRule != nil {
		in, out := &in.BypassDefaultRule, &out.BypassDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomMsg != nil {
		in, out := &in.CustomMsg, &out.CustomMsg
		*out = new(string)
		**out = **in
	}
	if in.DefaultRule != nil {
		in, out := &in.DefaultRule, &out.DefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LssDefaultRule != nil {
		in, out := &in.LssDefaultRule, &out.LssDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.PolicySetID != nil {
		in, out := &in.PolicySetID, &out.PolicySetID
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.ReauthDefaultRule != nil {
		in, out := &in.ReauthDefaultRule, &out.ReauthDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.ReauthIdleTimeout != nil {
		in, out := &in.ReauthIdleTimeout, &out.ReauthIdleTimeout
		*out = new(string)
		**out = **in
	}
	if in.ReauthTimeout != nil {
		in, out := &in.ReauthTimeout, &out.ReauthTimeout
		*out = new(string)
		**out = **in
	}
	if in.RuleOrder != nil {
		in, out := &in.RuleOrder, &out.RuleOrder
		*out = new(string)
		**out = **in
	}
	if in.ZpnCbiProfileID != nil {
		in, out := &in.ZpnCbiProfileID, &out.ZpnCbiProfileID
		*out = new(string)
		**out = **in
	}
	if in.ZpnInspectionProfileID != nil {
		in, out := &in.ZpnInspectionProfileID, &out.ZpnInspectionProfileID
		*out = new(string)
		**out = **in
	}
	if in.ZpnIsolationProfileID != nil {
		in, out := &in.ZpnIsolationProfileID, &out.ZpnIsolationProfileID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingRuleObservation.
func (in *ForwardingRuleObservation) DeepCopy() *ForwardingRuleObservation {
	if in == nil {
		return nil
	}
	out := new(ForwardingRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingRuleParameters) DeepCopyInto(out *ForwardingRuleParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.ActionID != nil {
		in, out := &in.ActionID, &out.ActionID
		*out = new(string)
		**out = **in
	}
	if in.BypassDefaultRule != nil {
		in, out := &in.BypassDefaultRule, &out.BypassDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ConditionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomMsg != nil {
		in, out := &in.CustomMsg, &out.CustomMsg
		*out = new(string)
		**out = **in
	}
	if in.DefaultRule != nil {
		in, out := &in.DefaultRule, &out.DefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LssDefaultRule != nil {
		in, out := &in.LssDefaultRule, &out.LssDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.PolicySetID != nil {
		in, out := &in.PolicySetID, &out.PolicySetID
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.ReauthDefaultRule != nil {
		in, out := &in.ReauthDefaultRule, &out.ReauthDefaultRule
		*out = new(bool)
		**out = **in
	}
	if in.ReauthIdleTimeout != nil {
		in, out := &in.ReauthIdleTimeout, &out.ReauthIdleTimeout
		*out = new(string)
		**out = **in
	}
	if in.ReauthTimeout != nil {
		in, out := &in.ReauthTimeout, &out.ReauthTimeout
		*out = new(string)
		**out = **in
	}
	if in.RuleOrder != nil {
		in, out := &in.RuleOrder, &out.RuleOrder
		*out = new(string)
		**out = **in
	}
	if in.ZpnCbiProfileID != nil {
		in, out := &in.ZpnCbiProfileID, &out.ZpnCbiProfileID
		*out = new(string)
		**out = **in
	}
	if in.ZpnInspectionProfileID != nil {
		in, out := &in.ZpnInspectionProfileID, &out.ZpnInspectionProfileID
		*out = new(string)
		**out = **in
	}
	if in.ZpnIsolationProfileID != nil {
		in, out := &in.ZpnIsolationProfileID, &out.ZpnIsolationProfileID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingRuleParameters.
func (in *ForwardingRuleParameters) DeepCopy() *ForwardingRuleParameters {
	if in == nil {
		return nil
	}
	out := new(ForwardingRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingRuleSpec) DeepCopyInto(out *ForwardingRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingRuleSpec.
func (in *ForwardingRuleSpec) DeepCopy() *ForwardingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ForwardingRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingRuleStatus) DeepCopyInto(out *ForwardingRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingRuleStatus.
func (in *ForwardingRuleStatus) DeepCopy() *ForwardingRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ForwardingRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperandsInitParameters) DeepCopyInto(out *OperandsInitParameters) {
	*out = *in
	if in.IdpID != nil {
		in, out := &in.IdpID, &out.IdpID
		*out = new(string)
		**out = **in
	}
	if in.Lhs != nil {
		in, out := &in.Lhs, &out.Lhs
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Rhs != nil {
		in, out := &in.Rhs, &out.Rhs
		*out = new(string)
		**out = **in
	}
	if in.RhsList != nil {
		in, out := &in.RhsList, &out.RhsList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperandsInitParameters.
func (in *OperandsInitParameters) DeepCopy() *OperandsInitParameters {
	if in == nil {
		return nil
	}
	out := new(OperandsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperandsObservation) DeepCopyInto(out *OperandsObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IdpID != nil {
		in, out := &in.IdpID, &out.IdpID
		*out = new(string)
		**out = **in
	}
	if in.Lhs != nil {
		in, out := &in.Lhs, &out.Lhs
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Rhs != nil {
		in, out := &in.Rhs, &out.Rhs
		*out = new(string)
		**out = **in
	}
	if in.RhsList != nil {
		in, out := &in.RhsList, &out.RhsList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperandsObservation.
func (in *OperandsObservation) DeepCopy() *OperandsObservation {
	if in == nil {
		return nil
	}
	out := new(OperandsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperandsParameters) DeepCopyInto(out *OperandsParameters) {
	*out = *in
	if in.IdpID != nil {
		in, out := &in.IdpID, &out.IdpID
		*out = new(string)
		**out = **in
	}
	if in.Lhs != nil {
		in, out := &in.Lhs, &out.Lhs
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObjectType != nil {
		in, out := &in.ObjectType, &out.ObjectType
		*out = new(string)
		**out = **in
	}
	if in.Rhs != nil {
		in, out := &in.Rhs, &out.Rhs
		*out = new(string)
		**out = **in
	}
	if in.RhsList != nil {
		in, out := &in.RhsList, &out.RhsList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperandsParameters.
func (in *OperandsParameters) DeepCopy() *OperandsParameters {
	if in == nil {
		return nil
	}
	out := new(OperandsParameters)
	in.DeepCopyInto(out)
	return out
}
