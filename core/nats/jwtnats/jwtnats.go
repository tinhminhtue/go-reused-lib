package jwtnats

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

func GetAccountPublicKey() string {
	// return viper.GetString("nats.account.public_key")
	return "AADJYH5MPADBUXFF6NROB3VDMANPTGBOMZJ474VYHFO4N7UANDHDZQ5P"
}

func GetAccountSigningKey(accountSigningSeedKey string) nkeys.KeyPair {
	// Content of the account signing key seed can come from a file or an environment variable as well

	accSeed := []byte(accountSigningSeedKey)
	// accSeed := []byte("SAAPZPXHT354ODAL7LSFFVXW4KNZTAAFDU6G43I2P4QF6FH7X5KDH47K44")
	accountSigningKey, err := nkeys.ParseDecoratedNKey(accSeed)
	if err != nil {
		panic(err)
	}
	return accountSigningKey
}

// func loadNKeyFromSeed(seed string) (nkeys.KeyPair, error) {
// 	// return nkeys.FromSeed([]byte(seed))
// 	return nkeys.ParseDecoratedNKey([]byte(seed))
// }

func ConnectUserDemo() {
	// Setup! Obtain the account signing key!
	accountPublicKey := GetAccountPublicKey()
	accountSigningKey := GetAccountSigningKey("SAAPZPXHT354ODAL7LSFFVXW4KNZTAAFDU6G43I2P4QF6FH7X5KDH47K44")
	// userPublicKey, userSeed, userKeyPair := generateUserKey() // userSeed is not used in this example
	userPublicKey, userNkeySeed, userKeyPair := generateUserKey() // userSeed is not used in this example
	// userPublicKey, _, userKeyPair := generateUserKey()

	// print userNkeySeed
	log.Println("userNkeySeed: ", string(userNkeySeed))
	userClaim := UserClaim{
		PubAllows: []string{"subject.foo"}, // only allow publishing to subject.foo
		SubAllows: []string{"subject.bar"}, // only allow subscribing to subject.bar
		Expires:   time.Now().Add(time.Hour).Unix(),
	}
	userJWT, err := generateUserJWT(userClaim, userPublicKey, accountPublicKey, accountSigningKey)
	if err != nil {
		panic(err)
	}
	// connect with System NKey to update user claims
	// adminJwt := viper.GetString("nats.admin.jwt")
	// adminSeed := viper.GetString("nats.admin.seed")
	// natsAdminOption := nats.UserJWTAndSeed(adminJwt, adminSeed)
	// ncSys, err := nats.Connect("127.0.0.1", natsAdminOption)

	// This code wont work: jwt update resulted in error - n/a - not account claim (only account claim is needed)
	ncSys, err := nats.Connect("127.0.0.1", nats.UserCredentials("admin.creds"))
	if err != nil {
		panic(err)
	}
	msg, err := ncSys.Request("$SYS.REQ.CLAIMS.UPDATE", []byte(userJWT), time.Second)
	if err != nil {
		if err.Error() != "nats: no responders available for request" {
			panic(err)
		}
	}
	log.Println("Updated user claims: ", string(msg.Data))

	// userJWT and userKeyPair can be used in conjunction with this nats.Option
	jwtAuthOption := nats.UserJWT(func() (string, error) {
		return userJWT, nil
	},
		func(bytes []byte) ([]byte, error) {
			return userKeyPair.Sign(bytes)
		},
	)

	// Alternatively you can create a creds file and use it as nats.Option
	// credsContent, err := jwt.FormatUserConfig(userJWT, userNkeySeed)
	// if err != nil {
	// 	panic(err)
	// }
	// os.WriteFile("my.creds", credsContent, 0644)
	// jwtAuthOption := nats.UserCredentials("my.creds")

	// use in a connection as desired
	nc, err := nats.Connect("nats://localhost:4222", jwtAuthOption)
	if err != nil {
		panic(err)
	}
	if nc.IsConnected() {
		defer nc.Close()
		log.Println("Connected to NATS")
	}
	// try to publish subject.foo
	err = nc.Publish("subject.foo", []byte("Hello World!"))
	if err != nil { // NOTE: This will fail because the user is not allowed to publish to subject.foo (but server will not return an error)
		log.Println("Failed to publish message")
	}
}
