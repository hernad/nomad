// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loader

import (
	"github.com/hernad/nomad/plugins/base"
	"github.com/hernad/nomad/plugins/device"
	"github.com/hernad/nomad/plugins/drivers"
)

var (
	// AgentSupportedApiVersions is the set of API versions supported by the
	// Nomad agent by plugin type.
	AgentSupportedApiVersions = map[string][]string{
		base.PluginTypeDevice: {device.ApiVersion010},
		base.PluginTypeDriver: {drivers.ApiVersion010},
	}
)
