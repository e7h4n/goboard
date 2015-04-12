package logical

import "gopkg.in/gorp.v1"

// Context contains database connection
type Context struct {
	DbMap *gorp.DbMap
}
