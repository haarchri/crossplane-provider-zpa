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

	// +kubebuilder:validation:Optional
	ControlType *string `json:"controlType,omitempty" tf:"control_type,omitempty"`

	// +kubebuilder:validation:Optional
	Count *string `json:"count,omitempty" tf:"count,omitempty"`
}

type CustomControlsObservation struct {
}

type CustomControlsParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	ActionValue *string `json:"actionValue,omitempty" tf:"action_value,omitempty"`

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type PredefinedControlsObservation struct {
}

type PredefinedControlsParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	ActionValue *string `json:"actionValue,omitempty" tf:"action_value,omitempty"`

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type ProfileObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProfileParameters struct {

	// +kubebuilder:validation:Optional
	AssociateAllControls *bool `json:"associateAllControls,omitempty" tf:"associate_all_controls,omitempty"`

	// +kubebuilder:validation:Optional
	CommonGlobalOverrideActionsConfig map[string]*string `json:"commonGlobalOverrideActionsConfig,omitempty" tf:"common_global_override_actions_config,omitempty"`

	// +kubebuilder:validation:Optional
	ControlsInfo []ControlsInfoParameters `json:"controlsInfo,omitempty" tf:"controls_info,omitempty"`

	// +kubebuilder:validation:Optional
	CustomControls []CustomControlsParameters `json:"customControls,omitempty" tf:"custom_controls,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	GlobalControlActions []*string `json:"globalControlActions,omitempty" tf:"global_control_actions,omitempty"`

	// +kubebuilder:validation:Optional
	IncarnationNumber *string `json:"incarnationNumber,omitempty" tf:"incarnation_number,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ParanoiaLevel *string `json:"paranoiaLevel,omitempty" tf:"paranoia_level,omitempty"`

	// +kubebuilder:validation:Optional
	PredefinedControls []PredefinedControlsParameters `json:"predefinedControls,omitempty" tf:"predefined_controls,omitempty"`

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

// Profile is the Schema for the Profiles API. <no value>
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
