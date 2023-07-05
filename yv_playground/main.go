package main

import (
	"fmt"
	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/keyset"
	"log"
)

const (
	// Change this. AWS KMS, Google Cloud KMS and HashiCorp Vault are supported out of the box.
	keyURI          = "gcp-kms://projects/tink-examples/locations/global/keyRings/foo/cryptoKeys/bar"
	credentialsPath = "credentials.json"
)

func main() {
	kh, err := keyset.NewHandle(aead.AES128CTRHMACSHA256KeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	kh1, err := keyset.NewHandle(aead.AES128GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(kh.String())
	fmt.Println(kh1.String())
}
