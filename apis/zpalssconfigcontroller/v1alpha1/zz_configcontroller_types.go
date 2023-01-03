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

type ConditionsObservation struct {
}

type ConditionsParameters struct {

	// +kubebuilder:validation:Optional
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// This signifies the various policy criteria.
	// +kubebuilder:validation:Optional
	Operands []OperandsParameters `json:"operands,omitempty" tf:"operands,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`
}

type ConfigControllerObservation struct {

	// +kubebuilder:validation:Optional
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PolicyRuleID *string `json:"policyRuleId,omitempty" tf:"policy_rule_id,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyRuleResource []PolicyRuleResourceObservation `json:"policyRuleResource,omitempty" tf:"policy_rule_resource,omitempty"`
}

type ConfigControllerParameters struct {

	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// App Connector Group(s) to be added to the LSS configuration
	// +kubebuilder:validation:Optional
	ConnectorGroups []ConnectorGroupsParameters `json:"connectorGroups,omitempty" tf:"connector_groups,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyRuleResource []PolicyRuleResourceParameters `json:"policyRuleResource,omitempty" tf:"policy_rule_resource,omitempty"`
}

type ConfigObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConfigParameters struct {

	// +kubebuilder:validation:Optional
	AuditMessage *string `json:"auditMessage,omitempty" tf:"audit_message,omitempty"`

	// Description of the LSS configuration
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this LSS configuration is enabled or not. Supported values: true, false
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Filter for the LSS configuration. Format given by the following API to get status codes: /mgmtconfig/v2/admin/lssConfig/statusCodes
	// +kubebuilder:validation:Optional
	Filter []*string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Format of the log type. Format given by the following API to get formats: /mgmtconfig/v2/admin/lssConfig/logType/formats
	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`

	// Host of the LSS configuration
	// +kubebuilder:validation:Required
	LssHost *string `json:"lssHost" tf:"lss_host,omitempty"`

	// Port of the LSS configuration
	// +kubebuilder:validation:Required
	LssPort *string `json:"lssPort" tf:"lss_port,omitempty"`

	// Name of the LSS configuration
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Log type of the LSS configuration
	// +kubebuilder:validation:Required
	SourceLogType *string `json:"sourceLogType" tf:"source_log_type,omitempty"`

	// +kubebuilder:validation:Optional
	UseTLS *bool `json:"useTls,omitempty" tf:"use_tls,omitempty"`
}

type ConnectorGroupsObservation struct {
}

type ConnectorGroupsParameters struct {

	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type OperandsObservation struct {
}

type OperandsParameters struct {

	// This is for specifying the policy critiera.
	// +kubebuilder:validation:Required
	ObjectType *string `json:"objectType" tf:"object_type,omitempty"`

	// This denotes a list of values for the given object type. The value depend upon the key. If rhs is defined this list will be ignored
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type PolicyRuleResourceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyRuleResourceParameters struct {

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
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

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

// ConfigControllerSpec defines the desired state of ConfigController
type ConfigControllerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigControllerParameters `json:"forProvider"`
}

// ConfigControllerStatus defines the observed state of ConfigController.
type ConfigControllerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigControllerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigController is the Schema for the ConfigControllers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type ConfigController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigControllerSpec   `json:"spec"`
	Status            ConfigControllerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigControllerList contains a list of ConfigControllers
type ConfigControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigController `json:"items"`
}

// Repository type metadata.
var (
	ConfigController_Kind             = "ConfigController"
	ConfigController_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConfigController_Kind}.String()
	ConfigController_KindAPIVersion   = ConfigController_Kind + "." + CRDGroupVersion.String()
	ConfigController_GroupVersionKind = CRDGroupVersion.WithKind(ConfigController_Kind)
)

func init() {
	SchemeBuilder.Register(&ConfigController{}, &ConfigControllerList{})
}