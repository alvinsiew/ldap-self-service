package web

import (
	"fmt"
	"ldap-self-service/internal/random"
	"ldap-self-service/internal/smtpss"
	"ldap-self-service/internal/yamlcustom"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func execute(u string, op string, np string) error {
	conf := yamlcustom.ParseYAML()
	userDN := conf.Ldap[0].UserDN
	ldapADDR := conf.Ldap[1].LDAP

	cmd := "ldappasswd -H " + ldapADDR + " -x -D cn=" + u + "," + userDN + " -w " + op + " -s " + np
	fmt.Println(cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out)
	fmt.Println(output)

	return err
}

func searchMail(u string) string {
	conf := yamlcustom.ParseYAML()
	userDN := conf.Ldap[0].UserDN
	ldapADDR := conf.Ldap[1].LDAP

	cmd := "ldapsearch -H " + ldapADDR + " -x -b cn=" + u + "," + userDN + " -LLL mail | grep mail | awk '{print $2}'"
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out)
	//fmt.Println(output)

	return output
}

//nolint
func FormHandler(w http.ResponseWriter, r *http.Request) {
	//nolint
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	//nolint
	fmt.Fprintf(w, "POST request successful\n")
	//nolint
	username := r.FormValue("username")

	oldPassword := r.FormValue("oldpassword")
	newPassword := r.FormValue("newpassword")

	result := execute(username, oldPassword, newPassword)
	success := "Successful"

	if result != nil {
		fmt.Fprintf(w, "Status = %s\n", result)
	} else {
		fmt.Fprintf(w, "Status = %s\n", success)
	}
}

//nolint
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	//nolint
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	//nolint
	fmt.Fprintf(w, "POST request successful\n")
	//nolint
	username := r.FormValue("username")

	randomPassword, err := random.GeneratePlainPassword(8)
	if err != nil {
		log.Fatal("Password Generation failed!")
	}

	conf := yamlcustom.ParseYAML()
	smtpUser := conf.Smtp[0].Username
	password := conf.Smtp[1].Password
	hostname := conf.Smtp[2].Hostname
	from := conf.Smtp[3].From
	msg := []byte("Subject: OSG LDAP Password Reset \r\n" +
		"\r\n" +
		"New Random Password: " + randomPassword + "\r\n")
	recipients := strings.Fields(searchMail(username))

	result := smtpss.PlainAuth(smtpUser, password, hostname, from, msg, recipients)

	sucess := "Successful"

	if result != nil {
		fmt.Fprintf(w, "Status = %s\n", result)
	} else {
		fmt.Fprintf(w, "Status = %s\n", sucess)
	}
}
