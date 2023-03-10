// Copyright 2022-2023 VMware, Inc.
//
// This product is licensed to you under the BSD-2 license (the "License").
// You may not use this product except in compliance with the BSD-2 License.
// This product may include a number of subcomponents with separate copyright
// notices and license terms. Your use of these subcomponents is subject to
// the terms and conditions of the subcomponent's license, as noted in the
// LICENSE file.
//
// SPDX-License-Identifier: BSD-2-Clause

package config

import (
	"fmt"

	"github.com/rdimitrov/go-tuf-metadata/metadata"
	"github.com/rdimitrov/go-tuf-metadata/metadata/fetcher"
)

type UpdaterConfig struct {
	// TUF configuration
	MaxRootRotations   int64
	MaxDelegations     int
	RootMaxLength      int64
	TimestampMaxLength int64
	SnapshotMaxLength  int64
	TargetsMaxLength   int64
	// Updater configuration
	Fetcher               fetcher.Fetcher
	LocalTrustedRootPath  string
	LocalTrustedRootBytes []byte
	LocalMetadataDir      string
	LocalTargetsDir       string
	RemoteMetadataURL     string
	RemoteTargetsURL      string
	DisableLocalCache     bool
	PrefixTargetsWithHash bool
}

// New creates a new UpdaterConfig instance used by the Updater to
// store configuration
func New(rootPath ...string) *UpdaterConfig {
	// if no rootPath is provided, default to looking for the root.json in the current working directory
	trustedRootPath := fmt.Sprintf("%s.json", metadata.ROOT)
	if len(rootPath) != 0 {
		trustedRootPath = rootPath[0]
	}
	return &UpdaterConfig{
		// TUF configuration
		MaxRootRotations:   32,
		MaxDelegations:     32,
		RootMaxLength:      512000,  // bytes
		TimestampMaxLength: 16384,   // bytes
		SnapshotMaxLength:  2000000, // bytes
		TargetsMaxLength:   5000000, // bytes
		// Updater configuration
		Fetcher:               &fetcher.DefaultFetcher{}, // use the default built-in download fetcher
		LocalTrustedRootPath:  trustedRootPath,           // local path to trusted root.json
		DisableLocalCache:     false,                     // enable local caching of trusted metadata
		PrefixTargetsWithHash: true,                      // use hash-prefixed target files with consistent snapshots
	}
}
