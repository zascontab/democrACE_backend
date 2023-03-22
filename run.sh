# Check if there is a built executable and delete it
if [ -f "./governance" ]; then
  rm ./governance
fi

# Build the Go app
go build

# Execute the binary, passing an optional flag argument
./governance $1