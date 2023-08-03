# The ECR Authorization Token data source allows the authorization token, 
# proxy endpoint, token expiration date, user name and password to be retrieved 
# for an ECR repository.
// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/ecr_authorization_token
data aws_ecr_authorization_token token {}

# Provides an Elastic Container Registry Repository.
// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository
resource aws_ecr_repository services {
  for_each     = toset(var.services)
  # name - (Required) Name of the repository.
  name         = each.key
  # force_delete - (Optional) If true, will delete the repository 
  # even if it contains images. Defaults to false.
  force_delete = true
}

# Manages the lifecycle of docker image in a registry. 
# You can upload images to a registry (= docker push) and also delete them again
// https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/registry_image
resource docker_registry_image services {
  for_each = toset(var.services)
  # name (String) The name of the Docker image.
  name     = "${aws_ecr_repository.services[each.key].repository_url}:latest"

  # Build an image with the docker_image resource and then push it to a registry:
  build {
    context    = "../.."
    dockerfile = "docker/Dockerfile.microservices"

    build_args = {
      service = each.key
    }
  }
}

output ecr_url {
  value = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${var.region}.amazonaws.com"
}
