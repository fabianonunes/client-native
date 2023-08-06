// Code generated with struct_equal_generator; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// Equal checks if two structs of type Bind are equal
//
// By default empty maps and slices are equal to nil:
//
//	var a, b Bind
//	equal := a.Equal(b)
//
// For more advanced use case you can configure these options (default values are shown):
//
//	var a, b Bind
//	equal := a.Equal(b,Options{
//		NilSameAsEmpty: true,
//	})
func (s Bind) Equal(t Bind, opts ...Options) bool {
	opt := getOptions(opts...)

	if !s.BindParams.Equal(t.BindParams, opt) {
		return false
	}

	if s.Address != t.Address {
		return false
	}

	if !equalPointers(s.Port, t.Port) {
		return false
	}

	if !equalPointers(s.PortRangeEnd, t.PortRangeEnd) {
		return false
	}

	return true
}

// Diff checks if two structs of type Bind are equal
//
// By default empty arrays, maps and slices are equal to nil:
//
//	var a, b Bind
//	diff := a.Diff(b)
//
// For more advanced use case you can configure the options (default values are shown):
//
//	var a, b Bind
//	equal := a.Diff(b,Options{
//		NilSameAsEmpty: true,
//	})
func (s Bind) Diff(t Bind, opts ...Options) map[string][]interface{} {
	opt := getOptions(opts...)

	diff := make(map[string][]interface{})
	if !s.BindParams.Equal(t.BindParams, opt) {
		diff["BindParams"] = []interface{}{s.BindParams, t.BindParams}
	}

	if s.Address != t.Address {
		diff["Address"] = []interface{}{s.Address, t.Address}
	}

	if !equalPointers(s.Port, t.Port) {
		diff["Port"] = []interface{}{s.Port, t.Port}
	}

	if !equalPointers(s.PortRangeEnd, t.PortRangeEnd) {
		diff["PortRangeEnd"] = []interface{}{s.PortRangeEnd, t.PortRangeEnd}
	}

	return diff
}