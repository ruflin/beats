// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package haproxy

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
<<<<<<< HEAD
	if err := asset.SetFields("metricbeat", "haproxy", asset.ModuleFieldsPri, AssetHaproxy); err != nil {
=======
	if err := asset.SetFields("metricbeat", "Haproxy", asset.ModuleFieldsPri, AssetHaproxy); err != nil {
>>>>>>> Introduce local fields generation
		panic(err)
	}
}

// AssetHaproxy returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/haproxy.
func AssetHaproxy() string {
<<<<<<< HEAD
	return "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"
=======
	return "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"
>>>>>>> Introduce local fields generation
}
