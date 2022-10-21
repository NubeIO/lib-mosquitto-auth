package mosquitto

import (
	"fmt"
	"github.com/NubeIO/lib-mosquitto-auth/utils"
	"os"
)

var acControlFilePath string = "/etc/mosquitto/ca_certificates/aclfile"
var caCertificateFilePath string = "/etc/mosquitto/ca_certificates/cacertfile.crt"
var passFilePath string = "/etc/mosquitto/ca_certificates/passfile"
var revocationListFilePath string = "/etc/mosquitto/ca_certificates/serverrevlistfile"
var serverKeyFile string = "/etc/mosquitto/certs/key.pem"
var serverCertificateFile string = "/etc/mosquitto/certs/cert.pem"

type Security struct {
	SSL                bool   `json:"ssl"`
	ClientVerification bool   `json:"clientVerification"`
	Password           string `json:"password"`
}

type AccessControl struct {
	Anonymous bool `json:"anonymous"`
}

type Config struct {
	Path          string        `json:"path"`
	Persistence   bool          `json:"persistence"`
	Security      Security      `json:"security"`
	AccessControl AccessControl `json:"accessControl"`
}

func WriteConfig(c *Config) error {
	file := "/etc/mosquitto/mosquitto.conf"
	if c.Path != "" {
		file = c.Path
	}
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	if c.Security.SSL {
		fmt.Fprintln(f, "listener 8883")
		fmt.Fprintln(f, "cafile "+caCertificateFilePath)
		fmt.Fprintln(f, "certfile "+serverCertificateFile)
		fmt.Fprintln(f, "keyfile  "+serverKeyFile)
		if c.Security.ClientVerification {
			fmt.Fprintln(f, "require_certificate true")

			if utils.Exists(revocationListFilePath) {
				fmt.Fprintln(f, "crlfile "+revocationListFilePath)
			}
		}

	} else {
		fmt.Fprintln(f, "listener 1883")
	}

	if c.Persistence {
		fmt.Fprintln(f, "persistence true")
	} else {
		fmt.Fprintln(f, "persistence false")
	}

	if !c.AccessControl.Anonymous {
		fmt.Fprintln(f, "allow_anonymous false")
	} else {
		fmt.Fprintln(f, "allow_anonymous true")
	}
	if utils.Exists(passFilePath) {
		fmt.Fprintln(f, "password_file "+passFilePath)
	}
	if utils.Exists(acControlFilePath) {
		fmt.Fprintln(f, "acl_file "+acControlFilePath)
	}
	return nil
}
