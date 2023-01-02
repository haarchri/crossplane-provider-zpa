/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TimeoutRuleConditionsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// This signifies the various policy criteria.
	// +kubebuilder:validation:Optional
	Operands []TimeoutRuleConditionsOperandsObservation `json:"operands,omitempty" tf:"operands,omitempty"`
}

type TimeoutRuleConditionsOperandsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TimeoutRuleConditionsOperandsParameters struct {

	// +kubebuilder:validation:Optional
	IdpID *string `json:"idpId,omitempty" tf:"idp_id,omitempty"`

	// This signifies the key for the object type. String ID example: id
	// +kubebuilder:validation:Required
	Lhs *string `json:"lhs" tf:"lhs,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This is for specifying the policy critiera.
	// +kubebuilder:validation:Required
	ObjectType *string `json:"objectType" tf:"object_type,omitempty"`

	// This denotes the value for the given object type. Its value depends upon the key.
	// +kubebuilder:validation:Optional
	Rhs *string `json:"rhs,omitempty" tf:"rhs,omitempty"`

	// This denotes a list of values for the given object type. The value depend upon the key. If rhs is defined this list will be ignored
	// +kubebuilder:validation:Optional
	RhsList []*string `json:"rhsList,omitempty" tf:"rhs_list,omitempty"`
}

type TimeoutRuleConditionsParameters struct {

	// +kubebuilder:validation:Optional
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// This signifies the various policy criteria.
	// +kubebuilder:validation:Optional
	Operands []TimeoutRuleConditionsOperandsParameters `json:"operands,omitempty" tf:"operands,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`
}

type TimeoutRuleObservation struct {

	// This is for proviidng the set of conditions for the policy.
	// +kubebuilder:validation:Optional
	Conditions []TimeoutRuleConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TimeoutRuleParameters struct {

	// This is for providing the rule action.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// This field defines the description of the server.
	// +kubebuilder:validation:Optional
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// +kubebuilder:validation:Optional
	BypassDefaultRule *bool `json:"bypassDefaultRule,omitempty" tf:"bypass_default_rule,omitempty"`

	// This is for proviidng the set of conditions for the policy.
	// +kubebuilder:validation:Optional
	Conditions []TimeoutRuleConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// This is for providing a customer message for the user.
	// +kubebuilder:validation:Optional
	CustomMsg *string `json:"customMsg,omitempty" tf:"custom_msg,omitempty"`

	// This is for providing a customer message for the user.
	// +kubebuilder:validation:Optional
	DefaultRule *bool `json:"defaultRule,omitempty" tf:"default_rule,omitempty"`

	// This is the description of the access policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	LssDefaultRule *bool `json:"lssDefaultRule,omitempty" tf:"lss_default_rule,omitempty"`

	// This is the name of the policy.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	PolicySetID *string `json:"policySetId,omitempty" tf:"policy_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	ReauthDefaultRule *bool `json:"reauthDefaultRule,omitempty" tf:"reauth_default_rule,omitempty"`

	// +kubebuilder:validation:Optional
	ReauthIdleTimeout *string `json:"reauthIdleTimeout,omitempty" tf:"reauth_idle_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	ReauthTimeout *string `json:"reauthTimeout,omitempty" tf:"reauth_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	RuleOrder *string `json:"ruleOrder,omitempty" tf:"rule_order,omitempty"`

	// +kubebuilder:validation:Optional
	ZpnInspectionProfileID *string `json:"zpnInspectionProfileId,omitempty" tf:"zpn_inspection_profile_id,omitempty"`
}

// TimeoutRuleSpec defines the desired state of TimeoutRule
type TimeoutRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TimeoutRuleParameters `json:"forProvider"`
}

// TimeoutRuleStatus defines the observed state of TimeoutRule.
type TimeoutRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TimeoutRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TimeoutRule is the Schema for the TimeoutRules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type TimeoutRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TimeoutRuleSpec   `json:"spec"`
	Status            TimeoutRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TimeoutRuleList contains a list of TimeoutRules
type TimeoutRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TimeoutRule `json:"items"`
}

// Repository type metadata.
var (
	TimeoutRule_Kind             = "TimeoutRule"
	TimeoutRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TimeoutRule_Kind}.String()
	TimeoutRule_KindAPIVersion   = TimeoutRule_Kind + "." + CRDGroupVersion.String()
	TimeoutRule_GroupVersionKind = CRDGroupVersion.WithKind(TimeoutRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TimeoutRule{}, &TimeoutRuleList{})
}
