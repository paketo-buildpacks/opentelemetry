package main

import (
	"os"

	"github.com/ThomasVitale/buildpacks-opentelemetry/opentelemetry"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
)

func main() {
	libpak.Main(
		opentelemetry.Detect{},
		opentelemetry.Build{Logger: bard.NewLogger(os.Stdout)},
	)
}
