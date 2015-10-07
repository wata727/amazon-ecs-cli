// Copyright 2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

package config

import (
	"strings"
	"testing"
)

func TestNewDefaultDestination(t *testing.T) {
	dest, err := newDefaultDestionation()
	if err != nil {
		t.Errorf("Error creating new config path: ", err)
	}

	if !strings.HasSuffix(dest.Path, "ecs") {
		t.Errorf("Invalid suffix for ecs config path: ", dest.Path)
	}

	if !dest.Mode.IsDir() {
		t.Errorf("user home directory expected to have directory in mode")
	}
}
