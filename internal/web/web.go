package web

import (
	"fmt"
	"ldap-self-service/internal/smtpss"
	"ldap-self-service/internal/yamlcustom"
	"net/http"
	"os/exec"
	"strings"
)

func execute(u string, op string, np string) error {
	conf := yamlcustom.ParseYAML()
	userDN := conf.Ldap[0].UserDN
	ldapADDR := conf.Ldap[1].LDAP

	out, err := exec.Command("ldappasswd", "-H", ldapADDR, "-x", "-D", "cn="+u+","+userDN, "-w", op, "-s", np).Output()

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

	// ldapsearch -H ldap://test-x -b cn=testuser,ou=users,dc=company,dc=com -LLL mail | grep mail | awk '{print $2}'
	out, err := exec.Command("ldapsearch", "-H", ldapADDR, "-x", "-b", "cn="+u+","+userDN, "-LLL", "mail", "|", "grep", "mail", "|", "awk", "'{print $2}'").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out)
	fmt.Println(output)

	return output
}

//nolint
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	username := r.FormValue("username")
	// ...

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
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	username := r.FormValue("username")
	// ...

	conf := yamlcustom.ParseYAML()
	smtpUser := conf.Smtp[0].Username
	password := conf.Smtp[1].Password
	hostname := conf.Smtp[2].Hostname
	from := conf.Smtp[3].From
	msg := []byte("test message")
	recipients := strings.Fields(searchMail(username))

	result := smtpss.PlainAuth(smtpUser, password, hostname, from, msg, recipients)
	// result := execute(username, oldPassword, newPassword)
	// fmt.Println(username, password)

	sucess := "Successful"

	if result != nil {
		fmt.Fprintf(w, "Status = %s\n", result)
	} else {
		fmt.Fprintf(w, "Status = %s\n", sucess)
	}
}
