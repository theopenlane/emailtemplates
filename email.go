package emailtemplates

import (
	"io"
	"net/url"
	"time"

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
	// TrustCenterNDAFullURL is the full URL for the trust center NDA signing page, if provided it will be used instead of TrustCenterURL with token appended
	TrustCenterNDAFullURL string
}

// TrustCenterNDASignedData contains the data needed to create a trust center NDA signed notification email
type TrustCenterNDASignedData struct {
	// OrganizationName is the name of the organization whose NDA was signed
	OrganizationName string
	// TrustCenterURL is the URL where the recipient can access the trust center
	TrustCenterURL string
}

// NewTrustCenterNDASignedEmail creates a new email message notifying the recipient that their NDA has been signed
// and they now have access to the organization's trust center resources.
// The attachment parameter is the signed NDA document to include as an email attachment.
func (c Config) NewTrustCenterNDASignedEmail(r Recipient, data TrustCenterNDASignedData, attachment io.Reader, fileName string) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	content, err := io.ReadAll(attachment)
	if err != nil {
		return nil, err
	}

	if fileName == "" {
		return nil, newMissingRequiredFieldError("filename")
	}

	emailData := TrustCenterNDASignedEmailData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		OrganizationName: data.OrganizationName,
		TrustCenterURL:   data.TrustCenterURL,
	}

	msg, err := trustCenterNDASigned(emailData)
	if err != nil {
		return nil, err
	}

	msg.AddAttachment(newman.NewAttachment(fileName, content))

	return msg, nil
}

// TrustCenterAuthData contains the data needed to create a trust center auth link email
type TrustCenterAuthData struct {
	// OrganizationName is the name of the organization granting access
	OrganizationName string
	// TrustCenterURL is the base URL for the trust center authentication page (token will be appended)
	TrustCenterURL string
	// TrustCenterAuthFullURL is the full URL for the trust center authentication page, if provided it will be used instead of TrustCenterURL with token appended
	TrustCenterAuthFullURL string
}

// QuestionnaireAuthData contains the data needed to create a questionnaire auth link email
type QuestionnaireAuthData struct {
	// CompanyName is the name of the company sending the assessment
	CompanyName string
	// AssessmentName is the name of the assessment/questionnaire
	AssessmentName string
	// QuestionnaireAuthFullURL is the full URL for the questionnaire authentication page, if provided it will be used instead of the configured questionnaire URL with token appended
	QuestionnaireAuthFullURL string
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

	emailData.TrustCenterNDAURL = data.TrustCenterNDAFullURL
	if emailData.TrustCenterNDAURL == "" {
		emailData.TrustCenterNDAURL, err = addTokenToURL(data.TrustCenterURL, token)
		if err != nil {
			return nil, err
		}
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

	emailData.TrustCenterAuthURL = data.TrustCenterAuthFullURL
	if emailData.TrustCenterAuthURL == "" {
		emailData.TrustCenterAuthURL, err = addTokenToURL(data.TrustCenterURL, token)
		if err != nil {
			return nil, err
		}
	}

	return trustCenterAuth(emailData)
}

// NewQuestionnaireAuthEmail creates a new email message with an authentication link to access a questionnaire.
// It takes a recipient, a security token, and questionnaire auth data, then generates an email
// with a tokenized URL that allows the recipient to authenticate and access the questionnaire directly.
func (c Config) NewQuestionnaireAuthEmail(r Recipient, token string, data QuestionnaireAuthData) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	emailData := QuestionnaireAuthEmailData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		CompanyName:    data.CompanyName,
		AssessmentName: data.AssessmentName,
	}

	if c.QuestionnaireEmail != "" {
		emailData.FromEmail = c.QuestionnaireEmail
	}

	emailData.QuestionnaireAuthURL = data.QuestionnaireAuthFullURL
	if emailData.QuestionnaireAuthURL == "" {
		var err error

		emailData.QuestionnaireAuthURL, err = addTokenToURL(c.URLS.Questionnaire, token)
		if err != nil {
			return nil, err
		}
	}

	return questionnaireAuth(emailData)
}

// BillingEmailChangedTemplateData includes the data needed to render the billing email templates
type BillingEmailChangedTemplateData struct {
	OrganizationName string
	OldEmail         string
	NewEmail         string
	ChangedAt        time.Time
}

// NewBillingEmailChangedEmail creates a new email message that is meant to notify orgs
// about changes to their billing email.
func (c Config) NewBillingEmailChangedEmail(r Recipient, data BillingEmailChangedTemplateData) (*newman.EmailMessage, error) {
	if err := c.ensureDefaults(); err != nil {
		return nil, err
	}

	emailData := BillingEmailChangedData{
		EmailData: EmailData{
			Config:    c,
			Recipient: r,
		},
		OrganizationName: data.OrganizationName,
		OldEmail:         data.OldEmail,
		NewEmail:         data.NewEmail,
		ChangedAt:        data.ChangedAt,
	}

	return billingEmailChanged(emailData)
}
