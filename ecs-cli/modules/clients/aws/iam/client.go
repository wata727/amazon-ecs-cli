// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package iam

import (
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients"
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/config"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

// IAMClient defines methods to interact with the IAM API interface.
type IAMClient interface {
}

// cloudWatchEventsClient implements CloudWatchEventsClient
type iamClient struct {
	client iamiface.IAMAPI
}

// NewClient creates an instance of cloudWatchEventsClient object.
func NewClient(params *config.CliParams) IAMClient {
	client := iam.New(params.Session)
	client.Handlers.Build.PushBackNamed(clients.CustomUserAgentHandler())
	return newClient(params, client)
}

func newClient(params *config.CliParams, client iamiface.IAMAPI) IAMClient {
	return &iamClient{
		client: client,
	}
}
