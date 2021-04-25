resource "aws_lb" "api" {
  name               = "${local.prefix}-main"
  load_balancer_type = "application"
  subnets = [
    aws_subnet.public_a.id,
    aws_subnet.public_b.id
  ]

  security_groups = [aws_security_group.lb.id]

  tags = local.common_tags
}

## target group is a group of servers that lb can  forward requests to. This is for HTTP call
resource "aws_lb_target_group" "api" {
  name        = "${local.prefix}-api"
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id
  target_type = "ip"
  port        = 8085

  health_check {
    path = "/v1/login/ui/"
  }
}

resource "aws_lb_listener" "api" {
  load_balancer_arn = aws_lb.api.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.api.arn
  }
}

#  for GRPC
# resource "aws_lb_target_group" "apig" {
#   name     = "${local.prefix}-apig"
#   protocol = "HTTP"
#   protocol_version = "GRPC"
#   vpc_id           = aws_vpc.main.id
#   target_type      = "ip"
#   port             = 8080

# }

# resource "aws_lb_listener" "apig" {
#   load_balancer_arn = aws_lb.api.arn
#   port              = 90

#   protocol         = "HTTPS"

#   default_action {
#     type             = "forward"
#     target_group_arn = aws_lb_target_group.apig.arn
#   }
# }


resource "aws_security_group" "lb" {
  description = "Allow access to Application Load Balancer"
  name        = "${local.prefix}-lb"
  vpc_id      = aws_vpc.main.id

  # for HTTP
  ingress {
    protocol    = "tcp"
    from_port   = 80
    to_port     = 80
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    protocol    = "tcp"
    from_port   = 8085
    to_port     = 8085
    cidr_blocks = ["0.0.0.0/0"]
  }


  # for GRPC

  # ingress {
  #   protocol    = "tcp"
  #   from_port   = 90
  #   to_port     = 90
  #   cidr_blocks = ["0.0.0.0/0"]
  # }

  # egress {
  #   protocol    = "tcp"
  #   from_port   = 8080
  #   to_port     = 8080
  #   cidr_blocks = ["0.0.0.0/0"]
  # }

  tags = local.common_tags
}
