package jwtnats

import (
	"time"

	"github.com/nats-io/jwt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

func GetAccountPublicKey() string {
	return "random_account_public_key"
}

func GetAccountSigningKey() nkeys.KeyPair {
	// Content of the account signing key seed can come from a file or an environment variable as well
	accSeed := []byte("SAAJGCAHPHHM6AVJJWQ2YAS3I4NETXMWVQSTCQMJ7VVTGAJF5UCN3IX7J4")
	accountSigningKey, err := nkeys.ParseDecoratedNKey(accSeed)
	if err != nil {
		panic(err)
	}
	return accountSigningKey
}

func RequestUser() {
	// Setup! Obtain the account signing key!
	accountPublicKey := GetAccountPublicKey()
	accountSigningKey := GetAccountSigningKey()
	// userPublicKey, userSeed, userKeyPair := generateUserKey() // userSeed is not used in this example
	userPublicKey, _, userKeyPair := generateUserKey()

	userJWT := generateUserJWT(userPublicKey, accountPublicKey, accountSigningKey)
	// userJWT and userKeyPair can be used in conjunction with this nats.Option
	var jwtAuthOption nats.Option
	jwtAuthOption = nats.UserJWT(func() (string, error) {
		return userJWT, nil
	},
		func(bytes []byte) ([]byte, error) {
			return userKeyPair.Sign(bytes)
		},
	)

	// Alternatively you can create a creds file and use it as nats.Option
	// credsContent, err := jwt.FormatUserConfig(userJWT, userSeed)
	// if err != nil {
	// 	panic(err)
	// }
	// os.WriteFile("my.creds", credsContent, 0644)
	// jwtAuthOption = nats.UserCredentials("my.creds")

	// use in a connection as desired
	nc, err := nats.Connect("nats://localhost:4222", jwtAuthOption)
	// ...
}

func generateUserKey() (userPublicKey string, userSeed []byte, userKeyPair nkeys.KeyPair) {
	kp, err := nkeys.CreateUser()
	if err != nil {
		return "", nil, nil
	}
	if userSeed, err = kp.Seed(); err != nil {
		return "", nil, nil
	} else if userPublicKey, err = kp.PublicKey(); err != nil {
		return "", nil, nil
	}
	return
}

func generateUserJWT(userPublicKey, accountPublicKey string, accountSigningKey nkeys.KeyPair) (userJWT string) {
	uc := jwt.NewUserClaims(userPublicKey)
	uc.Pub.Allow.Add("subject.foo")               // only allow publishing to subject.foo
	uc.Expires = time.Now().Add(time.Hour).Unix() // expire in an hour
	uc.IssuerAccount = accountPublicKey
	vr := jwt.ValidationResults{}
	uc.Validate(&vr)
	if vr.IsBlocking(true) {
		panic("Generated user claim is invalid")
	}
	var err error
	userJWT, err = uc.Encode(accountSigningKey)
	if err != nil {
		return ""
	}
	return
}
