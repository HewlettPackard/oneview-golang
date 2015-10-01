package oneview

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/codegangsta/cli"
	"github.com/docker/machine/drivers"
	"github.com/docker/machine/drivers/oneview/icsp"
	"github.com/docker/machine/drivers/oneview/ov"
	"github.com/docker/machine/log"
	"github.com/docker/machine/ssh"
	"github.com/docker/machine/state"
)

// Driver OneView driver structure
type Driver struct {
	*drivers.BaseDriver
	ClientICSP     *icsp.ICSPClient
	ClientOV       *ov.OVClient
	IloUser        string
	IloPassword    string
	IloPort        int
	OSBuildPlan    string
	SSHUser        string
	SSHPort        int
	SSHPublicKey   string
	ServerTemplate string
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
		cli.StringFlag{
			Name:   "oneview-ilo-user",
			Usage:  "ILO User id that is used during ICSP server creation.",
			Value:  "docker",
			EnvVar: "ONEVIEW_ILO_USER",
		},
		cli.StringFlag{
			Name:   "oneview-ilo-password",
			Usage:  "ILO password that is used during ICSP server creation.",
			Value:  "",
			EnvVar: "ONEVIEW_ILO_PASSWORD",
		},
		cli.IntFlag{
			Name:   "oneview-ilo-port",
			Usage:  "optional ILO port to use.",
			Value:  443,
			EnvVar: "ONEVIEW_ILO_PORT",
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

	d.ClientICSP = d.ClientICSP.NewICSPClient(flags.String("oneview-icsp-user"),
		flags.String("oneview-icsp-password"),
		flags.String("oneview-icsp-domain"),
		flags.String("oneview-icsp-endpoint"),
		flags.Bool("oneview-sslverify"),
		flags.Int("oneview-apiversion"))

	d.ClientOV = d.ClientOV.NewOVClient(flags.String("oneview-ov-user"),
		flags.String("oneview-ov-password"),
		flags.String("oneview-ov-domain"),
		flags.String("oneview-ov-endpoint"),
		flags.Bool("oneview-sslverify"),
		flags.Int("oneview-apiversion"))

	d.IloUser = flags.String("oneview-ilo-user")
	d.IloPassword = flags.String("oneview-ilo-password")
	d.IloPort = flags.Int("oneview-ilo-port")

	d.SSHUser = flags.String("oneview-ssh-user")
	d.SSHPort = flags.Int("oneview-ssh-port")

	d.ServerTemplate = flags.String("oneview-server-template")
	d.OSBuildPlan = flags.String("oneview-os-plan")
	// TODO : we should verify settings for each client

	// check for the ov endpoint
	if d.ClientOV.Endpoint == "" {
		return errors.New("Missing option --oneview-ov-endpoint or environment ONEVIEW_OV_ENDPOINT")
	}
	// check for the icsp endpoint
	if d.ClientICSP.Endpoint == "" {
		return errors.New("Missing option --oneview-icsp-endpoint or environment ONEVIEW_ICSP_ENDPOINT")
	}
	// check for the template name
	if d.ServerTemplate == "" {
		return errors.New("Missing option --oneview-server-template or environment ONEVIEW_SERVER_TEMPLATE")
	}
	return nil
}

// PreCreateCheck - pre create check
func (d *Driver) PreCreateCheck() error {
	log.Debug("PreCreateCheck...")
	// TODO: verify you can connect to ov
	// TODO: verify you can connect to icsp
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
	// create the server profile in oneview, we need a hostname and a template name

	// TODO: delete when done
	dJSON, err := json.Marshal(d)
	if err != nil {
		return err
	}
	log.Debugf("Creating machine -> %+v", bytes.NewBuffer(dJSON))

	log.Debugf("***> CreateMachine")
	if err := d.ClientOV.CreateMachine(d.MachineName, d.ServerTemplate); err != nil {
		return err
	}

	log.Debugf("***> GetProfileByName")
	profileMachineName, err := d.ClientOV.GetProfileByName(d.MachineName)
	if err != nil {
		return err
	}

	log.Debugf("***> check GetProfileByName")
	if profileMachineName.URI.IsNil() {
		err := fmt.Errorf("Attempting to get machine profile information, unable to find machine: %s", d.MachineName)
		return err
	}

	// get the server hardware associated with that test profile
	log.Debugf("***> GetServerHardware")
	machineBlade, err := d.ClientOV.GetServerHardware(profileMachineName.ServerHardwareURI)
	if machineBlade.URI.IsNil() {
		err := fmt.Errorf("Attempting to get machine blade information, unable to find machine: %s", d.MachineName)
		return err
	}
	log.Debugf("client 2 *******---> %+v", machineBlade.Client.APIKey)

	// power on the server, and leave it in that state
	log.Debugf("***> PowerOn Server")
	var pt *ov.PowerTask
	pt = pt.NewPowerTask(machineBlade)
	pt.Timeout = 46 // timeout is 20 sec
	err = pt.PowerExecutor(ov.P_ON)
	if err != nil {
		return err
	}

	// TODO: add the server to icsp, TestCreateServer
	// TODO: apply a build plan, TestApplyDeploymentJobs
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
// currently the only way i can see to get this is with sudo ifconfig|grep inet
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
	// TODO: implement power on with oneview, TestPowerState
	// TODO: implement icsp check for is in maintenance mode or started
}

// Stop - stop the docker machine target
func (d *Driver) Stop() error {
	log.Debug("Stop...")
	return fmt.Errorf("oneview driver does not support stop")
	// TODO: call power off from oneview, TestPowerState
}

// Remove - remove the docker machine target
//    Should remove the ICSP provisioned plan and the Server Profile from OV
func (d *Driver) Remove() error {
	log.Debug("Remove...")
	//TODO: remove the ssh keys
	//TODO: destroy the server in icsp
	//TODO: power off the blade, call Stop()
	//TODO: delete the server profile in ov : TestDeleteProfile
	return nil
}

// Restart - restart the target machine
func (d *Driver) Restart() error {
	log.Debug("Restarting...")
	//TODO: should power off / on the server, TestPowerState
	return nil
}

// Kill - kill the docker machine
func (d *Driver) Kill() error {
	log.Debug("Killing...")
	//TODO: implement power off , is there a force?
	return nil
}

// publicSSHKeyPath - get the path to public ssh key
func (d *Driver) publicSSHKeyPath() string {
	log.Debug("publicSSHKeyPath...")
	return d.GetSSHKeyPath() + ".pub"
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
