package ybx

import (
	"log"
	"sync"

	"github.com/tinhminhtue/go-reused-lib/core/dbx/dbcfg"
	"github.com/yugabyte/gocql"
)

type CqlSession struct {
	clusterConfig dbcfg.DBConfig
	session       *gocql.Session
}

// singleton with once
var (
	cqlSession *CqlSession
	once       sync.Once
)

func (s *CqlSession) GetClusterConfig() dbcfg.DBConfig {
	return s.clusterConfig
}

func (s *CqlSession) GetSession() *gocql.Session {
	return s.session
}

func GetCqlSession() *CqlSession {
	once.Do(func() {
		cqlSession = &CqlSession{}
	})
	return cqlSession
}

// InitSession init session
// remember call this function in main.go before using session
func (s *CqlSession) InitSession(dbCf dbcfg.DBConfig) error {
	s.clusterConfig = dbCf
	cluster := gocql.NewCluster(dbCf.Hosts...)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: dbCf.Username,
		Password: dbCf.Password,
	}
	cluster.SslOpts = &gocql.SslOptions{
		CaPath: dbCf.CaPath,
	}
	cluster.Keyspace = dbCf.Keyspace
	cluster.Consistency = dbCf.ConvertConsistencyStringToCqlConsistency()
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	s.session = session
	return nil

}
