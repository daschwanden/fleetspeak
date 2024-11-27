#!/bin/bash
export SPANNER_DATABASE=fleetspeak
export SPANNER_INSTANCE=fleetspeak-instance
export SPANNER_PROJECT=fleetspeak-spanner
export SPANNER_TOPIC=fleetspeak-server-messages
export SPANNER_SUBSCRIPTION=fleetspeak-server-messages-sub
cd ./fleetspeak
go test -race --timeout 2.5m github.com/google/fleetspeak/fleetspeak/src/server/spanner