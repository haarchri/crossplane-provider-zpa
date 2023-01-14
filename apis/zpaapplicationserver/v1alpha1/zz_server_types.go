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

type ServerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServerParameters struct {

	// Address. The address of the application server to be exported.
	// This field defines the domain or IP address of the server.
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// References to Group in zpaservergroup to populate appServerGroupIds.
	// +kubebuilder:validation:Optional
	AppServerGroupIDRefs []v1.Reference `json:"appServerGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of Group in zpaservergroup to populate appServerGroupIds.
	// +kubebuilder:validation:Optional
	AppServerGroupIDSelector *v1.Selector `json:"appServerGroupIdSelector,omitempty" tf:"-"`

	// This field defines the list of server group IDs.
	// This field defines the list of server groups IDs.
	// +crossplane:generate:reference:type=github.com/zscaler/crossplane-provider-zpa/apis/zpaservergroup/v1alpha1.Group
	// +crossplane:generate:reference:refFieldName=AppServerGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=AppServerGroupIDSelector
	// +kubebuilder:validation:Optional
	AppServerGroupIds []*string `json:"appServerGroupIds,omitempty" tf:"app_server_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigSpace *string `json:"configSpace,omitempty" tf:"config_space,omitempty"`

	// This field defines the description of the server.
	// This field defines the description of the server.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// This field defines the status of the server.
	// This field defines the status of the server.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Name. The name of the application server to be exported.
	// This field defines the name of the server.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerParameters `json:"forProvider"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API. Creates and manages ZPA Application Servers.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	Server_Kind             = "Server"
	Server_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Server_Kind}.String()
	Server_KindAPIVersion   = Server_Kind + "." + CRDGroupVersion.String()
	Server_GroupVersionKind = CRDGroupVersion.WithKind(Server_Kind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
