// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aip0131

import (
	"fmt"

	"github.com/googleapis/api-linter/lint"
	"github.com/jhump/protoreflect/desc"
)

// Get messages should have a properly named Request message.
var requestMessageName = &lint.MethodRule{
	Name: lint.NewRuleName("core", "0131", "request-message", "name"),
	URI:  "https://aip.dev/131#guidance",
	LintMethod: func(m *desc.MethodDescriptor) []lint.Problem {
		// We only care about Get methods for the purpose of this rule;
		// ignore everything else.
		if !isGetMethod(m) {
			return nil
		}

		// Rule check: Establish that for methods such as `GetFoo`, the request
		// message is named `GetFooRequest`.
		if got, want := m.GetInputType().GetName(), m.GetName()+"Request"; got != want {
			return []lint.Problem{{
				Message: fmt.Sprintf(
					"Get RPCs should have a request message named after the RPC, such as %q.",
					want,
				),
				Suggestion: want,
				Descriptor: m,
			}}
		}

		return nil
	},
}

// Get messages should use the resource as the response message
var responseMessageName = &lint.MethodRule{
	Name: lint.NewRuleName("core", "0131", "response-message", "name"),
	URI:  "https://aip.dev/131#guidance",
	LintMethod: func(m *desc.MethodDescriptor) []lint.Problem {
		// We only care about Get methods for the purpose of this rule;
		// ignore everything else.
		if !isGetMethod(m) {
			return nil
		}

		// Rule check: Establish that for methods such as `GetFoo`, the response
		// message is named `Foo`.
		if got, want := m.GetOutputType().GetName(), m.GetName()[3:]; got != want {
			return []lint.Problem{{
				Message: fmt.Sprintf(
					"Get RPCs should have the corresponding resource as the response message, such as %q.",
					want,
				),
				Suggestion: want,
				Descriptor: m,
			}}
		}

		return nil
	},
}