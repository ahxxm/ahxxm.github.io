package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"

	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"cloud.google.com/go/logging"
	"google.golang.org/api/option"
)

const cred = `lPzK7WpPdCEdhW1dUbH5SIXeRC_gZNz7JSMeiUj4jLAzFMur7iHuf-nbyv1beG1JSpLzD9MAhYDEzRuBWAbRP9zEuqV9SdOOSx8WxtLkwCqw4LWO5P7yutaJ8BAWDr8wzKCoHqi3rCqzmhb6jj84oFcmUb_Rsj7yrxO7UTjr7A4Cn6Uuqx3xwxzB9TZBPYHKH9d6o7dPJ4p33iXmBtoqwjuRcFQ8Cwc98D2_5ec0wAOWl4vahlVhKz6_G96W-3FiC-lVogdtGygqjOxDI3W-b602P6M8S1F7K3nEJHWRKJLYJudyNPeDcYH9xrBdzXFFK8sFvG7IKHYiseXN25lWVG9A7wBdmTmWazr_OU2AksZkVryTUuFca9Egva_igH_8mPeyw5n8sXGTydJ0mPYkLvwUaILZjmIqf-EQkhUurSii45c6RW95JFR-YXh2wzlFVYDGoCepIXl6khZ0WRWgNb9GDSAxxStqo0_p207eaqSKxWSAYUWuWjkQ-mhNVpWLhLiFvgf-u4VsxCMvmUgz-3ME7dvqPJ47l_5I1UF8ef1YBJtK1zAqJd4HEnZhI8csh_vTZHPQ8HEqW0-ZeVsedb0GAJhhOjcb-KDpgzbEw3HPOUGLJyaCtViHkfRzqRllaFY2RYt55v5jrFaESokZ54Ik-e6lpb_G78ncJ1lrree-_OWxO1WUofp-sS7w22__J41LMLpxmfuc7jA6M_H_WtGkmtiwWre4d2quCfmVnDpxg-k2yJnXx9XDpYREKoYpyPhWVfkZuUGlet7yBU0U0WDPu7jAg7lu6OcoddshlnCFDd2w5JDWh1QDQt8IaZjAcpv9S-PtTtA6rMKpMGcHoA30rj-37BUcskiCcAa1YaUW8AGI83ZO3ohsXLbv3tYru1MPJpZiWk3q1evw8EpLW75hNmMNAHNOVy6mYndL2noqOdYdA2kHPctNFkTJGx5eDStLXmjaQvq0JCxVHgwAAzs-lPhc57icCU1cXDlXScKCyJzhRVbHmJFD2M6BODAEZn_lnYjMOiY9hnMcvv_nzQW0EUBWrZlevsxTD2XQGeZdhtlwgHCzqSYZHI6UHwmhBw4VDSsCX9fOmWtz-qNhAKEe2893_P6n5ZeU3a1TdstSPZ9NM56aTc5_kNKk1qM8B1kmqtpRYepccn-NrgW2J5uAycbr1ynJltYLHAhc25b_rsoSDTb3ooyl2WlChc0BeEsCcG8NpRDkHQFJfvaIz31BT8iPSMpqUFJ-Co_Tk1iBG1dPBkV4s1aYkDXD2qwzbPCcJxG1k-eouURQn9kFByQit8gBLiYxMCRtYg0UkroZ9KjOTQt_E2EkKcQPrLj9SsWZ8Oc7Zm1A42PGkiIQl4EmtwlsN6T5NW5ub_aGvLsvYvtXJZ_DB7X_nohXuWoaznjRxLEE_6P6fTUcZO96HLWv-JX_1ByDWOSiwxlGCfXwe9sgUzQipqqQ7ZMFF2Tyexmg9yXi2DDO8f_tPqxm7ecxitc8JJ_2WuqQFtZ4EBI56AInwaR_x-IwMOqZ3MHcLc48mTxeNSP_H0qh9WjwXAOu8Rji1zfenwYymZXAq6l20-A4t19fFfC1v9plbVX101rcOAti1DYm2xsGsmQ0RwCnlTkzSm9gi-28iENeVArpDxhHL-SIQHkMmao9kSdPR-iLjeMD4Bk6OaBFrCz_5DDni_lepR6NQs7hOlAYYCWBtA_iFK8tXDBLDwvmPtWX9aXAH_F5eApTYKj9_7-lhU1i4qot2sOycWS-Xz5QSqDAD-OMCJxaL4zbvR4Esukby5UwP1j5BPDmeusxczPlQ1Z4acVMEGt7XYJvX-l6pozjoHHutUgVfeiX-J5V6RuJZ2g-UZ9tiHS-kyhqPLYCEIUFSTWp7PsqP3ma8cEM7buHHKQQp29lrkOFSGKBivRO54y_nxsGBZrZ26mlK9HH55HvQBYn3Uif968Oqzn6M-oTvg3AezxLTGt-dr6yKZAi3Xe0dK56MOrogM5w8ACdmi02Xx5ERzhN04eHsPckXmPAEn7kQXPplvWUdmdZfN29LP6rvR4TQaD-qS2DFSPsfCSVY5oKPQnPnwS1vFDViZaCRfPPhgyKAFjBuENHTr7lO7IHQNT1pR8oQVJ80CdWrUw1zOoJs6zuT1opMknAi2qUkR3_yjqh_viPaFiJnfRpSGx8HQXiwb0Ky_F5lY7-28ov06tbOMrF7lyGt66SR6bCQKVzV5PPxGvxoQMr-slW8oQct53m3XNyPSNYYA9sDD1fkX76mPCrN7B0wPle0s0pk0WQsz1PKZWYFPBfeXkugfhzu2PhL2qhXMDX8wApML9ALOzBbME7w203qlsLyZ6exze-SYh56RbwOLWOd0q7Zgq-QHGyUL1ar4sRJkr127DjTAc0P3aViJ823UBi5FHEWKm3JMe3Rom_PoUxdwdDlxMsqhE7jGszT-yaAAds9fpDcEvg1SFGqrmjiEcVrsM5h27AcL7s7z307nb-wOwAz8osfRkWnqTd5uUN1cNltwDIzgCtde1zXGFquS47Qk0FU2OakHccuv0dB5GyXBBYGXuDJPypbExeqWzCQA5TSEvhi8InHsvfVrVAuqJBsVmJKIAeajR4LuNkkt4-3hF_FB-BqgBR0qJR__hnFTX2w8yKwgXVbrzB7smINVFFuwipOTjBRa_jCAD-jX9iQC4cE8L2UpF1um5_3ID_l92MDw16B8E0UCDcPYwOJDPkp2kuNqe7vL2dKWerTP284drpLz3Wre8mZftaYB9b3U8flg_Q_jnjoVUZgMIn-BznA5Wlb0BjnPKhroYdhi3jvpuuleQvrX7gPa-LF3epoK4u1PdYJnvH66OiQ3u7TjkotADjZRL2CCDJuLYXewNvWxQiLsverSkJ92qBfGvdrbDwvSXdtB5YBwHhZqWVgM_LzVDxf1pKh6dszF3BrXv6JEFLguWe6HZHkRV-b5WjyBgALzln9xQ7Go7YREN1cmzQupYofTe1HlVLwLSaH-w_VkBU5jy9dTWGM-Qp3Hq7DiVjY79cg1iNRimz1a-Af3XewAOckswbHIJhQ5G0Cw7RDKDV6iYz`

var (
	hostname = ""
	credjson = ""
	projID   = "ahxxmblog"
	logger   *logging.Logger

	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	aesgcm    cipher.AEAD
	nonce     = []byte{100, 169, 67, 62, 174, 124, 204, 238, 226, 252, 14, 218}
	decodekey = "todo-32-byte-key-here"
)

// modified from: https://golang.org/src/crypto/cipher/example_test.go
func encrypt(text string) string {
	plaintext := []byte(text)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func decode(text string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(text)
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}

func init() {
	// GCM setup
	block, err := aes.NewCipher([]byte(decodekey))
	if err != nil {
		panic(err.Error())
	}
	a, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	aesgcm = a

	// Hostname
	if h, err := os.Hostname(); err != nil {
		hostname = h
	}

	// Logger
	ctx := context.Background()
	credjson = decode(cred)
	keyopt := option.WithCredentialsJSON([]byte(credjson))
	client, err := logging.NewClient(ctx, projID, keyopt)
	if err != nil {
		log.Fatalf("Failed to create logging client: %v", err)
	}
	client.OnError = func(err error) {
		log.Printf("client.OnError: %v", err)
	}
	logger = client.Logger(projID)
}

func writeEntry(data interface{}) {
	// Ensure the entry is written.
	defer logger.Flush()
	logger.Log(logging.Entry{
		Payload:  data,
		Severity: logging.Info,
	})
}

func handle(ctx context.Context, e events.APIGatewayProxyRequest) error {
	// TODO: select keys from e
	writeEntry(e)
	return nil
}

func main() {
	// to encrypt, `fmt.Println(encrypt(somejsonliteral))` here
	log.Printf("Logger started from host %v", hostname)
	lambda.Start(handle)
}
