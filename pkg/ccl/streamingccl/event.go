// Copyright 2020 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package streamingccl

import (
	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
)

// EventType enumerates all possible events emitted over a cluster stream.
type EventType int

const (
	// KVEvent indicates that the KV field of an event holds an updated KV which
	// needs to be ingested.
	KVEvent EventType = iota
	// SSTableEvent indicates that the SSTable field of an event holds an updated
	// SSTable which needs to be ingested.
	SSTableEvent
	// CheckpointEvent indicates that GetResolved will be meaningful. The resolved
	// timestamp indicates that all KVs have been emitted up to this timestamp.
	CheckpointEvent
	// GenerationEvent indicates that the stream should start ingesting with the
	// updated topology.
	GenerationEvent
)

// Event describes an event emitted by a cluster to cluster stream.  Its Type
// field indicates which other fields are meaningful.
type Event interface {
	// Type specifies which accessor will be meaningful.
	Type() EventType

	// GetKV returns a KV event if the EventType is KVEvent.
	GetKV() *roachpb.KeyValue

	// GetSSTable returns a SSTable event if the EventType is SSTable.
	GetSSTable() *roachpb.RangeFeedSSTable
	// GetResolved returns a resolved timestamp if the EventType is
	// CheckpointEvent. The resolved timestamp indicates that all KV events until
	// this time have been emitted.
	GetResolved() *hlc.Timestamp
}

// kvEvent is a key value pair that needs to be ingested.
type kvEvent struct {
	kv roachpb.KeyValue
}

var _ Event = kvEvent{}

// Type implements the Event interface.
func (kve kvEvent) Type() EventType {
	return KVEvent
}

// GetKV implements the Event interface.
func (kve kvEvent) GetKV() *roachpb.KeyValue {
	return &kve.kv
}

// GetSSTable implements the Event interface.
func (kve kvEvent) GetSSTable() *roachpb.RangeFeedSSTable {
	return nil
}

// GetResolved implements the Event interface.
func (kve kvEvent) GetResolved() *hlc.Timestamp {
	return nil
}

// sstableEvent is a sstable that needs to be ingested.
type sstableEvent struct {
	sst roachpb.RangeFeedSSTable
}

// Type implements the Event interface.
func (sste sstableEvent) Type() EventType {
	return SSTableEvent
}

// GetKV implements the Event interface.
func (sste sstableEvent) GetKV() *roachpb.KeyValue {
	return nil
}

// GetSSTable implements the Event interface.
func (sste sstableEvent) GetSSTable() *roachpb.RangeFeedSSTable {
	return &sste.sst
}

// GetResolved implements the Event interface.
func (sste sstableEvent) GetResolved() *hlc.Timestamp {
	return nil
}

var _ Event = sstableEvent{}

// checkpointEvent indicates that the stream has emitted every change for all
// keys in the span it is responsible for up until this timestamp.
type checkpointEvent struct {
	resolvedTimestamp hlc.Timestamp
}

var _ Event = checkpointEvent{}

// Type implements the Event interface.
func (ce checkpointEvent) Type() EventType {
	return CheckpointEvent
}

// GetKV implements the Event interface.
func (ce checkpointEvent) GetKV() *roachpb.KeyValue {
	return nil
}

// GetSSTable implements the Event interface.
func (ce checkpointEvent) GetSSTable() *roachpb.RangeFeedSSTable {
	return nil
}

// GetResolved implements the Event interface.
func (ce checkpointEvent) GetResolved() *hlc.Timestamp {
	return &ce.resolvedTimestamp
}

// generationEvent indicates that the topology of the stream has changed.
type generationEvent struct{}

var _ Event = generationEvent{}

// Type implements the Event interface.
func (ge generationEvent) Type() EventType {
	return GenerationEvent
}

// GetKV implements the Event interface.
func (ge generationEvent) GetKV() *roachpb.KeyValue {
	return nil
}

// GetSSTable implements the Event interface.
func (ge generationEvent) GetSSTable() *roachpb.RangeFeedSSTable {
	return nil
}

// GetResolved implements the Event interface.
func (ge generationEvent) GetResolved() *hlc.Timestamp {
	return nil
}

// MakeKVEvent creates an Event from a KV.
func MakeKVEvent(kv roachpb.KeyValue) Event {
	return kvEvent{kv: kv}
}

// MakeSSTableEvent creates an Event from a SSTable.
func MakeSSTableEvent(sst roachpb.RangeFeedSSTable) Event {
	return sstableEvent{sst: sst}
}

// MakeCheckpointEvent creates an Event from a resolved timestamp.
func MakeCheckpointEvent(resolvedTimestamp hlc.Timestamp) Event {
	return checkpointEvent{resolvedTimestamp: resolvedTimestamp}
}

// MakeGenerationEvent creates an GenerationEvent.
func MakeGenerationEvent() Event {
	return generationEvent{}
}
