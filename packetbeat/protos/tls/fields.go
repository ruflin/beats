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

package tls

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
<<<<<<< HEAD
	if err := asset.SetFields("packetbeat", "tls", asset.ModuleFieldsPri, AssetTls); err != nil {
=======
	if err := asset.SetFields("packetbeat", "Tls", asset.ModuleFieldsPri, AssetTls); err != nil {
>>>>>>> Introduce local fields generation
		panic(err)
	}
}

<<<<<<< HEAD
// AssetTls returns asset data.
// This is the base64 encoded gzipped contents of protos/tls.
func AssetTls() string {
=======
<<<<<<< HEAD
// Asset returns asset data
func Asset() string {
>>>>>>> Introduce local fields generation
	return "eJzsW99v47gRfs9fMUgf7g5Ye7u3zQHNQ4EgSdEt9u4W9S7aN4GmxhYvFKkjR/Zq//qCpCRLomQ7ibO9uzpPsS3NfB85P76h5Rk8YHUNJO0FAAmSeA2XH98vLi8AUrTciIKEVtfwtwsAgI/vFzNbIBcrwQE3qAhWAmVq5xdQ/3ftL5yBYjk2ht0fVQVew9rosqjf6V7fvWeDxgqt2vebex+w2mqTdt4fQdj8fcywMQR6BZShAw+F0aS5llBaTOedO/Azy4uGPbyZv728iIBlTKU2Yw+YcO0uJkwjkEutJTJ1HMh/Z0gZmhadwrUmwdylkDELS0QFtuQcrV2VEphK/bUWra2v6dkjw5QV7nZMgTSg4qYqCFPIdYrzmJFBW+bPZfFut7wdYAF87QBWRufAoDC4Ebq0zYVTkLyfJEfKdAzuUXFQg5sC9sp/WqoUjayEWkPwGcIDflYIetWzdynSS1hp06P77g60gUsS/AFp93F4DfiZUE2w5VKgooSjIZdTjDAx+GuJ9qSxZdFs0HjyrXX/QXDvQoWVlKEiDwEEWZR93qV1q8OaOzqAp1llKKWOWHRLAIyUAdhbCvaFwYEVgSPKwrKCbSZ41l2drbAZWiAdWeM6z0sV1iwtjVshysRYdHdJ2bIotCFMEy6KDI0dpceMYdXjyL0Xlhyr2myXg7CwFVJ6hNpFtw/SPtihuQUiZESFvX79ervdzgVTbK7N+jWzVqxVjorsa5J2VjDDciQ0w5fzzxnl8k/9N2d/ObgsOi9MQFXXgJMtkdt/2SzTzk2d9r0lq/HYJy6Msz5zZmciHbwKyzK+Cm2xGKc8TJ6IsiPo865jyYX3RqSYuvDeMRwSG0vE3gb5KpK4F4lQqQv7ODfhQH5GgJugzbQlZ9peTAJgRSFrr4lkFZqkyduk0zmfCygOGujkVgfDzGNoa8f+fJtP06oTMAnt4oXgo1pT1pS8pm8Fj69ArNoQeeU6GVOAeUEVWHJFbdSi6xnpxrUBi02uhKLijdp9hNtMr0vxMNZPRbreM1fhG09AGaMn71QL3OfhC8O+l9JdyOG2NBuEWyfl9NqwIqvg2/vb2++A+w8mscAO8DDzpymKtWJUGkyYXGsjKMtfmGbrEXYewy7lrIIl+m4lFKRiLYjJcZ6NjX1hhzwptFBkk5U2OaOvu33f3t9+B94/1O7n8C5UUfSlY5wYelu9+6IQ5kxBwcze0E1K5S9Jk5dl3ek6HuYWDYLEFUEDwIXiB+ZqxBIZjajHus/8DtWjZyysnx6OEFjv/MXOSibWGVpqjUd5Wyt4pd04wRHTsbJMHXg7nS9UN1b8qk4JU5TId7r0tOsWbIItBWHraUDPT4nUSLSmJUXmltWTSEWq8sQEIzl5FM3IVivWHsMx0BInpPRJiV9LBFXmSzdFahCpmw9XVU9BhCB35I1BW2g1GphcK4XcH2tsBWXTneirquCwG49Vwc8UoWMzSwT/p9oQpl130Neb/1tB+cmGIyamlC4Vx1D72EBaemnlevhkLkPI585+wI3cssoORehvQE92tqUjKd0qBJWS/l+Jj7qY/cHER3wgd1CC9ADe7m6cGLt9zewctw2kQVN/Hq1opI4qbw/Zf67+/Nd6Fxs7k+3ECCaTUPqf2lF8S9wtxje2Nlt3lAnXSlOyxJU2w4gMftP+fkRO79yyh/u753idPRHWi6gNk2KYrl0MbEUT1I+C4G+fQICfCxEPKo3zolxKwZMHrHYD2Gl1Smt2IFP7mxVwOC+xbg3H4vCvxc0ruFvcgDZwf3u3uDlMyYov4xt7KHgX4gs2krsLbTx+4xn2KyxhFO4NigmUTBIaxUhsMAkHX2MYD55xLsrlL8gJbnbm4CdnbnRnJ7AYtn1Onhu27SeZgg/3P9bFZvIU3ON+mr5rSBt0khtVq7GHbPs2Dmk7rktFZkyePapJ3QYzwPtfug29abNmSnw5yaHlzx1b3vyRfplMSiWerRE/KUFe2gvVM78HhW+Oio8pj0e5/lDbcVXI4LqZMYRqdnMPBq7zXCuffc+F4VLOQci0JW+9mZdEp/8fzENhbTnRdw7mxL0iQVVzBmFLNzyo1Jch/23fOTXOqfE7To2VUGs0hRHq+J4xHtSNxTy9io/Q9hKbUvrfWPjx7qoLMVrZtu1l7M3pvC7+cTN7c6zf769+OK3n769+GPiOxqn6LPdFxqnmNO08Tp3HqfM41WNzHqfO4xScx6knuzxrxj+EZjyPU+fUOKfGeZw61utvbZyKvEbTVMIzJuKfSgxlTt+zu8WrP1NaauVW833eMU92vwyCo54tZxINJc6JjZw+6pcBNyqABK4VMaHC09vhiR7vxBv14HCDpqrfNMhRbHrf9Y7kVIxt75j7986dO6HZ+YnB0SPtL+zt0/rdP2/eeofNl5ST9WHUPfR+JmOz5xa58e95nTh2ODvgvLdhBFmxtx1aGmqCE+MLTyyEaYY0cCZ5KZ2S9896MZvNL/4bAAD//3wPukM="
=======
// AssetTls returns asset data.
// This is the base64 encoded gzipped contents of protos/tls.
func AssetTls() string {
	return "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"
>>>>>>> Introduce local fields generation
}
