//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceprovisioningservices

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DpsCertificateClient contains the methods for the DpsCertificate group.
// Don't use this type directly, use NewDpsCertificateClient() instead.
type DpsCertificateClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDpsCertificateClient creates a new instance of DpsCertificateClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDpsCertificateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DpsCertificateClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DpsCertificateClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Add new certificate or update an existing certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-05
// resourceGroupName - Resource group identifier.
// provisioningServiceName - The name of the provisioning service.
// certificateName - The name of the certificate create or update.
// certificateDescription - The certificate body.
// options - DpsCertificateClientCreateOrUpdateOptions contains the optional parameters for the DpsCertificateClient.CreateOrUpdate
// method.
func (client *DpsCertificateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, provisioningServiceName string, certificateName string, certificateDescription CertificateResponse, options *DpsCertificateClientCreateOrUpdateOptions) (DpsCertificateClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, provisioningServiceName, certificateName, certificateDescription, options)
	if err != nil {
		return DpsCertificateClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DpsCertificateClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, provisioningServiceName string, certificateName string, certificateDescription CertificateResponse, options *DpsCertificateClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-05")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, certificateDescription)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DpsCertificateClient) createOrUpdateHandleResponse(resp *http.Response) (DpsCertificateClientCreateOrUpdateResponse, error) {
	result := DpsCertificateClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateResponse); err != nil {
		return DpsCertificateClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified certificate associated with the Provisioning Service
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-05
// resourceGroupName - Resource group identifier.
// ifMatch - ETag of the certificate
// provisioningServiceName - The name of the provisioning service.
// certificateName - This is a mandatory field, and is the logical name of the certificate that the provisioning service will
// access by.
// options - DpsCertificateClientDeleteOptions contains the optional parameters for the DpsCertificateClient.Delete method.
func (client *DpsCertificateClient) Delete(ctx context.Context, resourceGroupName string, ifMatch string, provisioningServiceName string, certificateName string, options *DpsCertificateClientDeleteOptions) (DpsCertificateClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ifMatch, provisioningServiceName, certificateName, options)
	if err != nil {
		return DpsCertificateClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DpsCertificateClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DpsCertificateClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DpsCertificateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ifMatch string, provisioningServiceName string, certificateName string, options *DpsCertificateClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CertificateName1 != nil {
		reqQP.Set("certificate.name", *options.CertificateName1)
	}
	if options != nil && options.CertificateRawBytes != nil {
		reqQP.Set("certificate.rawBytes", base64.StdEncoding.EncodeToString(options.CertificateRawBytes))
	}
	if options != nil && options.CertificateIsVerified != nil {
		reqQP.Set("certificate.isVerified", strconv.FormatBool(*options.CertificateIsVerified))
	}
	if options != nil && options.CertificatePurpose != nil {
		reqQP.Set("certificate.purpose", string(*options.CertificatePurpose))
	}
	if options != nil && options.CertificateCreated != nil {
		reqQP.Set("certificate.created", options.CertificateCreated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateLastUpdated != nil {
		reqQP.Set("certificate.lastUpdated", options.CertificateLastUpdated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateHasPrivateKey != nil {
		reqQP.Set("certificate.hasPrivateKey", strconv.FormatBool(*options.CertificateHasPrivateKey))
	}
	if options != nil && options.CertificateNonce != nil {
		reqQP.Set("certificate.nonce", *options.CertificateNonce)
	}
	reqQP.Set("api-version", "2022-02-05")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GenerateVerificationCode - Generate verification code for Proof of Possession.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-05
// certificateName - The mandatory logical name of the certificate, that the provisioning service uses to access.
// ifMatch - ETag of the certificate. This is required to update an existing certificate, and ignored while creating a brand
// new certificate.
// resourceGroupName - name of resource group.
// provisioningServiceName - Name of provisioning service.
// options - DpsCertificateClientGenerateVerificationCodeOptions contains the optional parameters for the DpsCertificateClient.GenerateVerificationCode
// method.
func (client *DpsCertificateClient) GenerateVerificationCode(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, options *DpsCertificateClientGenerateVerificationCodeOptions) (DpsCertificateClientGenerateVerificationCodeResponse, error) {
	req, err := client.generateVerificationCodeCreateRequest(ctx, certificateName, ifMatch, resourceGroupName, provisioningServiceName, options)
	if err != nil {
		return DpsCertificateClientGenerateVerificationCodeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateClientGenerateVerificationCodeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateClientGenerateVerificationCodeResponse{}, runtime.NewResponseError(resp)
	}
	return client.generateVerificationCodeHandleResponse(resp)
}

// generateVerificationCodeCreateRequest creates the GenerateVerificationCode request.
func (client *DpsCertificateClient) generateVerificationCodeCreateRequest(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, options *DpsCertificateClientGenerateVerificationCodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}/generateVerificationCode"
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CertificateName1 != nil {
		reqQP.Set("certificate.name", *options.CertificateName1)
	}
	if options != nil && options.CertificateRawBytes != nil {
		reqQP.Set("certificate.rawBytes", base64.StdEncoding.EncodeToString(options.CertificateRawBytes))
	}
	if options != nil && options.CertificateIsVerified != nil {
		reqQP.Set("certificate.isVerified", strconv.FormatBool(*options.CertificateIsVerified))
	}
	if options != nil && options.CertificatePurpose != nil {
		reqQP.Set("certificate.purpose", string(*options.CertificatePurpose))
	}
	if options != nil && options.CertificateCreated != nil {
		reqQP.Set("certificate.created", options.CertificateCreated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateLastUpdated != nil {
		reqQP.Set("certificate.lastUpdated", options.CertificateLastUpdated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateHasPrivateKey != nil {
		reqQP.Set("certificate.hasPrivateKey", strconv.FormatBool(*options.CertificateHasPrivateKey))
	}
	if options != nil && options.CertificateNonce != nil {
		reqQP.Set("certificate.nonce", *options.CertificateNonce)
	}
	reqQP.Set("api-version", "2022-02-05")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateVerificationCodeHandleResponse handles the GenerateVerificationCode response.
func (client *DpsCertificateClient) generateVerificationCodeHandleResponse(resp *http.Response) (DpsCertificateClientGenerateVerificationCodeResponse, error) {
	result := DpsCertificateClientGenerateVerificationCodeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VerificationCodeResponse); err != nil {
		return DpsCertificateClientGenerateVerificationCodeResponse{}, err
	}
	return result, nil
}

// Get - Get the certificate from the provisioning service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-05
// certificateName - Name of the certificate to retrieve.
// resourceGroupName - Resource group identifier.
// provisioningServiceName - Name of the provisioning service the certificate is associated with.
// options - DpsCertificateClientGetOptions contains the optional parameters for the DpsCertificateClient.Get method.
func (client *DpsCertificateClient) Get(ctx context.Context, certificateName string, resourceGroupName string, provisioningServiceName string, options *DpsCertificateClientGetOptions) (DpsCertificateClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, certificateName, resourceGroupName, provisioningServiceName, options)
	if err != nil {
		return DpsCertificateClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DpsCertificateClient) getCreateRequest(ctx context.Context, certificateName string, resourceGroupName string, provisioningServiceName string, options *DpsCertificateClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}"
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-05")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DpsCertificateClient) getHandleResponse(resp *http.Response) (DpsCertificateClientGetResponse, error) {
	result := DpsCertificateClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateResponse); err != nil {
		return DpsCertificateClientGetResponse{}, err
	}
	return result, nil
}

// List - Get all the certificates tied to the provisioning service.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-05
// resourceGroupName - Name of resource group.
// provisioningServiceName - Name of provisioning service to retrieve certificates for.
// options - DpsCertificateClientListOptions contains the optional parameters for the DpsCertificateClient.List method.
func (client *DpsCertificateClient) List(ctx context.Context, resourceGroupName string, provisioningServiceName string, options *DpsCertificateClientListOptions) (DpsCertificateClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, provisioningServiceName, options)
	if err != nil {
		return DpsCertificateClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DpsCertificateClient) listCreateRequest(ctx context.Context, resourceGroupName string, provisioningServiceName string, options *DpsCertificateClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-05")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DpsCertificateClient) listHandleResponse(resp *http.Response) (DpsCertificateClientListResponse, error) {
	result := DpsCertificateClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateListDescription); err != nil {
		return DpsCertificateClientListResponse{}, err
	}
	return result, nil
}

// VerifyCertificate - Verifies the certificate's private key possession by providing the leaf cert issued by the verifying
// pre uploaded certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-05
// certificateName - The mandatory logical name of the certificate, that the provisioning service uses to access.
// ifMatch - ETag of the certificate.
// resourceGroupName - Resource group name.
// provisioningServiceName - Provisioning service name.
// request - The name of the certificate
// options - DpsCertificateClientVerifyCertificateOptions contains the optional parameters for the DpsCertificateClient.VerifyCertificate
// method.
func (client *DpsCertificateClient) VerifyCertificate(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, request VerificationCodeRequest, options *DpsCertificateClientVerifyCertificateOptions) (DpsCertificateClientVerifyCertificateResponse, error) {
	req, err := client.verifyCertificateCreateRequest(ctx, certificateName, ifMatch, resourceGroupName, provisioningServiceName, request, options)
	if err != nil {
		return DpsCertificateClientVerifyCertificateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DpsCertificateClientVerifyCertificateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DpsCertificateClientVerifyCertificateResponse{}, runtime.NewResponseError(resp)
	}
	return client.verifyCertificateHandleResponse(resp)
}

// verifyCertificateCreateRequest creates the VerifyCertificate request.
func (client *DpsCertificateClient) verifyCertificateCreateRequest(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, request VerificationCodeRequest, options *DpsCertificateClientVerifyCertificateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}/verify"
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if provisioningServiceName == "" {
		return nil, errors.New("parameter provisioningServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{provisioningServiceName}", url.PathEscape(provisioningServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.CertificateName1 != nil {
		reqQP.Set("certificate.name", *options.CertificateName1)
	}
	if options != nil && options.CertificateRawBytes != nil {
		reqQP.Set("certificate.rawBytes", base64.StdEncoding.EncodeToString(options.CertificateRawBytes))
	}
	if options != nil && options.CertificateIsVerified != nil {
		reqQP.Set("certificate.isVerified", strconv.FormatBool(*options.CertificateIsVerified))
	}
	if options != nil && options.CertificatePurpose != nil {
		reqQP.Set("certificate.purpose", string(*options.CertificatePurpose))
	}
	if options != nil && options.CertificateCreated != nil {
		reqQP.Set("certificate.created", options.CertificateCreated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateLastUpdated != nil {
		reqQP.Set("certificate.lastUpdated", options.CertificateLastUpdated.Format(time.RFC3339Nano))
	}
	if options != nil && options.CertificateHasPrivateKey != nil {
		reqQP.Set("certificate.hasPrivateKey", strconv.FormatBool(*options.CertificateHasPrivateKey))
	}
	if options != nil && options.CertificateNonce != nil {
		reqQP.Set("certificate.nonce", *options.CertificateNonce)
	}
	reqQP.Set("api-version", "2022-02-05")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// verifyCertificateHandleResponse handles the VerifyCertificate response.
func (client *DpsCertificateClient) verifyCertificateHandleResponse(resp *http.Response) (DpsCertificateClientVerifyCertificateResponse, error) {
	result := DpsCertificateClientVerifyCertificateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateResponse); err != nil {
		return DpsCertificateClientVerifyCertificateResponse{}, err
	}
	return result, nil
}
