// Copyright 2024 Fabian `xx4h` Sylvester
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

package cmd

import (
	"testing"
)

func Test_newCmdConfigGet(t *testing.T) {
	h := newTestingHctl(t)

	var tests = map[string]cmdTest{
		"get existing option": {
			"config get completion.short_names",
			"(?m)^OPTION\\s+VALUE$\n^completion.short_names\\s+true",
			"",
		},
	}

	testCmd(t, h, tests)
}