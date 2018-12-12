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

package kibana

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
<<<<<<< HEAD
	if err := asset.SetFields("metricbeat", "kibana", asset.ModuleFieldsPri, AssetKibana); err != nil {
=======
	if err := asset.SetFields("metricbeat", "Kibana", asset.ModuleFieldsPri, AssetKibana); err != nil {
>>>>>>> Introduce local fields generation
		panic(err)
	}
}

// AssetKibana returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kibana.
func AssetKibana() string {
<<<<<<< HEAD
	return "eJzMmM1u4zYQx+9+ioEve0n0AD4UKNoCLYoExbZBD0VhjMWxxYYiVc7Iifv0BSnJkWXKllsLuzoEiT7+85vhfJB5hFc6rOBVb9DiAkC0GFrB8ud4Y7kAUMS515VoZ1fwzQIAoHkIpVO1oQUAF87LOnd2q3cr2KLhcNeTIWRawS4IM4lou+MV/LFkNssHWBYi1fLPBcBWk1G8itqPYLGkHlG45FAFHe/qqr2ToDrV6WuxoPDxbkpuVLK5WoejDqBV4Gv7KLokKEm8zjnrvX7qd3cN2fp8da3VyYMOEY1GHjypUIoVMPm9zik7+7DUO4+ND+JrGjy94GPPT21Z0OYELy8/fZ8kDj+TxK90eHN+CHWj2TP1zqy2it7vZ/cZSwK3be1/YtBWyFs0CUMdQuFYsnndDyZGQyAeLVeh3lApTzxMj6mJk/76HtkTLJD/xEc/YsUE5KRDe/KsnZ3sxjSItOppCNLvXArBsaFYrLhwkoTeOGcIh7JXsH8vSAryIAV1LmxqbRRoBjzaa+6lkQSlTifDPZIyLCehkSJlqGPInc1r78nGYWApD+JpJuPs7sZircsN+VCuudFkBXomoERFIC6Gr8m/DJ6dEEiBAhvv3pg8Q44WmKyCsjaiK0PAOvyKllzNJ4riwNPfNbH0Xo7KgMwU5oCAszk9xPSWgg5R3tNjzQTEghujuSDVl82Scau8y8cqeTilJkTql0aum02D56lB9IFC+7B6xrlqrcjgISuH33dgnKMhtd4ah8M6uEoI8EOwA8EORDugLZTaGM2UO6uGRju8kkrnD1lBOIzJpXhN4OliFpRHAjcWul5vdoIm2xyEzr+9kPhR2vkSZQVjH191AOC3YL1xAI1xOQqpkMZtdoUIp9WP2xAm9cXofwzcgQA2h+nIrP+htdGlli8G/oTvGTijgCvMKRKdxv/ZKcr+4s6phwkLUYW9Zar0rrgzAfclaoc22gvyaOV1RG0rvFeHAvjc9tbc1VZu6lQfUEpz21rnCNTHvGmd52aYvJGnnmkaztVBO5iBrCl0e8Y3snJcOcu0Dst+r/X73IpCEL1t0HyQ4X43T4p/uyePOzq6HimvDJjeiMH3ebCe8F2XdTkRa3Rfd/Ph9deo0C7H/zqnzn/qS++PvtrzcQBLI7dni6yp0nkON20LyOY65MS8yVyoJmOy8Of9V79Vj8ZGVj/dXf5z73pq5M7rASb0rEknHLhDs/iusdIzOH6MgPNBPY6U3h9PYupmdlgrzTL4t1d3jYUQJk9uuBa/ibyQmJX9wX05ojBhks9JegL3bwAAAP//xlKE7A=="
=======
	return "eJzEmM1u4zYQx+9+ioEve0n0ADoUKNoCLYoExbZBD0VhjMSxxYYiVc7Isfv0BfXhyDIVy7syVocgkez//8fhfIh5hFc6pvCqM7S4AhAthlJY/9rcWK8AFHHudSXa2RS+WwEAtA+hdKo2tALgwnnZ5M5u9S6FLRoOdz0ZQqYUMpIgzSSi7Y5T+GvNbNYPsC5EqvXfK4CtJqM4bdQfwWJJA6ZwybGiFHbe1VV3J8J1rjPUYkHh092Y3KRke3VLbnQArQJf20fRJUFJ4nXOyeDj45X315huSFjXWp096CFf6fjm/PjZB6gDXG1Z0OYELy+//Bi1DT+jtkIH+TrPC+neU1tFh+XW+owlgdt2/p8YtBXyFk3EqEcoHEsyufZFQh4sJkMgHi1XoWpQKU/Mi3Mw+T35T3ziaLI2WEaB9uRZO7s4Rkz3VJQWKy7cOM1a08w5QzgGumL6Z0FSkAcpqAfIam0UaAY8+bX34kiCUi+/GX1ShO0gNFLEjHqG3Nm89p5s01It5UE8zmSc3d1YLHWZkQ/lkhtNVmBgASUqAnFN+Nr8SeDZCYEUKJB598bkGXK0wGQVlLURXRkC1uFXtORqPlMUB57+rYll8OFGGZCZQi8VcDanhyY9paBjI+/psWYCYsHMaC5IDWWTaNwq7/KpShp3+hmR+q2V6/v76Hmslb+j0D7snnGu2igyeEzK8fd7MM7RkNpsjcNxHVwlBPgp+EDwgcYHtIVSG6OZcmfV2LTHK6l0/pgUhOOYfBSvGTx9zILyROCmQjfojU7QJNlR6PK7HyR+I+18iZLC1JevLgDgj+DeLgCNcTkKqZDGXXaFCMfVT4OcSX0z+p8DdyCA7DgfmfV/tDG61PLNwJ/wkIAzCrjCnBqi8/g/O0XJP9wv6mHGRlTh/SxWeleWMwP3pdEObXQQ5MnK64m6VrhUhwL43PXW3NVWbupU71BKc9da7xGo93nTLZ7bYfJGngbWNJ6ro3ZwB7K20O0F38TOceUs0yZs+1L797kThSB626B5J8P97j4p/v2ePO7otPSG8sqAGYwYPNwH6wkPuqzLmViT73U3HwB/bxS67fjKs959Dx5BPf6GdN8zZlCP+3aHgKQttnudMbpajiO0CZC4kNTGJOHP5begU2/MJrYgXuRf3EKeWrnLtIQZrWPWQQMWqNkfWpeB4fTbPFzOy2mk+GvqLKZ+dIa90iyj/+D011QIYfYAhWvxm8kLkZE1nJ8fRxRmDNR7kp7B/R8AAP//HqE3bQ=="
>>>>>>> Introduce local fields generation
}
