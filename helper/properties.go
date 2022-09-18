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

package helper

import (
	"fmt"
	"strings"

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak/bard"
	"github.com/paketo-buildpacks/libpak/bindings"
)

type Properties struct {
	Bindings libcnb.Bindings
	Logger   bard.Logger
}

func (p Properties) Execute() (map[string]string, error) {
	b, ok, err := bindings.ResolveOne(p.Bindings, bindings.OfType("OpenTelemetry"))
	if err != nil {
		return nil, fmt.Errorf("unable to resolve binding OpenTelemetry\n%w", err)
	} else if !ok {
		return nil, nil
	}

	p.Logger.Info("Configuring OpenTelemetry properties")

	e := make(map[string]string, len(b.Secret))
	for k, v := range b.Secret {
		s := strings.ToUpper(k)
		s = strings.ReplaceAll(s, "-", "_")
		s = strings.ReplaceAll(s, ".", "_")

		e[s] = v
	}

	return e, nil
}
