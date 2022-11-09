# v0.23.0

## Changes

- apigee:
  - Re-enable generation of this client.
- compute:
  - Make Id a on ExternalVpnGateway a pointer type.
- idtoken:
  - Add new package to support making requests with and validating Google ID
    tokens.
- slides:
  - Make int values of Range optional.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.22.0

## Depreciation Notice

- Package `google.golang.org/api/sql/v1beta4` has been deprecated as it was
  generated under the wrong name. This package will be removed in a future
  release. Please migrate to: `google.golang.org/api/sqladmin/v1beta4`.

## Changes

- Apigee client has temporarily been disabled.

- Updated custom search example to be in line with new API.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.21.0

- Disabled automatic switching to *.mtls.googleapis.com endpoints.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.20.0

- WithGRPCConnectionPool is a no-op for some APIs.

- correctly report Go version of runtime.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.19.0

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.18.0

- Add the WithClientCertSource option for mTLS (client TLS certificates), currently only supported for HTTP clients.

- Allow host:port for endpoint overrides, rather than requiring the full base URL (for google.golang.org/api clients).

- Make DialPool work with WithGRPCConn plus non-zero pool size [googleapis/google-cloud-go#1780](https://github.com/googleapis/google-cloud-go/issues/1780)

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.17.0

- Revert sqladmin package name back from sql to sqladmin. (#448)

- Various updates to autogenerated clients.

Internal:

- transport/grpc: add internal WithDialPool option for GAPIC clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.16.0

- Increase the default chunk size for uploads (e.g., for the storage package) to 16 MB.

- transport:
  - Automatically populate QuotaProject from the "quota_project_id" field in the JSON credentials file.
  - Add grpc.DialPool, which opens multiple grpc.ClientConns based on WithGRPCConnectionPool option.

- Added a check to prevent mixed calls to Add and AddWait.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.15.0

- monitoring/v3:
  - Rename Service to MService; revert APIService to Service.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.14.0

- Fix for setting custom HTTP headers in the absence of UserAgent.

- Add a client option for disabling telemetry such as OpenCensus.

- Performance improvements to google.golang.org/api/support/bundler.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.13.0

- Changes to how media path redirection is handled in generated code.

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.12.0

- Various updates to autogenerated clients.

# v0.11.0

- Various updates to autogenerated clients.

- Module information now indicates go 1.11 as oldest supported version.  As of
  October 1, 2019 versions 1.9 and 1.10 are no longer supported.

- Removed the following APIs which are no longer available via the discovery
  service: dfareporting/v2.8, prediction/*.

- The internal gensupport library has been relocated to the more idiomatic
  path internal/gensupport.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.10.0

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.9.0

- Small fix to chunking retry logic such that each chunk has its own retry
  deadline, instead of unbounded retries.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.8.0

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.7.0

- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.6.0

- Add support for GCP DirectPath.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.5.0

- Better support for google.api.HttpBody.
- Support for google.api.HttpBody in the healthcare API.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.4.0

- Includes a re-pin of opencensus, greatly reducing the transitive
dependency list.
- Deletes photoslibrary/v1. The photoslibrary team hopes to fully support Go in
the near future, but this autogenerated library is ready to be sunset. If you
rely on this client, please vendor this library at v0.3.2.
- Various updates to autogenerated clients.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.3.2

This patch releases re-builds the go.sum. This was not possible in the
previous release.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.3.1

This patch release removes github.com/golang/lint from the transitive
dependency list, resolving `go get -u` problems.

_Please note_: this release intentionally has a broken go.sum. Please use v0.3.2.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.3.0

go.mod modifications, including removal of go 1.12 statement and update of
opencensus dependency.

_Please note_: the release version is not indicative of an individual client's
stability or version.

# v0.2.0

General improvements.

_Please note:_ the release version is not indicative of an individual client's
stability or version.

# v0.1.0

Initial release along with Go module support.

_Please note:_ the release version is not indicative of an individual client's
stability or version.