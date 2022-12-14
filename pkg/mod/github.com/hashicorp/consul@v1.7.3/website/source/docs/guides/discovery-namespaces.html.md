---
name: '[Enterprise] Register and Discover Services within Namespaces'
content_length: 8
id: discovery-namespaces
products_used:
  - Consul
description: >-
  In this guide you will register and discover services within a namespace.
level: Implementation
---

!> **Warning:** This guide is a draft and has not been fully tested.

!> **Warning:** Consul 1.7 is currently a beta release.

Namespaces allow multiple teams within the same organization to share the same
Consul datacenter(s) by separating services, key/value pairs, and other Consul
data per team. This provides operators with the ability to more easily run
Consul as a service. Namespaces also enable operators to [delegate ACL
management](/consul/namespaces/secure-namespaces).

Any service that is not registered in a namespace will be added to the `default`
namespace. This means that all services are namespaced in Consul 1.7 and newer,
even if the operator has not created any namespaces.

By the end of this guide, you will register two services in the Consul catalog:
one in the `default` namespace and one in an operator-configured namespace.
After you have registered the services, you will then use the Consul CLI, API
and UI to discover all the services registered in the Consul catalog.

## Perquisites

To complete this guide you will need at least a [local dev
agent](/consul/getting-started/install) running Consul Enterprise 1.7 or newer.
Review the documentation for downloading the [Enterprise
binary](https://www.consul.io/docs/enterprise/index.html#applied-after-bootstrapping).
You can also use an existing Consul datacenter that is running Consul Enterprise
1.7 or newer.

You should have at least one namespace configured. Review the [namespace
management]() documentation or execute the following command to create a
namespace.

```shell
$ consul namespace create app-team
```

## Register services in namespaces

You can register services in a namespace by using your existing workflow and
adding namespace information to the registration. There are two ways to add a
service to a namespace:

- adding the `namespace` option to the service registration file.
- using the `namespace` flag with the API or CLI at registration time.

If you would like to migrate an existing service into a new namespace,
re-register the service with the new namespace information.

### Default namespace

To register a service in the `default` namespace, use your existing registration
workflow; you do not need to add namespace information. In the example below,
you will register the `mysql` service in the default namespace.

First, create a service registration file for the MySQL service and its sidecar
proxy.

```hcl
service {
  name = ???mysql"
  port = 9003
  connect {sidecar_proxy}
}
```
Next, register the service and its sidecar proxy using the Consul CLI by
specifying the registration file.

```shell
$ consul services register mysql.hcl
```

### App-team namespace

To register a service in a user-defined namespace, include the namespace in the
registration file, or pass it with a flag at registration time. In this guide,
we will include the namespace in the file.

First, create the service registration file named `wordpress.hcl`. Paste in the
following registration, which includes the service name and port, and a sidecar
proxy, along with the namespace.

```hcl
service {
  name = ???wordpress"
  port = 9003
  connect {sidecar_proxy}
  namespace = "app-team"
}
```

Next register the service and its sidecar proxy.

```shell
$ consul services register wordpress.hcl -namespace app-team
```

## Discover services

You can discover namespaced services using all the usual methods for service
discovery in Consul: the CLI, web UI, DNS interface, and HTTP API.

### Consul CLI

To get a list of services in the default namespace use the `consul catalog` CLI
command. You do not need to add the flag any discover services in the `default`
namespace.

```shell
$ consul catalog services
consul
mysql
mysql-proxy
```

Notice that you do not see services that are registered in the app-team
namespace.

Add the `-namepsace` flag to discover services within a user-created namespace.
In the example below, you will use the `-namespace` flag with the CLI to
discover all services registered in the app-team namespace.

```shell
$ consul catalog services -namespace app-team
consul
wordpress
wordpress-proxy
```

Notice that you do not see services that are registered in the default
namespace. To discover all services in the catalog, you will need to query all
Consul namespaces.

```shell
$ consul catalog services
consul
mysql
mysql-proxy
$ consul catalog services -namespace app-team
consul
wordpress
wordpress-proxy
```

### Consul UI

You can also view namespaced-services in the Consul UI. Select a namespace using
the drop-down menu at the top of the top navigation. Then go to the ???Services???
tab to see the services within the namespace.

Before you select a namespace the UI will list the services in the `default`
namespace.

![IMAGE FROM RFC! REPLACE ME AT BETA LAUNCH](/static/img/consul/namespaces/consul-namespace-dropdown.png)

### DNS Interface

~> **Note:** To default to the `namespace` parameter in the DNS query, you must
set the `prefer_namespace` option to `true` in the [agent's configuration]().
The new query structure will be, `service.namespace.consul`. This will disable
the ability to query by datacenter only. However, you can add both namespace and
datacenter to the query, `service.namespace.datacenter.consul`.

To discover the location of service instances, you can use the DNS interface.

```shell
$ dig 127.0.0.1 -p 8500 wordpress.service.app-team.consul
<output should show one service>
```

If you don???t specify a namespace in the query, you will get results from the
default namespace.

```shell
$ dig 127.0.0.1 -p 8500 wordpress.service.consul
<output should show no services>
```
### Consul HTTP API

The Consul HTTP API is more verbose than the DNS API; it allows you to discover
the service locations and additional metadata. To discover service information
within a namespace, add the `ns=` query parameter to the call.

```shell
curl http://127.0.0.1:8500/v1/catalog/service/wordpress?ns=app-team
<output shows one service>
```

## Summary

In this guide, you registered two services: the WordPress service in the
app-team namespace and the MySQL service in the `default` namespace. You then
used the Consul CLI to discover services in both namespaces.

You can use ACLs to secure access to data, including services, in namespaces.
After ACLs are enabled, you will be able to restrict access to the namespaces
and all the data registered in that namespace.