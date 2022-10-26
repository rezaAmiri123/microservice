package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

const (
	NoClientCert               = "NoClientCert"
	RequestClientCert          = "RequestClientCert"
	RequireAnyClientCert       = "RequireAnyClientCert"
	VerifyClientCertIfGiven    = "VerifyClientCertIfGiven"
	RequireAndVerifyClientCert = "RequireAndVerifyClientCert"
)

func GetClientAuthType(authType string) tls.ClientAuthType {
	switch authType {
	case NoClientCert:
		return tls.NoClientCert
	case RequestClientCert:
		return tls.RequestClientCert
	case RequireAnyClientCert:
		return tls.RequireAnyClientCert
	case VerifyClientCertIfGiven:
		return tls.VerifyClientCertIfGiven
	case RequireAndVerifyClientCert:
		return tls.RequireAndVerifyClientCert
	default:
		return tls.NoClientCert

	}
}

func SetupTLSConfig(cfg TLSConfig) (*tls.Config, error) {
	var err error
	tlsConfig := &tls.Config{}
	if cfg.CertFile != "" && cfg.KeyFile != "" {
		tlsConfig.Certificates = make([]tls.Certificate, 1)
		tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(
			cfg.CertFile,
			cfg.KeyFile,
		)
		if err != nil {
			return nil, err
		}
	}
	if cfg.CAFile != "" {
		b, err := ioutil.ReadFile(cfg.CAFile)
		if err != nil {
			return nil, err
		}
		ca := x509.NewCertPool()
		ok := ca.AppendCertsFromPEM([]byte(b))
		if !ok {
			return nil, fmt.Errorf(
				"failed to parse root certificate: %q",
				cfg.CAFile,
			)
		}
		if cfg.Server {
			tlsConfig.ClientCAs = ca
			tlsConfig.ClientAuth = cfg.ClientAuthType
		} else {
			tlsConfig.RootCAs = ca
		}
		tlsConfig.ServerName = cfg.ServerAddress
	}
	return tlsConfig, nil
}

type TLSConfig struct {
	CertFile       string
	KeyFile        string
	CAFile         string
	ServerAddress  string
	Server         bool
	ClientAuthType tls.ClientAuthType
}
