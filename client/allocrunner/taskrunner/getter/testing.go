// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package getter

import (
	"os"
	"path/filepath"
	"testing"

	cconfig "github.com/hernad/nomad/client/config"
	"github.com/hernad/nomad/helper/pointer"
	"github.com/hernad/nomad/helper/testlog"
	sconfig "github.com/hernad/nomad/nomad/structs/config"
	"github.com/shoenig/test/must"
)

// TestSandbox creates a real artifact downloader configured via the default
// artifact config. It is good enough for tests so no mock implementation exists.
func TestSandbox(t *testing.T) *Sandbox {
	defaultConfig := sconfig.DefaultArtifactConfig()
	defaultConfig.DecompressionSizeLimit = pointer.Of("1MB")
	defaultConfig.DecompressionFileCountLimit = pointer.Of(10)
	ac, err := cconfig.ArtifactConfigFromAgent(defaultConfig)
	must.NoError(t, err)
	return New(ac, testlog.HCLogger(t))
}

// SetupDir creates a directory suitable for testing artifact - i.e. it is
// owned by the user under which nomad runs.
//
// returns alloc_dir, task_dir
func SetupDir(t *testing.T) (string, string) {
	allocDir := t.TempDir()
	taskDir := filepath.Join(allocDir, "local")
	topDir := filepath.Dir(allocDir)

	must.NoError(t, os.Chmod(topDir, 0o755))

	must.NoError(t, os.Chmod(allocDir, 0o755))

	must.NoError(t, os.Mkdir(taskDir, 0o755))
	must.NoError(t, os.Chmod(taskDir, 0o755))
	return allocDir, taskDir
}
