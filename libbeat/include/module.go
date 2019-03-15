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
	if err := asset.SetFields("libbeat", "Module", asset.BeatFieldsPri, AssetModule); err != nil {
		panic(err)
	}
}

// AssetModule returns asset data.
// This is the base64 encoded gzipped contents of ./build/fields/module.yml.
func AssetModule() string {
	return "eJysl81u4zgMx+95CqJ7jrG97eawi0EzmBaDmRZo5xzQFp1oIouGpLRJn34gW7GU2kmctjkU0Af5o0jqL3cKa9rNoFC8ERMAJ52iGdz4IdSGn6UgAxU5FOhwAiDIFkbWTrKewX8TAIAfYRVKw1XrqTO1gEKQgHwHbkV+sGg2LPYu/c6CrGWTTQBKSUrY2aTxOwWNFYXYstrwbypcJkWzCEBbrGofbFiZbsPCQIj+9xMrAi6bOIIJSA3fmJeK2hNnLfgv+KIkWijZACsRojqIyYefdYE1Rw0gt6tpBugdhJka3So5Rbq5kkuDbaDObGhyjCG1dagLWnSnP43Z74/ZupjkVy5jJRZjaRUWK6lp4QmjYMEgSwzGsvAZpcJcKul2i1fW44DHrMZSQ6eNLVyvy8dyDC0l61GMg61v/Qc9EFysmy4NgjDfjwcuV7sG1qGzULBSVDgSrRq0a/5q2xUbtyhYl3I5gxKV9alEXazY7HnTTgwSLUgP3IUVT7k0vKnDTGpyoCCsHUpNJuZ1OE1JqoZNBstxBFXhki6neav0Kl1IfWM5Dvp+nMKclPWS+XQ/v5/BLb+AY6iw9jpr6f9eLJz7Bk+m24lFu7qm3QubNONH1Nz/7nyuoA0h23fuiq2LfXvbjgac3OmS025l07wM3hz2KpM0qJ8fbM/wYny9eQxTgIb8w9KOMipstqtUFvY9+aw0WzxNs55Gy+yg18NBTnf68dLIV+r3es6sCPXI9JYxI9ICJmXvc9lm+UaqPrJf0e7hvrr+Z379979X48K5f4SGAFKXbKqmQbuqrzc5GU2ObKz993RuwHFc7z5wDr9WolNIK39ayaLRWTU7CPp0nd+mu2YxfNUvukBJBmoWjet+ZT1qM6Cb7yU9sIBfd/M+yP+1NRafd6josQ9j0ZPZD8JY0JEUthJ1XgrHgYLmVlj3Sag1u+ZqfBoucTnMPPP4vDefndsjST330n6c2/oNChMuc5SXh25iwG9YjMLS/Q8xJATRN1ykArQd+9YHQkZbKjYOc3Xqvf8TAAD//1B1IPQ="
}
