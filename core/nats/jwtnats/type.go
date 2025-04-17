package jwtnats

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
	"github.com/sirupsen/logrus"
)

type NatsJwtConfig struct {
	ServerUrl                 string
	AdminJwt                  string // jwt is public key and claims
	AdminSeed                 string // seed is private key
	AccountPublicKey          string // as ID for the account
	AccountSigningSeedKey     string
	accountSigningSeedKeyPair nkeys.KeyPair
	TimeoutForJwtUpdate       time.Duration
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
	// update the account signing key
	Instance.Config.accountSigningSeedKeyPair = GetAccountSigningKey(Instance.Config.AccountSigningSeedKey)
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
	userJWT, err := generateUserJWT(userClaim, userPublicKey, Instance.Config.AccountPublicKey, Instance.Config.accountSigningSeedKeyPair)
	if err != nil {
		return userKeys, err
	}
	userKeys.JwtKey = userJWT
	userKeys.SeedKey = string(userSeed)

	// connect with System NKey to update user claims (This is no need, only required for Account level)

	// ncSys, err := cachedConnection()
	// if err != nil {
	// 	return userKeys, err
	// }
	// _, err = ncSys.Request("$SYS.REQ.CLAIMS.UPDATE", []byte(userJWT), Instance.Config.TimeoutForJwtUpdate)
	// if err != nil {
	// 	if err.Error() == "nats: no responders available for request" { // false positive error, nats server works fine for the update
	// 		err = nil
	// 	} else {
	// 		return userKeys, err
	// 	}
	// }

	return userKeys, nil
}

// this function is used to test connect to NATS server with the user keys
func TestUserKeys(userKeys UserJwtKeyPair) (err error) {
	// userJWT and userKeyPair can be used in conjunction with this nats.Option
	userKeyPair, err := nkeys.FromSeed([]byte(userKeys.SeedKey))
	if err != nil {
		return err
	}
	jwtAuthOption := nats.UserJWT(func() (string, error) {
		return userKeys.JwtKey, nil
	},
		func(bytes []byte) ([]byte, error) {
			return userKeyPair.Sign(bytes)
		},
	)

	// use in a connection as desired
	nc, err := nats.Connect(Instance.Config.ServerUrl, jwtAuthOption)
	if err != nil {
		panic(err)
	}
	if nc.IsConnected() {
		defer nc.Close()
		logrus.Println("Connected to NATS")
	}

	// test subscribe subject.foo
	_, err = nc.Subscribe("subject.foo", func(msg *nats.Msg) {
		logrus.Println("Received message: ", string(msg.Data))
	})
	if err != nil {
		return err
	}

	// try to publish subject.foo
	err = nc.Publish("subject.foo", []byte("Hello World!"))
	if err != nil { // NOTE: This will fail because the user is not allowed to publish to subject.foo (but server will not return an error)
		logrus.Println("Failed to publish message")
	}
	// wait for the message to be received
	time.Sleep(5 * time.Second)
	return nil
}
