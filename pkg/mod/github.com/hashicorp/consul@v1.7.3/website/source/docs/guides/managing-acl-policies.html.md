---
name: Managing ACL Policies
content_length: 15
id: managing-acl-policies
products_used:
  - Consul
description: >-
  In this guide, you'll learn how to discover the minimum privileges required to complete operations within your Consul datacenter and how to manage access using the operator-only implementation method.
level: Implementation
---

This guide is for Operators with the responsibility of creating and managing ACL tokens for a Consul datacenter. It includes several recommendations on how to discover the minimum privileges required to complete operations. Throughout the guide we'll provide examples and use cases that you can adapt to your environment, however, it does not include environment specific recommendations. After completing this guide, you will have a better understanding of how to effectively manage ACL policies and tokens. 

We expect operators to automate the policy and token generation process in production environments. Additionally, if you are using a container orchestrator, the process will vary even though the concepts in this guide will still be applicable. If you are using the official Consul-Kubernetes Helm chart to deploy Consul, use the [authentication method documentation](https://www.consul.io/docs/acl/acl-auth-methods.html) instead of generating policies manually or automating the methods here.

## Prerequisites

We provide high-level recommendations in this guide, however, we will not describe the command by command token generation process. To learn how to create tokens, read the [ACL bootstrapping guide](https://learn.hashicorp.com/consul/security-networking/production-acls).

This guide assumes the `default_policy` of `deny` is set on all agents, in accordance to the [security model documentation](https://www.consul.io/docs/internals/security.html#secure-configuration).

## Security and Usability 

The examples in this guide illustrate how to create multiple policies that can be used to accomplish the same task. For example, using an exact match resource rule, is the most secure. It grants the least privileges necessary to accomplish the task. Generally, creating policies and tokens with the least privileges will result in more policy definitions. Alternatively, for a simplified process, the prefix resources rules can apply to zero-to-many objects. The trade-off of a less complicated token creation process is wider potential blast radius on token or workload compromise. 

## Discover Required Privileges 

After bootstrapping the ACL system and configuring Consul agents with tokens, you will need to create tokens to complete any additional task within the datacenter including registering services. 

Before discovering the minimum privileges, it's important to understand the basic components of a token. A rule is a specific privilege and the basic unit of a token. Rules are combined to create policies. There are two main parts of a rule, the resource and policy disposition. The resource is the object that the rule applies to and the policy dispositions dictates privileges. The example below applies to any service object named "web". The policy disposition grants read privileges. 

![ACL Rule Diagram](/static/img/consul/ACL-rule.png 'ACL Diagram with rules')

To discover the minimum privileges required for a specific operation, we have three recommendations. 

First, focus on the data in your environment that needs to be secured. Ensure your sensitive data has policies that are specific and limited. Since policies can be combined to create tokens, you will usually write more policies for sensitive data. Sensitive data could be a specific application or a set of values in the key-value store. 

Second, reference the Consul docs, both the [rules page](https://www.consul.io/docs/acl/acl-rules.html) and [API pages](https://www.consul.io/api/index.html), often to understand the required privileges for any given operation. 

The rules documentation explains the 11 rule resources. The following four resource types are critical for any operating datacenter with ACLs enabled. 


| Rule              | Summary
--------------------|--------------------------------------------------------
|`acl`              |ACL rules grant privileges for ACL operations including to create,update, or view tokens and policies.
|`node` and  `node_prefix` |Node rules grant privileges for node-level registration, including adding agents to the datacenter and catalog.
| `service` and `service_prefix`   | Service rules grant privileges for service-level registration, including adding services to the catalog. 
| `operator`        | The operator grants privileges for datacenter operations, including interacting with Raft. 


On the API pages, each endpoint will have a table that includes required ACLs. The node health endpoint, shown below, requires node and service read to view all checks for the specified node. 

![API Health Endpoint](/static/img/consul/api-endpoint.png 'Screenshot of Health endpoint page')

Finally, before using a token in production, you should test that it has the correct privileges. If the token being used is missing a privilege, then Consul will issue a 403 permission denied error. 

```sh
$ consul operator raft list-peers
Error getting peers: Failed to retrieve raft configuration: Unexpected response code: 403 (rpc error making call: Permission denied)
``` 

### Operations Example: Listing Consul Agents

To view all Consul agents in the datacenter, you can use the `consul members` CLI command. You will need at least read privileges for every agent in the datacenter. Depending on your threat model you could have either of the following policies. 

An individual rule and policy for each agent: this requires you to add a new policy to the token every time a new agent is added to the datacenter. Your token would have as many policies as agents in the datacenter. This method is ideal for a static environment, and for maximum security.

```hcl
agent "server-one" 
  { policy = "read"
  }
```

If you have a dynamic or very large environment, you may want to consider creating one policy that can apply to all agents, to reduce the Operation Team???s workload. 

```hcl
agent_prefix "" 
  { policy = "read"
  }
```

## Secure Access Control: Operator-Only Access

The most secure access control implementation restricts tokens with `acl="write"` policies to only one or a few trusted operators. Tokens with the policy `acl = "write"` grant the holder unlimited privileges, because they can generate tokens with any other resource and policy. The operators are responsible for creating all other policies and tokens that grant limited access to the datacenter. We refer to this implementation as the operator-only implementation. This implementation type is the most secure, and most complex to manage.  

For this implementation type operators are responsible for managing policies and tokens for: 
- service registration
- Connect proxy registration
- intention management
- agent management
- API, CLI, and UI access

For static or non-containerized workflows, this implementation type is straightforward and the operator's workload scales linearly. If the workflow is dynamic, implementing an automation tool will ensure the operator's workload does not scale exponentially. 

!> If you need to share token generation responsibilities with other teams, anyone with the ability to create tokens (`acl = "write"`) effectively has unrestricted access to the datacenter. That individual will have the ability to access or delete any data in Consul including ACL tokens, Connect Intentions and all Catalog and KV data. 

### Operator-Only Implementation Example 

In this following example Operators retain responsibility for service token management, but delegate access control between Connect-enabled services to the security team. 

-> Note: The service registration examples describe the token generation process for per-service tokens, therefore, they are only applicable if you are able to create per-service tokens. 

Initially, the Operator creates a single token with intention management privileges for the Security Team, and another service token for the Developer. For intention management the `intention` policy disposition should be included in the service rule.

```hcl
service "wordpress" 
  { policy = "read" 
    intentions = "write"
  }
```

This enables the security team to create an intention that allows the `wordpress` service to open new connections to the upstream `mysql` service.

```sh
consul intention create wordpress mysql
``` 

The Security Team responsible for managing intentions would need to create allow intentions when the `default_policy` is set to `deny`, since deny all will be inherited from the ACL configuration.

In this implementation type, Developers are responsible for requesting a token for their service. For a Connect enabled service, the Operator would need to create a policy that provides write privileges for the service and proxy, and read privileges for all services and nodes to enable discovery of other upstream dependencies.

```hcl
service "mysql" {
  policy = "write"
}

service "mysql-sidecar-proxy" {
  policy = "write"
}

service_prefix "" {
    policy = "read"
}

node_prefix "" {
    policy = "read"
}
```

With the token the Developer can then register the service with Consul using the `register` command.

```sh
$ consul services register mysql
``` 

The Operator example above illustrates creating policies on the security spectrum. The first example, using an exact match resource rule, is the most secure. It grants the least privileges necessary to accomplish the task. Generally, creating policies and tokens with the least privileges will result in more policy definitions. However, this will help you create the most secure environment. Alternatively, the prefix rule can apply to zero-to-many objects. The trade-off of a less complicated token creation process is security. Note, this applies to all rules not just for agents.  

## Next Steps

After setting up access control processes, you will need to implement a token rotation policy. If you are using third-party tool to generate tokens, such as Vault, Consul ACL tokens will adhere to the TTLs set in that third party tool. If you are manually rotating tokens or need to revoke access, you can delete a token at any time with the [API](https://www.consul.io/api/acl/tokens.html#delete-a-token). 
