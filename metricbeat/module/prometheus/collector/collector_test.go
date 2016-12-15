// +build !integration

package collector

import (
	"testing"

	"github.com/docker/docker/pkg/testutil/assert"
	"github.com/elastic/beats/libbeat/common"
)

func TestDecodeLine(t *testing.T) {
	tests := []struct {
		Line   string
		Output common.MapStr
	}{
		{
			Line: `http_request_duration_microseconds{handler="query",quantile="0.99"} 17`,
			Output: common.MapStr{
				"http_request_duration_microseconds{handler=query,quantile=0.99}": common.MapStr{
					"value":    int64(17),
					"handler":  "query",
					"key":      "http_request_duration_microseconds",
					"quantile": 0.99,
				},
			},
		},
		{
			Line:   `http_request_duration_microseconds{handler="query",quantile="0.99"} NaN`,
			Output: nil,
		},
		{
			Line: `http_request_duration_microseconds{handler="query",quantile="0.99"} 13.2`,
			Output: common.MapStr{
				"http_request_duration_microseconds{handler=query,quantile=0.99}": common.MapStr{
					"value":    13.2,
					"handler":  "query",
					"key":      "http_request_duration_microseconds",
					"quantile": 0.99,
				},
			},
		},
	}

	for _, test := range tests {
		output := decodeLine(test.Line)
		assert.DeepEqual(t, output, test.Output)
	}
}
