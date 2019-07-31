Terraform Provider ScaleFT
==================

Maintainers
-----------

This provider plugin is maintained by the Splunk Cloud Security team.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html)
-	[Go](https://golang.org/doc/install)

Build
---------------------

```
go build -o terraform-provider-scaleft
```

Usage
---------------------

This plugin requires a couple of inputs to run: ScaleFT API key, secret and team name. Set them as environment variables.
```
SCALEFT_KEY_SECRET = < secret here >
SCALEFT_KEY = < key here >
SCALEFT_TEAM = < team name>
```

Using the provider
----------------------
There are tree main resources that you need to provision:
1. Project - project is an authorization scope. It associates a collection of resources with a set of configurations, including RBAC and access policies.
1. Enrollment token - using enrollment token, you can add servers to a project.
1. Groups - in order to give access to a project, you need to assign a group to it.


Sample terraform plan:

```
resource "scaleft_project" "demo-stack" {
  project_name = "tf-test"
}

resource "scaleft_enrollment_token" "test-token" {
  project_name = "${scaleft_project.demo-stack.project_name}"
  description = "Token for X"
}

resource "scaleft_enrollment_token" "test-import-token" {
  project_name = "${scaleft_project.demo-stack.project_name}"
  description = "Token for Y"
}

resource "scaleft_assign_group" "group-assignment" {
  project_name = "${scaleft_project.demo-stack.project_name}"
  group_name = "cloud-sre"
  server_access = true
  server_admin = true
  create_server_group = false
}

resource "scaleft_assign_group" "dev-group-assignment" {
  project_name = "${scaleft_project.demo-stack.project_name}"
  group_name = "cloud-support"
  server_access = true
  server_admin = true
}


```

## Resources
### Project

Example usage:
```
resource "scaleft_project" "demo-stack" {
  project_name = "tf-test"
}
```
Parameters:
* project_name (Required) - name of the project.
* next_unix_uid (Optional - Default: 60101) - ScaleFT will start assigning Unix user IDs from this value
* next_unix_gid (Optional - Default: 63001) - ScaleFT will start assigning Unix group IDs from this value

### Enrollment token
Enrollment is the process where the ScaleFT agent configures a server to be managed by a specific project. An enrollment token is a base64 encoded object with metadata that the ScaleFT Agent can configure itself from.  

Example usage:
```
resource "scaleft_token" "stack-x-token" {
  project_name = "tf-test"
  description = "Token for X stack"
}
```
Parameters:
* project_name (Required) - name of the project.
* Description (Required) - free form text field to provide description. You will NOT be able to change it later without recreating a token.

### Create Group
If groups is not synced from Okta, you may need to create in ScaleFT.
```
resource "scaleft_create_group" "test-tf-group" {
  name = "test-tf-group"
}
```
Parameters:
* name (Required) - name for the ScaleFT group.

NOTE: group is created with basic access_user access. It does not give any privileges in ScaleFT console.
Creation of groups with access_admin and reporting_user is currently not supported in the provider.

### Assign group to project
In order to give access to project, you need to assign Okta group to a project. Use "server_access" and "server_admin" parameters to control access level.

Example usage:
```
resource "scaleft_assign_group" "sg-cloud-group-access" {
  project_name = "tf-test"
  group_name = "cloud-ro"
  server_access = true
  server_admin = false
  create_server_group = true
}
```
Parameters:
* project_name (Required) - name of the project.
* server_access (bool) (Required) - Whether users in this group have access permissions on the servers
in this project.
* server_admin (bool) (Required) - Whether users in this group have sudo permissions on the servers in this project.
* create_server_group (bool) (Optional - Default: true) - will make ScaleFT synchronize group name to linux box. To avoid naming collision, group created by ScaleFT will have prefix of "sft_"
