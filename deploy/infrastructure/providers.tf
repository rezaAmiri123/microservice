terraform {
  required_version = "~> 1.2.0"

  backend local {
    path = "./infrastructure.tfstate"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.31.0"
    }

    docker = {
      source  = "kreuzwerker/docker"
      version = "2.21.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }

    helm = {
      source = "hashicorp/helm"
      version = "~> 2.6.0"
    }
  }
}

// https://registry.terraform.io/providers/hashicorp/aws/latest/docs
provider aws {
  region = var.region

  default_tags {
    tags = {
      Application = "MallBots"
    }
  }
}

# The Docker provider is used to interact with Docker containers and images. 
# It uses the Docker API to manage the lifecycle of Docker containers. 
# Because the Docker provider uses the Docker API, 
# it is immediately compatible not only with single server Docker but Swarm 
# and any additional Docker-compatible API hosts.
// https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs
provider "docker" {
  registry_auth {
    # address (String) Address of the registry
    address  = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${var.region}.amazonaws.com"
    # username (String) Username for the registry. 
    # Defaults to DOCKER_REGISTRY_USER env variable if set.
    username = data.aws_ecr_authorization_token.token.user_name
    # password (String, Sensitive) Password for the registry. 
    # Defaults to DOCKER_REGISTRY_PASS env variable if set.
    password = data.aws_ecr_authorization_token.token.password
  }
}
# The Kubernetes (K8S) provider is used to interact 
# with the resources supported by Kubernetes. 
# The provider needs to be configured with the proper credentials 
# before it can be used.
// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs
provider kubernetes {
  # host - (Optional) The hostname (in form of URI) of the Kubernetes API. 
  # Can be sourced from KUBE_HOST.
  host                   = module.eks.cluster_endpoint
  # cluster_ca_certificate - (Optional) PEM-encoded root certificates bundle 
  # for TLS authentication. Can be sourced from KUBE_CLUSTER_CA_CERT_DATA
  cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
  # Some cloud providers have short-lived authentication tokens 
  # that can expire relatively quickly. To ensure the Kubernetes provider 
  # is receiving valid credentials, an exec-based plugin can be used to fetch 
  # a new token before initializing the provider. For example, on EKS, 
  # the command eks get-token can be used
  exec {
    api_version = "client.authentication.k8s.io/v1beta1"
    command     = "aws"
    # This requires the awscli to be installed locally where Terraform is executed
    args = ["eks", "get-token", "--region", var.region, "--cluster-name", module.eks.cluster_id]
  }
}
# The Helm provider is used to deploy software packages in Kubernetes. 
# The provider needs to be configured with the proper credentials 
# before it can be used.
# https://registry.terraform.io/providers/hashicorp/helm/latest/docs
provider helm {
  # You must have a Kubernetes cluster available. We support version 1.14.0 or higher.
  # NOTE: The provider does not use the KUBECONFIG environment variable by default. 
  # See the attribute reference below for the environment variables 
  # that map to provider block attributes.
  kubernetes {
    host                   = module.eks.cluster_endpoint
    cluster_ca_certificate = base64decode(module.eks.cluster_certificate_authority_data)
    exec {
      api_version = "client.authentication.k8s.io/v1beta1"
      command     = "aws"
      # This requires the awscli to be installed locally where Terraform is executed
      args = ["eks", "get-token", "--region", var.region, "--cluster-name", module.eks.cluster_id]
    }
  }
}
