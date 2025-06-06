variable "aws_region" {
  default = "us-east-1"
}

variable "container_image" {
  description = "Docker image URI (e.g. your ECR image)"
  type        = string
}