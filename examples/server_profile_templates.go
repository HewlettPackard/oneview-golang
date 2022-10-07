package main

import (
	"fmt"

	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
)

func main() {

	config, config_err := ov.LoadConfigFile("config.json")
	if config_err != nil {
		fmt.Println(config_err)
	}
	var (
		ClientOV                          *ov.OVClient
		server_profile_template_name      = "Test SPT"
		server_profile_template_name_auto = config.ServerProfileTemplateConfig.ServerPrpofileTemplateName
		enclosure_group_name              = config.ServerProfileTemplateConfig.EnclosureGroupName
		server_hardware_type_name         = config.ServerProfileTemplateConfig.ServerHardwareTypeName
		scope                             = "Auto-Scope"
	)

	ovc := ClientOV.NewOVClient(
		config.OVCred.UserName,
		config.OVCred.Password,
		config.OVCred.Domain,
		config.OVCred.Endpoint,
		config.OVCred.SslVerify,
		config.OVCred.ApiVersion,
		config.OVCred.IfMatch)

	server_hardware_type, err := ovc.GetServerHardwareTypeByName(server_hardware_type_name)

	enc_grp, err := ovc.GetEnclosureGroupByName(enclosure_group_name)

	conn_settings := ov.ConnectionSettings{

		ManageConnections: true,
	}

	initialScopeUris := new([]utils.Nstring)
	scp, scperr := ovc.GetScopeByName(scope)

	if scperr != nil {
		*initialScopeUris = append(*initialScopeUris, scp.URI)
	}

	aa := ov.AdministratorAccount{
		DeleteAdministratorAccount: utils.GetBoolPointer(false),
		Password:                   "password123",
	}

	la := []ov.LocalAccounts{}
	la = append(la, ov.LocalAccounts{
		UserName:                 "Test",
		DisplayName:              "test",
		Password:                 "passoriutuytguytuytuytuytd",
		UserConfigPriv:           utils.GetBoolPointer(true),
		RemoteConsolePriv:        utils.GetBoolPointer(true),
		VirtualMediaPriv:         utils.GetBoolPointer(true),
		VirtualPowerAndResetPriv: utils.GetBoolPointer(true),
		ILOConfigPriv:            utils.GetBoolPointer(true),
	})

	duc := []string{"OU=US,OU=Users,OU=Accounts,dc=Subdomain,dc=example,dc=com",
		"ou=People,o=example.com"}

	d := ov.Directory{
		DirectoryAuthentication:    "defaultSchema",
		DirectoryGenericLDAP:       utils.GetBoolPointer(false),
		DirectoryServerAddress:     "ldap.example.com",
		DirectoryServerPort:        636,
		DirectoryServerCertificate: "-----BEGIN CERTIFICATE-----\nMIIBozCCAQwCCQCWGqL41Y6YKTANBgkqhkiG9w0BAQUFADAWMRQwEgYDVQQDEwtD\nb21tb24gTmFtZTAeFw0xNzA3MTQxOTQzMjZaFw0xODA3MTQxOTQzMjZaMBYxFDAS\nBgNVBAMTC0NvbW1vbiBOYW1lMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCf\nCNTrU4AZF044Rtu8jiGR6Ce1u9K6GJE+60VCau2y4A2z5B5kKA2XnyP+2JpLRRA8\n8PjEyVJuL1fJomGF74L305j6ucetXZGcEy26XNyKFOtsBeoHtjkISYNTMxikjvC1\nXHctTYds0D6Q6u7igkN9ew8ngn61LInFqb6dLm+CmQIDAQABMA0GCSqGSIb3DQEB\nBQUAA4GBAFVOQ8zXFNHdXVa045onbkx8pgM2zK5VQ69YFtlAymFDWaS7a5M+96JW\n2c3001GDGZcW6fGqW+PEyu3COImRlEhEHaZKs511I7RfckMzZ3s7wPrQrC8WQLqI\ntiZtCWfUX7tto7YDdmfol7bHiaSbrLUv4H/B7iS9FGemA+nrghCK\n-----END CERTIFICATE-----",
		DirectoryUserContext:       duc,
		IloObjectDistinguishedName: "service",
		Password:                   "kjhkjhkjhkjh0",
		KerberosAuthentication:     utils.GetBoolPointer(false),
	}

	dg := []ov.DirectoryGroups{}
	dg = append(dg, ov.DirectoryGroups{
		GroupDN:                  "ilos.example.com,ou=Groups,o=example.com",
		GroupSID:                 "S-1-5-11",
		UserConfigPriv:           utils.GetBoolPointer(false),
		RemoteConsolePriv:        utils.GetBoolPointer(true),
		VirtualMediaPriv:         utils.GetBoolPointer(true),
		VirtualPowerAndResetPriv: utils.GetBoolPointer(true),
		ILOConfigPriv:            utils.GetBoolPointer(false),
	})

	km := ov.KeyManager{
		PrimaryServerAddress:   "192.0.2.91",
		PrimaryServerPort:      9000,
		SecondaryServerAddress: "192.0.2.92",
		SecondaryServerPort:    9000,
		RedundancyRequired:     utils.GetBoolPointer(true),
		GroupName:              "GRP",
		CertificateName:        "Local CA",
		LoginName:              "deployment",
		Password:               "Passw0rd",
	}

	host := ov.IloHostName{
		HostName: "{serverProfileName}",
	}
	mps := ov.MpSettings{
		AdministratorAccount: aa,
		LocalAccounts:        la,
		Directory:            d,
		DirectoryGroups:      dg,
		KeyManager:           km,
		IloHostName:          host,
	}

	mp := ov.ManagementProcessors{
		ManageMp:  true,
		MpSetting: mps,
	}

	server_profile_template := ov.ServerProfile{
		Type:                  "ServerProfileTemplateV8",
		Name:                  server_profile_template_name,
		EnclosureGroupURI:     enc_grp.URI,
		ServerHardwareTypeURI: server_hardware_type.URI,
		ConnectionSettings:    conn_settings,
		InitialScopeUris:      *initialScopeUris,
		ManagementProcessors:  mp,
	}

	server_profile_template_auto := ov.ServerProfile{
		Type:                  "ServerProfileTemplateV8",
		Name:                  server_profile_template_name_auto,
		EnclosureGroupURI:     enc_grp.URI,
		ServerHardwareTypeURI: server_hardware_type.URI,
		ConnectionSettings:    conn_settings,
		InitialScopeUris:      *initialScopeUris,
	}

	err = ovc.CreateProfileTemplate(server_profile_template)
	err = ovc.CreateProfileTemplate(server_profile_template_auto)
	if err != nil {
		fmt.Println("Server Profile Template Creation Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template Created---------------#")
	}
	sort := "name:asc"
	spt_list, err := ovc.GetProfileTemplates("", "", "", sort, "")
	if err != nil {
		fmt.Println("Server Profile Template Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template List---------------#")

		for i := 0; i < len(spt_list.Members); i++ {
			fmt.Println(spt_list.Members[i].Name)
		}
	}

	spt, err := ovc.GetProfileTemplateByName(server_profile_template_name)
	if err != nil {
		fmt.Println("Server Profile Template Retrieval By Name Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template by Name---------------#")
		fmt.Println(spt.Name)
	}

	fmt.Println("Server Profile Template refresh using PATCH request")
	options := new([]ov.Options)
	*options = append(*options, ov.Options{"replace", "/refreshState", "RefreshPending"})

	err = ovc.PatchServerProfileTemplate(spt, *options) //patchRequest)
	if err != nil {
		fmt.Println("Refresh failed", err)
	}

	spt.Name = "Renamed Test SPT"
	err = ovc.UpdateProfileTemplate(spt)
	if err != nil {
		fmt.Println("Server Profile Template Updation Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template Updated---------------#")
	}

	spt_list, err = ovc.GetProfileTemplates("", "", "", sort, "")
	if err != nil {
		fmt.Println("Server Profile Template Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template List---------------#")

		for i := 0; i < len(spt_list.Members); i++ {
			fmt.Println(spt_list.Members[i].Name)
		}
	}

	err = ovc.DeleteProfileTemplate(spt.Name)
	if err != nil {
		fmt.Println("Server Profile Template Delete Failed: ", err)
	} else {
		fmt.Println("#----------------Server Profile Template Deleted---------------#")
	}
}
