variable vpc_cidr_block {
  description = "VPC CIDR block"
  type        = string
  default     = "10.0.0.0/16"
}

variable vpc_public_subnets {
  description = "List of public subnet CIDR blocks"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable vpc_private_subnets {
  description = "List of private subnet CIDR blocks"
  type        = list(string)
  default     = ["10.0.3.0/24", "10.0.4.0/24"]
}

variable vpc_database_subnets {
  description = "List of database subnet CIDR blocks"
  type        = list(string)
  default     = ["10.0.5.0/24", "10.0.6.0/24"]
}

# The Availability Zones data source allows access to the list of AWS Availability Zones 
# which can be accessed by an AWS account within the region configured in the provider.

# This is different from the aws_availability_zone (singular) data source, 
# which provides some details about a specific availability zone.
# Note
# When Local Zones are enabled in a region, by default the API 
# and this data source include both Local Zones and Availability Zones. 
# To return only Availability Zones,
// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/availability_zones
data aws_availability_zones available {
  state = "available"
}

# With Amazon Virtual Private Cloud (Amazon VPC), you can launch AWS resources 
# in a logically isolated virtual network that you've defined. 
# This virtual network closely resembles a traditional network 
# that you'd operate in your own data center, 
# with the benefits of using the scalable infrastructure of AWS.
# Examples
# https://github.com/terraform-aws-modules/terraform-aws-vpc/blob/master/examples/
# https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html
// https://registry.terraform.io/modules/terraform-aws-modules/vpc/aws/3.14.3
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.14.0"

  name = "${var.project}-vpc"

  cidr = var.vpc_cidr_block
  azs  = slice(data.aws_availability_zones.available.names, 0, 2)
  # azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]

  # private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  private_subnets  = var.vpc_private_subnets
  public_subnets   = var.vpc_public_subnets
  database_subnets = var.vpc_database_subnets

  # Should be true if you want to provision NAT Gateways for each of your private networks
  enable_nat_gateway = true
  # Should be true if you want to provision a single shared NAT Gateway across 
  # all of your private networks
  single_nat_gateway = true

  # Should be true to enable DNS support in the VPC
  enable_dns_support   = true
  # Should be true to enable DNS hostnames in the VPC
  enable_dns_hostnames = true

  // Allows public access to the database
  create_database_subnet_group           = true
  create_database_subnet_route_table     = true
  create_database_internet_gateway_route = true

  # Additional tags for the public subnets
  public_subnet_tags = {
    "kubernetes.io/cluster/${var.project}" = "shared"
    "kubernetes.io/role/elb"               = 1
  }

  # Additional tags for the private subnets
  private_subnet_tags = {
    "kubernetes.io/cluster/${var.project}" = "shared"
    "kubernetes.io/role/internal-elb"      = 1
  }
}

output vpc_cidr_block {
  value = var.vpc_cidr_block
}
