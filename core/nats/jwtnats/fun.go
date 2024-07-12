package jwtnats

import (
	"github.com/nats-io/jwt"
	"github.com/nats-io/nkeys"
)

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
	userKeyPair = kp
	return
}

func generateUserJWT(userClaim UserClaim, userPublicKey string, accountPublicKey string, accountSigningKey nkeys.KeyPair) (userJWT string, err error) {
	uc := jwt.NewUserClaims(userPublicKey)
	for _, pubAllow := range userClaim.PubAllows {
		uc.Pub.Allow.Add(pubAllow)
	}
	for _, subAllow := range userClaim.SubAllows {
		uc.Sub.Allow.Add(subAllow)
	}
	uc.Expires = userClaim.Expires
	// uc.Pub.Allow.Add("subject.foo") // only allow publishing to subject.foo
	// uc.Expires = time.Now().Add(time.Hour).Unix() // expire in an hour
	uc.IssuerAccount = accountPublicKey
	vr := jwt.ValidationResults{}
	uc.Validate(&vr)
	if vr.IsBlocking(true) {
		panic("Generated user claim is invalid")
	}
	userJWT, err = uc.Encode(accountSigningKey)
	if err != nil {
		return "", err
	}
	return
}
