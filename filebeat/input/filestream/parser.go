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

package filestream

import (
	"errors"
	"fmt"
	"io"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/reader"
	"github.com/elastic/beats/v7/libbeat/reader/multiline"
	"github.com/elastic/beats/v7/libbeat/reader/readfile"
	"github.com/elastic/beats/v7/libbeat/reader/readjson"
)

var (
	ErrNoSuchParser = errors.New("no such parser")
)

// parser transforms or translates the Content attribute of a Message.
// They are able to aggregate two or more Messages into a single one.
type parser interface {
	io.Closer
	Next() (reader.Message, error)
}

type parserConfig struct {
	maxBytes       int
	lineTerminator readfile.LineTerminator
}

type postProcesser interface {
	PostProcess(*reader.Message)
	Name() string
}

func newParsers(in reader.Reader, pCfg parserConfig, c []common.ConfigNamespace) (parser, error) {
	p := in

	parserCheck := make(map[string]int)
	for _, ns := range c {
		name := ns.Name()
		switch name {
		case "multiline":
			parserCheck["multiline"]++
			var config multiline.Config
			cfg := ns.Config()
			err := cfg.Unpack(&config)
			if err != nil {
				return nil, fmt.Errorf("error while parsing multiline parser config: %+v", err)
			}
			p, err = multiline.New(p, "\n", pCfg.maxBytes, &config)
			if err != nil {
				return nil, fmt.Errorf("error while creating multiline parser: %+v", err)
			}
		case "ndjson":
			parserCheck["ndjson"]++
			var config readjson.Config
			cfg := ns.Config()
			err := cfg.Unpack(&config)
			if err != nil {
				return nil, fmt.Errorf("error while parsing ndjson parser config: %+v", err)
			}
			p = readjson.NewJSONReader(p, &config)
		default:
			return nil, fmt.Errorf("%s: %s", ErrNoSuchParser, name)
		}
	}

	// This is a temporary check. In the long run configuring multiple parsers with the same
	// type is going to be supported.
	if count, ok := parserCheck["multiline"]; ok && count > 1 {
		return nil, fmt.Errorf("only one parser is allowed for multiline, got %d", count)
	}
	if count, ok := parserCheck["ndjson"]; ok && count > 1 {
		return nil, fmt.Errorf("only one parser is allowed for ndjson, got %d", count)
	}

	return p, nil
}

func newPostProcessors(c []common.ConfigNamespace) []postProcesser {
	var pp []postProcesser
	for _, ns := range c {
		name := ns.Name()
		switch name {
		case "ndjson":
			var config readjson.Config
			cfg := ns.Config()
			cfg.Unpack(&config)
			pp = append(pp, readjson.NewJSONPostProcessor(&config))
		default:
			continue
		}
	}

	return pp
}

func validateParserConfig(pCfg parserConfig, c []common.ConfigNamespace) error {
	for _, ns := range c {
		name := ns.Name()
		switch name {
		case "multiline":
			var config multiline.Config
			cfg := ns.Config()
			err := cfg.Unpack(&config)
			if err != nil {
				return fmt.Errorf("error while parsing multiline parser config: %+v", err)
			}
		case "ndjson":
			var config readjson.Config
			cfg := ns.Config()
			err := cfg.Unpack(&config)
			if err != nil {
				return fmt.Errorf("error while parsing ndjson parser config: %+v", err)
			}
		default:
			return fmt.Errorf("%s: %s", ErrNoSuchParser, name)
		}
	}

	return nil
}
