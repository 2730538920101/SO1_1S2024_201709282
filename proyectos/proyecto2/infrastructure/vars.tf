variable "project_id" {
  description = "Google Cloud Project ID"
}

variable "region" {
  description = "Google Cloud region"
  default     = "us-central1"
}

variable "cluster_name" {
  description = "Name of the Kubernetes cluster"
  default     = "my-cluster"
}

variable "node_count" {
  description = "Number of nodes in the Kubernetes cluster"
  default     = 1
}

variable "machine_type" {
  description = "Machine type for Kubernetes nodes"
  default     = "n1-standard-1"
}

variable "private_subnet_cidr" {
  description = "CIDR block for the private subnet"
  default     = "10.0.0.0/24"
}

variable "public_subnet_cidr" {
  description = "CIDR block for the public subnet"
  default     = "10.0.1.0/24"
}

variable "gke_username" {
  default     = ""
  description = "gke username"
}

variable "gke_password" {
  default     = ""
  description = "gke password"
}

variable "gke_num_nodes" {
  default     = 2
  description = "number of gke nodes"
}