# Amazon Relational Database Service (Amazon RDS) is a web service 
# that makes it easier to set up, operate, and scale a relational database in the cloud. 
# It provides cost-efficient, resizeable capacity 
# for an industry-standard relational database 
# and manages common database administration tasks. 
# Amazon Aurora is a fully managed relational database engine 
# that's built for the cloud and compatible with MySQL and PostgreSQL. 
# Amazon Aurora is part of Amazon RDS.
variable db_instance_type {
  description = "RDS serverless instance type"
  type        = string
  default     = "db.serverless"
}

variable db_family {
  description = "RDS serverless family"
  type        = string
  default     = "aurora-postgresql13"
}

variable db_username {
  description = "User name for the RDS PostgreSQL database"
  type        = string
}

# Information about an RDS engine version.
// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/rds_engine_version
data aws_rds_engine_version postgres {
  # engine - (Required) DB engine. Engine values include aurora, aurora-mysql, 
  # aurora-postgresql, docdb, mariadb, mysql, neptune, oracle-ee, oracle-se, 
  # oracle-se1, oracle-se2, postgres, sqlserver-ee, sqlserver-ex, sqlserver-se, 
  # and sqlserver-web.
  engine = "aurora-postgresql"
}

// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_parameter_group
resource aws_db_parameter_group postgres {
  # name - (Optional, Forces new resource) The name of the DB parameter group. 
  # If omitted, Terraform will assign a random, unique name.
  name   = "${var.project}-parameter-group"
  # family - (Required, Forces new resource) The family of the DB parameter group.
  family = var.db_family
}

// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/rds_cluster_parameter_group
resource aws_rds_cluster_parameter_group postgres {
  # name - (Optional, Forces new resource) The name of the DB cluster parameter group. 
  # If omitted, Terraform will assign a random, unique name.
  name   = "${var.project}-cluster-parameter-group"
  # family - (Required) The family of the DB cluster parameter group.
  family = var.db_family
}

// https://registry.terraform.io/modules/terraform-aws-modules/rds-aurora/aws/7.5.1
module "db" {
  source  = "terraform-aws-modules/rds-aurora/aws"
  version = "~> 7.5.0"

  name                   = "${var.project}-db-cluster"
  # nstance type to use at master instance. Note: if autoscaling_enabled is true, 
  # this will be the same instance class used on instances created by autoscaling
  instance_class         = var.db_instance_type
  # The name of the database engine to be used for this DB cluster. Defaults to aurora. 
  # Valid Values: aurora, aurora-mysql, aurora-postgresql
  engine                 = data.aws_rds_engine_version.postgres.engine
  # The database engine mode. Valid values: global, multimaster, parallelquery, 
  # provisioned, serverless. Defaults to: provisioned
  engine_mode            = "provisioned"
  # The database engine version. Updating this argument results in an outage
  engine_version         = data.aws_rds_engine_version.postgres.version
  # Username for the master DB user
  master_username        = var.db_username
  # Determines whether to create random password for RDS primary cluster
  create_random_password = true
  port                   = 5432

  db_parameter_group_name         = aws_db_parameter_group.postgres.id
  db_cluster_parameter_group_name = aws_rds_cluster_parameter_group.postgres.id

  # Map of cluster instances and any specific/overriding attributes to be created
  instances = {
    primary = {}
  }

  # Map of nested attributes with serverless v2 scaling properties. 
  # Only valid when engine_mode is set to provisioned
  serverlessv2_scaling_configuration = {
    min_capacity = 1
    max_capacity = 5
  }

  vpc_id  = module.vpc.vpc_id
  subnets = module.vpc.database_subnets

  # A list of CIDR blocks which are allowed to access the database
  allowed_cidr_blocks    = module.vpc.private_subnets_cidr_blocks
  db_subnet_group_name   = module.vpc.database_subnet_group_name
  # Determines whether to create the database subnet group or use existing
  create_db_subnet_group = false
  create_security_group  = false
  vpc_security_group_ids = [module.security_group.security_group_id]

  # Specifies whether any cluster modifications are applied immediately, 
  # or during the next maintenance window. Default is false
  apply_immediately   = true
  # Determines whether a final snapshot is created before the cluster is deleted. 
  # If true is specified, no snapshot is created
  skip_final_snapshot = true

  # Determines whether instances are publicly accessible. Default false
  # This should never be set for a real production database!
  publicly_accessible = true
}

output db_endpoint {
  value = module.db.cluster_endpoint
}

output db_port {
  value = module.db.cluster_port
}

output db_username {
  value     = module.db.cluster_master_username
  sensitive = true
}

output db_password {
  value     = module.db.cluster_master_password
  sensitive = true
}

output db_conn {
  value     = "postgres://${module.db.cluster_master_username}:${module.db.cluster_master_password}@${module.db.cluster_endpoint}:${module.db.cluster_port}"
  sensitive = true
}
