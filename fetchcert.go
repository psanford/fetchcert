package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

var asPem = flag.Bool("pem", false, "Print cert(s) in PEM format (default is to print cert details)")

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalf("usage: %s <host|host:port|http[s]://url>", os.Args[0])
	}

	dst := args[0]

	if strings.HasPrefix(dst, "http://") || strings.HasPrefix(dst, "https://") {
		u, err := url.Parse(dst)
		if err != nil {
			log.Fatalf("Failed to parse url: %s", err)
		}
		dst = u.Host
	}

	conf := tls.Config{
		InsecureSkipVerify: true,
	}
	conf.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
		for _, cert := range rawCerts {
			pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE", Bytes: cert})
		}

		var certs []*x509.Certificate
		for i := range rawCerts {
			fmt.Println("========================================")
			cert, err := x509.ParseCertificate(rawCerts[i])
			if err != nil {
				log.Fatalf("Parse cert err: %s", err)
			}
			certs = append(certs, cert)

			fmt.Printf("Serial Number : %s\n", cert.SerialNumber)
			fmt.Printf("Subject: %s\n", cert.Subject)
			fmt.Printf("Issuer: %s\n", cert.Issuer)
			fmt.Printf("Not Before: %s\n", cert.NotBefore)
			fmt.Printf("Not After: %s\n", cert.NotAfter)
			fmt.Printf("Subject Alt Names: %s\n", strings.Join(cert.DNSNames, ";"))
		}
		fmt.Println("========================================")

		opts := x509.VerifyOptions{
			Roots:         conf.RootCAs,
			CurrentTime:   time.Now(),
			DNSName:       conf.ServerName,
			Intermediates: x509.NewCertPool(),
		}
		for _, cert := range certs[1:] {
			opts.Intermediates.AddCert(cert)
		}
		var err error
		verifiedChains, err = certs[0].Verify(opts)
		if err != nil {
			fmt.Printf("Untrusted certificate: %s\n", err)
			return err
		}

		return nil
	}

	if strings.Index(dst, ":") < 0 {
		dst += ":443"
	}

	fmt.Printf("Dail %s\n", dst)
	_, err := tls.Dial("tcp", dst, &conf)
	if err != nil {
		log.Fatalf("Dail err: %s", err)
	}
}
