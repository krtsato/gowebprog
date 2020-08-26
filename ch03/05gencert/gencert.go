package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)

	// 認証局 CA から与えられたと仮定する
	serialNumber, _ := rand.Int(rand.Reader, max)

	// 証明書の識別名
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// 証明書の構成を設定する
	// X.509 とは公開鍵基盤 PKI 用の標準規格
	template := x509.Certificate{
		SerialNumber: serialNumber, // 認証局 CA が発行する一意の番号
		Subject:      subject,
		NotBefore:    time.Now(),                           // 今から
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 1 年間有効
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")}, // 127.0.0.1 の IP だけ効力をもつ
	}

	// RSA 秘密鍵の生成
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// 構造体 Ccertificate と公開鍵，秘密鍵などから byte データのスライスを生成する
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	// 証明書データを符号化して cert.pem ファイルを作成
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// RSA 秘密鍵を PEM 符号化して key.pem を作成
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
