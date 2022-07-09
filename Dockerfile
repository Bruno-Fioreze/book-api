# Build Stage
# First pull Golang image
FROM golang:1.17-alpine as build-env
 
# Set environment variable
ENV APP_NAME books
ENV CMD_PATH main.go
ENV APP_PATH $GOPATH/src/$APP_NAME/

# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
 
# # Budild application
# RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# # Run Stage
# FROM alpine:3.14
 
# # Copy only required data into this image
# COPY --from=build-env /$APP_NAME .
 
# # Expose application port
# EXPOSE 5000
RUN chmod +x $GOPATH/src/$APP_NAME/entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]