// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by client-gen. DO NOT EDIT.
package privateendpointclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/metrics"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/utils"
)

type Client struct {
	*armnetwork.PrivateEndpointsClient
	subscriptionID string
	tracer         tracing.Tracer
}

func New(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (Interface, error) {
	if options == nil {
		options = utils.GetDefaultOption()
	}
	tr := options.TracingProvider.NewTracer(utils.ModuleName, utils.ModuleVersion)

	client, err := armnetwork.NewPrivateEndpointsClient(subscriptionID, credential, options)
	if err != nil {
		return nil, err
	}
	return &Client{
		PrivateEndpointsClient: client,
		subscriptionID:         subscriptionID,
		tracer:                 tr,
	}, nil
}

const GetOperationName = "PrivateEndpointsClient.Get"

// Get gets the PrivateEndpoint
func (client *Client) Get(ctx context.Context, resourceGroupName string, resourceName string, expand *string) (result *armnetwork.PrivateEndpoint, err error) {
	var ops *armnetwork.PrivateEndpointsClientGetOptions
	if expand != nil {
		ops = &armnetwork.PrivateEndpointsClientGetOptions{Expand: expand}
	}
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "PrivateEndpoint", "get")
	defer func() { metricsCtx.Observe(ctx, err) }()
	ctx, endSpan := runtime.StartSpan(ctx, GetOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := client.PrivateEndpointsClient.Get(ctx, resourceGroupName, resourceName, ops)
	if err != nil {
		return nil, err
	}
	//handle statuscode
	return &resp.PrivateEndpoint, nil
}

const CreateOrUpdateOperationName = "PrivateEndpointsClient.Create"

// CreateOrUpdate creates or updates a PrivateEndpoint.
func (client *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, resource armnetwork.PrivateEndpoint) (result *armnetwork.PrivateEndpoint, err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "PrivateEndpoint", "create_or_update")
	defer func() { metricsCtx.Observe(ctx, err) }()
	ctx, endSpan := runtime.StartSpan(ctx, CreateOrUpdateOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := utils.NewPollerWrapper(client.PrivateEndpointsClient.BeginCreateOrUpdate(ctx, resourceGroupName, resourceName, resource, nil)).WaitforPollerResp(ctx)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return &resp.PrivateEndpoint, nil
	}
	return nil, nil
}
