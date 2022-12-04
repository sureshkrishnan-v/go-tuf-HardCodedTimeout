package metadata

import (
	"encoding/json"
	"sync"
	"time"
)

// Generic type constraint
type Roles interface {
	RootType | SnapshotType | TimestampType | TargetsType
}

// Define version of the TUF specification
const (
	SPECIFICATION_VERSION = "1.0.31"
)

// Define top level role names
const (
	ROOT      = "root"
	SNAPSHOT  = "snapshot"
	TARGETS   = "targets"
	TIMESTAMP = "timestamp"
)

type Metadata[T Roles] struct {
	Signed     T           `json:"signed"`
	Signatures []Signature `json:"signatures"`
}

type Signature struct {
	KeyID     string   `json:"keyid"`
	Signature HexBytes `json:"sig"`
}

type RootType struct {
	Type               string           `json:"_type"`
	SpecVersion        string           `json:"spec_version"`
	ConsistentSnapshot bool             `json:"consistent_snapshot"`
	Version            int64            `json:"version"`
	Expires            time.Time        `json:"expires"`
	Keys               map[string]*Key  `json:"keys"`
	Roles              map[string]*Role `json:"roles"`
	Custom             json.RawMessage  `json:"custom,omitempty"`
}

type SnapshotType struct {
	Type        string               `json:"_type"`
	SpecVersion string               `json:"spec_version"`
	Version     int64                `json:"version"`
	Expires     time.Time            `json:"expires"`
	Meta        map[string]MetaFiles `json:"meta"`
	Custom      json.RawMessage      `json:"custom,omitempty"`
}

type TargetsType struct {
	Type        string                 `json:"_type"`
	SpecVersion string                 `json:"spec_version"`
	Version     int64                  `json:"version"`
	Expires     time.Time              `json:"expires"`
	Targets     map[string]TargetFiles `json:"targets"`
	Delegations *Delegations           `json:"delegations,omitempty"`
	Custom      json.RawMessage        `json:"custom,omitempty"`
}

type TimestampType struct {
	Type        string               `json:"_type"`
	SpecVersion string               `json:"spec_version"`
	Version     int64                `json:"version"`
	Expires     time.Time            `json:"expires"`
	Meta        map[string]MetaFiles `json:"meta"`
	Custom      json.RawMessage      `json:"custom,omitempty"`
}

type Key struct {
	Type   KeyType         `json:"keytype"`
	Scheme KeyScheme       `json:"scheme"`
	Value  json.RawMessage `json:"keyval"`
	Custom json.RawMessage `json:"custom,omitempty"`
	id     string
	idOnce sync.Once
}

type Role struct {
	KeyIDs    []string `json:"keyids"`
	Threshold int      `json:"threshold"`
}

type HexBytes []byte

type KeyType string

type KeyScheme string

type Hashes map[string]HexBytes

type MetaFiles struct {
	Length  int64           `json:"length,omitempty"`
	Hashes  Hashes          `json:"hashes,omitempty"`
	Version int64           `json:"version"`
	Custom  json.RawMessage `json:"custom,omitempty"`
}

type TargetFiles struct {
	Length int64           `json:"length"`
	Hashes Hashes          `json:"hashes"`
	Custom json.RawMessage `json:"custom,omitempty"`
}

type Delegations struct {
	Keys  map[string]*Key `json:"keys"`
	Roles []DelegatedRole `json:"roles"`
}

type DelegatedRole struct {
	Name             string   `json:"name"`
	KeyIDs           []string `json:"keyids"`
	Threshold        int      `json:"threshold"`
	Terminating      bool     `json:"terminating"`
	PathHashPrefixes []string `json:"path_hash_prefixes,omitempty"`
	Paths            []string `json:"paths"`
}