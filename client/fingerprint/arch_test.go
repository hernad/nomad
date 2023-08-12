// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fingerprint

import (
	"testing"

	"github.com/hernad/nomad/ci"
	"github.com/hernad/nomad/client/config"
	"github.com/hernad/nomad/helper/testlog"
	"github.com/hernad/nomad/nomad/structs"
)

func TestArchFingerprint(t *testing.T) {
	ci.Parallel(t)

	f := NewArchFingerprint(testlog.HCLogger(t))
	node := &structs.Node{
		Attributes: make(map[string]string),
	}

	request := &FingerprintRequest{Config: &config.Config{}, Node: node}
	var response FingerprintResponse
	err := f.Fingerprint(request, &response)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if !response.Detected {
		t.Fatalf("expected response to be applicable")
	}

	assertNodeAttributeContains(t, response.Attributes, "cpu.arch")
}
