package raft

import (
	"time"
)

// Config provides any necessary configuraiton to
// the Raft server
type Config struct {
	// Time in follower state without a leader before we attempt an election
	HeartbeatTimeout time.Duration

	// Time in candidate state without a leader before we attempt an election
	ElectionTimeout time.Duration

	// Time without an Apply() operation before we heartbeat to ensure
	// a timely commit. Should be far less than HeartbeatTimeout to ensure
	// we don't lose leadership.
	CommitTimeout time.Duration

	// MaxAppendEntries controls the maximum number of append entries
	// to send at once. We want to strike a balance between efficiency
	// and avoiding waste if the follower is going to reject because of
	// an inconsistent log
	MaxAppendEntries int
}

func DefaultConfig() *Config {
	return &Config{
		HeartbeatTimeout: 100 * time.Millisecond,
		ElectionTimeout:  150 * time.Millisecond,
		CommitTimeout:    5 * time.Millisecond,
		MaxAppendEntries: 16,
	}
}
