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

type SRVRecordSetInitParameters struct {

	// (String) The name of the record set. The zone argument will be appended to this value to create the full record path.
	// The name of the record set. The `zone` argument will be appended to this value to create the full record path.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set) Can be specified multiple times for each SRV record. (see below for nested schema)
	// Can be specified multiple times for each SRV record.
	Srv []SrvInitParameters `json:"srv,omitempty" tf:"srv,omitempty"`

	// (Number) The TTL of the record set. Defaults to 3600.
	// The TTL of the record set. Defaults to `3600`.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type SRVRecordSetObservation struct {

	// (String) Always set to the fully qualified domain name of the record set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the record set. The zone argument will be appended to this value to create the full record path.
	// The name of the record set. The `zone` argument will be appended to this value to create the full record path.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set) Can be specified multiple times for each SRV record. (see below for nested schema)
	// Can be specified multiple times for each SRV record.
	Srv []SrvObservation `json:"srv,omitempty" tf:"srv,omitempty"`

	// (Number) The TTL of the record set. Defaults to 3600.
	// The TTL of the record set. Defaults to `3600`.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (String) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	// DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SRVRecordSetParameters struct {

	// (String) The name of the record set. The zone argument will be appended to this value to create the full record path.
	// The name of the record set. The `zone` argument will be appended to this value to create the full record path.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block Set) Can be specified multiple times for each SRV record. (see below for nested schema)
	// Can be specified multiple times for each SRV record.
	// +kubebuilder:validation:Optional
	Srv []SrvParameters `json:"srv,omitempty" tf:"srv,omitempty"`

	// (Number) The TTL of the record set. Defaults to 3600.
	// The TTL of the record set. Defaults to `3600`.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (String) DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	// DNS zone the record set belongs to. It must be an FQDN, that is, include the trailing dot.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type SrvInitParameters struct {

	// (Number) The port for the service on the target.
	// The port for the service on the target.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Number) The priority for the record.
	// The priority for the record.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (String) The FQDN of the target, include the trailing dot.
	// The FQDN of the target, include the trailing dot.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// (Number) The weight for the record.
	// The weight for the record.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type SrvObservation struct {

	// (Number) The port for the service on the target.
	// The port for the service on the target.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Number) The priority for the record.
	// The priority for the record.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (String) The FQDN of the target, include the trailing dot.
	// The FQDN of the target, include the trailing dot.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// (Number) The weight for the record.
	// The weight for the record.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type SrvParameters struct {

	// (Number) The port for the service on the target.
	// The port for the service on the target.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// (Number) The priority for the record.
	// The priority for the record.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// (String) The FQDN of the target, include the trailing dot.
	// The FQDN of the target, include the trailing dot.
	// +kubebuilder:validation:Optional
	Target *string `json:"target" tf:"target,omitempty"`

	// (Number) The weight for the record.
	// The weight for the record.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// SRVRecordSetSpec defines the desired state of SRVRecordSet
type SRVRecordSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SRVRecordSetParameters `json:"forProvider"`
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
	InitProvider SRVRecordSetInitParameters `json:"initProvider,omitempty"`
}

// SRVRecordSetStatus defines the observed state of SRVRecordSet.
type SRVRecordSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SRVRecordSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SRVRecordSet is the Schema for the SRVRecordSets API. Creates an SRV type DNS record set.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dns}
type SRVRecordSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SRVRecordSetSpec   `json:"spec"`
	Status SRVRecordSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SRVRecordSetList contains a list of SRVRecordSets
type SRVRecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SRVRecordSet `json:"items"`
}

// Repository type metadata.
var (
	SRVRecordSet_Kind             = "SRVRecordSet"
	SRVRecordSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SRVRecordSet_Kind}.String()
	SRVRecordSet_KindAPIVersion   = SRVRecordSet_Kind + "." + CRDGroupVersion.String()
	SRVRecordSet_GroupVersionKind = CRDGroupVersion.WithKind(SRVRecordSet_Kind)
)

func init() {
	SchemeBuilder.Register(&SRVRecordSet{}, &SRVRecordSetList{})
}
