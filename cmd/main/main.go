package main

import (
	"os"

	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
	"github.com/paketo-buildpacks/opentelemetry/opentelemetry"
)

func main() {
	libpak.Main(
		opentelemetry.Detect{},
		opentelemetry.Build{Logger: bard.NewLogger(os.Stdout)},
	)
}
