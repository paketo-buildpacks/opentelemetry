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
	"os"
	"testing"

	"github.com/ThomasVitale/buildpacks-opentelemetry/opentelemetry"
	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx    libcnb.DetectContext
		detect opentelemetry.Detect
	)

	context("BP_OPENTELEMETRY_ENABLED is not set", func() {
		it.Before(func() {
			Expect(os.Unsetenv("BP_OPENTELEMETRY_ENABLED")).To(Succeed())
		})

		it("fails without BP_OPENTELEMETRY_ENABLED", func() {
			Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{}))
		})
	})

	context("BP_OPENTELEMETRY_ENABLED is set", func() {
		it.Before(func() {
			Expect(os.Setenv("BP_OPENTELEMETRY_ENABLED", "true")).To(Succeed())
		})

		it.After(func() {
			Expect(os.Unsetenv("BP_OPENTELEMETRY_ENABLED")).To(Succeed())
		})

		it("passes with BP_OPENTELEMETRY_ENABLED set to true", func() {
			Expect(detect.Detect(ctx)).To(Equal(libcnb.DetectResult{
				Pass: true,
				Plans: []libcnb.BuildPlan{
					{
						Provides: []libcnb.BuildPlanProvide{
							{Name: "opentelemetry-java"},
						},
						Requires: []libcnb.BuildPlanRequire{
							{Name: "opentelemetry-java"},
							{Name: "jvm-application"},
						},
					},
				},
			}))
		})
	})
}
