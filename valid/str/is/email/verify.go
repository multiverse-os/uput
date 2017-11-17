package verify

// https://www.webdigi.co.uk/blog/2009/how-to-check-if-an-email-address-exists-without-sending-an-email/
// TODO:
// Cache DNS lookup and mail sever results
// Throttle per server? Use all MX servers?
// Better error handling? Custom error interface? http://blog.golang.org/error-handling-and-go

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

// VerifyAddress checks an address
// returns: AddressOK , SMTP MSG, Error
func VerifyAddress(emailAddress string) (bool, string, error) {

	// Parse the address into a user and a host part
	parts := strings.Split(emailAddress, "@")
	if len(parts) != 2 {
		return false, "", fmt.Errorf("verifyAddress: Can not parse this email address: %q", emailAddress)
	}
	log.Println(parts)

	// Lookup an MX server for this address
	// Pick one at random? Or??
	servers, err := net.LookupMX(parts[1])
	if err != nil {
		return false, "", fmt.Errorf("verifyAddress: LookupMX failed: %q", err)
	}
	//for _, s := range servers {
	//	fmt.Println(s.Host)
	//}
	server := servers[0].Host
	log.Println("Mail server:", server)

	// Connect to the remote SMTP server.
	c, err := smtp.Dial(fmt.Sprintf("%s:%d", server, 25))
	if err != nil {
		return false, "", fmt.Errorf("verifyAddress: smtp Dial failed: %q", err)
	}

	// Be polite, say HELLO
	if err = c.Hello("emailcaptain.com"); err != nil {
		return false, err.Error(), fmt.Errorf("verifyAddress: smtp Hello failed: %q", err)
	}

	// First try the smtp package's verify()
	if err := c.Verify(emailAddress); err != nil {
		log.Printf("verifyAddress: Go's verify() belives it's invalid: %q\n", err)
	} else {
		log.Printf("verifyAddress: Go's verify() belives it's Valid!\n")
	}

	// Set the sender and recipient first
	if err := c.Mail("emailtest@emailcaptain.com"); err != nil { // FIXME =)
		return false, err.Error(), fmt.Errorf("verifyAddress: smtp Mail FROM failed: %q", err)
	}
	if err := c.Rcpt(emailAddress); err != nil {
		return false, err.Error(), fmt.Errorf("verifyAddress: smtp RCPT TO failed: %q", err)
	}

	return true, "", nil

}
