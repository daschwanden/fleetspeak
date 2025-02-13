#!/bin/bash
# Use as following:
#   docker build -t fleetspeak:test -f Dockerfile.dev .
#   docker run -it --rm -v /home/user/.config/gcloud/:/root/.config/gcloud --entrypoint=/bin/bash fleetspeak:test
#   protoc -I=. -I=/usr/include --include_imports --descriptor_set_out=fleetspeak1.pb ./fleetspeak/src/common/proto/fleetspeak/common.proto ./fleetspeak/src/server/proto/fleetspeak_server/broadcasts.proto ./fleetspeak/src/server/proto/fleetspeak_server/resource.proto
#   ./test.sh
export SPANNER_DATABASE=fleetspeak
export SPANNER_INSTANCE=fleetspeak-instance
export SPANNER_PROJECT=fleetspeak-spanner
export SPANNER_TOPIC=fleetspeak-server-messages
export SPANNER_SUBSCRIPTION=fleetspeak-server-messages-sub
cd ./fleetspeak
go test -race --timeout 2.5m github.com/google/fleetspeak/fleetspeak/src/server/spanner