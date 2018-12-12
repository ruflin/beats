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

package traefik

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
<<<<<<< HEAD
	if err := asset.SetFields("filebeat", "traefik", asset.ModuleFieldsPri, AssetTraefik); err != nil {
=======
	if err := asset.SetFields("filebeat", "Traefik", asset.ModuleFieldsPri, AssetTraefik); err != nil {
>>>>>>> Introduce local fields generation
		panic(err)
	}
}

// AssetTraefik returns asset data.
// This is the base64 encoded gzipped contents of module/traefik.
func AssetTraefik() string {
<<<<<<< HEAD
	return "eJyslb9upDAQxvt9ilH6IEV31RbXRIp0xTWn9MjBA2ut8XDjcSLe/mR2SQiyCZB1t3jn+33zx/Y9nLE/grDC2pwPAGLE4hHuni9f7g4AGn3FphND7gi/DgAAf0gHi1ATQ6fYG9eAnBCuQWCpgdpY9MUBoDZotT8OcffgVItTXlzSd3iEhil01y8JZFxPgxTUTG2eF9eUOeWqqkLv3z+n0Av4uB7JiTLOXxFDCaZWLoTo6N1MytDUVPDIpdHoxNQG+dN/Rodn7N+I9WxvwWdcv/1g7e/TIzz8fPgBF4b0QPWwUVmDTpKeGP8F9FJWFGb/GB1Zcs02O88nBBfaF+Ro4ErwSXzN5ASdLuPP2xVkcKBaHAswYmILdNLIi6rO0Udgu9NG0kRgO3q4EuDthIxjVcAMk/WmWE+MpR2S7kuPToqXXtAnXSpr1HynU3I6wkmkKxh9R85jEbWSMq1pWF2qKhwwMzItCZam22jBU+AKC6U1fz6ba8HD+ckOSh4c44pE3Bpmi3KiedtXFnvocJFUWJVuZhQXEmVbEJvGODUPXQOMtstXZG/I7ck4HbpupC6DWVakt3b382R7URJ8SmedjxqZM1fzyn5nNNbgVTO/ptcNdzkEbm19/ozlfcwfUMg8eVNJja+mmndjObVkehed1EH+YCU3N5MWEZSS2AggX9TB2tRlNgWl9/fQlkHpudlDmislZ6xByjwee8arIifGoZPvVOv6NjVIxRd6H9jghPvSeErdNrvAXyiOaEvVcJ6/j8wqfVyIjSF3o8ouib2X1Uh/q0ZmpWbZ3a6FOcHD/wAAAP//R4Rz0Q=="
=======
	return "eJysl8Fu4zgMhu95CqL3Gih2TznspYsCe9jLoHdDtWlHE1v0UFQ7fvuBHDt1DSpx3OgWB/y/n6RE2Y9wxH4PwgYre9wBiJUG9/DwenrysAMo0RdsO7Hk9vDPDgDgfypDg1ARQ2fYW1eDHBDGIGiohso26LMdQGWxKf1+iHsEZ1qc8+KSvsM91EyhG58oyLheBimomNo0L645c841RYHenx9r6Av4uJ7JibHOj4ihBHMrJ0J0dDajGZqbKgObiPry52StIVcv/rjgLq5/RzmganA2OkIn3GeqgeCRc1uiE1tZZNXHEfsP4vI2K//5wcGPl2d4+vvpLzgxpJ+sFY1FJ6qnNyr73KOT7K0X9GtrUxG3RvagBV0x+3pAcKF9Q472BoHJp0d+RwZG35HzOJjTS8n4K6CXvKCwSGx7Q78aGwlexVdMTtCVefyp4gV/L32twZsWp2JMjLhtSr13pjhGE4GbLR5UB4GbycAoDx8HZJzqAXY4ix+Gy5mrRI9aEsxtp5ozjTXLndMZOezBU+ACM1OW/HWIxNXa+nTs4nALmD5nyc6kwTEuU+LWMFuUAy2P7TXgQaTLxsJmqsKqdBPtv5AoNxmxra0zy9A1wGg7f0f2qWF6JWM9dA15mgx5QeWt3R2LfRLIvBgJXtNZ56NC5sQIX9nvhMYavKmX43zd5s6HwFtbnz5jaR/Lmx4Sd/NcssR3Wyy7cTk1Nb2Ezvmkmp+0LPoGii5zhlh3F4gqM0E6I8Xh+xBdZoIo83ADQ1WZEKRJ3Aggn1WhabSxPwfld+o++SsbIKLuswci6uI2IK9nvYV0uXz63NhCWiqpM6ZGSrw8bBkvBTmxDp18p1rju0mNlF3R+8SG+EWQW0/abbMJfEVxQjdUaB88G5BJpc8Lsbbk7lTZS2Lnslrp79XIpNQiu/u1MCW4+xMAAP//WNoXGA=="
>>>>>>> Introduce local fields generation
}
