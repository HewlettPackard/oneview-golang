package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
)

func main() {
	var (
		ClientOV                *ov.OVClient
		server_certificate_ip                 = "172.18.13.11"
		server_certificate_name               = "new_test_certificate"
		new_cert_base64data     utils.Nstring = "---BEGIN CERTIFICATE----END CERTIFICATE------"
	)
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1600,
		"")

	server_cert, err := ovc.GetServerCertificateByIp(server_certificate_ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(server_cert)
	}

	server_cert.CertificateDetails[0].AliasName = server_certificate_name
	server_cert.Type = ""     // The type field in certificate is not required in POST call, so making it empty
	fmt.Println(server_cert.CertificateDetails[0].AliasName)

	er := ovc.CreateServerCertificate(server_cert)
	if er != nil {
		fmt.Println("............... Adding Server Certificate Failed:", er)
	} else {
		fmt.Println(".... Adding Server Certificate Success")
	}
	fmt.Println("#................... Server Certificate by Name ...............#")
	server_certn, err := ovc.GetServerCertificateByName(server_certificate_name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(server_certn)
	}

	certificateDetails := new([]ov.CertificateDetail)
	certificateDetail_new := ov.CertificateDetail{Type: "CertificateDetailV2", AliasName: server_certificate_name, Base64Data: new_cert_base64data}
	*certificateDetails = append(*certificateDetails, certificateDetail_new)
	server_certn.CertificateDetails = *certificateDetails
	err = ovc.UpdateServerCertificate(server_certn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#.................... Server Certificate after Updating ...........#")
		server_cert_after_update, err := ovc.GetServerCertificateByName(server_certificate_name)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("..............Server Certificate Successfully updated.........")
			fmt.Println(server_cert_after_update)
		}
	}

	err = ovc.DeleteServerCertificate(server_certificate_name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#...................... Deleted Server Certificate Successfully .....#")
	}

}
