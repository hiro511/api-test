package config

type configure struct {
	dbBasePath string
}

var c = configure{}

func init() {
	setDefault()
}

func setDefault() {
	// set defaults
}

func SetDBBasePath(v string) {
	c.dbBasePath = v
}

func DBBasePath() string {
	return c.dbBasePath
}
