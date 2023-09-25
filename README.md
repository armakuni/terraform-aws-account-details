# AWS Account Details Module

A module that provides various details about the AWS account.
This module exists to provide input values for other module rather than them using `data` providers themselves.
This is particularly useful when using [Terragrunt](https://terragrunt.gruntwork.io/).

<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_canonical_user_id.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/canonical_user_id) | data source |

## Inputs

No inputs.

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_canonical_user_id"></a> [canonical\_user\_id](#output\_canonical\_user\_id) | The output from aws\_canonical\_user |
<!-- END_TF_DOCS -->
