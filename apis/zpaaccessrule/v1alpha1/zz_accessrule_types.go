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

type AccessRuleInitParameters struct {

	// This is for providing the rule action. Supported values: ALLOW, DENY
	// This is for providing the rule action.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ID of an app connector group resource
	// This field defines the description of the server.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// List of app-connector IDs.
	AppConnectorGroups []AppConnectorGroupsInitParameters `json:"appConnectorGroups,omitempty" tf:"app_connector_groups,omitempty"`

	// List of the server group IDs.
	AppServerGroups []AppServerGroupsInitParameters `json:"appServerGroups,omitempty" tf:"app_server_groups,omitempty"`

	BypassDefaultRule *bool `json:"bypassDefaultRule,omitempty" tf:"bypass_default_rule,omitempty"`

	// This is for proviidng the set of conditions for the policy.
	Conditions []ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// This is for providing a customer message for the user.
	// This is for providing a customer message for the user.
	CustomMsg *string `json:"customMsg,omitempty" tf:"custom_msg,omitempty"`

	// This is for providing a customer message for the user.
	DefaultRule *bool `json:"defaultRule,omitempty" tf:"default_rule,omitempty"`

	// This is the description of the access policy rule.
	// This is the description of the access policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	LssDefaultRule *bool `json:"lssDefaultRule,omitempty" tf:"lss_default_rule,omitempty"`

	// This is the name of the policy rule.
	// This is the name of the policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Supported values: AND, OR
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Use zpa_policy_type data source to retrieve the necessary policy Set ID policy_set_id
	PolicySetID *string `json:"policySetId,omitempty" tf:"policy_set_id,omitempty"`

	// Supported values: ACCESS_POLICY or GLOBAL_POLICY
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	ReauthDefaultRule *bool `json:"reauthDefaultRule,omitempty" tf:"reauth_default_rule,omitempty"`

	ReauthIdleTimeout *string `json:"reauthIdleTimeout,omitempty" tf:"reauth_idle_timeout,omitempty"`

	ReauthTimeout *string `json:"reauthTimeout,omitempty" tf:"reauth_timeout,omitempty"`

	RuleOrder *string `json:"ruleOrder,omitempty" tf:"rule_order,omitempty"`

	// The ID of an app connector group resource
	ZpnCbiProfileID *string `json:"zpnCbiProfileId,omitempty" tf:"zpn_cbi_profile_id,omitempty"`

	// The ID of an app connector group resource
	ZpnInspectionProfileID *string `json:"zpnInspectionProfileId,omitempty" tf:"zpn_inspection_profile_id,omitempty"`

	// The ID of an app connector group resource
	ZpnIsolationProfileID *string `json:"zpnIsolationProfileId,omitempty" tf:"zpn_isolation_profile_id,omitempty"`
}

type AccessRuleObservation struct {

	// This is for providing the rule action. Supported values: ALLOW, DENY
	// This is for providing the rule action.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The ID of an app connector group resource
	// This field defines the description of the server.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// List of app-connector IDs.
	AppConnectorGroups []AppConnectorGroupsObservation `json:"appConnectorGroups,omitempty" tf:"app_connector_groups,omitempty"`

	// List of the server group IDs.
	AppServerGroups []AppServerGroupsObservation `json:"appServerGroups,omitempty" tf:"app_server_groups,omitempty"`

	BypassDefaultRule *bool `json:"bypassDefaultRule,omitempty" tf:"bypass_default_rule,omitempty"`

	// This is for proviidng the set of conditions for the policy.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// This is for providing a customer message for the user.
	// This is for providing a customer message for the user.
	CustomMsg *string `json:"customMsg,omitempty" tf:"custom_msg,omitempty"`

	// This is for providing a customer message for the user.
	DefaultRule *bool `json:"defaultRule,omitempty" tf:"default_rule,omitempty"`

	// This is the description of the access policy rule.
	// This is the description of the access policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of an app connector group resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LssDefaultRule *bool `json:"lssDefaultRule,omitempty" tf:"lss_default_rule,omitempty"`

	// This is the name of the policy rule.
	// This is the name of the policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Supported values: AND, OR
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Use zpa_policy_type data source to retrieve the necessary policy Set ID policy_set_id
	PolicySetID *string `json:"policySetId,omitempty" tf:"policy_set_id,omitempty"`

	// Supported values: ACCESS_POLICY or GLOBAL_POLICY
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	ReauthDefaultRule *bool `json:"reauthDefaultRule,omitempty" tf:"reauth_default_rule,omitempty"`

	ReauthIdleTimeout *string `json:"reauthIdleTimeout,omitempty" tf:"reauth_idle_timeout,omitempty"`

	ReauthTimeout *string `json:"reauthTimeout,omitempty" tf:"reauth_timeout,omitempty"`

	RuleOrder *string `json:"ruleOrder,omitempty" tf:"rule_order,omitempty"`

	// The ID of an app connector group resource
	ZpnCbiProfileID *string `json:"zpnCbiProfileId,omitempty" tf:"zpn_cbi_profile_id,omitempty"`

	// The ID of an app connector group resource
	ZpnInspectionProfileID *string `json:"zpnInspectionProfileId,omitempty" tf:"zpn_inspection_profile_id,omitempty"`

	// The ID of an app connector group resource
	ZpnIsolationProfileID *string `json:"zpnIsolationProfileId,omitempty" tf:"zpn_isolation_profile_id,omitempty"`
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
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

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
	ZpnCbiProfileID *string `json:"zpnCbiProfileId,omitempty" tf:"zpn_cbi_profile_id,omitempty"`

	// The ID of an app connector group resource
	// +kubebuilder:validation:Optional
	ZpnInspectionProfileID *string `json:"zpnInspectionProfileId,omitempty" tf:"zpn_inspection_profile_id,omitempty"`

	// The ID of an app connector group resource
	// +kubebuilder:validation:Optional
	ZpnIsolationProfileID *string `json:"zpnIsolationProfileId,omitempty" tf:"zpn_isolation_profile_id,omitempty"`
}

type AppConnectorGroupsInitParameters struct {

	// The ID of an app connector group resource
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppConnectorGroupsObservation struct {

	// The ID of an app connector group resource
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppConnectorGroupsParameters struct {

	// The ID of an app connector group resource
	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppServerGroupsInitParameters struct {

	// The ID of an app connector group resource
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppServerGroupsObservation struct {

	// The ID of an app connector group resource
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppServerGroupsParameters struct {

	// The ID of an app connector group resource
	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConditionsInitParameters struct {

	// Supported values: true or false
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// Operands block must be repeated if multiple per object_type conditions are to be added to the rule.
	// This signifies the various policy criteria.
	Operands []OperandsInitParameters `json:"operands,omitempty" tf:"operands,omitempty"`

	// Supported values: AND, OR
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type ConditionsObservation struct {

	// The ID of an app connector group resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Supported values: true or false
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// Operands block must be repeated if multiple per object_type conditions are to be added to the rule.
	// This signifies the various policy criteria.
	Operands []OperandsObservation `json:"operands,omitempty" tf:"operands,omitempty"`

	// Supported values: AND, OR
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
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
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type OperandsInitParameters struct {
	IdpID *string `json:"idpId,omitempty" tf:"idp_id,omitempty"`

	// Trusted Network (network_id) required when object_type = "TRUSTED_NETWORK". Use zpa_trusted_network data source to retrieve the network_id
	// This signifies the key for the object type. String ID example: id
	Lhs *string `json:"lhs,omitempty" tf:"lhs,omitempty"`

	// This is the name of the policy rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This is for specifying the policy critiera. For posture profile the supported value is:  TRUSTED_NETWORK
	// This is for specifying the policy critiera.
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// Required when object_type = "TRUSTED_NETWORK". Supported values are:
	// This denotes the value for the given object type. Its value depends upon the key.
	Rhs *string `json:"rhs,omitempty" tf:"rhs,omitempty"`

	// This denotes a list of values for the given object type. The value depend upon the key. If rhs is defined this list will be ignored
	RhsList []*string `json:"rhsList,omitempty" tf:"rhs_list,omitempty"`
}

type OperandsObservation struct {

	// The ID of an app connector group resource
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IdpID *string `json:"idpId,omitempty" tf:"idp_id,omitempty"`

	// Trusted Network (network_id) required when object_type = "TRUSTED_NETWORK". Use zpa_trusted_network data source to retrieve the network_id
	// This signifies the key for the object type. String ID example: id
	Lhs *string `json:"lhs,omitempty" tf:"lhs,omitempty"`

	// This is the name of the policy rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This is for specifying the policy critiera. For posture profile the supported value is:  TRUSTED_NETWORK
	// This is for specifying the policy critiera.
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

	// Required when object_type = "TRUSTED_NETWORK". Supported values are:
	// This denotes the value for the given object type. Its value depends upon the key.
	Rhs *string `json:"rhs,omitempty" tf:"rhs,omitempty"`

	// This denotes a list of values for the given object type. The value depend upon the key. If rhs is defined this list will be ignored
	RhsList []*string `json:"rhsList,omitempty" tf:"rhs_list,omitempty"`
}

type OperandsParameters struct {

	// +kubebuilder:validation:Optional
	IdpID *string `json:"idpId,omitempty" tf:"idp_id,omitempty"`

	// Trusted Network (network_id) required when object_type = "TRUSTED_NETWORK". Use zpa_trusted_network data source to retrieve the network_id
	// This signifies the key for the object type. String ID example: id
	// +kubebuilder:validation:Optional
	Lhs *string `json:"lhs,omitempty" tf:"lhs,omitempty"`

	// This is the name of the policy rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// This is for specifying the policy critiera. For posture profile the supported value is:  TRUSTED_NETWORK
	// This is for specifying the policy critiera.
	// +kubebuilder:validation:Optional
	ObjectType *string `json:"objectType,omitempty" tf:"object_type,omitempty"`

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
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AccessRuleInitParameters `json:"initProvider,omitempty"`
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
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   AccessRuleSpec   `json:"spec"`
	Status AccessRuleStatus `json:"status,omitempty"`
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
