// Copyright 2018-2026 the original author or authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/paketo-buildpacks/libdependency/retrieve"
	"github.com/paketo-buildpacks/libdependency/upstream"
	"github.com/paketo-buildpacks/libdependency/versionology"
	"github.com/paketo-buildpacks/packit/v2/cargo"
)

const (
	id   = "opentelemetry-java"
	name = "OpenTelemetry Java Agent"
	purl = "pkg:generic/opentelemetry-java"

	org  = "open-telemetry"
	repo = "opentelemetry-java-instrumentation"

	licenseType = "Apache-2.0"
	licenseURI  = "https://github.com/open-telemetry/opentelemetry-java-instrumentation/blob/main/LICENSE"
)

var (
	assetPattern   = regexp.MustCompile(`^opentelemetry-javaagent\.jar$`)
	versionPattern = regexp.MustCompile(`^v(\d+\.\d+\.\d+)$`)
)

type asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type release struct {
	TagName    string  `json:"tag_name"`
	Prerelease bool    `json:"prerelease"`
	Assets     []asset `json:"assets"`
}

type version struct {
	version *semver.Version
	tag     string
	uri     string
}

func (v version) Version() *semver.Version {
	return v.version
}

func main() {
	retrieve.NewMetadata(id, getAllVersions, generateMetadata)
}

func getAllVersions() (versionology.VersionFetcherArray, error) {
	releases, err := fetchReleases(org, repo)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch releases\n%w", err)
	}

	var versions versionology.VersionFetcherArray
	for _, r := range releases {
		m := versionPattern.FindStringSubmatch(r.TagName)
		if m == nil {
			fmt.Printf("Skipping %s: tag does not match\n", r.TagName)
			continue
		}

		v, err := semver.NewVersion(m[1])
		if err != nil {
			fmt.Printf("Skipping %s: unable to parse version\n", r.TagName)
			continue
		}

		uri := findAsset(r.Assets)
		if uri == "" {
			fmt.Printf("Skipping %s: no matching asset\n", r.TagName)
			continue
		}

		versions = append(versions, version{version: v, tag: r.TagName, uri: uri})
	}

	return versions, nil
}

func generateMetadata(versionFetcher versionology.VersionFetcher) ([]versionology.Dependency, error) {
	v, ok := versionFetcher.(version)
	if !ok {
		return nil, fmt.Errorf("unexpected version type %T", versionFetcher)
	}

	versionString := v.version.String()

	checksum, err := upstream.GetSHA256OfRemoteFile(v.uri)
	if err != nil {
		return nil, fmt.Errorf("unable to checksum %s\n%w", v.uri, err)
	}

	source := fmt.Sprintf("https://github.com/%s/%s/archive/refs/tags/%s.tar.gz", org, repo, v.tag)
	sourceChecksum, err := upstream.GetSHA256OfRemoteFile(source)
	if err != nil {
		return nil, fmt.Errorf("unable to checksum %s\n%w", source, err)
	}

	dependency := cargo.ConfigMetadataDependency{
		Checksum: fmt.Sprintf("sha256:%s", checksum),
		CPE:      fmt.Sprintf("cpe:2.3:a:open-telemetry:opentelemetry-java-agent:%s:*:*:*:*:*:*:*", versionString),
		ID:       id,
		Licenses: []interface{}{
			map[string]string{
				"type": licenseType,
				"uri":  licenseURI,
			},
		},
		Name:           name,
		PURL:           fmt.Sprintf("%s@%s", purl, versionString),
		SHA256:         checksum,
		Source:         source,
		SourceChecksum: fmt.Sprintf("sha256:%s", sourceChecksum),
		Stacks:         []string{"*"},
		URI:            v.uri,
		Version:        versionString,
	}

	return versionology.NewDependencyArray(dependency, "")
}

func findAsset(assets []asset) string {
	for _, a := range assets {
		if assetPattern.MatchString(a.Name) {
			return a.BrowserDownloadURL
		}
	}

	return ""
}

func fetchReleases(owner, repo string) ([]release, error) {
	var all []release
	for page := 1; ; page++ {
		releases, err := fetchPage(owner, repo, page)
		if err != nil {
			return nil, err
		}

		if len(releases) == 0 {
			break
		}

		for _, r := range releases {
			if r.Prerelease {
				continue
			}
			all = append(all, r)
		}
	}

	return all, nil
}

func fetchPage(owner, repo string, page int) ([]release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases?page=%d&per_page=100", owner, repo, page)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", strings.Join([]string{"Bearer", token}, " "))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API returned status %d for %s", resp.StatusCode, url)
	}

	var releases []release
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return nil, err
	}

	return releases, nil
}
