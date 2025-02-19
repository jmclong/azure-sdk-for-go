//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package redisenterprise

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/redisenterprise/mgmt/2022-01-01/redisenterprise"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessKeyType = original.AccessKeyType

const (
	AccessKeyTypePrimary   AccessKeyType = original.AccessKeyTypePrimary
	AccessKeyTypeSecondary AccessKeyType = original.AccessKeyTypeSecondary
)

type ActionType = original.ActionType

const (
	ActionTypeInternal ActionType = original.ActionTypeInternal
)

type AofFrequency = original.AofFrequency

const (
	AofFrequencyAlways AofFrequency = original.AofFrequencyAlways
	AofFrequencyOnes   AofFrequency = original.AofFrequencyOnes
)

type ClusteringPolicy = original.ClusteringPolicy

const (
	ClusteringPolicyEnterpriseCluster ClusteringPolicy = original.ClusteringPolicyEnterpriseCluster
	ClusteringPolicyOSSCluster        ClusteringPolicy = original.ClusteringPolicyOSSCluster
)

type EvictionPolicy = original.EvictionPolicy

const (
	EvictionPolicyAllKeysLFU     EvictionPolicy = original.EvictionPolicyAllKeysLFU
	EvictionPolicyAllKeysLRU     EvictionPolicy = original.EvictionPolicyAllKeysLRU
	EvictionPolicyAllKeysRandom  EvictionPolicy = original.EvictionPolicyAllKeysRandom
	EvictionPolicyNoEviction     EvictionPolicy = original.EvictionPolicyNoEviction
	EvictionPolicyVolatileLFU    EvictionPolicy = original.EvictionPolicyVolatileLFU
	EvictionPolicyVolatileLRU    EvictionPolicy = original.EvictionPolicyVolatileLRU
	EvictionPolicyVolatileRandom EvictionPolicy = original.EvictionPolicyVolatileRandom
	EvictionPolicyVolatileTTL    EvictionPolicy = original.EvictionPolicyVolatileTTL
)

type LinkState = original.LinkState

const (
	LinkStateLinked       LinkState = original.LinkStateLinked
	LinkStateLinkFailed   LinkState = original.LinkStateLinkFailed
	LinkStateLinking      LinkState = original.LinkStateLinking
	LinkStateUnlinkFailed LinkState = original.LinkStateUnlinkFailed
	LinkStateUnlinking    LinkState = original.LinkStateUnlinking
)

type Origin = original.Origin

const (
	OriginSystem     Origin = original.OriginSystem
	OriginUser       Origin = original.OriginUser
	OriginUsersystem Origin = original.OriginUsersystem
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatus

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusApproved
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusPending
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = original.PrivateEndpointServiceConnectionStatusRejected
)

type Protocol = original.Protocol

const (
	ProtocolEncrypted Protocol = original.ProtocolEncrypted
	ProtocolPlaintext Protocol = original.ProtocolPlaintext
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type RdbFrequency = original.RdbFrequency

const (
	RdbFrequencyOneh    RdbFrequency = original.RdbFrequencyOneh
	RdbFrequencyOneTwoh RdbFrequency = original.RdbFrequencyOneTwoh
	RdbFrequencySixh    RdbFrequency = original.RdbFrequencySixh
)

type ResourceState = original.ResourceState

const (
	ResourceStateCreateFailed  ResourceState = original.ResourceStateCreateFailed
	ResourceStateCreating      ResourceState = original.ResourceStateCreating
	ResourceStateDeleteFailed  ResourceState = original.ResourceStateDeleteFailed
	ResourceStateDeleting      ResourceState = original.ResourceStateDeleting
	ResourceStateDisabled      ResourceState = original.ResourceStateDisabled
	ResourceStateDisableFailed ResourceState = original.ResourceStateDisableFailed
	ResourceStateDisabling     ResourceState = original.ResourceStateDisabling
	ResourceStateEnableFailed  ResourceState = original.ResourceStateEnableFailed
	ResourceStateEnabling      ResourceState = original.ResourceStateEnabling
	ResourceStateRunning       ResourceState = original.ResourceStateRunning
	ResourceStateUpdateFailed  ResourceState = original.ResourceStateUpdateFailed
	ResourceStateUpdating      ResourceState = original.ResourceStateUpdating
)

type SkuName = original.SkuName

const (
	SkuNameEnterpriseE10        SkuName = original.SkuNameEnterpriseE10
	SkuNameEnterpriseE100       SkuName = original.SkuNameEnterpriseE100
	SkuNameEnterpriseE20        SkuName = original.SkuNameEnterpriseE20
	SkuNameEnterpriseE50        SkuName = original.SkuNameEnterpriseE50
	SkuNameEnterpriseFlashF1500 SkuName = original.SkuNameEnterpriseFlashF1500
	SkuNameEnterpriseFlashF300  SkuName = original.SkuNameEnterpriseFlashF300
	SkuNameEnterpriseFlashF700  SkuName = original.SkuNameEnterpriseFlashF700
)

type TLSVersion = original.TLSVersion

const (
	TLSVersionOneFullStopOne  TLSVersion = original.TLSVersionOneFullStopOne
	TLSVersionOneFullStopTwo  TLSVersion = original.TLSVersionOneFullStopTwo
	TLSVersionOneFullStopZero TLSVersion = original.TLSVersionOneFullStopZero
)

type AccessKeys = original.AccessKeys
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type Client = original.Client
type Cluster = original.Cluster
type ClusterList = original.ClusterList
type ClusterListIterator = original.ClusterListIterator
type ClusterListPage = original.ClusterListPage
type ClusterProperties = original.ClusterProperties
type ClusterUpdate = original.ClusterUpdate
type CreateFuture = original.CreateFuture
type Database = original.Database
type DatabaseList = original.DatabaseList
type DatabaseListIterator = original.DatabaseListIterator
type DatabaseListPage = original.DatabaseListPage
type DatabaseProperties = original.DatabaseProperties
type DatabasePropertiesGeoReplication = original.DatabasePropertiesGeoReplication
type DatabaseUpdate = original.DatabaseUpdate
type DatabasesClient = original.DatabasesClient
type DatabasesCreateFuture = original.DatabasesCreateFuture
type DatabasesDeleteFuture = original.DatabasesDeleteFuture
type DatabasesExportFuture = original.DatabasesExportFuture
type DatabasesForceUnlinkFuture = original.DatabasesForceUnlinkFuture
type DatabasesImportFuture = original.DatabasesImportFuture
type DatabasesRegenerateKeyFuture = original.DatabasesRegenerateKeyFuture
type DatabasesUpdateFuture = original.DatabasesUpdateFuture
type DeleteFuture = original.DeleteFuture
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type ExportClusterParameters = original.ExportClusterParameters
type ForceUnlinkParameters = original.ForceUnlinkParameters
type ImportClusterParameters = original.ImportClusterParameters
type LinkedDatabase = original.LinkedDatabase
type Module = original.Module
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationStatus = original.OperationStatus
type OperationsClient = original.OperationsClient
type OperationsStatusClient = original.OperationsStatusClient
type Persistence = original.Persistence
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsPutFuture = original.PrivateEndpointConnectionsPutFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourceListResult = original.PrivateLinkResourceListResult
type PrivateLinkResourceProperties = original.PrivateLinkResourceProperties
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type UpdateFuture = original.UpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterListIterator(page ClusterListPage) ClusterListIterator {
	return original.NewClusterListIterator(page)
}
func NewClusterListPage(cur ClusterList, getNextPage func(context.Context, ClusterList) (ClusterList, error)) ClusterListPage {
	return original.NewClusterListPage(cur, getNextPage)
}
func NewDatabaseListIterator(page DatabaseListPage) DatabaseListIterator {
	return original.NewDatabaseListIterator(page)
}
func NewDatabaseListPage(cur DatabaseList, getNextPage func(context.Context, DatabaseList) (DatabaseList, error)) DatabaseListPage {
	return original.NewDatabaseListPage(cur, getNextPage)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsStatusClient(subscriptionID string) OperationsStatusClient {
	return original.NewOperationsStatusClient(subscriptionID)
}
func NewOperationsStatusClientWithBaseURI(baseURI string, subscriptionID string) OperationsStatusClient {
	return original.NewOperationsStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessKeyTypeValues() []AccessKeyType {
	return original.PossibleAccessKeyTypeValues()
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleAofFrequencyValues() []AofFrequency {
	return original.PossibleAofFrequencyValues()
}
func PossibleClusteringPolicyValues() []ClusteringPolicy {
	return original.PossibleClusteringPolicyValues()
}
func PossibleEvictionPolicyValues() []EvictionPolicy {
	return original.PossibleEvictionPolicyValues()
}
func PossibleLinkStateValues() []LinkState {
	return original.PossibleLinkStateValues()
}
func PossibleOriginValues() []Origin {
	return original.PossibleOriginValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return original.PossiblePrivateEndpointServiceConnectionStatusValues()
}
func PossibleProtocolValues() []Protocol {
	return original.PossibleProtocolValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRdbFrequencyValues() []RdbFrequency {
	return original.PossibleRdbFrequencyValues()
}
func PossibleResourceStateValues() []ResourceState {
	return original.PossibleResourceStateValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleTLSVersionValues() []TLSVersion {
	return original.PossibleTLSVersionValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
