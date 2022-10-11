/*
 * Copyright 2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package helper_test

import (
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/opentelemetry/helper"
)

func testProperties(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		p helper.Properties
	)

	it("does not contribute properties if no binding exists", func() {
		Expect(p.Execute()).To(BeNil())
	})

	it("contributes properties if binding exists as a config tree", func() {
		p.Bindings = libcnb.Bindings{
			{
				Name: "test-binding",
				Type: "OpenTelemetry",
				Secret: map[string]string{
					"otel.test.my-key":            "test-value",
					"otel.testagain.my-other-key": "other test value",
				},
			},
		}

		Expect(p.Execute()).To(Equal(map[string]string{
			"OTEL_TEST_MY_KEY":            "test-value",
			"OTEL_TESTAGAIN_MY_OTHER_KEY": "other test value",
		}))
	})
}
