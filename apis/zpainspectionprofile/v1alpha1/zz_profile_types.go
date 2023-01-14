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

type ControlsInfoObservation struct {
}

type ControlsInfoParameters struct {

	// Control types. Supported Values: CUSTOM, PREDEFINED, ZSCALER
	// +kubebuilder:validation:Optional
	ControlType *string `json:"controlType,omitempty" tf:"control_type,omitempty"`

	// Control information counts Long
	// +kubebuilder:validation:Optional
	Count *string `json:"count,omitempty" tf:"count,omitempty"`
}

type CustomControlsObservation struct {
}

type CustomControlsParameters struct {

	// The action of the predefined control. Supported values: PASS, BLOCK and REDIRECT
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to REDIRECT.
	// +kubebuilder:validation:Optional
	ActionValue *string `json:"actionValue,omitempty" tf:"action_value,omitempty"`

	// ID of the predefined control
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type PredefinedControlsObservation struct {
}

type PredefinedControlsParameters struct {

	// The action of the predefined control. Supported values: PASS, BLOCK and REDIRECT
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to REDIRECT.
	// +kubebuilder:validation:Optional
	ActionValue *string `json:"actionValue,omitempty" tf:"action_value,omitempty"`

	// ID of the predefined control
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type ProfileObservation struct {

	// ID of the predefined control
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProfileParameters struct {

	// +kubebuilder:validation:Optional
	AssociateAllControls *bool `json:"associateAllControls,omitempty" tf:"associate_all_controls,omitempty"`

	// +kubebuilder:validation:Optional
	CommonGlobalOverrideActionsConfig map[string]*string `json:"commonGlobalOverrideActionsConfig,omitempty" tf:"common_global_override_actions_config,omitempty"`

	// Types for custom controls
	// +kubebuilder:validation:Optional
	ControlsInfo []ControlsInfoParameters `json:"controlsInfo,omitempty" tf:"controls_info,omitempty"`

	// Types for custom controls
	// +kubebuilder:validation:Optional
	CustomControls []CustomControlsParameters `json:"customControls,omitempty" tf:"custom_controls,omitempty"`

	// Description of the inspection profile.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	GlobalControlActions []*string `json:"globalControlActions,omitempty" tf:"global_control_actions,omitempty"`

	// +kubebuilder:validation:Optional
	IncarnationNumber *string `json:"incarnationNumber,omitempty" tf:"incarnation_number,omitempty"`

	// The name of the inspection profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
	// +kubebuilder:validation:Optional
	ParanoiaLevel *string `json:"paranoiaLevel,omitempty" tf:"paranoia_level,omitempty"`

	// The predefined controls. The default predefined control Preprocessors are mandatory and injected in the request by default. Individual predefined_controls can be set by using the data source data_source_zpa_predefined_controls or by group using the data source zpa_inspection_all_predefined_controls.
	// +kubebuilder:validation:Optional
	PredefinedControls []PredefinedControlsParameters `json:"predefinedControls,omitempty" tf:"predefined_controls,omitempty"`

	// = "OWASP_CRS/3.3.0"
	// +kubebuilder:validation:Optional
	PredefinedControlsVersion *string `json:"predefinedControlsVersion,omitempty" tf:"predefined_controls_version,omitempty"`
}

// ProfileSpec defines the desired state of Profile
type ProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProfileParameters `json:"forProvider"`
}

// ProfileStatus defines the observed state of Profile.
type ProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Profile is the Schema for the Profiles API. Creates and manages Inspection Profile in Zscaler Private Access cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProfileSpec   `json:"spec"`
	Status            ProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileList contains a list of Profiles
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

// Repository type metadata.
var (
	Profile_Kind             = "Profile"
	Profile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Profile_Kind}.String()
	Profile_KindAPIVersion   = Profile_Kind + "." + CRDGroupVersion.String()
	Profile_GroupVersionKind = CRDGroupVersion.WithKind(Profile_Kind)
)

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
