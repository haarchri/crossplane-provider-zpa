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

type AccessRuleObservation struct {

	// This is for proviidng the set of conditions for the policy.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// The ID of an app connector group resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AccessRuleParameters struct {

	// This is for providing the rule action. Supported values: ALLOW, DENY
	// This is for providing the rule action.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ID of an app connector group resource
	// This field defines the description of the server.
	// +kubebuilder:validation:Optional
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// List of app-connector IDs.
	// +kubebuilder:validation:Optional
	AppConnectorGroups []AppConnectorGroupsParameters `json:"appConnectorGroups,omitempty" tf:"app_connector_groups,omitempty"`

	// List of the server group IDs.
	// +kubebuilder:validation:Optional
	AppServerGroups []AppServerGroupsParameters `json:"appServerGroups,omitempty" tf:"app_server_groups,omitempty"`

	// +kubebuilder:validation:Optional
	BypassDefaultRule *bool `json:"bypassDefaultRule,omitempty" tf:"bypass_default_rule,omitempty"`

	// This is for proviidng the set of conditions for the policy.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// This is for providing a customer message for the user.
	// This is for providing a customer message for the user.
	// +kubebuilder:validation:Optional
	CustomMsg *string `json:"customMsg,omitempty" tf:"custom_msg,omitempty"`

	// This is for providing a customer message for the user.
	// +kubebuilder:validation:Optional
	DefaultRule *bool `json:"defaultRule,omitempty" tf:"default_rule,omitempty"`

	// This is the description of the access policy rule.
	// This is the description of the access policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	LssDefaultRule *bool `json:"lssDefaultRule,omitempty" tf:"lss_default_rule,omitempty"`

	// This is the name of the policy rule.
	// This is the name of the policy.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Supported values: AND, OR
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Use zpa_policy_type data source to retrieve the necessary policy Set ID policy_set_id
	// +kubebuilder:validation:Optional
	PolicySetID *string `json:"policySetId,omitempty" tf:"policy_set_id,omitempty"`

	// Supported values: ACCESS_POLICY or GLOBAL_POLICY
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

	// The ID of an app connector group resource
	// +kubebuilder:validation:Optional
	ZpnInspectionProfileID *string `json:"zpnInspectionProfileId,omitempty" tf:"zpn_inspection_profile_id,omitempty"`
}

type AppConnectorGroupsObservation struct {
}

type AppConnectorGroupsParameters struct {

	// The ID of an app connector group resource
	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppServerGroupsObservation struct {
}

type AppServerGroupsParameters struct {

	// The ID of an app connector group resource
	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConditionsObservation struct {

	// The ID of an app connector group resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Operands block must be repeated if multiple per object_type conditions are to be added to the rule.
	// This signifies the various policy criteria.
	// +kubebuilder:validation:Optional
	Operands []OperandsObservation `json:"operands,omitempty" tf:"operands,omitempty"`
}

type ConditionsParameters struct {

	// Supported values: true or false
	// +kubebuilder:validation:Optional
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// Operands block must be repeated if multiple per object_type conditions are to be added to the rule.
	// This signifies the various policy criteria.
	// +kubebuilder:validation:Optional
	Operands []OperandsParameters `json:"operands,omitempty" tf:"operands,omitempty"`

	// Supported values: AND, OR
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`
}

type OperandsObservation struct {

	// The ID of an app connector group resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OperandsParameters struct {

	// +kubebuilder:validation:Optional
	IdpID *string `json:"idpId,omitempty" tf:"idp_id,omitempty"`

	// Trusted Network (network_id) required when object_type = "TRUSTED_NETWORK". Use zpa_trusted_network data source to retrieve the network_id
	// This signifies the key for the object type. String ID example: id
	// +kubebuilder:validation:Required
	Lhs *string `json:"lhs" tf:"lhs,omitempty"`

	// This is the name of the policy rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This is for specifying the policy critiera. For posture profile the supported value is:  TRUSTED_NETWORK
	// This is for specifying the policy critiera.
	// +kubebuilder:validation:Required
	ObjectType *string `json:"objectType" tf:"object_type,omitempty"`

	// Required when object_type = "TRUSTED_NETWORK". Supported values are:
	// This denotes the value for the given object type. Its value depends upon the key.
	// +kubebuilder:validation:Optional
	Rhs *string `json:"rhs,omitempty" tf:"rhs,omitempty"`

	// This denotes a list of values for the given object type. The value depend upon the key. If rhs is defined this list will be ignored
	// +kubebuilder:validation:Optional
	RhsList []*string `json:"rhsList,omitempty" tf:"rhs_list,omitempty"`
}

// AccessRuleSpec defines the desired state of AccessRule
type AccessRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessRuleParameters `json:"forProvider"`
}

// AccessRuleStatus defines the observed state of AccessRule.
type AccessRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessRule is the Schema for the AccessRules API. Creates and manages ZPA Policy Access Rule with Trusted Networks conditions.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type AccessRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessRuleSpec   `json:"spec"`
	Status            AccessRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessRuleList contains a list of AccessRules
type AccessRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessRule `json:"items"`
}

// Repository type metadata.
var (
	AccessRule_Kind             = "AccessRule"
	AccessRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessRule_Kind}.String()
	AccessRule_KindAPIVersion   = AccessRule_Kind + "." + CRDGroupVersion.String()
	AccessRule_GroupVersionKind = CRDGroupVersion.WithKind(AccessRule_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessRule{}, &AccessRuleList{})
}
