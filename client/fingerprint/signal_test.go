// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fingerprint

import (
	"testing"

	"github.com/hernad/nomad/ci"
	"github.com/hernad/nomad/helper/testlog"
	"github.com/hernad/nomad/nomad/structs"
)

func TestSignalFingerprint(t *testing.T) {
	ci.Parallel(t)

	fp := NewSignalFingerprint(testlog.HCLogger(t))
	node := &structs.Node{
		Attributes: make(map[string]string),
	}

	response := assertFingerprintOK(t, fp, node)
	assertNodeAttributeContains(t, response.Attributes, "os.signals")
}
