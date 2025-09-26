package emailtemplates

import (
	"net/url"

	"github.com/theopenlane/newman"
)

// NewVerifyEmail returns a new email message based on the config values and the provided recipient and token
func (c Config) NewVerifyEmail(r Recipient, token string) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	data := VerifyEmailData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
	}

	var err error

	data.URLS.Verify, err = addTokenToURL(c.URLS.Verify, token)
	if err != nil {
		return nil, err
	}

	return verify(data)
}

// NewWelcomeEmail returns a new email message based on the config values and the provided recipient and organization name
func (c Config) NewWelcomeEmail(r Recipient) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	data := WelcomeData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
	}

	return welcome(data)
}

// InviteTemplateData includes the data needed to render the invite email templates
type InviteTemplateData struct {
	InviterName      string
	OrganizationName string
	Role             string
}

// NewInviteEmail returns a new email message based on the config values and the provided recipient and invite data
func (c Config) NewInviteEmail(r Recipient, i InviteTemplateData, token string) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	data := c.newInvite(r, i)

	data.Recipient = r

	var err error

	data.URLS.Invite, err = addTokenToURL(c.URLS.Invite, token)
	if err != nil {
		return nil, err
	}

	return invite(data)
}

// NewInviteEmail returns a new email message based on the config values and the provided recipient and invite data
func (c Config) NewInviteAcceptedEmail(r Recipient, i InviteTemplateData) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	data := c.newInvite(r, i)

	return inviteAccepted(data)
}

// newInvite creates new invite data for use in the invite emails
func (c Config) newInvite(r Recipient, i InviteTemplateData) InviteData {
	data := InviteData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		InviterName:      i.InviterName,
		OrganizationName: i.OrganizationName,
		Role:             i.Role,
	}

	return data
}

// NewPasswordResetRequestEmail returns a new email message based on the config values and the provided recipient and token
func (c Config) NewPasswordResetRequestEmail(r Recipient, token string) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

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

// NewPasswordResetSuccessEmail returns  a new email message based on the config values and the provided recipient
func (c Config) NewPasswordResetSuccessEmail(r Recipient) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	data := ResetSuccessData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
	}

	return passwordResetSuccess(data)
}

// NewSubscriberEmail returns a new email message based on the config values and the provided recipient, organization name, and token
func (c Config) NewSubscriberEmail(r Recipient, organizationName, token string) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	data := SubscriberEmailData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		OrganizationName: organizationName,
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

// NewVerifyBillingEmail returns a new email message based on the config values and the provided recipient and token
func (c Config) NewVerifyBillingEmail(r Recipient, token string) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	data := VerifyBillingEmailData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
	}

	var err error

	data.URLS.VerifyBilling, err = addTokenToURL(c.URLS.VerifyBilling, token)
	if err != nil {
		return nil, err
	}

	return verifyBilling(data)
}

// TrustCenterNDARequestData contains the data needed to create a trust center NDA request email
type TrustCenterNDARequestData struct {
	// OrganizationName is the name of the organization requesting the NDA signature
	OrganizationName string
	// TrustCenterURL is the base URL for the trust center NDA signing page (token will be appended)
	TrustCenterURL string
}

// TrustCenterAuthData contains the data needed to create a trust center auth link email
type TrustCenterAuthData struct {
	// OrganizationName is the name of the organization granting access
	OrganizationName string
	// TrustCenterURL is the base URL for the trust center authentication page (token will be appended)
	TrustCenterURL string
}

// NewTrustCenterNDARequestEmail creates a new email message for requesting an NDA signature to access the trust center.
// It takes a recipient, a security token, and trust center NDA request data, then generates an email
// with a tokenized URL that allows the recipient to sign the NDA and gain access to protected trust center resources.
func (c Config) NewTrustCenterNDARequestEmail(r Recipient, token string, data TrustCenterNDARequestData) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	var err error
	emailData := TrustCenterNDARequestEmailData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		OrganizationName: data.OrganizationName,
	}

	emailData.TrustCenterNDAURL, err = addTokenToURL(data.TrustCenterURL, token)
	if err != nil {
		return nil, err
	}
	return trustCenterNDARequest(emailData)
}

// NewTrustCenterAuthEmail creates a new email message with an authentication link to access the trust center.
// It takes a recipient, a security token, and trust center auth data, then generates an email
// with a tokenized URL that allows the recipient to authenticate and access trust center resources directly.
func (c Config) NewTrustCenterAuthEmail(r Recipient, token string, data TrustCenterAuthData) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	var err error
	emailData := TrustCenterAuthEmailData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		OrganizationName: data.OrganizationName,
	}

	emailData.TrustCenterAuthURL, err = addTokenToURL(data.TrustCenterURL, token)
	if err != nil {
		return nil, err
	}
	return trustCenterAuth(emailData)
}
