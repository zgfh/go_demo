package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net"
	"net/mail"
	"net/smtp"
	"net/textproto"
	"time"
)

func sendMail(servername, sender, password, recevices, subj, body string,RequireTLS bool) error {
	var c *smtp.Client
	host, port, err := net.SplitHostPort(servername)
	if err != nil {
		return fmt.Errorf("invalid address: %s", err)
	}

	if port == "465" {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}

		conn, err := tls.Dial("tcp", servername, tlsconfig)
		if err != nil {
			return err
		}
		c, err = smtp.NewClient(conn, host)
		if err != nil {
			return err
		}

	} else {
		// Connect to the SMTP smarthost.
		c, err = smtp.Dial(servername)
		if err != nil {
			return err
		}
	}
	defer func() {
		if err := c.Quit(); err != nil {
			fmt.Errorf( "failed to close SMTP connection",  err)
		}
	}()


	// Global Config guarantees RequireTLS is not nil
	if RequireTLS {
		if ok, _ := c.Extension("STARTTLS"); !ok {
			return  fmt.Errorf("require_tls: true (default), but %q does not advertise the STARTTLS extension", servername)
		}

		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}

		if err := c.StartTLS(tlsconfig); err != nil {
			return  fmt.Errorf("starttls failed: %s", err)
		}
	}
	c.Auth(smtp.PlainAuth("", sender, password, host))

	addrs, err := mail.ParseAddressList(sender)
	if err != nil {
		return  fmt.Errorf("parsing from addresses: %s", err)
	}
	if len(addrs) != 1 {
		return  fmt.Errorf("must be exactly one from address")
	}
	if err := c.Mail(addrs[0].Address); err != nil {
		return  fmt.Errorf("sending mail from: %s", err)
	}
	addrs, err = mail.ParseAddressList(recevices)
	if err != nil {
		return  fmt.Errorf("parsing to addresses: %s", err)
	}
	for _, addr := range addrs {
		if err := c.Rcpt(addr.Address); err != nil {
			return  fmt.Errorf("sending rcpt to: %s", err)
		}
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	headers := make(map[string]string)
	headers["From"] = sender
	headers["To"] = recevices
	headers["Subject"] = subj

	for header, value := range headers {
		if err != nil {
			return fmt.Errorf("executing %q header template: %s", header, err)
		}
		fmt.Fprintf(wc, "%s: %s\r\n", header, mime.QEncoding.Encode("utf-8", value))
	}

	buffer := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(buffer)

	fmt.Fprintf(wc, "Date: %s\r\n", time.Now().Format(time.RFC1123Z))
	fmt.Fprintf(wc, "Content-Type: multipart/alternative;  boundary=%s\r\n", multipartWriter.Boundary())
	fmt.Fprintf(wc, "MIME-Version: 1.0\r\n")

	// TODO: Add some useful headers here, such as URL of the alertmanager
	// and active/resolved.
	fmt.Fprintf(wc, "\r\n")

	if len(body) > 0 {
		// Text template
		w, err := multipartWriter.CreatePart(textproto.MIMEHeader{
			"Content-Transfer-Encoding": {"quoted-printable"},
			"Content-Type":              {"text/plain; charset=UTF-8"},
		})
		if err != nil {
			return  fmt.Errorf("creating part for text template: %s", err)
		}
		qw := quotedprintable.NewWriter(w)
		_, err = qw.Write([]byte(body))
		if err != nil {
			return  err
		}
		err = qw.Close()
		if err != nil {
			return  err
		}
	}


	err = multipartWriter.Close()
	if err != nil {
		return  fmt.Errorf("failed to close multipartWriter: %v", err)
	}

	_, err = wc.Write(buffer.Bytes())
	if err != nil {
		return  fmt.Errorf("failed to write body buffer: %v", err)
	}
	return nil
}

func main() {
	err := sendMail("smtp.exmail.qq.com:465", "zzg@daozzg.com", "h6egcFCcHkppm33G", "zzg@daozzg.com,zg.zhu@daocloud.io", "test", "test\n\n\n\nssss",false)
	if err != nil {
		fmt.Printf("send mail failed: %s \n", err)
	}
}
