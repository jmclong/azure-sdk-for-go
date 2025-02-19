//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azblob

import (
	"context"
	"io"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/shared"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
)

// ClientOptions contains the optional parameters when creating a Client.
type ClientOptions struct {
	azcore.ClientOptions
}

// Client represents a URL to an Azure Storage blob; the blob may be a block blob, append blob, or page blob.
type Client struct {
	svc *service.Client
}

// NewClient creates a BlobClient object using the specified URL, Azure AD credential, and options.
func NewClient(serviceURL string, cred azcore.TokenCredential, options *ClientOptions) (*Client, error) {
	var clientOptions *service.ClientOptions
	if options != nil {
		clientOptions = &service.ClientOptions{ClientOptions: options.ClientOptions}
	}
	svcClient, err := service.NewClient(serviceURL, cred, clientOptions)
	if err != nil {
		return nil, err
	}

	return &Client{
		svc: svcClient,
	}, nil
}

// NewClientWithNoCredential creates a BlobClient object using the specified URL and options.
func NewClientWithNoCredential(serviceURL string, options *ClientOptions) (*Client, error) {
	var clientOptions *service.ClientOptions
	if options != nil {
		clientOptions = &service.ClientOptions{ClientOptions: options.ClientOptions}
	}
	svcClient, err := service.NewClientWithNoCredential(serviceURL, clientOptions)
	if err != nil {
		return nil, err
	}

	return &Client{
		svc: svcClient,
	}, nil
}

// NewClientWithSharedKey creates a BlobClient object using the specified URL, shared key, and options.
func NewClientWithSharedKey(serviceURL string, cred *SharedKeyCredential, options *ClientOptions) (*Client, error) {
	svcClient, err := service.NewClientWithSharedKey(serviceURL, cred, (*service.ClientOptions)(options))
	if err != nil {
		return nil, err
	}

	return &Client{
		svc: svcClient,
	}, nil
}

// NewClientFromConnectionString creates BlobClient from a connection String
func NewClientFromConnectionString(connectionString string, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}
	containerClient, err := service.NewClientFromConnectionString(connectionString, (*service.ClientOptions)(options))
	if err != nil {
		return nil, err
	}
	return &Client{
		svc: containerClient,
	}, nil
}

// URL returns the URL endpoint used by the BlobClient object.
func (c *Client) URL() string {
	return c.svc.URL()
}

// CreateContainer is a lifecycle method to creates a new container under the specified account.
// If the container with the same name already exists, a ResourceExistsError will be raised.
// This method returns a client with which to interact with the newly created container.
func (c *Client) CreateContainer(ctx context.Context, containerName string, o *CreateContainerOptions) (CreateContainerResponse, error) {
	return c.svc.CreateContainer(ctx, containerName, o)
}

// DeleteContainer is a lifecycle method that marks the specified container for deletion.
// The container and any blobs contained within it are later deleted during garbage collection.
// If the container is not found, a ResourceNotFoundError will be raised.
func (c *Client) DeleteContainer(ctx context.Context, containerName string, o *DeleteContainerOptions) (DeleteContainerResponse, error) {
	return c.svc.DeleteContainer(ctx, containerName, o)
}

// DeleteBlob marks the specified blob or snapshot for deletion. The blob is later deleted during garbage collection.
// Note that deleting a blob also deletes all its snapshots.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/delete-blob.
func (c *Client) DeleteBlob(ctx context.Context, containerName string, blobName string, o *DeleteBlobOptions) (DeleteBlobResponse, error) {
	return c.svc.NewContainerClient(containerName).NewBlobClient(blobName).Delete(ctx, o)
}

// NewListBlobsPager returns a pager for blobs starting from the specified Marker. Use an empty
// Marker to start enumeration from the beginning. Blob names are returned in lexicographic order.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/list-blobs.
func (c *Client) NewListBlobsPager(containerName string, o *ListBlobsOptions) *runtime.Pager[ListBlobsResponse] {
	return c.svc.NewContainerClient(containerName).NewListBlobsFlatPager(o)
}

// NewListContainersPager operation returns a pager of the containers under the specified account.
// Use an empty Marker to start enumeration from the beginning. Container names are returned in lexicographic order.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/list-containers2.
func (c *Client) NewListContainersPager(o *ListContainersOptions) *runtime.Pager[ListContainersResponse] {
	return c.svc.NewListContainersPager(o)
}

// UploadBuffer uploads a buffer in blocks to a block blob.
func (c *Client) UploadBuffer(ctx context.Context, containerName string, blobName string, b []byte, o *UploadBufferOptions) (UploadBufferResponse, error) {
	return c.svc.NewContainerClient(containerName).NewBlockBlobClient(blobName).UploadBuffer(ctx, b, o)
}

// UploadFile uploads a file in blocks to a block blob.
func (c *Client) UploadFile(ctx context.Context, containerName string, blobName string, file *os.File, o *UploadFileOptions) (UploadFileResponse, error) {
	return c.svc.NewContainerClient(containerName).NewBlockBlobClient(blobName).UploadFile(ctx, file, o)
}

// UploadStream copies the file held in io.Reader to the Blob at blockBlobClient.
// A Context deadline or cancellation will cause this to error.
func (c *Client) UploadStream(ctx context.Context, containerName string, blobName string, body io.Reader, o *UploadStreamOptions) (UploadStreamResponse, error) {
	return c.svc.NewContainerClient(containerName).NewBlockBlobClient(blobName).UploadStream(ctx, body, o)
}

// DownloadToStream reads a range of bytes from a blob. The response also includes the blob's properties and metadata.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/get-blob.
func (c *Client) DownloadToStream(ctx context.Context, containerName string, blobName string, o *DownloadToStreamOptions) (DownloadToStreamResponse, error) {
	o = shared.CopyOptions(o)
	return c.svc.NewContainerClient(containerName).NewBlobClient(blobName).DownloadToStream(ctx, o)
}

// DownloadToWriterAt downloads an Azure blob to a WriterAt in parallel.
func (c *Client) DownloadToWriterAt(ctx context.Context, containerName string, blobName string, writer io.WriterAt, o *DownloadToWriterAtOptions) error {
	return c.svc.NewContainerClient(containerName).NewBlobClient(blobName).DownloadToWriterAt(ctx, writer, o)
}

// DownloadToBuffer downloads an Azure blob to a buffer with parallel.
func (c *Client) DownloadToBuffer(ctx context.Context, containerName string, blobName string, _bytes []byte, o *DownloadToBufferOptions) error {
	return c.svc.NewContainerClient(containerName).NewBlobClient(blobName).DownloadToBuffer(ctx, shared.NewBytesWriter(_bytes), o)
}

// DownloadToFile downloads an Azure blob to a local file.
// The file would be truncated if the size doesn't match.
func (c *Client) DownloadToFile(ctx context.Context, containerName string, blobName string, file *os.File, o *DownloadToFileOptions) error {
	return c.svc.NewContainerClient(containerName).NewBlobClient(blobName).DownloadToFile(ctx, file, o)
}

// ServiceClient returns the underlying *service.Client for this client.
func (c *Client) ServiceClient() *service.Client {
	return c.svc
}
