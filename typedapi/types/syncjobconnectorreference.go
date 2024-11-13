// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// SyncJobConnectorReference type.
//
// https://github.com/elastic/elasticsearch-specification/blob/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827/specification/connector/_types/SyncJob.ts#L31-L40
type SyncJobConnectorReference struct {
	Configuration ConnectorConfiguration `json:"configuration"`
	Filtering     FilteringRules         `json:"filtering"`
	Id            string                 `json:"id"`
	IndexName     string                 `json:"index_name"`
	Language      *string                `json:"language,omitempty"`
	Pipeline      *IngestPipelineParams  `json:"pipeline,omitempty"`
	ServiceType   string                 `json:"service_type"`
	SyncCursor    json.RawMessage        `json:"sync_cursor,omitempty"`
}

func (s *SyncJobConnectorReference) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "configuration":
			if err := dec.Decode(&s.Configuration); err != nil {
				return fmt.Errorf("%s | %w", "Configuration", err)
			}

		case "filtering":
			if err := dec.Decode(&s.Filtering); err != nil {
				return fmt.Errorf("%s | %w", "Filtering", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "index_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "IndexName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexName = o

		case "language":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Language", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Language = &o

		case "pipeline":
			if err := dec.Decode(&s.Pipeline); err != nil {
				return fmt.Errorf("%s | %w", "Pipeline", err)
			}

		case "service_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ServiceType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ServiceType = o

		case "sync_cursor":
			if err := dec.Decode(&s.SyncCursor); err != nil {
				return fmt.Errorf("%s | %w", "SyncCursor", err)
			}

		}
	}
	return nil
}

// NewSyncJobConnectorReference returns a SyncJobConnectorReference.
func NewSyncJobConnectorReference() *SyncJobConnectorReference {
	r := &SyncJobConnectorReference{}

	return r
}