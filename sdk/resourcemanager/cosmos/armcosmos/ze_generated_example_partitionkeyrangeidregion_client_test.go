//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBPKeyRangeIdRegionGetMetrics.json
func ExamplePartitionKeyRangeIDRegionClient_ListMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewPartitionKeyRangeIDRegionClient("<subscription-id>", cred, nil)
	res, err := client.ListMetrics(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<region>",
		"<database-rid>",
		"<collection-rid>",
		"<partition-key-range-id>",
		"<filter>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PartitionKeyRangeIDRegionClientListMetricsResult)
}