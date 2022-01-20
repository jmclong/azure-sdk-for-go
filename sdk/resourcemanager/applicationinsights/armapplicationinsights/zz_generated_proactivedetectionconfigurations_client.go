//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ProactiveDetectionConfigurationsClient contains the methods for the ProactiveDetectionConfigurations group.
// Don't use this type directly, use NewProactiveDetectionConfigurationsClient() instead.
type ProactiveDetectionConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProactiveDetectionConfigurationsClient creates a new instance of ProactiveDetectionConfigurationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProactiveDetectionConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProactiveDetectionConfigurationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ProactiveDetectionConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get the ProactiveDetection configuration for this configuration id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// configurationID - The ProactiveDetection configuration ID. This is unique within a Application Insights component.
// options - ProactiveDetectionConfigurationsClientGetOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.Get
// method.
func (client *ProactiveDetectionConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, options *ProactiveDetectionConfigurationsClientGetOptions) (ProactiveDetectionConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, configurationID, options)
	if err != nil {
		return ProactiveDetectionConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProactiveDetectionConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProactiveDetectionConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProactiveDetectionConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, options *ProactiveDetectionConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationID == "" {
		return nil, errors.New("parameter configurationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ConfigurationId}", url.PathEscape(configurationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProactiveDetectionConfigurationsClient) getHandleResponse(resp *http.Response) (ProactiveDetectionConfigurationsClientGetResponse, error) {
	result := ProactiveDetectionConfigurationsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentProactiveDetectionConfiguration); err != nil {
		return ProactiveDetectionConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of ProactiveDetection configurations of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - ProactiveDetectionConfigurationsClientListOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.List
// method.
func (client *ProactiveDetectionConfigurationsClient) List(ctx context.Context, resourceGroupName string, resourceName string, options *ProactiveDetectionConfigurationsClientListOptions) (ProactiveDetectionConfigurationsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ProactiveDetectionConfigurationsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProactiveDetectionConfigurationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProactiveDetectionConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ProactiveDetectionConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ProactiveDetectionConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProactiveDetectionConfigurationsClient) listHandleResponse(resp *http.Response) (ProactiveDetectionConfigurationsClientListResponse, error) {
	result := ProactiveDetectionConfigurationsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentProactiveDetectionConfigurationArray); err != nil {
		return ProactiveDetectionConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// Update - Update the ProactiveDetection configuration for this configuration id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// configurationID - The ProactiveDetection configuration ID. This is unique within a Application Insights component.
// proactiveDetectionProperties - Properties that need to be specified to update the ProactiveDetection configuration.
// options - ProactiveDetectionConfigurationsClientUpdateOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.Update
// method.
func (client *ProactiveDetectionConfigurationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, proactiveDetectionProperties ComponentProactiveDetectionConfiguration, options *ProactiveDetectionConfigurationsClientUpdateOptions) (ProactiveDetectionConfigurationsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, configurationID, proactiveDetectionProperties, options)
	if err != nil {
		return ProactiveDetectionConfigurationsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProactiveDetectionConfigurationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProactiveDetectionConfigurationsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProactiveDetectionConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, proactiveDetectionProperties ComponentProactiveDetectionConfiguration, options *ProactiveDetectionConfigurationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationID == "" {
		return nil, errors.New("parameter configurationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ConfigurationId}", url.PathEscape(configurationID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, proactiveDetectionProperties)
}

// updateHandleResponse handles the Update response.
func (client *ProactiveDetectionConfigurationsClient) updateHandleResponse(resp *http.Response) (ProactiveDetectionConfigurationsClientUpdateResponse, error) {
	result := ProactiveDetectionConfigurationsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentProactiveDetectionConfiguration); err != nil {
		return ProactiveDetectionConfigurationsClientUpdateResponse{}, err
	}
	return result, nil
}