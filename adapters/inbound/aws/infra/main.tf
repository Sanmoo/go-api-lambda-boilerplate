provider "aws" {
  region = "us-east-1"
}

data "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"
}

data "terraform_remote_state" "common_infra" {
  backend = "s3"
  config = {
    bucket = "deployer-terraform-states",
    key    = "media-tracker/common-infra",
    region = "us-east-1"
  }
}

data "aws_apigatewayv2_api" "media_tracker_api" {
  api_id = data.terraform_remote_state.common_infra.outputs.media_tracker_api_id
}

// define a map of strings to strings
locals {
  lambdas = ["books", "movies", "electronic-games", "non-electronic-games", "tv-series"]
}

resource "aws_lambda_function" "lambda" {
  for_each = toset(local.lambdas)
  filename      = "${path.module}/artifacts/${each.value}.zip"
  function_name = "media-tracker-${each.value}-dev"
  role          = data.aws_iam_role.iam_for_lambda.arn
  handler       = "run"

  source_code_hash = filesha256("${path.module}/../${each.value}/bootstrap")

  runtime = "provided.al2"
  
  logging_config {
    log_format = "JSON"
    application_log_level = "DEBUG"
    system_log_level = "DEBUG"
  }
  
  depends_on = [ aws_cloudwatch_log_group.lambda_log_group ]
}

resource "aws_cloudwatch_log_group" "lambda_log_group" {
  for_each = toset(local.lambdas)
  name     = "/aws/lambda/media-tracker-${each.value}-dev"
  retention_in_days = 3
}

resource "aws_apigatewayv2_integration" "lambda_integration" {
  for_each               = toset(local.lambdas)
  api_id                 = data.aws_apigatewayv2_api.media_tracker_api.id
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.lambda[each.value].invoke_arn
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "books_lambda_base_route" {
  for_each = toset(local.lambdas)

  api_id             = data.aws_apigatewayv2_api.media_tracker_api.id
  route_key          = "ANY /${each.value}"
  target             = "integrations/${aws_apigatewayv2_integration.lambda_integration[each.value].id}"
  authorizer_id      = data.terraform_remote_state.common_infra.outputs.media_tracker_authorizer_id
  authorization_type = "JWT"
}

resource "aws_apigatewayv2_route" "books_lambda_route" {
  for_each = toset(local.lambdas)

  api_id             = data.aws_apigatewayv2_api.media_tracker_api.id
  route_key          = "ANY /${each.value}/{proxy+}"
  target             = "integrations/${aws_apigatewayv2_integration.lambda_integration[each.value].id}"
  authorizer_id      = data.terraform_remote_state.common_infra.outputs.media_tracker_authorizer_id
  authorization_type = "JWT"
}

data "aws_caller_identity" "current" {}

resource "aws_lambda_permission" "apigateway" {
  for_each = toset(local.lambdas)

  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda[each.value].arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:us-east-1:${data.aws_caller_identity.current.account_id}:${data.aws_apigatewayv2_api.media_tracker_api.id}/*/*"
}