// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !ent
// +build !ent

package raftutil

import "github.com/hernad/nomad/nomad/state"

func insertEnterpriseState(m map[string][]interface{}, state *state.StateStore) {
}
