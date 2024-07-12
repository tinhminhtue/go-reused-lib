package jwtnats

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
	"github.com/sirupsen/logrus"
)

type NatsJwtConfig struct {
	ServerUrl             string
	AdminJwt              string // jwt is public key and claims
	AdminSeed             string // seed is private key
	AccountPublicKey      string // as ID for the account
	AccountSigningSeedKey nkeys.KeyPair
	TimeoutForJwtUpdate   time.Duration
}

type NatsJwtProvider struct {
	Config    NatsJwtConfig
	AdminConn *nats.Conn
}

type UserJwtKeyPair struct {
	JwtKey  string
	SeedKey string
}

type UserClaim struct {
	// array of string, allowed subjects for publishing
	PubAllows []string `json:"pubAllows"`
	// array of string, allowed subjects for subscribing
	SubAllows []string `json:"subAllows"`
	// Timeouts for the user claim
	Expires int64 `json:"exp"`
}

// default global instance
var Instance = NatsJwtProvider{}

// Setup the global instance, call this fuction before using the global instance
func Setup(config NatsJwtConfig) {
	Instance.Config = config
}

func cachedConnection() (*nats.Conn, error) {
	// userJWT and userKeyPair can be used in conjunction with this nats.Option
	// var jwtAuthOption nats.Option
	if Instance.AdminConn == nil || Instance.AdminConn.IsClosed() {
		natsAdminOption := nats.UserJWTAndSeed(Instance.Config.AdminJwt, Instance.Config.AdminSeed)
		ncSys, err := nats.Connect(Instance.Config.ServerUrl, natsAdminOption)
		if err != nil {
			return nil, err
		}
		Instance.AdminConn = ncSys
		return ncSys, nil
	} else {
		return Instance.AdminConn, nil
	}
}

func RequestUser(userClaim UserClaim) (userKeys UserJwtKeyPair, err error) {
	userPublicKey, userSeed, _ := generateUserKey()
	userJWT, err := generateUserJWT(userClaim, userPublicKey, Instance.Config.AccountPublicKey, Instance.Config.AccountSigningSeedKey)
	if err != nil {
		return userKeys, err
	}
	userKeys.JwtKey = userJWT
	userKeys.SeedKey = string(userSeed)
	// connect with System NKey to update user claims
	ncSys, err := cachedConnection()
	if err != nil {
		return userKeys, err
	}
	msg, err := ncSys.Request("$SYS.REQ.CLAIMS.UPDATE", []byte(userJWT), Instance.Config.TimeoutForJwtUpdate)
	if err != nil {
		return userKeys, err
	}
	logrus.Infof("Nats Jwt Claims Response: %s", string(msg.Data))

	return userKeys, nil
}
