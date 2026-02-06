package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

const (
	keyLen    = 18
	secretLen = 36
	pad       = '+'

	devVersion = "dev"

	about = `minio-keygen (%s, %s)
Generates a MinIO access key and secret. You can redirect the output into an .env file.
`
)

func main() {
	flag.Usage = usage
	flag.Parse()
	enc := base64.URLEncoding.WithPadding(pad)
	accessKey, err := generateKey(enc, keyLen, true)
	if err != nil {
		exitWithError("failed to generate access key", err)
	}
	secretKey, err := generateKey(enc, secretLen, false)
	if err != nil {
		exitWithError("failed to generate secret key", err)
	}

	fmt.Printf(
		"MINIO_ACCESS_KEY=%s\nMINIO_SECRET_KEY=%s\n",
		accessKey,
		secretKey,
	)
}

func usage() {
	var (
		info, ok = debug.ReadBuildInfo()
		ver      = devVersion
	)
	if ok {
		ver = info.Main.Version
	}
	fmt.Printf(about, ver, runtime.Version())
}

func generateKey(encoder *base64.Encoding, length int, upper bool) (string, error) {
	data := make([]byte, length)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}

	encoded := encoder.EncodeToString(data)
	if upper {
		encoded = strings.ToUpper(encoded)
	}

	return encoded, nil
}

func exitWithError(message string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
	flag.Usage()
	os.Exit(1)
}
