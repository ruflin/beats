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

package docker

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "docker", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
