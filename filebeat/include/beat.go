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
	if err := asset.SetFields("filebeat", "Beat", asset.BeatFieldsPri, AssetBeat); err != nil {
		panic(err)
	}
}

// AssetBeat returns asset data.
// This is the base64 encoded gzipped contents of ./build/fields/beat.yml.
func AssetBeat() string {
	return "eJy0VbFy2zAM3f0V2LI08u6hS+8ydWt2m5IgileKUADIjv6+R1Kya9lxclc1k0UC7z3gIeAz/MZxB57sBkCdetzBT7LQOI9QUVAMugGoUSp2vToKO/i+AQD4QUGNCxJzc7h3AaXYADQOfS27TYp7hmA6TBRFDCt6o226AdCxx11UcCKupzPGt8Ex1jtojBecTu8IiH+vLWbuhqmDU+uqFrTNUuBkBBhNXcBr6ySrSjUl2THMlEJ+UISoCZTSYRJ5ZnghBnw3XR87c9geDW892a2MotgVnuyhuK1TaOAKC1PXjCJrFfsrocKEelMyWcAjBj3XDVuQ+B0D74ikphHUK3Gegv0nGzJm0sPYEyvW2QtRwypgdCFElNF0a3UoTm5GhFOLIXXBBTubjhwLl29QmQAlwpNoTYM+AXH6jcxP1/Jc6Actoq4bOcrDF7qSAFJhS78uXlkMyEbxelCdQGplHsqj8QOC9Fi5xmF95miI0/0hUhyAkghwIR1mcsEqHVKTDl+cxxKNxqY0zk7jflV1klYIvg0YKlxtQmZACENXImdBTia6xVyMkhaGqZx3Oq43pRMg4LuyqeJ8Jl9iZ3p2xE7H+1Lm29WkzICzMZnnUTcEjxgz9t6U6Nfcoe3QmfzvYkofncpEj0357zJmooWMnqlCkaJnsrze8ogCIsHsxwR/Z3M23lj5DOz+i5NSZ4Z7VreqfcEoPQXBYnqB9x6DXbyaxjszi4jP1zK1pHosylFxDuqcZZN1pt11RTsI8t7Y/NpfSCzT0E8n86uev+ZEknPNtynLpEtaM3i/jz/PN0sHr+TFJSWoxV8Z9/x+4Ox57U1QoK3Ry+r9eBPN1B3Vg8dPLMgAV6EP2x5Hfa+uQ1HT9V8Crxij4A/Q/wQAAP//fzIAxw=="
}
