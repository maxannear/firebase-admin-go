// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package internal contains functionality that is only accessible from within the Admin SDK.
package internal

import (
	"go/build"
	"path/filepath"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// AuthConfig represents the configuration of Firebase Auth service.
type AuthConfig struct {
	Opts      []option.ClientOption
	Creds     *google.DefaultCredentials
	ProjectID string
}

// StorageConfig represents the configuration of Google Cloud Storage service.
type StorageConfig struct {
	Opts   []option.ClientOption
	Bucket string
}

// Resource returns the absolute path to the specified test resource file.
func Resource(name string) string {
	p := []string{build.Default.GOPATH, "src", "firebase.google.com", "go", "testdata", name}
	return filepath.Join(p...)
}
