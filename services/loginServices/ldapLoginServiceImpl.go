package loginservices

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
	"github.com/gofiber/fiber/v2"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
)

const (
    ldapServer   = "ldap.forumsys.com"
    ldapPort     = 389
    ldapBindDN   = "cn=read-only-admin,dc=example,dc=com"
    ldapPassword = "password"
    ldapSearchDN = "dc=example,dc=com"
)

func (ldapLoginServiceImpl *LdapLoginServiceImpl) AuthUsingLDAP(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin) (bool, *loginmodel.UserLDAPData, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
    if err != nil {
        return false, nil, err
    }
    defer l.Close()

	err = l.Bind(ldapBindDN, ldapPassword)
	if err != nil {
		return false, nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		ldapSearchDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", reqLogin.Username),
		[]string{"uid", "cn", "sn", "mail","telephoneNumber"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	
	if err != nil {
		return false, nil, err
	}

	if len(sr.Entries) == 0 {
		return false, nil, fmt.Errorf("user not found")
	}
	entry := sr.Entries[0]

	// verify user password by binding to user dn (with user password)
	err = l.Bind(entry.DN, reqLogin.Password)
	if err != nil {
		return false, nil, err
	}

	// (optional) store data
	data := new(loginmodel.UserLDAPData)
	data.ID = reqLogin.Username

	for _, attr := range entry.Attributes {
		switch attr.Name {
		case "sn":
			data.Name = attr.Values[0]
		case "mail":
			data.Email = attr.Values[0]
		case "cn":
			data.FullName = attr.Values[0]
		case "telephoneNumber":
			data.TelephoneNumber = attr.Values[0]
		}
	}
	return true, data, nil
}