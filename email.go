package emailtemplates

import (
	"net/url"

	"github.com/theopenlane/newman"
)

// NewVerifyEmail returns a new email message based on the config values
func (c Config) NewVerifyEmail(r Recipient, token string) (*newman.EmailMessage, error) {
	data := VerifyEmailData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r

	var err error

	data.VerifyURL, err = addTokenToURL(c.VerifyURL, token)
	if err != nil {
		return nil, err
	}

	return verify(data)
}

// NewWelcomeEmail returns a new email message based on the config values
func (c Config) NewWelcomeEmail(r Recipient, org string) (*newman.EmailMessage, error) {
	data := WelcomeData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r
	data.Organization = org

	return welcome(data)
}

// NewInviteEmail returns a new email message based on the config values
func (c Config) NewInviteEmail(r Recipient, inviterName, org, role, token string) (*newman.EmailMessage, error) {
	data := c.newInvite(r, inviterName, org, role)

	var err error

	data.InviteURL, err = addTokenToURL(c.InviteURL, token)
	if err != nil {
		return nil, err
	}

	return invite(data)
}

// NewInviteEmail returns a new email message based on the config values
func (c Config) NewInviteAcceptedEmail(r Recipient, inviterName string, org string, role string) (*newman.EmailMessage, error) {
	data := c.newInvite(r, inviterName, org, role)

	return inviteAccepted(data)
}

// newInvite creates new invite data for use in the invite emails
func (c Config) newInvite(r Recipient, inviterName string, org string, role string) InviteData {
	data := InviteData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r
	data.InviterName = inviterName
	data.OrganizationName = org
	data.Role = role

	return data
}

// NewPasswordResetRequestEmail returns a new email message based on the config values
func (c Config) NewPasswordResetRequestEmail(r Recipient, token string) (*newman.EmailMessage, error) {
	data := ResetRequestData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r

	var err error

	data.ResetURL, err = addTokenToURL(c.ResetURL, token)
	if err != nil {
		return nil, err
	}

	return passwordResetRequest(data)
}

// NewPasswordResetSuccessEmail returns  a new email message based on the config values
func (c Config) NewPasswordResetSuccessEmail(r Recipient) (*newman.EmailMessage, error) {
	data := ResetSuccessData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r

	return passwordResetSuccess(data)
}

// NewSubscriberEmail returns a new email message based on the config values
func (c Config) NewSubscriberEmail(r Recipient, org, token string) (*newman.EmailMessage, error) {
	data := SubscriberEmailData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.OrganizationName = org

	var err error

	data.VerifySubscriberURL, err = addTokenToURL(c.VerifySubscriberURL, token)
	if err != nil {
		return nil, err
	}

	return subscribe(data)
}

func addTokenToURL(baseURL, token string) (string, error) {
	if token == "" {
		return "", newMissingRequiredFieldError("token")
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	url := base.ResolveReference(&url.URL{RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}
