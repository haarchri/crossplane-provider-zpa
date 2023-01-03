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

type SegmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SegmentParameters struct {

	// Indicates whether users can bypass ZPA to access applications.
	// +kubebuilder:validation:Optional
	BypassType *string `json:"bypassType,omitempty" tf:"bypass_type,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigSpace *string `json:"configSpace,omitempty" tf:"config_space,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultIdleTimeout *string `json:"defaultIdleTimeout,omitempty" tf:"default_idle_timeout,omitempty"`

	// Description of the application.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of domains and IPs.
	// +kubebuilder:validation:Required
	DomainNames []*string `json:"domainNames" tf:"domain_names,omitempty"`

	// Whether Double Encryption is enabled or disabled for the app.
	// +kubebuilder:validation:Optional
	DoubleEncrypt *bool `json:"doubleEncrypt,omitempty" tf:"double_encrypt,omitempty"`

	// Whether this application is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	// +kubebuilder:validation:Optional
	HealthReporting *string `json:"healthReporting,omitempty" tf:"health_reporting,omitempty"`

	// +kubebuilder:validation:Optional
	IPAnchored *bool `json:"ipAnchored,omitempty" tf:"ip_anchored,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpAccessType *string `json:"icmpAccessType,omitempty" tf:"icmp_access_type,omitempty"`

	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	// +kubebuilder:validation:Optional
	IsCnameEnabled *bool `json:"isCnameEnabled,omitempty" tf:"is_cname_enabled,omitempty"`

	// Name of the application.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PassiveHealthEnabled *bool `json:"passiveHealthEnabled,omitempty" tf:"passive_health_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	SegmentGroupID *string `json:"segmentGroupId,omitempty" tf:"segment_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SegmentGroupName *string `json:"segmentGroupName,omitempty" tf:"segment_group_name,omitempty"`

	// List of the server group IDs.
	// +kubebuilder:validation:Required
	ServerGroups []ServerGroupsParameters `json:"serverGroups" tf:"server_groups,omitempty"`

	// tcp port range
	// +kubebuilder:validation:Optional
	TCPPortRange []TCPPortRangeParameters `json:"tcpPortRange,omitempty" tf:"tcp_port_range,omitempty"`

	// TCP port ranges used to access the app.
	// +kubebuilder:validation:Optional
	TCPPortRanges []*string `json:"tcpPortRanges,omitempty" tf:"tcp_port_ranges,omitempty"`

	// udp port range
	// +kubebuilder:validation:Optional
	UDPPortRange []UDPPortRangeParameters `json:"udpPortRange,omitempty" tf:"udp_port_range,omitempty"`

	// UDP port ranges used to access the app.
	// +kubebuilder:validation:Optional
	UDPPortRanges []*string `json:"udpPortRanges,omitempty" tf:"udp_port_ranges,omitempty"`
}

type ServerGroupsObservation struct {
}

type ServerGroupsParameters struct {

	// +kubebuilder:validation:Required
	ID []*string `json:"id" tf:"id,omitempty"`
}

type TCPPortRangeObservation struct {
}

type TCPPortRangeParameters struct {

	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from"`

	// +kubebuilder:validation:Optional
	To *string `json:"to,omitempty" tf:"to"`
}

type UDPPortRangeObservation struct {
}

type UDPPortRangeParameters struct {

	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from"`

	// +kubebuilder:validation:Optional
	To *string `json:"to,omitempty" tf:"to"`
}

// SegmentSpec defines the desired state of Segment
type SegmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SegmentParameters `json:"forProvider"`
}

// SegmentStatus defines the observed state of Segment.
type SegmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SegmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Segment is the Schema for the Segments API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type Segment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SegmentSpec   `json:"spec"`
	Status            SegmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SegmentList contains a list of Segments
type SegmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Segment `json:"items"`
}

// Repository type metadata.
var (
	Segment_Kind             = "Segment"
	Segment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Segment_Kind}.String()
	Segment_KindAPIVersion   = Segment_Kind + "." + CRDGroupVersion.String()
	Segment_GroupVersionKind = CRDGroupVersion.WithKind(Segment_Kind)
)

func init() {
	SchemeBuilder.Register(&Segment{}, &SegmentList{})
}