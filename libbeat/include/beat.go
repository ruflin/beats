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

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("libbeat", "Beat", asset.BeatFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded gzipped contents of ./build/fields/beat.yml.
func AssetBeat() string {
	return "eJyEUltSwzAM/M8p9gLtAfzRDxhmuAWjJkoqcOyMrZQpp2f8AJI2E/JnZbW70qo54INvBmcmbQAVtWzwVF4dxzbIpOKdwakBgGfvlMRFtH4cvct96IVtF0FXEktnyxAHshZ8ZafQ28Tx2KDCTOY5wNHIBjSw0+PFR03v/Au5wyRfnz50tbby8lrx8D30wpWlWTEnZ0eVkb+8WxOTFYq1MpFeTDF6Dx5lCFT0NMy8Zi+zrGj9+Z1braXyePt3klMtLjY7Rw5ovetlmENeZxG7m49D8GFlYAh+nvZFXlLTT15tURQ3gLpOEpYsxPU+BdhSzPvNOjm/9C0zXLpJDn6L2wnu2Pqzlhq3gny4jscQ0xEtcbv5ZdLNs3sk3rzRe/bvAAAA///mH/vl"
}
