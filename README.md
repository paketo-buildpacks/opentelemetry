# `ghcr.io/ThomasVitale/buildpacks-opentelemetry`

The Paketo OpenTelemetry Buildpack is a Cloud Native Buildpack that contributes and configures the OpenTelemetry Agent.

## Behavior

This buildpack will participate if all the following conditions are met

* The `$BP_OPENTELEMETRY_ENABLED` is set to a truthy value (i.e. `true`, `t`, `1` ignoring case)

The buildpack will do the following for Java applications:

* Contributes the OpenTelemetry Java agent to a layer and configures `$JAVA_TOOL_OPTIONS` to use it

## Usage

TBD

## Bindings

The buildpack optionally accepts the following bindings:

### Type: `dependency-mapping`

| Key                   | Value   | Description                                                                                       |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------- |
| `<dependency-digest>` | `<uri>` | If needed, the buildpack will fetch the dependency with digest `<dependency-digest>` from `<uri>` |

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
