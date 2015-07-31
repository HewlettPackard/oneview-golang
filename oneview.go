package oneview

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/codegangsta/cli"
	"github.com/docker/machine/drivers"
	"github.com/docker/machine/drivers/oneview/ov"
	"github.com/docker/machine/log"
	"github.com/docker/machine/ssh"
	"github.com/docker/machine/state"
)

// Driver OneView driver structure
type Driver struct {
	*drivers.BaseDriver
	SSHUser        string
	SSHPort        int
	SSHPublicKey   string
	ClientICSP     *ov.ICSPClient
	ClientOV       *ov.OVClient
	ServerTemplate string
	OSBuildPlan    string
}

const (
	defaultTimeout = 1 * time.Second
)

func init() {
	drivers.Register("oneview", &drivers.RegisteredDriver{
		New:            NewDriver,
		GetCreateFlags: GetCreateFlags,
	})
}

// GetCreateFlags registers the flags this driver adds to
// "docker hosts create"
// --oneview-ov-user        : String User to OneView
// --oneview-ov-password    : String Password to OneView
// --oneview-ov-endpoint    : String url end point, base path
//
// --oneview-icsp-user      : String User to ICSP
// --oneview-icsp-password  : String Password to ICSP
// --oneview-icsp-endpoint  : String url end point, base path
//
// --oneview-sslverify      : Bool false means no https verification
// --oneview-apiversion     : Int version of api to use 120 is default
//
func GetCreateFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "oneview-ov-user",
			Usage:  "User Name to OneView Server",
			Value:  "",
			EnvVar: "ONEVIEW_OV_USER",
		},
		cli.StringFlag{
			Name:   "oneview-ov-password",
			Usage:  "Password to OneView Server",
			Value:  "",
			EnvVar: "ONEVIEW_OV_PASSWORD",
		},
		cli.StringFlag{
			Name:   "oneview-ov-domain",
			Usage:  "Domain to OneView Server",
			Value:  "LOCAL",
			EnvVar: "ONEVIEW_OV_DOMAIN",
		},
		cli.StringFlag{
			Name:   "oneview-ov-endpoint",
			Usage:  "OneView Server URL Endpoint",
			Value:  "",
			EnvVar: "ONEVIEW_OV_ENDPOINT",
		},
		cli.StringFlag{
			Name:   "oneview-icsp-user",
			Usage:  "User Name to OneView Insight Controller",
			Value:  "",
			EnvVar: "ONEVIEW_ICSP_USER",
		},
		cli.StringFlag{
			Name:   "oneview-icsp-password",
			Usage:  "Password to OneView Insight Controller",
			Value:  "",
			EnvVar: "ONEVIEW_ICSP_PASSWORD",
		},
		cli.StringFlag{
			Name:   "oneview-icsp-domain",
			Usage:  "Domain to OneView Insight Controller",
			Value:  "LOCAL",
			EnvVar: "ONEVIEW_ICSP_DOMAIN",
		},
		cli.StringFlag{
			Name:   "oneview-icsp-endpoint",
			Usage:  "OneView Insight Controller URL Endpoint",
			Value:  "",
			EnvVar: "ONEVIEW_ICSP_ENDPOINT",
		},
		cli.BoolFlag{
			Name:   "oneview-sslverify",
			Usage:  "SSH private key path",
			EnvVar: "ONEVIEW_SSLVERIFY",
		},
		cli.IntFlag{
			Name:   "oneview-apiversion",
			Usage:  "OneView API Version",
			Value:  120,
			EnvVar: "ONEVIEW_APIVERSION",
		},
		cli.StringFlag{
			Name:   "oneview-ssh-user",
			Usage:  "OneView build plan ssh user account",
			Value:  "root",
			EnvVar: "ONEVIEW_SSH_USER",
		},
		cli.IntFlag{
			Name:   "oneview-ssh-port",
			Usage:  "OneView build plan ssh host port",
			Value:  22,
			EnvVar: "ONEVIEW_SSH_PORT",
		},
		cli.StringFlag{
			Name:   "oneview-server-template",
			Usage:  "OneView server template to use for blade provisioning, see OneView Server Template for setup.",
			Value:  "DOCKER_1.8_OVTEMP",
			EnvVar: "ONEVIEW_SERVER_TEMPLATE",
		},
		cli.StringFlag{
			Name:   "oneview-os-plan",
			Usage:  "OneView ICSP OS Build plan to use for OS provisioning, see ICS OS Plan for setup.",
			Value:  "RHEL71_DOCKER_1.8",
			EnvVar: "ONEVIEW_OS_PLAN",
		},
	}
}

// NewDriver - create a OneView object driver
func NewDriver(machineName string, storePath string, caCert string, privateKey string) (drivers.Driver, error) {
	inner := drivers.NewBaseDriver(machineName, storePath, caCert, privateKey)
	return &Driver{BaseDriver: inner}, nil
}

// DriverName - get the name of the driver
func (d *Driver) DriverName() string {
	log.Debug("DriverName...")
	return "oneview"
}

// GetSSHHostname - gets the hostname that docker-machine connects to
func (d *Driver) GetSSHHostname() (string, error) {
	log.Debug("GetSSHHostname...")
	return d.GetIP()
}

// GetSSHUsername - gets the ssh user that will be connected to
func (d *Driver) GetSSHUsername() string {
	log.Debug("GetSSHUsername...")
	return d.SSHUser
}

// SetConfigFromFlags - gets the cli configuration flags
func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	log.Debug("SetConfigFromFlags...")
	d.ClientICSP = &ov.ICSPClient{
		ov.Client{
			User:       flags.String("oneview-icsp-user"),
			Password:   flags.String("oneview-icsp-password"),
			Domain:     flags.String("oneview-icsp-domain"),
			Endpoint:   flags.String("oneview-icsp-endpoint"),
			SSLVerify:  flags.Bool("oneview-sslverify"),
			APIVersion: flags.Int("oneview-apiversion"),
			APIKey:     "none",
		},
	}

	d.ClientOV = &ov.OVClient{
		ov.Client{
			User:       flags.String("oneview-ov-user"),
			Password:   flags.String("oneview-ov-password"),
			Domain:     flags.String("oneview-ov-domain"),
			Endpoint:   flags.String("oneview-ov-endpoint"),
			SSLVerify:  flags.Bool("oneview-sslverify"),
			APIVersion: flags.Int("oneview-apiversion"),
			APIKey:     "none",
		},
	}

	d.SSHUser = flags.String("oneview-ssh-user")
	d.SSHPort = flags.Int("oneview-ssh-port")

	d.ServerTemplate = flags.String("oneview-server-template")
	d.OSBuildPlan = flags.String("oneview-os-plan")
	// TODO : we should verify settings for each client
	return nil
}

// PreCreateCheck - pre create check
func (d *Driver) PreCreateCheck() error {
	log.Debug("PreCreateCheck...")
	return nil
}

// Create - create server for docker
func (d *Driver) Create() error {
	log.Infof("Generating SSH keys...")
	if err := d.createKeyPair(); err != nil {
		return fmt.Errorf("unable to create key pair: %s", err)
	}

	log.Debugf("ICSP Endpoint is: %s", d.ClientICSP.Endpoint)
	log.Debugf("OV Endpoint is: %s", d.ClientOV.Endpoint)

	return nil
}

// GetURL - get docker url
func (d *Driver) GetURL() (string, error) {
	log.Debug("GetURL...")
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("tcp://%s:2376", ip), nil
}

// GetIP - get server host or ip address
// TODO: we need to get ip of server from icsp or ov??
func (d *Driver) GetIP() (string, error) {
	log.Debug("GetIP...")
	if d.ClientICSP.Endpoint == "" {
		return "", fmt.Errorf("IP address is not set")
	}
	return d.ClientICSP.Endpoint, nil
}

// GetState - get the running state of the target machine
func (d *Driver) GetState() (state.State, error) {
	log.Debug("GetState...")
	//  addr := fmt.Sprintf("%s:%d", d.IPAddress, d.SSHPort)
	//  _, err := net.DialTimeout("tcp", addr, defaultTimeout)
	//  var st state.State
	//  if err != nil {
	//    st = state.Stopped
	//  } else {
	//    st = state.Running
	//  }
	return state.Stopped, nil
}

// Start - start the docker machine target
func (d *Driver) Start() error {
	log.Debug("Start...")
	return fmt.Errorf("oneview driver does not support start")
}

// Stop - stop the docker machine target
func (d *Driver) Stop() error {
	log.Debug("Stop...")
	return fmt.Errorf("oneview driver does not support stop")
}

// Remove - remove the docker machine target
//    Should remove the ICSP provisioned plan and the Server Profile from OV
func (d *Driver) Remove() error {
	log.Debug("Remove...")
	return nil
}

// Restart - restart the target machine
func (d *Driver) Restart() error {
	log.Debug("Restarting...")

	return nil
}

// Kill - kill the docker machine
func (d *Driver) Kill() error {
	log.Debug("Killing...")

	return nil
}

// publicSSHKeyPath - get the path to public ssh key
func (d *Driver) publicSSHKeyPath() string {
	log.Debug("publicSSHKeyPath...")
	return ""
}

// /////////  HELPLERS /////////////

// createKeyPair - generate key files needed
func (d *Driver) createKeyPair() error {

	if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
		return err
	}

	publicKey, err := ioutil.ReadFile(d.GetSSHKeyPath() + ".pub")
	if err != nil {
		return err
	}

	log.Debugf("created keys => %s", string(publicKey))
	d.SSHPublicKey = string(publicKey)
	return nil
}
