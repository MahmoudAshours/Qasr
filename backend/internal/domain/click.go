package model

import "time"

type Click struct {
	Slug       string    `bson:"slug"`
	Timestamp  time.Time `bson:"timestamp"`
	IP         string    `bson:"ip"`
	Country    string    `bson:"country"`
	City       string    `bson:"city"`
	Timezone   string    `bson:"timezone"`
	UserAgent  string    `bson:"user_agent"`
	Browser    string    `bson:"browser"`
	DeviceType string    `bson:"device_type"`
	IsBot      bool      `bson:"is_bot"`
	Referrer   string    `bson:"referrer"`
	Language   string    `bson:"language"`
}
