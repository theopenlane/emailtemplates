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

	data.URLS.Verify, err = addTokenToURL(c.URLS.Verify, token)
	if err != nil {
		return nil, err
	}

	return verify(data)
}

// NewWelcomeEmail returns a new email message based on the config values
func (c Config) NewWelcomeEmail(r Recipient, org string) (*newman.EmailMessage, error) {
	data := WelcomeData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		Organization: org,
	}

	return welcome(data)
}

// NewInviteEmail returns a new email message based on the config values
func (c Config) NewInviteEmail(r Recipient, inviterName, org, role, token string) (*newman.EmailMessage, error) {
	data := c.newInvite(r, inviterName, org, role)

	data.Recipient = r

	var err error

	data.URLS.Invite, err = addTokenToURL(c.URLS.Invite, token)
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
			Config:    c,
			Recipient: r,
		},
		InviterName:      inviterName,
		OrganizationName: org,
		Role:             role,
	}

	return data
}

// NewPasswordResetRequestEmail returns a new email message based on the config values
func (c Config) NewPasswordResetRequestEmail(r Recipient, token string) (*newman.EmailMessage, error) {
	data := ResetRequestData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
	}

	var err error

	data.URLS.PasswordReset, err = addTokenToURL(c.URLS.PasswordReset, token)
	if err != nil {
		return nil, err
	}

	return passwordResetRequest(data)
}

// NewPasswordResetSuccessEmail returns  a new email message based on the config values
func (c Config) NewPasswordResetSuccessEmail(r Recipient) (*newman.EmailMessage, error) {
	data := ResetSuccessData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
	}

	return passwordResetSuccess(data)
}

// NewSubscriberEmail returns a new email message based on the config values
func (c Config) NewSubscriberEmail(r Recipient, org, token string) (*newman.EmailMessage, error) {
	data := SubscriberEmailData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		OrganizationName: org,
	}

	var err error

	data.URLS.VerifySubscriber, err = addTokenToURL(c.URLS.VerifySubscriber, token)
	if err != nil {
		return nil, err
	}

	return subscribe(data)
}

// addTokenToURL adds a token to the URL as a query parameter
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
