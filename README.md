# protobuf-bid-request
# Takes a json bid request and converts to protobuf

# First clone this repo
git clone git@github.com:jonSteck/protobuf-bid-request.git
cd protobuf-bid-request

# Set your $GOPATH
source gvp in

# Download Dependencies
gpm

# Compile
go get ./...

# Run program
bid
