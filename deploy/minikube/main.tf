# locals {
#   services             = var.services
#   vpc_cidr_block       = data.terraform_remote_state.infra.outputs.vpc_cidr_block
#   allowed_cidr_block   = data.terraform_remote_state.infra.outputs.allowed_cidr_block
#   region               = data.terraform_remote_state.infra.outputs.region
#   aws_ecr_url          = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.terraform_remote_state.infra.outputs.region}.amazonaws.com"
#   eks_cluster_id       = data.terraform_remote_state.infra.outputs.eks_cluster_id
#   eks_vpc_cni_role_arn = data.terraform_remote_state.infra.outputs.eks_vpc_cni_role_arn
#   project              = var.project
#   db_conn              = var.db_conn
#   db_host              = var.db_endpoint
#   db_port              = var.db_port
# }
variable "host" {
  type = string
}

variable "client_certificate" {
  type = string
}

variable "client_key" {
  type = string
}

variable "cluster_ca_certificate" {
  type = string
}

variable services {
  description = "List of MallBots microservices"
  type        = list(string)
  default     = ["baskets", "cosec", "customers", "depot", "ordering", "notifications", "payments", "search", "stores"]
}

variable project {
  description = "Project name"
  type        = string
  default     = "mallbots"
}

variable db_endpoint {
  description = "database endpoint"
  type        = string
  # default     = "mallbots"
}

variable db_port {
  description = "database port"
  type        = string
  # default     = "mallbots"
}

variable image_repo {
  description = "database port"
  type        = string
  # default     = "mallbots"
}

// https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/string
resource random_string suffix {
  length  = 8
  special = false
}


output project {
  description = "Project name"
  value       = var.project
}

output services {
  value = var.services
}
