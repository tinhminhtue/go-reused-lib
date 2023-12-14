package dbcfg

import "github.com/yugabyte/gocql"

type DBConfig struct {
	Hosts       []string
	Username    string
	Password    string
	CaPath      string
	Keyspace    string
	Consistency string
}

func (dbCfg *DBConfig) ConvertConsistencyStringToCqlConsistency() gocql.Consistency {
	switch dbCfg.Consistency {
	case "Any":
		return gocql.Any
	case "One":
		return gocql.One
	case "Two":
		return gocql.Two
	case "Three":
		return gocql.Three
	case "Quorum":
		return gocql.Quorum
	case "All":
		return gocql.All
	case "LocalQuorum":
		return gocql.LocalQuorum
	case "EachQuorum":
		return gocql.EachQuorum
	// case "Serial":
	// 	return gocql.Serial
	// case "LocalSerial":
	// 	return gocql.LocalSerial
	case "LocalOne":
		return gocql.LocalOne
	default:
		return gocql.Quorum
	}
}
