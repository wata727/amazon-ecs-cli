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

package cloudwatchevents

import (
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients"
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/config"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
)

// CloudWatchEventsClient defines methods to interact with the CloudWatchEvents API interface.
type CloudWatchEventsClient interface{}

// cloudWatchEventsClient implements CloudWatchEventsClient
type cloudWatchEventsClient struct {
	client cloudwatcheventsiface.CloudWatchEventsAPI
}

// NewClient creates an instance of cloudWatchEventsClient object.
func NewClient(params *config.CliParams) CloudWatchEventsClient {
	client := cloudwatchevents.New(params.Session)
	client.Handlers.Build.PushBackNamed(clients.CustomUserAgentHandler())
	return newClient(params, client)
}

func newClient(params *config.CliParams, client cloudwatcheventsiface.CloudWatchEventsAPI) CloudWatchEventsClient {
	return &cloudWatchEventsClient{
		client: client,
	}
}
