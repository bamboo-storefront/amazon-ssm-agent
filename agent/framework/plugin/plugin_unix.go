// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// +build darwin freebsd linux netbsd openbsd

package plugin

import (
	"github.com/aws/amazon-ssm-agent/agent/context"
	"github.com/aws/amazon-ssm-agent/agent/log"
)

// IsPluginSupportedForCurrentPlatform always returns true because currently, there is no plugin that particular
// linux version doesn't support while other linux version does.
func IsPluginSupportedForCurrentPlatform(log log.T, pluginID string) (bool, string) {
	return true, ""
}

// loadPlatformDependentPlugins registers platform dependent plugins
func loadPlatformDependentPlugins(context context.T) PluginRegistry {
	return PluginRegistry{}
}
