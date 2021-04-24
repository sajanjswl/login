[
    {
        "name": "api",
        "image": "${app_image}",
        "essential": true,
        "memoryReservation": 256,
        "environment": [
            {"name": "DB_HOST", "value": "${db_host}"},
            {"name": "DB_NAME", "value": "${db_name}"},
            {"name": "DB_USER", "value": "${db_user}"},
            {"name": "DB_PASS", "value": "${db_pass}"},
            {"name": "DB_PORT", "value": "5432"},

            {"name": "GRPC_NETWORK_TYPE", "value": "tcp"},
            {"name": "GRPC_HOST", "value": "localhost"}, 
            {"name": "GRPC_PORT", "value": "8000"},
            {"name": "REST_HOST", "value": "localhost"},
            {"name": "REST_PORT", "value": "9000"},
            {"name": "DB_DIALECT", "value": "postgres"},

            {"name": "TOKEN_KEY", "value": "something"},
            {"name": "TOKEN_ISSUER ", "value": "Tesla"},
            {"name": "ACCESS_TOKEN_ALIVE_TIME", "value": "1000000"},
            {"name": "REFRESH_TOKEN_ALIVE_TIME", "value": "100000"},

            {"name": "OTP_LENGTH", "value": "6"},
            {"name": "OTP_EXPIRE_TIME", "value": "1"},

            {"name": "OTP_SENDER", "value": "sjnjaiswal2@gmail.com"},
            {"name": "LOGIN_UI", "value": "/v1/login/ui/"},
            {"name": "TEMPLATES_HTML", "value": "./templates"},


            {"name": "GOOGLE_LOGIN_ENDPOINT", "value": "/auth/google/login"},
            {"name": "GOOGLE_CALLBACK_ENDPOINT", "value": "/auth/google/callback"},
            {"name": "GOOGLE_REDIRECT_URL", "value": "http://localhost:9000/auth/google/callback"},
            {"name": "GOOGLE_CLIENT_ID", "value": "997183094499-fidl2a9ol2gutlmjbpaadrnndbura862.apps.googleusercontent.com"},
            {"name": "GOOGLE_CLIENT_SECRET", "value": "-7rHfJbWJW-RS50McT6g8iEG"},

            {"name": "FACEBOOK_LOGIN_END_POINT", "value": "/auth/facebook/login"},
            {"name": "FACEBOOK_CALLBACK_END_POINT", "value": "/auth/facebook/callback"},
            {"name": "FAEBOOK_CLIENT_ID", "value": "694838257754877"},
            {"name": "FACEBOOK_CLIENT_SECRET", "value": "a5d062c0b2cf2082c50610365ff5ff57"},
            {"name": "FACEBOOK_REDIRECT_URL", "value": "http://localhost:9000/auth/facebook/callback"},

            {"name": "EXPIRE_TIME_OUTH_COOKIE", "value": "20"}, 
            {"name": "USER_BLOCKED_RESET_TIME", "value": "10"},

            {"name": "AWS_ACCESS_KEY_ID", "value": "AKIATY2HYWVHFCHKYQMM"},
            {"name": "AWS_SECRET_ACCESS_KEY", "value": "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"},
            {"name": "AWS_REGION", "value": "ap-south-1"},
            {"name": "AWS_SMTP_USER", "value": "AKIATY2HYWVHB2NUXW5Y"},
            {"name": "AWS_SMTP_PASS", "value": "BA63wwpaOaJVWeFdYTlqFVUGatbJye4px+oTwYmEAT4M"},
            {"name": "AWS_HOST", "value": "email-smtp.ap-south-1.amazonaws.com"},
            {"name": "AWS_PORT", "value": "587"},
            {"name": "AWS_SENDER_EMAIL", "value": "sjnjaiswal2@gmail.com"}
        ],
        "environmentFiles": [
                {
                    "value": "arn:aws:s3:::user-service-aws-terraform-ecs-environmentfile/.env",
                    "type": "s3"
                }
            ],
        "logConfiguration": {
            "logDriver": "awslogs",
            "options": {
                "awslogs-group": "${log_group_name}",
                "awslogs-region": "${log_group_region}",
                "awslogs-stream-prefix": "api"
            }
        },
        "portMappings": [
            {
                "containerPort": 8000,
                "hostPort": 8000
            },
             {
              "containerPort": 9000,
                "hostPort": 9000
             }
        ],
        "mountPoints": [
            {
                "readOnly": false,
                "containerPath": "/vol/web",
                "sourceVolume": "static"
            }
        ]
    },
    {
        "name": "proxy",
        "image": "${proxy_image}",
        "essential": true,
        "portMappings": [
            {
                "containerPort": 8080,
                "hostPort": 8080
            },
            {
             "containerPort": 8085,
                "hostPort": 8085
            }
        ],
        "memoryReservation": 256,
        "environment": [
            {"name": "APP_HOST", "value": "127.0.0.1"},
            {"name": "REST_PORT", "value": "9000"},
            {"name": "GRPC_PORT", "value": "8000"},
             {"name": "GRPC_LISTEN_PORT", "value": "8080"},
              {"name": "REST_LISTEN_PORT", "value": "8085"}
        ],
        "logConfiguration": {
            "logDriver": "awslogs",
            "options": {
                "awslogs-group": "${log_group_name}",
                "awslogs-region": "${log_group_region}",
                "awslogs-stream-prefix": "proxy"
            }
        },
        "mountPoints": [
            {
                "readOnly": true,
                "containerPath": "/vol/static",
                "sourceVolume": "static"
            }
        ]
    }
]
