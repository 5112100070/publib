package dummyrds

// Mocker mapping
type Mocker map[string]mock

type dummydis struct {
	config Config
}

// Config for dummy elastic
type Config struct {
	MockingMap Mocker
}

// Mock result
type mock struct {
	IsError bool
	Result  interface{}
}
