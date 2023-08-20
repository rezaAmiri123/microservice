# Amazon container image registries
# When you deploy AWS Amazon EKS add-ons to your cluster, 
# your nodes pull the required container images from the registry specified 
# in the installation mechanism for the add-on, such as an installation manifest 
# or a Helm values.yaml file. The images are pulled 
# from an Amazon EKS Amazon ECR private repository. 
# Amazon EKS replicates the images to a repository in each Amazon EKS supported AWS Region. 
# Your nodes can pull the container image 
# over the internet from any of the following registries. 
# Alternatively, your nodes can pull the image over Amazon's network 
# if you created an interface VPC endpoint for Amazon ECR (AWS PrivateLink) in your VPC. 
# The registries require authentication with an AWS IAM account. 
# Your nodes authenticate using the Amazon EKS node IAM role, 
# which has the permissions in the AmazonEC2ContainerRegistryReadOnly managed IAM policy 
# associated to it.
# https://docs.aws.amazon.com/eks/latest/userguide/add-ons-images.html
variable lb_image_repository {
  description = "AWS Load Balancer image host See: https://docs.aws.amazon.com/eks/latest/userguide/add-ons-images.html"
  type        = string
  default     = "602401143452.dkr.ecr.us-east-1.amazonaws.com"
}

variable lb_service_account_name {
  type    = string
  default = "aws-load-balancer-controller"
}

# A service account provides an identity for processes that run in a Pod.
# https://kubernetes.io/docs/admin/service-accounts-admin/
// https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/service_account_v1
resource kubernetes_service_account_v1 lb_service_account {
  # metadata - (Required) Standard service account's metadata.
  metadata {
    # name - (Optional) Name of the service account, must be unique. Cannot be updated.
    name      = var.lb_service_account_name
    # namespace - (Optional) Namespace defines the space within 
    # which name of the service account must be unique.
    namespace = "kube-system"
    # labels - (Optional) Map of string keys and values that can be used to organize 
    # and categorize (scope and select) the service account. 
    # May match selectors of replication controllers and services. 
    labels    = {
      "app.kubernetes.io/name"      = "aws-load-balancer-controller"
      "app.kubernetes.io/component" = "controller"
    }
    # annotations - (Optional) An unstructured key value map stored 
    # with the service account that may be used to store arbitrary metadata.
    annotations = {
      "eks.amazonaws.com/role-arn" = module.vpc_cni_irsa.iam_role_arn
    }
    # Note
    # By default, the provider ignores any annotations whose key names end 
    # with kubernetes.io. This is necessary because such annotations 
    # can be mutated by server-side components and consequently cause a perpetual diff 
    # in the Terraform plan output. If you explicitly specify any such annotations 
    # in the configuration template then Terraform will consider 
    # these as normal resource attributes and manage them as expected 
    # (while still avoiding the perpetual diff problem)
  }
}

# A Release is an instance of a chart running in a Kubernetes cluster.

# A Chart is a Helm package. It contains all of the resource definitions necessary 
# to run an application, tool, or service inside of a Kubernetes cluster.

# helm_release describes the desired status of a chart in a kubernetes cluster.
// https://registry.terraform.io/providers/hashicorp/helm/latest/docs/resources/release
resource helm_release lb {
  # name - (Required) Release name.
  name = "load-balancer"
  # repository - (Optional) Repository URL where to locate the requested chart.
  repository = "https://aws.github.io/eks-charts"
  # chart - (Required) Chart name to be installed. The chart name can be local path, 
  # a URL to a chart, or the name of the chart if repository is specified. 
  # It is also possible to use the <repository>/<chart> format here 
  # if you are running Terraform on a system that the repository has been added to 
  # with helm repo add but this is not recommended.
  chart      = "aws-load-balancer-controller"
  # namespace - (Optional) The namespace to install the release into. Defaults to default.
  namespace  = "kube-system"

  # set - (Optional) Value block with custom values to be merged with the values yaml.
  # https://github.com/aws/eks-charts/tree/master/stable/aws-load-balancer-controller
  # Kubernetes cluster name
  set {
    name  = "clusterName"
    value = var.project
  }

  # If true, create a new service account
  set {
    name  = "serviceAccount.create"
    value = "false"
  }
  # Service account to be used
  set {
    name  = "serviceAccount.name"
    value = var.lb_service_account_name
  }
  # The AWS region for the kubernetes cluster
  set {
    name  = "region"
    value = var.region
  }
  # default: public.ecr.aws/eks/aws-load-balancer-controller
  // See: https://docs.aws.amazon.com/eks/latest/userguide/add-ons-images.html
  set {
    name  = "image.repository"
    value = "${var.lb_image_repository}/amazon/aws-load-balancer-controller"
  }
  # The VPC ID for the Kubernetes cluster
  set {
    name  = "vpcId"
    value = module.vpc.vpc_id
  }

  depends_on = [
    module.eks,
    kubernetes_service_account_v1.lb_service_account
  ]
}
