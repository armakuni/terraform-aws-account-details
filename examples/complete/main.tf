provider "aws" {}

module "account_details" {
  source = "../../"
}

output "canonical_user_id" {
  value = module.account_details.canonical_user_id
}
