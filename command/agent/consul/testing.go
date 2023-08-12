// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consul

import (
	"context"

	"github.com/hernad/nomad/client/serviceregistration"
	"github.com/hernad/nomad/nomad/structs"
)

func NoopRestarter() serviceregistration.WorkloadRestarter {
	return noopRestarter{}
}

type noopRestarter struct{}

func (noopRestarter) Restart(ctx context.Context, event *structs.TaskEvent, failure bool) error {
	return nil
}
