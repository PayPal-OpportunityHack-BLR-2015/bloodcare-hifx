package services

import (
	"github.com/PayPal-OpportunityHack-BLR-2015/bloodcare-hifx/admin-dash/app"
	"github.com/gocql/gocql"
)

// Cassandra service
type Cassandra struct {
	session *gocql.Session
}

// NewCassandra returns a cassandra object
func NewCassandra(host, keyspace string) *Cassandra {
	// connect to the cluster
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	csession, err := cluster.CreateSession()
	app.Chk(err)
	return &Cassandra{session: csession}

}

func (c *Cassandra) Query(query string, values ...interface{}) *gocql.Query {
	return c.session.Query(query, values...)
}
