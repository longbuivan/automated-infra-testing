

output "raw_data_s3_id" {
  value = aws_s3_bucket.raw_data_s3.id
}

output "raw_s3_bucket_region" {
   value = aws_s3_bucket.raw_data_s3.region
}