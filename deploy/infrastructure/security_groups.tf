# Terraform module which creates EC2 security group within VPC on AWS.
# A security group controls the traffic 
# that is allowed to reach and leave the resources that it is associated with. 
# For example, after you associate a security group with an EC2 instance, 
# it controls the inbound and outbound traffic for the instance. 
# You can associate a security group only with resources in the VPC for 
# which it is created.
# Examples
# https://github.com/terraform-aws-modules/terraform-aws-security-group/tree/v4.13.0/examples
# https://docs.aws.amazon.com/vpc/latest/userguide/vpc-security-groups.html
// https://registry.terraform.io/modules/terraform-aws-modules/security-group/aws/4.13.0
module security_group {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.13.0"

  name   = "${var.project}-sg"
  vpc_id = module.vpc.vpc_id

  # List of ingress rules to create where 'cidr_blocks' is used
  ingress_with_cidr_blocks = [
    {
      from_port   = 1
      protocol    = "TCP"
      to_port     = 65365
      cidr_blocks = "${var.allowed_cidr_block},${var.vpc_cidr_block}"
    },
    {
      from_port   = -1
      protocol    = "icmp"
      to_port     = -1
      cidr_blocks = "${var.allowed_cidr_block},${var.vpc_cidr_block}"
    }
  ]

  egress_with_cidr_blocks = [
    {
      from_port   = 0
      protocol    = "-1"
      to_port     = 0
      cidr_blocks = "0.0.0.0/0"
    }
  ]

  # A mapping of tags to assign to security group
  tags = {
    "kubernetes.io/cluster/${var.project}": "shared"
  }
}
