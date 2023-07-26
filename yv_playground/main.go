package main

import (
	"encoding/base64"
	"fmt"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/mac"
	"log"
)

func main() {
	kh, err := keyset.NewHandle(mac.HMACSHA256Tag256KeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	// TODO: save the keyset to a safe location. DO NOT hardcode it in source code.
	// Consider encrypting it with a remote key in Cloud KMS, AWS KMS or HashiCorp Vault.
	// See https://github.com/google/tink/blob/master/docs/GOLANG-HOWTO.md#storing-and-loading-existing-keysets.
	m, err := mac.New(kh)
	if err != nil {
		log.Fatal(err)
	}

	msg := []byte("This is the data being encrypted")
	tag, err := m.ComputeMAC(msg)
	if err != nil {
		log.Fatal(err)
	}

	if m.VerifyMAC(tag, msg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Key %s\n,", kh.String())
	fmt.Printf("Message %s\n", msg)
	fmt.Printf("Authentication Tag: %s\n", base64.StdEncoding.EncodeToString(tag))
}
