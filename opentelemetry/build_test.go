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

package opentelemetry_test

import (
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/opentelemetry/opentelemetry"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx libcnb.BuildContext
	)

	it("contributes Java agent API <= 0.6", func() {
		ctx.Plan.Entries = append(ctx.Plan.Entries, libcnb.BuildpackPlanEntry{Name: "opentelemetry-java"})
		ctx.Buildpack.Metadata = map[string]interface{}{
			"dependencies": []map[string]interface{}{
				{
					"id":      "opentelemetry-java",
					"version": "1.15.0",
					"stacks":  []interface{}{"test-stack-id"},
				},
			},
		}
		ctx.Buildpack.API = "0.6"
		ctx.StackID = "test-stack-id"

		result, err := opentelemetry.Build{}.Build(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Layers).To(HaveLen(1))
		Expect(result.Layers[0].Name()).To(Equal("opentelemetry-java"))
	})

	it("contributes Java agent API >= 0.7", func() {
		ctx.Plan.Entries = append(ctx.Plan.Entries, libcnb.BuildpackPlanEntry{Name: "opentelemetry-java"})
		ctx.Buildpack.Metadata = map[string]interface{}{
			"dependencies": []map[string]interface{}{
				{
					"id":      "opentelemetry-java",
					"version": "1.15.0",
					"stacks":  []interface{}{"test-stack-id"},
					"cpes":    []interface{}{"cpe:2.3:a:open-telemetry:opentelemetry-java-agent:1.15.0:*:*:*:*:*:*:*"},
					"purl":    "pkg:generic/opentelemetry-java@1.15.0",
				},
			},
		}
		ctx.Buildpack.API = "0.7"
		ctx.StackID = "test-stack-id"

		result, err := opentelemetry.Build{}.Build(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Layers).To(HaveLen(1))
		Expect(result.Layers[0].Name()).To(Equal("opentelemetry-java"))
	})

}
