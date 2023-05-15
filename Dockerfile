FROM golang:1.18-alpine

# Set up environment variables
ENV APP_HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

# Set up working directory and copy the code
WORKDIR $APP_HOME
COPY . $APP_HOME

# Copy the config file
COPY resources/config.json $APP_HOME/

# Install dependencies
RUN go mod download

# Build the app
RUN go build -o todo .

# Expose port 443 and start the app with HTTPS
EXPOSE 8080
CMD ["./todo"]
