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

type EdgeGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the version profile.
	VersionProfileVisibilityScope *string `json:"versionProfileVisibilityScope,omitempty" tf:"version_profile_visibility_scope,omitempty"`
}

type EdgeGroupParameters struct {

	// +kubebuilder:validation:Optional
	CityCountry *string `json:"cityCountry,omitempty" tf:"city_country,omitempty"`

	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// Description of the Service Edge Group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this Service Edge Group is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable or disable public access for the Service Edge Group.
	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// Latitude for the Service Edge Group.
	// +kubebuilder:validation:Required
	Latitude *string `json:"latitude" tf:"latitude,omitempty"`

	// Location for the Service Edge Group.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Longitude for the Service Edge Group.
	// +kubebuilder:validation:Required
	Longitude *string `json:"longitude" tf:"longitude,omitempty"`

	// Name of the Service Edge Group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Whether the default version profile of the App Connector Group is applied or overridden.
	// +kubebuilder:validation:Optional
	OverrideVersionProfile *bool `json:"overrideVersionProfile,omitempty" tf:"override_version_profile,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceEdges []ServiceEdgesParameters `json:"serviceEdges,omitempty" tf:"service_edges,omitempty"`

	// +kubebuilder:validation:Optional
	TrustedNetworks []TrustedNetworksParameters `json:"trustedNetworks,omitempty" tf:"trusted_networks,omitempty"`

	// Service Edges in this group will attempt to update to a newer version of the software during this specified day.
	// +kubebuilder:validation:Optional
	UpgradeDay *string `json:"upgradeDay,omitempty" tf:"upgrade_day,omitempty"`

	// Service Edges in this group will attempt to update to a newer version of the software during this specified time.
	// +kubebuilder:validation:Optional
	UpgradeTimeInSecs *string `json:"upgradeTimeInSecs,omitempty" tf:"upgrade_time_in_secs,omitempty"`

	// ID of the version profile.
	// +kubebuilder:validation:Optional
	VersionProfileID *string `json:"versionProfileId,omitempty" tf:"version_profile_id,omitempty"`

	// ID of the version profile.
	// +kubebuilder:validation:Optional
	VersionProfileName *string `json:"versionProfileName,omitempty" tf:"version_profile_name,omitempty"`
}

type ServiceEdgesObservation struct {
}

type ServiceEdgesParameters struct {

	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type TrustedNetworksObservation struct {
}

type TrustedNetworksParameters struct {

	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

// EdgeGroupSpec defines the desired state of EdgeGroup
type EdgeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EdgeGroupParameters `json:"forProvider"`
}

// EdgeGroupStatus defines the observed state of EdgeGroup.
type EdgeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EdgeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeGroup is the Schema for the EdgeGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type EdgeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeGroupSpec   `json:"spec"`
	Status            EdgeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeGroupList contains a list of EdgeGroups
type EdgeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeGroup `json:"items"`
}

// Repository type metadata.
var (
	EdgeGroup_Kind             = "EdgeGroup"
	EdgeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EdgeGroup_Kind}.String()
	EdgeGroup_KindAPIVersion   = EdgeGroup_Kind + "." + CRDGroupVersion.String()
	EdgeGroup_GroupVersionKind = CRDGroupVersion.WithKind(EdgeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&EdgeGroup{}, &EdgeGroupList{})
}
