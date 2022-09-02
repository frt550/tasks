#!/bin/bash

# Start the first process
./grpc &

# Start the second process
./rest &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?

