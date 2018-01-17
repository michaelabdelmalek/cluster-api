/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package cluster

import (
	"fmt"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalCluster = builders.NewInternalResource(
		"clusters",
		func() runtime.Object { return &Cluster{} },
		func() runtime.Object { return &ClusterList{} },
	)
	InternalClusterStatus = builders.NewInternalResourceStatus(
		"clusters",
		func() runtime.Object { return &Cluster{} },
		func() runtime.Object { return &ClusterList{} },
	)
	InternalMachine = builders.NewInternalResource(
		"machines",
		func() runtime.Object { return &Machine{} },
		func() runtime.Object { return &MachineList{} },
	)
	InternalMachineStatus = builders.NewInternalResourceStatus(
		"machines",
		func() runtime.Object { return &Machine{} },
		func() runtime.Object { return &MachineList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("cluster.k8s.io").WithKinds(
		InternalCluster,
		InternalClusterStatus,
		InternalMachine,
		InternalMachineStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ClusterSpec
	Status ClusterStatus
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Machine struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   MachineSpec
	Status MachineStatus
}

type ClusterStatus struct {
	APIEndpoints   []APIEndpoint
	ErrorReason    string
	ErrorMessage   string
	ProviderStatus string
}

type MachineStatus struct {
	NodeRef      *corev1.ObjectReference
	LastUpdated  metav1.Time
	Ready        bool
	ErrorReason  *string
	ErrorMessage *string
}

type APIEndpoint struct {
	Host string
	Port int
}

type MachineSpec struct {
	metav1.ObjectMeta
	ProviderConfig string
	Roles          []string
	Versions       MachineVersionInfo
	ConfigSource   *corev1.NodeConfigSource
}

type ClusterSpec struct {
	ClusterNetwork ClusterNetworkingConfig
	ProviderConfig string
}

type MachineVersionInfo struct {
	Kubelet          string
	ControlPlane     string
	ContainerRuntime ContainerRuntimeInfo
}

type ClusterNetworkingConfig struct {
	Services  NetworkRanges
	Pods      NetworkRanges
	DNSDomain string
}

type ContainerRuntimeInfo struct {
	Name    string
	Version string
}

type NetworkRanges struct {
	CIDRBlocks []string
}

//
// Cluster Functions and Structs
//
// +k8s:deepcopy-gen=false
type ClusterStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type ClusterStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Cluster
}

func (Cluster) NewStatus() interface{} {
	return ClusterStatus{}
}

func (pc *Cluster) GetStatus() interface{} {
	return pc.Status
}

func (pc *Cluster) SetStatus(s interface{}) {
	pc.Status = s.(ClusterStatus)
}

func (pc *Cluster) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Cluster) SetSpec(s interface{}) {
	pc.Spec = s.(ClusterSpec)
}

func (pc *Cluster) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Cluster) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Cluster) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Cluster.
// +k8s:deepcopy-gen=false
type ClusterRegistry interface {
	ListClusters(ctx request.Context, options *internalversion.ListOptions) (*ClusterList, error)
	GetCluster(ctx request.Context, id string, options *metav1.GetOptions) (*Cluster, error)
	CreateCluster(ctx request.Context, id *Cluster) (*Cluster, error)
	UpdateCluster(ctx request.Context, id *Cluster) (*Cluster, error)
	DeleteCluster(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewClusterRegistry(sp builders.StandardStorageProvider) ClusterRegistry {
	return &storageCluster{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageCluster struct {
	builders.StandardStorageProvider
}

func (s *storageCluster) ListClusters(ctx request.Context, options *internalversion.ListOptions) (*ClusterList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ClusterList), err
}

func (s *storageCluster) GetCluster(ctx request.Context, id string, options *metav1.GetOptions) (*Cluster, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Cluster), nil
}

func (s *storageCluster) CreateCluster(ctx request.Context, object *Cluster) (*Cluster, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, false)
	if err != nil {
		return nil, err
	}
	return obj.(*Cluster), nil
}

func (s *storageCluster) UpdateCluster(ctx request.Context, object *Cluster) (*Cluster, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object, builders.Scheme))
	if err != nil {
		return nil, err
	}
	return obj.(*Cluster), nil
}

func (s *storageCluster) DeleteCluster(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}

//
// Machine Functions and Structs
//
// +k8s:deepcopy-gen=false
type MachineStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type MachineStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MachineList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Machine
}

func (Machine) NewStatus() interface{} {
	return MachineStatus{}
}

func (pc *Machine) GetStatus() interface{} {
	return pc.Status
}

func (pc *Machine) SetStatus(s interface{}) {
	pc.Status = s.(MachineStatus)
}

func (pc *Machine) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Machine) SetSpec(s interface{}) {
	pc.Spec = s.(MachineSpec)
}

func (pc *Machine) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Machine) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Machine) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Machine.
// +k8s:deepcopy-gen=false
type MachineRegistry interface {
	ListMachines(ctx request.Context, options *internalversion.ListOptions) (*MachineList, error)
	GetMachine(ctx request.Context, id string, options *metav1.GetOptions) (*Machine, error)
	CreateMachine(ctx request.Context, id *Machine) (*Machine, error)
	UpdateMachine(ctx request.Context, id *Machine) (*Machine, error)
	DeleteMachine(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewMachineRegistry(sp builders.StandardStorageProvider) MachineRegistry {
	return &storageMachine{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageMachine struct {
	builders.StandardStorageProvider
}

func (s *storageMachine) ListMachines(ctx request.Context, options *internalversion.ListOptions) (*MachineList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*MachineList), err
}

func (s *storageMachine) GetMachine(ctx request.Context, id string, options *metav1.GetOptions) (*Machine, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Machine), nil
}

func (s *storageMachine) CreateMachine(ctx request.Context, object *Machine) (*Machine, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, false)
	if err != nil {
		return nil, err
	}
	return obj.(*Machine), nil
}

func (s *storageMachine) UpdateMachine(ctx request.Context, object *Machine) (*Machine, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object, builders.Scheme))
	if err != nil {
		return nil, err
	}
	return obj.(*Machine), nil
}

func (s *storageMachine) DeleteMachine(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}
