// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"testing"

	"github.com/hernad/nomad/acl"
	"github.com/hernad/nomad/ci"
	"github.com/hernad/nomad/client/config"
	"github.com/hernad/nomad/helper/pointer"
	"github.com/hernad/nomad/nomad"
	"github.com/hernad/nomad/nomad/mock"
	"github.com/hernad/nomad/nomad/structs"
	"github.com/hernad/nomad/testutil"
	"github.com/shoenig/test/must"
)

func TestNodeMeta_ACL(t *testing.T) {
	ci.Parallel(t)

	s, _, cleanupS := nomad.TestACLServer(t, nil)
	defer cleanupS()
	testutil.WaitForLeader(t, s.RPC)

	c1, cleanup := TestClient(t, func(c *config.Config) {
		c.ACLEnabled = true
		c.Servers = []string{s.GetConfig().RPCAddr.String()}
	})
	defer cleanup()

	// Dynamic node metadata endpoints should fail without auth
	applyReq := &structs.NodeMetaApplyRequest{
		NodeID: c1.NodeID(),
		Meta: map[string]*string{
			"foo": pointer.Of("bar"),
		},
	}

	resp := structs.NodeMetaResponse{}
	err := c1.ClientRPC("NodeMeta.Apply", applyReq, &resp)
	must.ErrorContains(t, err, structs.ErrPermissionDenied.Error())

	readReq := &structs.NodeSpecificRequest{
		NodeID: c1.NodeID(),
	}
	err = c1.ClientRPC("NodeMeta.Read", readReq, &resp)
	must.ErrorContains(t, err, structs.ErrPermissionDenied.Error())

	// Create a token to make it work
	policyGood := mock.NodePolicy(acl.PolicyWrite)
	tokenGood := mock.CreatePolicyAndToken(t, s.State(), 1009, "meta", policyGood)

	applyReq.AuthToken = tokenGood.SecretID
	err = c1.ClientRPC("NodeMeta.Apply", applyReq, &resp)
	must.NoError(t, err)
	must.Eq(t, "bar", resp.Meta["foo"])

	readReq.AuthToken = tokenGood.SecretID
	err = c1.ClientRPC("NodeMeta.Read", readReq, &resp)
	must.NoError(t, err)
}

func TestNodeMeta_Validation(t *testing.T) {
	ci.Parallel(t)

	s, cleanupS := nomad.TestServer(t, nil)
	defer cleanupS()
	testutil.WaitForLeader(t, s.RPC)

	c1, cleanup := TestClient(t, func(c *config.Config) {
		c.Servers = []string{s.GetConfig().RPCAddr.String()}
	})
	defer cleanup()

	applyReq := &structs.NodeMetaApplyRequest{
		NodeID: c1.NodeID(),
		Meta:   map[string]*string{},
	}

	resp := struct{}{}

	// An empty map is an error
	err := c1.ClientRPC("NodeMeta.Apply", applyReq, &resp)
	must.ErrorContains(t, err, "missing required Meta")

	// empty keys are prohibited
	applyReq.Meta[""] = pointer.Of("bad")
	err = c1.ClientRPC("NodeMeta.Apply", applyReq, &resp)
	must.ErrorContains(t, err, "empty")

	// * is prohibited in keys
	delete(applyReq.Meta, "")
	applyReq.Meta["*"] = pointer.Of("bad")
	err = c1.ClientRPC("NodeMeta.Apply", applyReq, &resp)
	must.ErrorContains(t, err, "*")
}
