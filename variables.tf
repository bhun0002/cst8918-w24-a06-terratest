# Define config variables
variable "labelPrefix" {
  type        = string
  description = "Your college username. This will form the beginning of various resource names."
  default     = "bhun0002-Go-Test"
}

variable "region" {
  default = "westus3"
}

variable "admin_username" {
  type        = string
  default     = "azureadmin"
  description = "The username for the local user account on the VM."
}

variable "delete_timeout" {
  description = "Timeout duration for resource deletion"
  type        = string
  default     = "30m"
}
