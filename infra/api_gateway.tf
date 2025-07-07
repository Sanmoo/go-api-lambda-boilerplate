resource "aws_apigatewayv2_api" "media_tracker_api" {
  name          = "media-tracker-api"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_authorizer" "media_tracker_authorizer" {
  api_id          = aws_apigatewayv2_api.media_tracker_api.id
  name            = "media-tracker-authorizer"
  authorizer_type = "JWT"
  identity_sources = [
    "$request.header.Authorization",
  ]

  jwt_configuration {
    audience = ["https://media-tracker-api-gateway"]
    issuer   = "https://dev-4kd22ihe5nq17n4h.us.auth0.com/"
  }
}

data "aws_caller_identity" "current" {}

output "media_tracker_api_id" {
  value = aws_apigatewayv2_api.media_tracker_api.id
}

output "media_tracker_authorizer_id" {
  value = aws_apigatewayv2_authorizer.media_tracker_authorizer.id
}

resource "aws_apigatewayv2_stage" "default" {
  api_id = aws_apigatewayv2_api.media_tracker_api.id
  name   = "$default"
  auto_deploy = true
}