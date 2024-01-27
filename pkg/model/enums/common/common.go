package common

type DeviceOS int32

const (
	DeviceOS__UNKNOWN DeviceOS = iota
	DeviceOS__IOS
	DeviceOS__Android
	DeviceOS__Web
)
