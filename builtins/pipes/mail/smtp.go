package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/smtp"
	"strings"
	"time"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/stdio"
)

// NewMail creates a new stdio.Io object for sending mail
func NewMail(parameters string) (stdio.Io, error) {
	m := new(Mail)
	//m.response = bytes.NewBuffer(m.b)
	params := strings.SplitN(parameters, ":", 2)

	switch len(params) {
	case 1:
		m.subject = "Sent from " + app.Name // TODO: pick this from config
	case 2:
		m.subject = params[1]
	default:
		return nil, fmt.Errorf("Unable to parse parameters. Expecting 'email addresses[:subject]', instead got '%s'", parameters)
	}

	m.recipients = strings.Split(params[0], ";")
	if len(m.recipients) == 0 {
		return nil, fmt.Errorf("No recipients found in Stdio.Io parameters")
	}

	m.smtpAuth, m.useAuth = getSmtpAuth()
	if m.useAuth {
		m.bodyAuth = bytes.NewBuffer(m.bAuth)
		err := setSubject(m.bodyAuth, m.subject)
		if err != nil {
			return nil, err
		}
		return m, nil
	}

	domain, err := getDomain(m.recipients[0])
	if err != nil {
		return nil, fmt.Errorf("Unable to find any host names: %s", err.Error())
	}

	mx, err := net.LookupMX(domain)
	if err != nil || len(mx) == 0 {
		mx = []*net.MX{{Host: domain, Pref: 1}}
	}

	var (
		serverName string
		conn       net.Conn
	)

	for i := range mx {
		for _, port := range ports {
			addr := fmt.Sprintf("%s:%d", mx[i].Host, port)
			debug.Log("mail: connecting to", addr, "....")

			conn, err = net.DialTimeout("tcp", addr, 2*time.Second)
			if err != nil {
				continue
			}

			m.client, err = smtp.NewClient(conn, addr)

			if err == nil {
				serverName = mx[i].Host
				goto connected
			}
			debug.Log("mail: connection failed:", err)
		}
	}

	return nil, fmt.Errorf("Unable to connect on any domains nor port numbers. Last error: %s", err.Error())

connected:
	debug.Log("mail: connected!")

	useTls, _ := m.client.Extension("STARTTLS")
	if useTls {
		debug.Log("mail: using TLS")
		tlsConfig := new(tls.Config)
		tlsConfig.InsecureSkipVerify = allowInsecure()
		tlsConfig.ServerName = serverName
		err := m.client.StartTLS(tlsConfig)
		if err != nil {
			return nil, fmt.Errorf("Unable to start TLS session: %s", err.Error())
		}

	} else if !allowPlainText() {
		return nil, fmt.Errorf("%s doesn't allow TLS connections and murex set to disallow clear text emails", serverName)
	}

	// Set the sender....
	err = m.client.Mail(senderAddr())
	if err != nil {
		return nil, fmt.Errorf("Cannot set sender name in FROM field: %s", err.Error())
	}

	// ...and the recipient(s)
	for _, addr := range m.recipients {
		debug.Log("mail: Adding rcpt:", addr)
		err = m.client.Rcpt(addr)
		if err != nil {
			return nil, fmt.Errorf("Error adding recipient '%s': %s", addr, err.Error())
		}
	}

	// Instantiate the email body.
	m.body, err = m.client.Data()
	if err != nil {
		return nil, fmt.Errorf("Unable to instantiate the email body: %s", err.Error())
	}

	err = setSubject(m.body, m.subject)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Mail) send() {
	if m.useAuth {
		err := smtp.SendMail(
			fmt.Sprintf("%s:%d", smtpAuthHost(), smtpAuthPort()),
			m.smtpAuth, senderAddr(), m.recipients, m.bAuth,
		)
		debug.Log("mail smtp.SendMail() error:", err.Error())
		return
	}

	err := m.body.Close()
	if err != nil {
		debug.Log("mail: m.body.Close() error:", err)
		//m.response.WriteString(err.Error())
	}

	// Send the QUIT command and close the connection.
	err = m.client.Quit()
	if err != nil && err != io.EOF {
		debug.Log("mail: m.client.Quit() error:", err)
	}
}
