[![Build Status](https://ci-vic.vmware.com/api/badges/vmware/vic/status.svg)](https://ci-vic.vmware.com/vmware/vic) [![codecov](https://codecov.io/gh/vmware/vic/branch/master/graph/badge.svg)](https://codecov.io/gh/vmware/vic) [![Download](https://img.shields.io/badge/download-latest-blue.svg)](https://github.com/vmware/vic/releases/latest) [![Go Report Card](https://goreportcard.com/badge/github.com/vmware/vic)](https://goreportcard.com/report/github.com/vmware/vic)

---

# HashiCorp Fork Changes

This fork exists only to solve hashicorp/go-discover#94. `vmware/vic` imports `logrus` with upper case `Sirupsen` which conflicts with almost all other libraries at this point and goes against the [general recommendation for interop](https://github.com/sirupsen/logrus#case-sensitivity). The issues is tracked in vmware/vic#8263 once the upstream has found a solution to that, this fork will be removed. Given the ugliness of upstrea having to update (maybe even raise issues and PR against) lots of deps or modify their imports in the vendor dir we are forking for now to solve our immediate problem.

---

# vSphere Integrated Containers Engine

vSphere Integrated Containers Engine (VIC Engine) is a container runtime for vSphere, allowing developers familiar with Docker to develop in containers and deploy them alongside traditional VM-based workloads on vSphere clusters, and allowing for these workloads to be managed through the vSphere UI in a way familiar to existing vSphere admins.

See [VIC Engine Architecture](doc/design/arch/arch.md) for a high level overview.

## Support and Troubleshooting

For general questions, visit the [vSphere Integrated Containers Engine channel](https://vmwarecode.slack.com/messages/vic-engine). If you do not have an @vmware.com or @emc.com email address, sign up at https://code.vmware.com/join to get an invitation.

## Project Status

[commands]:https://vmware.github.io/vic-product/assets/files/html/1.3/vic_app_dev/container_operations.html
[ctrust]:https://docs.docker.com/engine/security/trust/content_trust
[wizard]:https://vmware.github.io/vic-product/assets/files/html/1.3/vic_vsphere_admin/deploy_vch_client.html
[dch]:https://vmware.github.io/vic-product/assets/files/html/1.3/vic_vsphere_admin/deploy_vch_dchphoton.html

VIC Engine now provides:
* support for most of the Docker commands for core container, image, volume and network lifecycle operations. Several `docker compose` commands are also supported. See the complete list of supported commands [here][commands].
* vCenter support, leveraging DRS for initial placement. vMotion is also supported.
* volume support for standard datastores such as vSAN and iSCSI datastores. NFS shares are also supported. See [--volume-store](doc/user/usage.md#configuring-volumes-in-a-virtual-container-host) - SIOC is not integrated but can be set as normal.
* direct mapping of vSphere networks [--container-network](doc/user/usage.md#exposing-vsphere-networks-within-a-virtual-container-host) - NIOC is not integrated but can be set as normal.
* dual-mode management - IP addresses are reported as normal via vSphere UI, guest shutdown via the UI will trigger delivery of container STOPSIGNAL, restart will relaunch container process.
* client authentication - basic authentication via client certificates known as _tlsverify_.
* integration with the VIC Management Portal (Admiral) for Docker image [content trust][ctrust].
* integration with the vSphere Platform Services Controller (PSC) for Single Sign-on (SSO) for docker commands such as `docker login`.
* an install wizard in the vSphere HTML5 client, as a more interactive alternative to installing via the command line. See details [here][wizard].
* support for a standard Docker Container Host (DCH) deployed and managed as a container on VIC Engine. This can be used to run docker commands that are not currently supported by VIC Engine (`docker build, docker push`). See details [here][dch].

We are working hard to add functionality while building out our [foundation](doc/design/arch/arch.md#port-layer-abstractions) so continue to watch the repo for new features. Initial focus is on the production end of the CI pipeline, building backwards towards developer laptop scenarios.

## Installing

After building the binaries (see the [Building](#building) section), pick up the correct binary based on your OS, and install the Virtual Container Host (VCH) with the following command. For Linux:

```shell
bin/vic-machine-linux create --target <target-host>[/datacenter] --image-store <datastore name> --name <vch-name> --user <username> --password <password> --thumbprint <certificate thumbprint> --compute-resource <cluster or resource pool name> --tls-cname <FQDN, *.wildcard.domain, or static IP>
```

See `vic-machine-$OS create --help` for usage information. A more in-depth example can be found [here](doc/user/usage.md#deploying-a-virtual-container-host).

## Deleting

The installed VCH can be deleted using `vic-machine-$OS delete`.

See `vic-machine-$OS delete --help` for usage information. A more in-depth example can be found [here](doc/user/usage.md#deleting-a-virtual-container-host).

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting changes and the contribution workflow.

## Building

Building the project is done with a combination of make and containers, with gcr.io/eminent-nation-87317/vic-build-image:tdnf being the common container base. This requires Docker installed. If gcr.io is not accessible you can follow the steps in [Dockerfile](infra/build-image/Dockerfile.tdnf) to build this image.

To build as closely as possible to the formal build, you need the Go 1.8 toolchain and Drone.io installed:
```shell
drone exec
```

To build inside a Docker container:
```shell
docker run -v $GOPATH/bin:/go/bin -v $(pwd):/go/src/github.com/vmware/vic gcr.io/eminent-nation-87317/vic-build-image:tdnf make all
```

To build directly, you also need the Go 1.8 toolchain installed:
```shell
make all
```

There are three primary components generated by a full build, found in `$BIN` (the ./bin directory by default). The make targets used are the following:
 1. vic-machine - `make vic-machine`
 2. appliance.iso - `make appliance`
 3. bootstrap.iso - `make bootstrap`

## Building binaries for development

Some of the project binaries can only be built on Linux.  If you are developing on a Mac or Windows OS, then the easiest way to facilitate a build is by utilizing the project's Vagrantfile.  The Vagrantfile will share the directory where the file is executed and set the GOPATH based on that share.

To build the component binaries, ensure `GOPATH` is set, then issue the following command in the root directory:
```shell
make components
```
This will install required tools and build the component binaries `tether-linux`, `rpctool` and server binaries `docker-engine-server`, `port-layer-server`.  The binaries will be created in the `$BIN` directory, ./bin by default.

To run unit tests after a successful build, issue the following:
```shell
make test
```

Running "make" every time causes Go dependency regeneration for each component, so that "make" can rebuild only those components that are changed.  However, such regeneration may take significant amount of time when it is not really needed. To fight that developers can use cached dependencies that can be enabled by defining the environment variable VIC_CACHE_DEPS. As soon as it is set, infra/scripts/go-deps.sh will read cached version of dependencies if those exist.

```shell
export VIC_CACHE_DEPS=1
```

This is important to note that as soon as you add a new package or an internal project dependency that didn't exist before, those dependencies
should be regenerated to reflect latest changes. It can be done just by running:

```shell
make cleandeps
```

After that next "make" run will regenerate dependencies from scratch.

To enable generation of non-stripped binaries, the following environment variable can be set:

```shell
export VIC_DEBUG_BUILD=true
```

## Updating the appliance with newly built binaries
After building any of the binaries for the appliance VM (vicadmin, vic-init, port-layer-server, or the docker personality), run `make push` to replace the binaries on your VCH with the newly built ones.

`make push` will prompt you for information that it needs, or you can set your `GOVC` environment variables, as well as `VIC_NAME` (name of your VCH) and `VIC_KEY` with a path to your SSH key in order to run `make push` noninteractively.

Replace individual components with one of: `make push-portlayer`, `make push-vicadmin`, `make push-docker`, or `make push-vic-init`.


## Managing vendor/ directory

To build the VIC Engine dependencies, ensure `GOPATH` is set, then issue the following.
```shell
make gvt vendor
```

This will install the [gvt](https://github.com/FiloSottile/gvt) utility and retrieve the build dependencies via `gvt restore`.

## Building the ISOs

The component binaries above are packaged into ISO files, appliance.iso and bootstrap.iso, that are used by the installer. The generation of the ISOs is split into the following targets:

`iso-base, appliance-staging, bootstrap-staging, appliance, and bootstrap`

Generation of the ISOs involves authoring a new root filesystem, meaning running a package manager (currently tdnf) and packing/unpacking archives. This is done with gcr.io/eminent-nation-87317/vic-build-image:tdnf being the build container. This requires Docker installed. If gcr.io is not accessible you can follow the steps in [Dockerfile](infra/build-image/Dockerfile.tdnf) to build this image. To generate the ISOs:

```shell
make isos
```

The appliance and bootstrap ISOs are bootable CD images used to start the VMs that make up VIC Engine. To build the image, issue the following.

```shell
docker run -v $(pwd):/go/src/github.com/vmware/vic gcr.io/eminent-nation-87317/vic-build-image:tdnf make isos
```

The iso image will be created in `$BIN`

## Building Custom ISOs
Please reference [this document](isos/base/repos/README.md) to build your custom isos.

## Building with CI

[dronevic]:https://ci-vic.vmware.com/vmware/vic
[dronesrc]:https://github.com/drone/drone
[dronecli]:http://docs.drone.io/cli-installation/

PRs to this repository will trigger builds on our [Drone CI][dronevic].

To build locally with Drone:

Ensure that you have Docker 1.6 or higher installed.
Install the [Drone command line tools][dronecli].
From the root directory of the `vic` repository run `drone exec`

## Common Build Problems
1. Builds may fail when building either the appliance.iso or bootstrap.iso with the error: `cap_set_file failed - Operation not supported`

   *Cause:* Some Ubuntu and Debian based systems ship with a defective `aufs` driver, which Docker uses as its default backing store.  This driver does not support extended file capabilities such as `cap_set_file`

   *Solution:* Edit the `/etc/default/docker` file, add the option `--storage-driver=overlay` to the `DOCKER_OPTS` settings, and restart Docker.

2. `go vet` fails when doing a `make all`

    *Cause:* Apparently some caching takes place in `$GOPATH/pkg/linux_amd64/github.com/vmware/vic` and can cause `go vet` to fail when evaluating outdated files in this cache.

    *Solution:* Delete everything under `$GOPATH/pkg/linux_amd64/github.com/vmware/vic` and re-run `make all`.

3.  `vic-machine upgrade` integration tests fail due to `BUILD_NUMBER` being set incorrectly when building locally

    *Cause:* `vic-machine` checks the build number of its binary to determine upgrade status and a locally-built `vic-machine` binary may not have the `BUILD_NUMBER` set correctly. Upon running `vic-machine upgrade`, it may fail with the message `foo-VCH has same or newer version x than installer version y. No upgrade is available.`

    *Solution:* Set `BUILD_NUMBER` to a high number at the top of the `Makefile` - `BUILD_NUMBER ?= 9999999999`. Then, re-build binaries - `sudo make distclean && sudo make clean && sudo make all` and run `vic-machine upgrade` with the new binary.

## Integration Tests

[VIC Engine Integration Test Suite](tests/README.md) includes instructions to run locally.

## Debugging with DLV

[VIC Engine DLV Debugging with DLV](infra/dlv/README.md) includes instruction on how to use dlv.

## License

VIC Engine is available under the [Apache 2 license](LICENSE).
