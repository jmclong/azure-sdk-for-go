//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azblob_test

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/lease"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

// var headersToIgnoreForLease = []string {"X-Ms-Proposed-Lease-Id", "X-Ms-Lease-Id"}
var proposedLeaseIDs = []*string{to.Ptr("c820a799-76d7-4ee2-6e15-546f19325c2c"), to.Ptr("326cc5e1-746e-4af8-4811-a50e6629a8ca")}

func (s *azblobTestSuite) TestContainerAcquireLease() {
	_require := require.New(s.T())
	testName := s.T().Name()
	//ignoreHeaders(_context.recording, headersToIgnoreForLease)

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	containerLeaseClient, _ := lease.NewContainerClient(containerClient, &lease.ContainerClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})

	ctx := context.Background()
	acquireLeaseResponse, err := containerLeaseClient.AcquireLease(ctx, &lease.ContainerAcquireOptions{Duration: to.Ptr[int32](60)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(*acquireLeaseResponse.LeaseID, *containerLeaseClient.LeaseID())

	_, err = containerLeaseClient.ReleaseLease(ctx, nil)
	_require.Nil(err)
}

func (s *azblobTestSuite) TestContainerDeleteContainerWithoutLeaseId() {
	_require := require.New(s.T())
	testName := s.T().Name()

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	containerLeaseClient, _ := lease.NewContainerClient(containerClient, &lease.ContainerClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})

	ctx := context.Background()
	acquireLeaseResponse, err := containerLeaseClient.AcquireLease(ctx, &lease.ContainerAcquireOptions{Duration: to.Ptr[int32](60)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(*acquireLeaseResponse.LeaseID, *containerLeaseClient.LeaseID())

	_, err = containerClient.Delete(ctx, nil)
	_require.NotNil(err)

	leaseID := containerLeaseClient.LeaseID()
	_, err = containerClient.Delete(ctx, &container.DeleteOptions{
		AccessConditions: &container.AccessConditions{
			LeaseAccessConditions: &container.LeaseAccessConditions{
				LeaseID: leaseID,
			},
		},
	})
	_require.Nil(err)
}

func (s *azblobTestSuite) TestContainerReleaseLease() {
	_require := require.New(s.T())
	testName := s.T().Name()

	//ignoreHeaders(_context.recording, headersToIgnoreForLease)

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	containerLeaseClient, _ := lease.NewContainerClient(containerClient, &lease.ContainerClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})

	ctx := context.Background()
	acquireLeaseResponse, err := containerLeaseClient.AcquireLease(ctx, &lease.ContainerAcquireOptions{Duration: to.Ptr[int32](60)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(*acquireLeaseResponse.LeaseID, *containerLeaseClient.LeaseID())

	_, err = containerClient.Delete(ctx, nil)
	_require.NotNil(err)

	_, err = containerLeaseClient.ReleaseLease(ctx, nil)
	_require.Nil(err)

	_, err = containerClient.Delete(ctx, nil)
	_require.Nil(err)
}

func (s *azblobTestSuite) TestContainerRenewLease() {
	_require := require.New(s.T())
	testName := s.T().Name()

	//ignoreHeaders(_context.recording, headersToIgnoreForLease)

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	containerLeaseClient, _ := lease.NewContainerClient(containerClient, &lease.ContainerClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})

	ctx := context.Background()
	acquireLeaseResponse, err := containerLeaseClient.AcquireLease(ctx, &lease.ContainerAcquireOptions{Duration: to.Ptr[int32](15)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(*acquireLeaseResponse.LeaseID, *containerLeaseClient.LeaseID())

	_, err = containerLeaseClient.RenewLease(ctx, nil)
	_require.Nil(err)

	_, err = containerLeaseClient.ReleaseLease(ctx, nil)
	_require.Nil(err)
}

func (s *azblobTestSuite) TestContainerChangeLease() {
	_require := require.New(s.T())
	testName := s.T().Name()

	//ignoreHeaders(_context.recording, headersToIgnoreForLease)

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	containerLeaseClient, _ := lease.NewContainerClient(containerClient, &lease.ContainerClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})

	ctx := context.Background()
	acquireLeaseResponse, err := containerLeaseClient.AcquireLease(ctx, &lease.ContainerAcquireOptions{Duration: to.Ptr[int32](15)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(*acquireLeaseResponse.LeaseID, *containerLeaseClient.LeaseID())

	changeLeaseResp, err := containerLeaseClient.ChangeLease(ctx, &lease.ContainerChangeOptions{
		ProposedLeaseID: proposedLeaseIDs[1],
	})
	_require.Nil(err)
	_require.EqualValues(changeLeaseResp.LeaseID, proposedLeaseIDs[1])
	_require.EqualValues(containerLeaseClient.LeaseID(), proposedLeaseIDs[1])

	_, err = containerLeaseClient.RenewLease(ctx, nil)
	_require.Nil(err)

	_, err = containerLeaseClient.ReleaseLease(ctx, nil)
	_require.Nil(err)
}

func (s *azblobTestSuite) TestBlobAcquireLease() {
	_require := require.New(s.T())
	testName := s.T().Name()

	//ignoreHeaders(_context.recording, headersToIgnoreForLease)

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	blobName := generateBlobName(testName)
	bbClient := createNewBlockBlob(_require, blobName, containerClient)
	blobLeaseClient, err := lease.NewBlobClient(bbClient, &lease.BlobClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})
	_require.NoError(err)

	ctx := context.Background()
	acquireLeaseResponse, err := blobLeaseClient.AcquireLease(ctx, &lease.BlobAcquireOptions{Duration: to.Ptr[int32](60)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(acquireLeaseResponse.LeaseID, blobLeaseClient.LeaseID())

	_, err = blobLeaseClient.ReleaseLease(ctx, nil)
	_require.Nil(err)
}

func (s *azblobTestSuite) TestDeleteBlobWithoutLeaseId() {
	_require := require.New(s.T())
	testName := s.T().Name()

	//ignoreHeaders(_context.recording, headersToIgnoreForLease)

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	blobName := generateBlobName(testName)
	bbClient := createNewBlockBlob(_require, blobName, containerClient)
	blobLeaseClient, err := lease.NewBlobClient(bbClient, &lease.BlobClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})
	_require.NoError(err)

	ctx := context.Background()
	acquireLeaseResponse, err := blobLeaseClient.AcquireLease(ctx, &lease.BlobAcquireOptions{Duration: to.Ptr[int32](60)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(acquireLeaseResponse.LeaseID, blobLeaseClient.LeaseID())

	_, err = bbClient.Delete(ctx, nil)
	_require.NotNil(err)

	leaseID := blobLeaseClient.LeaseID()
	_, err = bbClient.Delete(ctx, &blob.DeleteOptions{
		AccessConditions: &blob.AccessConditions{
			LeaseAccessConditions: &blob.LeaseAccessConditions{
				LeaseID: leaseID,
			},
		},
	})
	_require.Nil(err)
}

func (s *azblobTestSuite) TestBlobReleaseLease() {
	_require := require.New(s.T())
	testName := s.T().Name()

	//ignoreHeaders(_context.recording, headersToIgnoreForLease)

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	blobName := generateBlobName(testName)
	bbClient := createNewBlockBlob(_require, blobName, containerClient)
	blobLeaseClient, _ := lease.NewBlobClient(bbClient, &lease.BlobClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})

	ctx := context.Background()
	acquireLeaseResponse, err := blobLeaseClient.AcquireLease(ctx, &lease.BlobAcquireOptions{Duration: to.Ptr[int32](60)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(acquireLeaseResponse.LeaseID, blobLeaseClient.LeaseID())

	_, err = bbClient.Delete(ctx, nil)
	_require.NotNil(err)

	_, err = blobLeaseClient.ReleaseLease(ctx, nil)
	_require.Nil(err)

	_, err = bbClient.Delete(ctx, nil)
	_require.Nil(err)
}

func (s *azblobTestSuite) TestBlobRenewLease() {
	_require := require.New(s.T())
	testName := s.T().Name()
	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	blobName := generateBlobName(testName)
	bbClient := createNewBlockBlob(_require, blobName, containerClient)
	blobLeaseClient, _ := lease.NewBlobClient(bbClient, &lease.BlobClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})

	ctx := context.Background()
	acquireLeaseResponse, err := blobLeaseClient.AcquireLease(ctx, &lease.BlobAcquireOptions{Duration: to.Ptr[int32](15)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.EqualValues(acquireLeaseResponse.LeaseID, blobLeaseClient.LeaseID())

	_, err = blobLeaseClient.RenewLease(ctx, nil)
	_require.Nil(err)

	_, err = blobLeaseClient.ReleaseLease(ctx, nil)
	_require.Nil(err)
}

func (s *azblobTestSuite) TestBlobChangeLease() {
	_require := require.New(s.T())
	testName := s.T().Name()

	//ignoreHeaders(_context.recording, headersToIgnoreForLease)

	svcClient, err := getServiceClient(s.T(), testAccountDefault, nil)
	_require.NoError(err)

	containerName := generateContainerName(testName)
	containerClient := createNewContainer(_require, containerName, svcClient)
	defer deleteContainer(_require, containerClient)

	blobName := generateBlobName(testName)
	bbClient := createNewBlockBlob(_require, blobName, containerClient)
	blobLeaseClient, _ := lease.NewBlobClient(bbClient, &lease.BlobClientOptions{
		LeaseID: proposedLeaseIDs[0],
	})

	ctx := context.Background()
	acquireLeaseResponse, err := blobLeaseClient.AcquireLease(ctx, &lease.BlobAcquireOptions{Duration: to.Ptr[int32](15)})
	_require.Nil(err)
	_require.NotNil(acquireLeaseResponse.LeaseID)
	_require.Equal(*acquireLeaseResponse.LeaseID, *proposedLeaseIDs[0])

	changeLeaseResp, err := blobLeaseClient.ChangeLease(ctx, &lease.BlobChangeOptions{
		ProposedLeaseID: proposedLeaseIDs[1],
	})
	_require.Nil(err)
	_require.Equal(*changeLeaseResp.LeaseID, *proposedLeaseIDs[1])

	_, err = blobLeaseClient.RenewLease(ctx, nil)
	_require.Nil(err)

	_, err = blobLeaseClient.ReleaseLease(ctx, nil)
	_require.Nil(err)
}
