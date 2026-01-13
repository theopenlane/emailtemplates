package emailtemplates

import (
	"fmt"

	"github.com/theopenlane/newman"
)

// Email subject lines
const (
	welcomeSubject               = "Welcome to %s!"
	verifyEmailSubject           = "Please verify your email address to login to %s"
	inviteSubject                = "Join Your Teammate %s on %s!"
	passwordResetRequestSubject  = "%s Password Reset - Action Required"
	passwordResetSuccessSubject  = "%s Password Reset Confirmation"
	inviteAcceptedSubject        = "You've been added to an Organization on %s"
	subscribedSubject            = "You've been subscribed to %s"
	verifyBillingSubject         = "Please verify the billing email for %s to ensure your account stays up to date"
	trustCenterNDARequestSubject = "%s Trust Center NDA Request"
	trustCenterAuthSubject       = "Access %s's Trust Center"
	questionnaireAuthSubject     = "Access %s Questionnaire from %s"
	billingEmailChangedSubject   = "Billing Email Changed for %s"
)

// Config includes fields that are common to all the email builders that are configurable
type Config struct {
	// CompanyName is the name of the company that is sending the email
	CompanyName string `koanf:"companyname" json:"companyname" default:""`
	// CompanyAddress is the address of the company that is sending the email, included in the footer
	CompanyAddress string `koanf:"companyaddress" json:"companyaddress" default:""`
	// Corporation is the official corporation name that is sending the email, included in the footer
	Corporation string `koanf:"corporation" json:"corporation" default:""`
	// Year is the year that the email is being sent, included in the footer for the copyright year
	Year int `koanf:"year" json:"year" default:""`
	// FromEmail is the email address that the email is sent from
	FromEmail string `koanf:"fromemail" json:"fromemail" default:"" domain:"inherit" domainPrefix:"no-reply@mail"`
	// SupportEmail is the email address that the recipient can contact for support
	SupportEmail string `koanf:"supportemail" json:"supportemail" default:"" domain:"inherit" domainPrefix:"support@"`
	// QuestionnaireEmail is the email address for questionnaire/assessment related emails.
	// If not provided, the FromEmail will be used as before
	QuestionnaireEmail string `koanf:"questionnaireemail" json:"questionnaireemail" default:"" domain:"inherit" domainPrefix:"questionnaire@"`
	// LogoURL is the URL to the company logo that is included in the email if provided
	LogoURL string `koanf:"logourl" json:"logourl" default:""`
	// URLS includes URLs that are used in the email templates
	URLS URLConfig `koanf:"urls" json:"urls"`
	// TemplatesPath is the path to the email templates to override the default templates
	TemplatesPath string `koanf:"templatespath" json:"templatespath" default:""`
}

// URLConfig includes urls that are used in the email templates
type URLConfig struct {
	// Root is the root domain for the email
	Root string `koanf:"root" json:"root" default:"" domain:"inherit" domainPrefix:"https://www"`
	// Product is the product domain for the email, usually the main UI where a user logs in
	Product string `koanf:"product" json:"product" default:"" domain:"inherit" domainPrefix:"https://console"`
	// Docs is the docs domain for the email, where a user can find documentation
	Docs string `koanf:"docs" json:"docs" default:"" domain:"inherit" domainPrefix:"https://docs"`
	// Verify is the URL to verify an email address
	Verify string `koanf:"verify" json:"verify" default:"" domain:"inherit" domainPrefix:"https://console" domainSuffix:"/verify"`
	// Invite is the URL to accept an invite to an organization
	Invite string `koanf:"invite" json:"invite" default:"" domain:"inherit" domainPrefix:"https://console" domainSuffix:"/invite"`
	// PasswordReset is the URL to reset a password
	PasswordReset string `koanf:"reset" json:"reset" default:"" domain:"inherit" domainPrefix:"https://console" domainSuffix:"/password-reset"`
	// VerifySubscriber is the URL to verify a subscriber for an organization
	VerifySubscriber string `koanf:"verifysubscriber" json:"verifysubscriber" default:"" domain:"inherit" domainPrefix:"https://console" domainSuffix:"/subscriber-verify"`
	// VerifyBilling is the URL to verify a billing account
	VerifyBilling string `koanf:"verifybilling" json:"verifybilling" default:"" domain:"inherit" domainPrefix:"https://console" domainSuffix:"/verify-billing"`
	// Questionnaire is the URL to access a questionnaire
	Questionnaire string `koanf:"questionnaire" json:"questionnaire" default:"" domain:"inherit" domainPrefix:"https://console" domainSuffix:"/questionnaire"`
}

// EmailData includes data fields that are common to all the email builders
type EmailData struct {
	Config
	// Subject is the subject line of the email
	Subject string `json:"subject"`
	// Recipient is the person who will receive the email
	Recipient Recipient `json:"recipient"`
}

// Recipient includes fields for the recipient of the email
type Recipient struct {
	// Email is the email address of the recipient
	Email string `json:"email"`
	// FirstName is the first name of the recipient
	FirstName string `json:"first_name"`
	// LastName is the last name of the recipient
	LastName string `json:"last_name"`
}

// WelcomeData includes fields for the welcome email
type WelcomeData struct {
	EmailData
}

// VerifyEmailData includes fields for the verify email
type VerifyEmailData struct {
	EmailData
}

// SubscriberEmailData includes fields for the subscriber email
type SubscriberEmailData struct {
	EmailData
	// Organization is the name of the organization that the user is subscribing to
	OrganizationName string `json:"organization_name"`
}

// VerifyBillingEmailData includes fields for the verify billing email
type VerifyBillingEmailData struct {
	EmailData
	// Organization is the name of the organization that the user is subscribing to
	OrganizationName string `json:"organization_name"`
}

// InviteData includes fields for the invite email
type InviteData struct {
	EmailData
	// InviterName is the name of the person who is inviting the recipient
	InviterName string `json:"inviter_name"`
	// OrganizationName is the name of the organization that the user is being invited to
	OrganizationName string `json:"organization_name"`
	// Role is the role that the user is being invited to join the organization as
	Role string `json:"role"`
}

// ResetRequestData includes fields for the password reset request email
type ResetRequestData struct {
	EmailData
}

// ResetSuccessData includes fields for the password reset success email
type ResetSuccessData struct {
	EmailData
}

// TrustCenterNDARequestEmailData includes fields for the trust center NDA request email
type TrustCenterNDARequestEmailData struct {
	EmailData
	// OrganizationName is the name of the organization requesting the NDA signature
	OrganizationName string `json:"organization_name"`
	// TrustCenterNDAURL is the URL where the recipient can sign the NDA to access the trust center
	TrustCenterNDAURL string `json:"trust_center_nda_url"`
}

// TrustCenterAuthEmailData includes fields for the trust center auth link email
type TrustCenterAuthEmailData struct {
	EmailData
	// OrganizationName is the name of the organization granting access
	OrganizationName string `json:"organization_name"`
	// TrustCenterAuthURL is the URL where the recipient can authenticate to access the trust center
	TrustCenterAuthURL string `json:"trust_center_auth_url"`
}

// QuestionnaireAuthEmailData includes fields for the questionnaire auth link email
type QuestionnaireAuthEmailData struct {
	EmailData
	// CompanyName is the name of the company sending the assessment
	CompanyName string `json:"company_name"`
	// AssessmentName is the name of the assessment/questionnaire
	AssessmentName string `json:"assessment_name"`
	// QuestionnaireAuthURL is the URL where the recipient can authenticate to access the questionnaire
	QuestionnaireAuthURL string `json:"questionnaire_auth_url"`
}

// BillingEmailChangedData includes fields for the billing email changed notification
type BillingEmailChangedData struct {
	EmailData
	OrganizationName string `json:"organization_name"`
	OldEmail         string `json:"old_email"`
	NewEmail         string `json:"new_email"`
}

// Build validates and creates a new email from pre-rendered templates
func (e EmailData) Build(text, html string) (*newman.EmailMessage, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}

	opts :=
		[]newman.MessageOption{
			newman.WithTo([]string{e.Recipient.Email}),
			newman.WithFrom(e.FromEmail),
			newman.WithSubject(e.Subject),
			newman.WithHTML(html),
			newman.WithText(text),
		}

	return newman.NewEmailMessageWithOptions(opts...), nil
}

// Validate that all required data is present to assemble a sendable email
func (e EmailData) Validate() error {
	switch {
	case e.Subject == "":
		return newMissingRequiredFieldError("subject")
	case e.Recipient.Email == "":
		return newMissingRequiredFieldError("email")
	}

	return nil
}

// verify creates a new email to verify an email address
func verify(data VerifyEmailData) (*newman.EmailMessage, error) {
	text, html, err := Render("verify_email", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(verifyEmailSubject, data.CompanyName)

	return data.Build(text, html)
}

// welcome creates a new email to welcome a new user
func welcome(data WelcomeData) (*newman.EmailMessage, error) {
	text, html, err := Render("welcome", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(welcomeSubject, data.CompanyName)

	return data.Build(text, html)
}

// invite creates a new email to invite a user to an organization
func invite(data InviteData) (*newman.EmailMessage, error) {
	text, html, err := Render("invite", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(inviteSubject, data.InviterName, data.CompanyName)

	return data.Build(text, html)
}

// inviteAccepted creates a new email to notify a user that their invite has been accepted
func inviteAccepted(data InviteData) (*newman.EmailMessage, error) {
	text, html, err := Render("invite_joined", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(inviteAcceptedSubject, data.CompanyName)

	return data.Build(text, html)
}

// passwordResetRequest creates a new email to request a password reset
func passwordResetRequest(data ResetRequestData) (*newman.EmailMessage, error) {
	text, html, err := Render("password_reset_request", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(passwordResetRequestSubject, data.CompanyName)

	return data.Build(text, html)
}

// passwordResetSuccess creates a new email to confirm a password reset
func passwordResetSuccess(data ResetSuccessData) (*newman.EmailMessage, error) {
	text, html, err := Render("password_reset_success", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(passwordResetSuccessSubject, data.CompanyName)

	return data.Build(text, html)
}

// subscribe creates a new email to confirm a subscription
func subscribe(data SubscriberEmailData) (*newman.EmailMessage, error) {
	text, html, err := Render("subscribe", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(subscribedSubject, data.CompanyName)

	return data.Build(text, html)
}

// verifyBilling creates a new email to verify a billing account
func verifyBilling(data VerifyBillingEmailData) (*newman.EmailMessage, error) {
	text, html, err := Render("verify_billing", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(verifyBillingSubject, data.CompanyName)

	return data.Build(text, html)
}

// trustCenterNDARequest creates a new email to request an NDA for the trust center
func trustCenterNDARequest(data TrustCenterNDARequestEmailData) (*newman.EmailMessage, error) {
	text, html, err := Render("trust_center_nda_request", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(trustCenterNDARequestSubject, data.OrganizationName)

	return data.Build(text, html)
}

// trustCenterAuth creates a new email with an auth link for the trust center
func trustCenterAuth(data TrustCenterAuthEmailData) (*newman.EmailMessage, error) {
	text, html, err := Render("trust_center_auth", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(trustCenterAuthSubject, data.OrganizationName)

	return data.Build(text, html)
}

// questionnaireAuth creates a new email with an auth link for the questionnaire
func questionnaireAuth(data QuestionnaireAuthEmailData) (*newman.EmailMessage, error) {
	text, html, err := Render("questionnaire_auth", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(questionnaireAuthSubject, data.AssessmentName, data.CompanyName)

	return data.Build(text, html)
}

// billingEmailChanged creates a new email to notify about a billing email change
func billingEmailChanged(data BillingEmailChangedData) (*newman.EmailMessage, error) {
	text, html, err := Render("billing_email_changed", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(billingEmailChangedSubject, data.OrganizationName)
	return data.Build(text, html)
}
