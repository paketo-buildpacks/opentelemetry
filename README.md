# Cloud Native Buildpacks - OpenTelemetry

The OpenTelemetry Buildpack is a Cloud Native Buildpack that contributes and configures the OpenTelemetry Agent.

## Behavior

This buildpack will participate if all the following conditions are met

* The `$BP_OPENTELEMETRY_ENABLED` is set to a truthy value (i.e. `true`, `t`, `1` ignoring case)

At build time, the buildpack will do the following for Java applications:

* Contributes the OpenTelemetry Java agent to a layer and configures `$JAVA_TOOL_OPTIONS` to use it.

At run time, the buildpack will do the following for Java applications:

* If a binding with `type` of `opentelemetry` exists, it uses the file from the binding to configure the OpenTelemetry Java agent via the `OTEL_JAVAAGENT_CONFIGURATION_FILE` environment variable. This is recommended when the configuration involves secrets of any kind.

## Testing

Given a Java application, you can package it by running the following command from the current folder.

```shell
pack build demo \
  --path test/demo \
  --buildpack paketo-buildpacks/java \
  --buildpack . \
  --builder paketobuildpacks/builder:base \
  --verbose --trust-builder \
  -e BP_JVM_VERSION=17 -e BP_OPENTELEMETRY_ENABLED=true
```

Next, run it as follows.

```shell
docker run --rm -p 8080:8080 demo
```

Check the logs and verify the OpenTelemetry Agent was included.

```log
Picked up JAVA_TOOL_OPTIONS: -Djava.security.properties=/layers/paketo-buildpacks_bellsoft-liberica/java-security-properties/java-security.properties -XX:+ExitOnOutOfMemoryError -javaagent:/layers/paketo-buildpacks_opentelemetry/opentelemetry-java/opentelemetry-javaagent.jar -XX:ActiveProcessorCount=4 -XX:MaxDirectMemorySize=10M -Xmx4352309K -XX:MaxMetaspaceSize=95562K -XX:ReservedCodeCacheSize=240M -Xss1M -XX:+UnlockDiagnosticVMOptions -XX:NativeMemoryTracking=summary -XX:+PrintNMTStatistics

[otel.javaagent 2022-09-18 19:10:16:624 +0000] [main] INFO io.opentelemetry.javaagent.tooling.VersionLogger - opentelemetry-javaagent - version: 1.18.0
```

## Usage

Once you enable the OpenTelemetry buildpack, you can configure it at run-time via environment variables, as described in the [project documentation](https://opentelemetry.io/docs/instrumentation/java/automatic/agent-config/) or via a configuration file provided to the container via a binding.

## Bindings

The buildpack optionally accepts the following bindings:

### Type: `dependency-mapping`

| Key                   | Value   | Description                                                                                       |
| --------------------- | ------- | ------------------------------------------------------------------------------------------------- |
| `<dependency-digest>` | `<uri>` | If needed, the buildpack will fetch the dependency with digest `<dependency-digest>` from `<uri>` |

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0
