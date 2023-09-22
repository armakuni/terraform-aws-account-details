output "canonical_user_id" {
  value       = data.aws_canonical_user_id.current
  description = "The output from aws_canonical_user"
}
