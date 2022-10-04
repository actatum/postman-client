// Package collections provides types/client for making requests to /collections.
package collections

import "time"

// Possible values for merge fork strategies.
const (
	MergeStrategyDeleteSource                = "deleteSource"
	MergeStrategyUpdateSourceWithDestination = "updateSourceWithDestination"
)

// Collection ...
type Collection struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Owner     string    `json:"owner,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UID       string    `json:"uid,omitempty"`
	IsPublic  bool      `json:"isPublic,omitempty"`
	Fork      Fork      `json:"fork,omitempty"`
}

// CollectionDetails ...
type CollectionDetails struct {
	Info  Info   `json:"info,omitempty"`
	Items []Item `json:"item,omitempty"`
}

// Info ...
type Info struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	PostmanID   string    `json:"_postman_id,omitempty"`
	Schema      string    `json:"schema,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	Fork        Fork      `json:"fork,omitempty"`
}

// Fork ...
type Fork struct {
	Label     string    `json:"label,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	From      string    `json:"from,omitempty"`
}

// Item ...
type Item struct {
	Name     string      `json:"name,omitempty"`
	ID       string      `json:"id,omitempty"`
	Events   []Event     `json:"event,omitempty"`
	Request  Request     `json:"request,omitempty"`
	Response interface{} `json:"response,omitempty"`
}

// Event ...
type Event struct {
	Listen string `json:"listen,omitempty"`
	Script Script `json:"script,omitempty"`
}

// Script ...
type Script struct {
	ID   string   `json:"id,omitempty"`
	Exec []string `json:"exec,omitempty"`
	Type string   `json:"type,omitempty"`
}

// Request ...
type Request struct {
	URL     string   `json:"url,omitempty"`
	Method  string   `json:"method,omitempty"`
	Headers []Header `json:"header,omitempty"`
}

// Header ...
type Header struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// MergeForkRequest ...
type MergeForkRequest struct {
	Strategy    string `json:"strategy,omitempty"`
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
}

type collectionWrapper struct {
	Collection  Collection   `json:"collection,omitempty"`
	Collections []Collection `json:"collections,omitempty"`
}

type collectionDetailsWrapper struct {
	Details CollectionDetails `json:"collection"`
}
