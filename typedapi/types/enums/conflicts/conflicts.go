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

// Package conflicts
package conflicts

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/48e2d9de9de2911b8cb1cf715e4bc0a2b1f4b827/specification/_types/common.ts#L184-L193
type Conflicts struct {
	Name string
}

var (
	Abort = Conflicts{"abort"}

	Proceed = Conflicts{"proceed"}
)

func (c Conflicts) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *Conflicts) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "abort":
		*c = Abort
	case "proceed":
		*c = Proceed
	default:
		*c = Conflicts{string(text)}
	}

	return nil
}

func (c Conflicts) String() string {
	return c.Name
}
