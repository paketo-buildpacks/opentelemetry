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

package opentelemetry

import (
	"github.com/buildpacks/libcnb"

	"github.com/paketo-buildpacks/libpak/bindings"
)

const (
	// BindingType is used to resolve bindings containing an OpenTelemetry configuration file for the Java Agent
	BindingType = "opentelemetry"
)

func getConfigurationFileFromBindings(binds libcnb.Bindings) string {
	var path string
	for _, bind := range bindings.Resolve(binds, bindings.OfType(BindingType)) {
		for k := range bind.Secret {
			if p, ok := bind.SecretFilePath(k); ok {
				path = p
			}
		}
	}
	return path
}
