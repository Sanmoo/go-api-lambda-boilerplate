provider "aws" {
  region = "us-east-1"
}

data "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"
}

resource "aws_lambda_function" "books_lambda" {
  filename      = "${path.module}/artifacts/books.zip"
  function_name = "media-tracker-books-dev"
  role          = data.aws_iam_role.iam_for_lambda.arn
  handler       = "run"

  source_code_hash = filesha256("${path.module}/../books/lambda.go")

  runtime = "provided.al2"
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

resource "aws_apigatewayv2_integration" "books_lambda_integration" {
  api_id                 = data.aws_apigatewayv2_api.media_tracker_api.id
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.books_lambda.invoke_arn
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "books_lambda_route" {
  api_id             = data.aws_apigatewayv2_api.media_tracker_api.id
  route_key          = "GET /books"
  target             = "integrations/${aws_apigatewayv2_integration.books_lambda_integration.id}"
  authorizer_id      = data.terraform_remote_state.common_infra.outputs.media_tracker_authorizer_id
  authorization_type = "JWT"
}

data "aws_caller_identity" "current" {}

resource "aws_lambda_permission" "apigateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.books_lambda.arn
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:us-east-1:${data.aws_caller_identity.current.account_id}:${data.aws_apigatewayv2_api.media_tracker_api.id}/*/*"
}