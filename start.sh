#!/bin/sh

# Start Nginx in the background
nginx &

# Wait for Nginx to fully start
sleep 2

# Run the Go application
exec /main
