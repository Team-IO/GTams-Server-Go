package entities

import "github.com/satori/go.uuid"

type Installation struct {
	Id uuid.UUID
	LastSeen int64
	Version string
	McVersion string
	Branding string
	Language string
}

type Terminal struct {
	Id uint64
	Owner uuid.UUID
	Online bool
}

type Player struct {
	Id uuid.UUID
	Online bool
	Funds uint64
	Name string
}

