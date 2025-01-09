// Copyright 2024 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanner

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/spanner"

	"github.com/google/fleetspeak/fleetspeak/src/common"
	"github.com/google/fleetspeak/fleetspeak/src/server/db"
	"github.com/google/fleetspeak/fleetspeak/src/server/ids"

	"google.golang.org/api/iterator"

	log "github.com/golang/glog"

	fspb "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"
	spb "github.com/google/fleetspeak/fleetspeak/src/server/proto/fleetspeak_server"
	anypb "google.golang.org/protobuf/types/known/anypb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

// dbBroadcast matches the schema of the broadcasts table.
type dbBroadcast struct {
	broadcastID           []byte
	sourceServiceName     string
	messageType           string
	expirationTimeSeconds sql.NullInt64
	expirationTimeNanos   sql.NullInt64
	dataTypeURL           sql.NullString
	dataValue             []byte
	sent                  uint64
	allocated             uint64
	messageLimit          uint64
}

type broadcast struct {
	PrimaryKey   *spanner.Key
	Broadcast    *spb.Broadcast
	Sent         uint64
	Allocated    uint64
	MessageLimit uint64
}

func fromBroadcastProto(b *spb.Broadcast) (*dbBroadcast, error) {
	if b == nil {
		return nil, errors.New("cannot convert nil Broadcast")
	}
	id, err := ids.BytesToBroadcastID(b.BroadcastId)
	if err != nil {
		return nil, err
	}
	if b.Source == nil {
		return nil, fmt.Errorf("Broadcast must have Source. Get: %v", b)
	}

	res := dbBroadcast{
		broadcastID:       id.Bytes(),
		sourceServiceName: b.Source.ServiceName,
		messageType:       b.MessageType,
	}
	if b.ExpirationTime != nil {
		res.expirationTimeSeconds = sql.NullInt64{Int64: b.ExpirationTime.Seconds, Valid: true}
		res.expirationTimeNanos = sql.NullInt64{Int64: int64(b.ExpirationTime.Nanos), Valid: true}
	}
	if b.Data != nil {
		res.dataTypeURL = sql.NullString{String: b.Data.TypeUrl, Valid: true}
		res.dataValue = b.Data.Value
	}
	return &res, nil
}

func toBroadcastProto(b *dbBroadcast) (*spb.Broadcast, error) {
	bid, err := ids.BytesToBroadcastID(b.broadcastID)
	if err != nil {
		return nil, err
	}
	ret := &spb.Broadcast{
		BroadcastId: bid.Bytes(),
		Source:      &fspb.Address{ServiceName: b.sourceServiceName},
		MessageType: b.messageType,
	}
	if b.expirationTimeSeconds.Valid && b.expirationTimeNanos.Valid {
		ret.ExpirationTime = &tspb.Timestamp{
			Seconds: b.expirationTimeSeconds.Int64,
			Nanos:   int32(b.expirationTimeNanos.Int64),
		}
	}
	if b.dataTypeURL.Valid {
		ret.Data = &anypb.Any{
			TypeUrl: b.dataTypeURL.String,
			Value:   b.dataValue,
		}
	}
	return ret, nil
}

// CreateBroadcast implements db.BroadcastStore.
func (d *Datastore) CreateBroadcast(ctx context.Context, b *spb.Broadcast, limit uint64) error {
	log.Error("----------- broadcaststore: CreateBroadcast() called")
	dbB, err := fromBroadcastProto(b)
	if err != nil {
		return err
	}
	dbB.messageLimit = limit
	return nil
}

// SetBroadcastLimit implements db.BroadcastStore.
func (d *Datastore) SetBroadcastLimit(ctx context.Context, id ids.BroadcastID, limit uint64) error {
	log.Error("----------- broadcaststore: SetBroadcastLimit() called")
	return nil
}

// SaveBroadcastMessage implements db.BroadcastStore.
func (d *Datastore) SaveBroadcastMessage(ctx context.Context, msg *fspb.Message, bID ids.BroadcastID, cID common.ClientID, aID ids.AllocationID) error {
    log.Error("----------- broadcaststore: SaveBroadcastMessage() called")
	return nil
}

// ListActiveBroadcasts implements db.BroadcastStore.
func (d *Datastore) ListActiveBroadcasts(ctx context.Context) ([]*db.BroadcastInfo, error) {
	//log.Error("+++ broadcaststore: ListActiveBroadcasts() called")
	now := db.NowProto()
	stmt := spanner.Statement{
		SQL: "SELECT " +
			"b.Broadcast, " +
			"b.Sent, " +
			"b.MessageLimit " +
			"FROM Broadcasts AS b " +
			"WHERE b.Sent < b.MessageLimit " +
			"AND (b.Broadcast.expiration_time IS NULL OR (b.Broadcast.expiration_time.seconds > @nowSec OR (b.Broadcast.expiration_time.seconds = @nowSec AND b.Broadcast.expiration_time.nanos > @nowNano)))",
		Params: map[string]interface{}{
			"nowSec":  int64(now.Seconds),
			"nowNano": int64(now.Nanos),
		},
	}

	var ret []*db.BroadcastInfo
	txn := d.dbClient.Single()
	defer txn.Close()
	iter := txn.Query(ctx, stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			return ret, nil
		}
		if err != nil {
			return ret, err
		}
		var broadcast *spb.Broadcast
		var sent, messageLimit uint64
		if err := row.Columns(&broadcast, &sent, &messageLimit); err != nil {
			return ret, err
		}
		ret = append(ret, &db.BroadcastInfo{
			Broadcast: broadcast,
			Sent:      sent,
			Limit:     messageLimit,
		})
	}
}

// ListSentBroadcasts implements db.BroadcastStore.
func (d *Datastore) ListSentBroadcasts(ctx context.Context, id common.ClientID) ([]ids.BroadcastID, error) {
	log.Error("+++ broadcaststore: ListSentBroadcasts() called")
	var res []ids.BroadcastID
	ro := d.dbClient.ReadOnlyTransaction()
	defer ro.Close()
	krp := spanner.Key{id.Bytes()}.AsPrefix()
	iter := ro.Read(ctx, d.broadcastSent, krp, []string{"BroadcastID"})
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return res, err
		}
		var b []byte
		if err := row.Columns(&b); err != nil {
			return res, err
		}
		bid, err := ids.BytesToBroadcastID(b)
		if err != nil {
			return res, err
		}
		res = append(res, bid)
	}
	return res, nil
}

// CreateAllocation implements db.BroadcastStore.
func (d *Datastore) CreateAllocation(ctx context.Context, id ids.BroadcastID, frac float32, expiry time.Time) (*db.AllocationInfo, error) {
	log.Error("----------- broadcaststore: CreateAllocation() called")
	var ret *db.AllocationInfo
	return ret, nil
}

// CleanupAllocation implements db.BroadcastStore.
func (d *Datastore) CleanupAllocation(ctx context.Context, bID ids.BroadcastID, aID ids.AllocationID) error {
	log.Error("----------- broadcaststore: CleanupAllocation() called")
	return nil
}
