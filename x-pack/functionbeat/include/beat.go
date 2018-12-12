// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("functionbeat", "Beat", asset.BeatFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded gzipped contents of ./build/fields/beat.yml.
func AssetBeat() string {
	return "eJzSVchOrbRSSCvNSy7JzM9LSk0s4VJQKMksyUm1UnBDFU1JLU4uyiwAiVhxKSikZabmpBRbcQECAAD//xC0Fsg="
}
