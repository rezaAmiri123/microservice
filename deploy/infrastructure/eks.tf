# https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html

variable eks_node_instance_types {
  description = "EC2 instance types to use for EKS nodes"
  type        = list(string)
  default     = ["t3.small"]
}
# Terraform module to create an Elastic Kubernetes (EKS) cluster and associated resources
// https://registry.terraform.io/modules/terraform-aws-modules/eks/aws/18.29.0
module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "~> 18.29.0"
  # Name of the EKS cluster
  cluster_name    = var.project
  # Kubernetes <major>.<minor> version to use for the EKS cluster (i.e.: 1.22)
  cluster_version = "1.22"
  # Indicates whether or not the Amazon EKS private API server endpoint is enabled
  cluster_endpoint_private_access       = true
  # Indicates whether or not the Amazon EKS public API server endpoint is enabled
  cluster_endpoint_public_access        = true
  # List of additional, externally created security group IDs to attach 
  # to the cluster control plane
  cluster_additional_security_group_ids = [module.security_group.security_group_id]
  # List of CIDR blocks which can access the Amazon EKS public API server endpoint
  cluster_endpoint_public_access_cidrs  = [var.allowed_cidr_block]
  
  # ID of the VPC where the cluster and its nodes will be provisioned
  vpc_id     = module.vpc.vpc_id
  # A list of subnet IDs where the nodes/node groups will be provisioned. 
  # If control_plane_subnet_ids is not provided, 
  # the EKS cluster control plane (ENIs) will be provisioned in these subnets
  subnet_ids = module.vpc.private_subnets

  # Map of attribute maps for all EKS cluster addons enabled
  cluster_addons = {
    # CoreDNS is a flexible, extensible DNS server that can serve 
    # as the Kubernetes cluster DNS. When you launch an Amazon EKS cluster 
    # with at least one node, two replicas of the CoreDNS image are deployed by default, 
    # regardless of the number of nodes deployed in your cluster. 
    # The CoreDNS Pods provide name resolution for all Pods in the cluster. 
    # The CoreDNS Pods can be deployed to Fargate nodes 
    # if your cluster includes an AWS Fargate profile with a namespace 
    # that matches the namespace for the CoreDNS deployment
    coredns = {
      resolve_conflicts = "OVERWRITE"
    }
    # The kube-proxy add-on is deployed on each Amazon EC2 node in your Amazon EKS cluster. 
    # It maintains network rules on your nodes and enables network communication 
    # to your Pods. The add-on isn't deployed to Fargate nodes in your cluster.
    kube-proxy = {}
    # The Amazon VPC CNI plugin for Kubernetes add-on is deployed 
    # on each Amazon EC2 node in your Amazon EKS cluster. 
    # The add-on creates elastic network interfaces and attaches them 
    # to your Amazon EC2 nodes. The add-on also assigns a private IPv4 
    # or IPv6 address from your VPC to each Pod and service.
    vpc-cni    = {
      resolve_conflicts = "OVERWRITE"
    }
  }

  # Determines whether to create an OpenID Connect Provider for EKS to enable IRSA
  enable_irsa = true
  # IAM Roles for Service Accounts (IRSA) is a feature of AWS 
  # which allows you to make use of IAM roles at the pod level 
  # by combining an OpenID Connect (OIDC) identity provider 
  # and Kubernetes service account annotations.

  # Map of EKS managed node group default configurations
  # https://docs.aws.amazon.com/eks/latest/userguide/create-managed-node-group.html
  eks_managed_node_group_defaults = {
    ami_type                              = "AL2_x86_64"
    disk_size                             = 10
    instance_types                        = var.eks_node_instance_types
    create_launch_template                = false
    launch_template_name                  = ""
    attach_cluster_primary_security_group = true
    iam_role_attach_cni_policy            = true
    vpc_security_group_ids                = [module.security_group.security_group_id]
  }

  # Map of EKS managed node group definitions to create
  eks_managed_node_groups = {
    primary = {
      name = "${var.project}-nodes"

      min_size     = 2
      max_size     = 5
      desired_size = 2
    }
  }
}

# Amazon VPC Container Network Interface
# Amazon EKS implements cluster networking through 
# the Amazon VPC Container Network Interface(VPC CNI) plugin. 
# The CNI plugin allows Kubernetes Pods to have the same IP address 
# as they do on the VPC network.

# IAM Roles for Service Accounts (IRSA) is a feature of AWS 
# which allows you to make use of IAM roles at the pod level 
# by combining an OpenID Connect (OIDC) identity provider 
# and Kubernetes service account annotations.

# Creates an IAM role which can be assumed by AWS EKS ServiceAccounts 
# with optional policies for commonly used controllers/custom resources within EKS

# This module supports multiple ServiceAccounts across multiple clusters 
# and/or namespaces. This allows for a single IAM role to be used when an application 
# may span multiple clusters (e.g. for DR) or multiple namespaces 
# (e.g. for canary deployments). For example, to create an IAM role named my-app 
# that can be assumed from the ServiceAccount named my-app-staging 
# in the namespace default and canary in a cluster in us-east-1; 
# and also the ServiceAccount name my-app-staging in the namespace default 
# in a cluster in ap-southeast-1
// https://registry.terraform.io/modules/terraform-aws-modules/iam/aws/5.3.1/submodules/iam-role-for-service-accounts-eks
module "vpc_cni_irsa" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-role-for-service-accounts-eks"
  version = "~> 5.4.0"

  # IAM role name prefix
  role_name_prefix                       = "vpc-cni-irsa-"
  # Determines whether to attach the Load Balancer Controller policy to the role
  attach_load_balancer_controller_policy = true

  # Your cluster has an OpenID Connect (OIDC) issuer URL associated with it. 
  # To use AWS Identity and Access Management (IAM) roles for service accounts, 
  # an IAM OIDC provider must exist for your cluster's OIDC issuer URL.
  # Map of OIDC providers where each provider map should contain the provider, 
  # provider_arn, and namespace_service_accounts
  oidc_providers = {
    main = {
      provider_arn               = module.eks.oidc_provider_arn
      namespace_service_accounts = ["kube-system:${var.lb_service_account_name}"]
    }
  }
}

output eks_cluster_id {
  description = "EKS cluster ID"
  value       = module.eks.cluster_id
}

output eks_endpoint {
  description = "Endpoint for EKS control plane"
  value       = module.eks.cluster_endpoint
}

output eks_certificate_authority_data {
  value = module.eks.cluster_certificate_authority_data
}

output eks_vpc_cni_role_arn {
  value = module.vpc_cni_irsa.iam_role_arn
}
