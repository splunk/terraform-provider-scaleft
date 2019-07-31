resource "scaleft_project" "demo-project" {
  project_name = "tf-test"
  next_unix_uid = 60120
  next_unix_gid = 63020
}

resource "scaleft_enrollment_token" "test-token" {
  project_name = "${scaleft_project.demo-project.project_name}"
  description  = "Token for X"
}

resource "scaleft_enrollment_token" "test-import-token" {
  project_name = "${scaleft_project.demo-project.project_name}"
  description  = "Token for Y"
}

resource "scaleft_assign_group" "group-assignment" {
  project_name        = "${scaleft_project.demo-project.project_name}"
  group_name          = "cloud-sre"
  server_access       = true
  server_admin        = true
  create_server_group = false
}

resource "scaleft_assign_group" "dev-group-assignment" {
  project_name  = "${scaleft_project.demo-project.project_name}"
  group_name    = "cloud-support"
  server_access = true
  server_admin  = false
}

resource "scaleft_create_group" "test-tf-group" {
  name = "test-tf-group"
}

resource "scaleft_assign_group" "test-sft-group-assignment" {
  project_name = "${scaleft_project.demo-project.project_name}"
  group_name   = "${scaleft_create_group.test-tf-group.name}"
}

output "enrollment_token" {
  value = "${scaleft_enrollment_token.test-token.token_value}"
}
