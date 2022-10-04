// Package monitors provides types/client for making requests to /monitors.
package monitors

import "time"

// Monitor ...
type Monitor struct {
	ID             string        `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	UID            string        `json:"uid,omitempty"`
	Owner          int           `json:"owner,omitempty"`
	Collection     string        `json:"collection,omitempty"`
	Environment    string        `json:"environment,omitempty"`
	CollectionUID  string        `json:"collectionUid,omitempty"`
	EnvironmentUID string        `json:"environmentUid,omitempty"`
	Options        Options       `json:"options,omitempty"`
	Notifications  Notifications `json:"notifications,omitempty"`
	Distribution   []interface{} `json:"distribution,omitempty"`
	Schedule       Schedule      `json:"schedule,omitempty"`
	LastRun        Run           `json:"lastRun,omitempty"`
	Stats          Stats         `json:"stats,omitempty"`
}

// Options ...
type Options struct {
	StrictSSL       bool `json:"strictSSL,omitempty"`
	FollowRedirects bool `json:"followRedirects,omitempty"`
	RequestTimeout  int  `json:"requestTimeout,omitempty"`
	RequestDelay    int  `json:"requestDelay,omitempty"`
}

// Notifications ...
type Notifications struct {
	OnError   Notification `json:"onError,omitempty"`
	OnFailure Notification `json:"onFailure,omitempty"`
}

// Notification ...
type Notification struct {
	Email string `json:"email,omitempty"`
}

// Schedule ...
type Schedule struct {
	Cron     string    `json:"cron,omitempty"`
	Timezone string    `json:"timezone,omitempty"`
	NextRun  time.Time `json:"nextRun,omitempty"`
}

// Run ...
type Run struct {
	Status     string    `json:"status,omitempty"`
	StartedAt  time.Time `json:"startedAt,omitempty"`
	FinishedAt time.Time `json:"finishedAt,omitempty"`
}

// Stats ...
type Stats struct {
	Assertions Assertions `json:"assertions,omitempty"`
	Requests   Requests   `json:"total,omitempty"`
}

// Assertions ...
type Assertions struct {
	Total  int `json:"total,omitempty"`
	Failed int `json:"failed,omitempty"`
}

// Requests ...
type Requests struct {
	Total int `json:"total,omitempty"`
}

type monitorWrapper struct {
	Monitor  Monitor   `json:"monitor,omitempty"`
	Monitors []Monitor `json:"monitors,omitempty"`
}

type runWrapper struct {
	Run Run `json:"run,omitempty"`
}
