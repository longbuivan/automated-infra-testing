resource "aws_s3_bucket" "raw_data_s3" {
  bucket = "${var.environment}-raw-data-s3"
  acl    = "private"
  tags   = local.tags
  versioning {
    enabled = false
  }
}

resource "aws_s3_bucket_public_access_block" "raw_data_s3_acl" {
  bucket                  = aws_s3_bucket.raw_data_s3.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}