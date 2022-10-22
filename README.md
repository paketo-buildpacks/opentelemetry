# Paketo Buildpack for OpenTelemetry

## `gcr.io/paketo-buildpacks/opentelemetry`

The OpenTelemetry Buildpack is a Cloud Native Buildpack that contributes and configures the OpenTelemetry Agent.

## Behavior

This buildpack will participate if all the following conditions are met

* The `$BP_OPENTELEMETRY_ENABLED` is set to a truthy value (i.e. `true`, `t`, `1` ignoring case)

At build time, the buildpack will do the following for Java applications:

* Contributes the OpenTelemetry Java agent to a layer and configures `$JAVA_TOOL_OPTIONS` to use it.
* By default, the agent is configured to be disabled (`OTEL_JAVAAGENT_ENABLED=false`).
* By default, the metrics exporting feature of the agent is configured to be disabled (`OTEL_METRICS_EXPORTER=none`).

At run time, the buildpack will do the following for Java applications:

* If a binding with `type` of `opentelemetry` exists, it uses the config tree from the binding to configure the OpenTelemetry Java agent. This is recommended when the configuration involves secrets of any kind.

## Usage

Once you enable the OpenTelemetry buildpack at build-time, you can configure it at run-time via environment variables, as described in the [project documentation](https://opentelemetry.io/docs/instrumentation/java/automatic/agent-config/) or by passing configuration properties to the container via a binding.

By default, the following configuration is applied to the OpenTelemetry Java Agent at run time.

* `OTEL_JAVAAGENT_ENABLED=false`
* `OTEL_METRICS_EXPORTER=none`

When using a binding to configure the agent, properties are expected to be passed via a config tree format,
where each key is the name of a file and the value its content. Keys can follow the environment variable format or the
system property format, as described in the [project documentation](https://opentelemetry.io/docs/instrumentation/java/automatic/agent-config/).

## Bindings

The buildpack optionally accepts the following bindings:

### Type: `opentelemetry`

| Key                   | Value   | Description                                                                                       |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------- |
| `<otel-property-key>` | `<otel-property-value>` | Properties are expected to be passed via a config tree format, where each key is the name of a file and the value its content. That's the default behavior when passing key/value pairs via a Kubernetes Secret. Keys can follow the environment variable format or the system property format, as described in the [project documentation](https://opentelemetry.io/docs/instrumentation/java/automatic/agent-config/).  |

### Type: `dependency-mapping`

| Key                   | Value   | Description                                                                                       |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------- |
| `<dependency-digest>` | `<uri>` | If needed, the buildpack will fetch the dependency with digest `<dependency-digest>` from `<uri>` |

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
