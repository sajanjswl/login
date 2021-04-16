
FROM nginx:alpine

COPY ./nginx.conf /etc/nginx/nginx.conf

ENV GRPC_LISTEN_PORT=8080
ENV REST_LISTEN_PORT=8085
ENV APP_HOST=user
ENV GRPC_PORT=8000
ENV REST_PORT=9000

CMD ["nginx", "-g", "daemon off;"]
