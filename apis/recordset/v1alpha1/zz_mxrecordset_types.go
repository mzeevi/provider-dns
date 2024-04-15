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

type MXRecordSetInitParameters struct {

	// (Block Set) Can be specified multiple times for each MX record. (see below for nested schema)
	// Can be specified multiple times for each MX record.
	Mx []MxInitParameters `json:"mx,omitempty" tf:"mx,omitempty"`

	// (String) The name of the record set. The zone argument will be appended to this value to create the full record path.
	// The name of the record set. The `zone` argument will be appended to this value to create the full record path.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) The TTL of the record set. Defaults to 3600.
	// The TTL of the record set. Defaults to `3600`.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type MXRecordSetObservation struct {

	// (String) Always set to the fully qualified domain name of the record set
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block Set) Can be specified multiple times for each MX record. (see below for nested schema)
	// Can be specified multiple times for each MX record.
	Mx []MxObservation `json:"mx,omitempty" tf:"mx,omitempty"`

	// (String) The name of the record set. The zone argument will be appended to this value to create the full record path.
	// The name of the record set. The `zone` argument will be appended to this value to create the full record path.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) The TTL of the record set. Defaults to 3600.
	// The TTL of the record set. Defaults to `3600`.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (String) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	// DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type MXRecordSetParameters struct {

	// (Block Set) Can be specified multiple times for each MX record. (see below for nested schema)
	// Can be specified multiple times for each MX record.
	// +kubebuilder:validation:Optional
	Mx []MxParameters `json:"mx,omitempty" tf:"mx,omitempty"`

	// (String) The name of the record set. The zone argument will be appended to this value to create the full record path.
	// The name of the record set. The `zone` argument will be appended to this value to create the full record path.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Number) The TTL of the record set. Defaults to 3600.
	// The TTL of the record set. Defaults to `3600`.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (String) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	// DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type MxInitParameters struct {

	// (String) The FQDN of the mail exchange, include the trailing dot.
	// The FQDN of the mail exchange, include the trailing dot.
	Exchange *string `json:"exchange,omitempty" tf:"exchange,omitempty"`

	// (Number) The preference for the record.
	// The preference for the record.
	Preference *float64 `json:"preference,omitempty" tf:"preference,omitempty"`
}

type MxObservation struct {

	// (String) The FQDN of the mail exchange, include the trailing dot.
	// The FQDN of the mail exchange, include the trailing dot.
	Exchange *string `json:"exchange,omitempty" tf:"exchange,omitempty"`

	// (Number) The preference for the record.
	// The preference for the record.
	Preference *float64 `json:"preference,omitempty" tf:"preference,omitempty"`
}

type MxParameters struct {

	// (String) The FQDN of the mail exchange, include the trailing dot.
	// The FQDN of the mail exchange, include the trailing dot.
	// +kubebuilder:validation:Optional
	Exchange *string `json:"exchange" tf:"exchange,omitempty"`

	// (Number) The preference for the record.
	// The preference for the record.
	// +kubebuilder:validation:Optional
	Preference *float64 `json:"preference" tf:"preference,omitempty"`
}

// MXRecordSetSpec defines the desired state of MXRecordSet
type MXRecordSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MXRecordSetParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MXRecordSetInitParameters `json:"initProvider,omitempty"`
}

// MXRecordSetStatus defines the observed state of MXRecordSet.
type MXRecordSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MXRecordSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MXRecordSet is the Schema for the MXRecordSets API. Creates an MX type DNS record set.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dns}
type MXRecordSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MXRecordSetSpec   `json:"spec"`
	Status            MXRecordSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MXRecordSetList contains a list of MXRecordSets
type MXRecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MXRecordSet `json:"items"`
}

// Repository type metadata.
var (
	MXRecordSet_Kind             = "MXRecordSet"
	MXRecordSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MXRecordSet_Kind}.String()
	MXRecordSet_KindAPIVersion   = MXRecordSet_Kind + "." + CRDGroupVersion.String()
	MXRecordSet_GroupVersionKind = CRDGroupVersion.WithKind(MXRecordSet_Kind)
)

func init() {
	SchemeBuilder.Register(&MXRecordSet{}, &MXRecordSetList{})
}
