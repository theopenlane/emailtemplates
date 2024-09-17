package emailtemplates

import (
	"fmt"

	"github.com/theopenlane/newman"
)

// Email subject lines
const (
	WelcomeSubject              = "Welcome to %s!"
	VerifyEmailSubject          = "Please verify your email address to login to %s"
	InviteSubject               = "Join Your Teammate %s on %s!"
	PasswordResetRequestSubject = "%s Password Reset - Action Required"
	PasswordResetSuccessSubject = "%s Password Reset Confirmation"
	InviteAcceptedSubject       = "You've been added to an Organization on %s"
	SubscribedSubject           = "You've been subscribed to %s"
)

// Config includes fields that are common to all the email builders that are configurable
type Config struct {
	CompanyName    string `koanf:"companyName" json:"companyName" default:"Openlane"`
	CompanyAddress string `koanf:"companyAddress" json:"companyAddress" default:"5150 Broadway St &middot; San Antonio, TX 78209"`
	Corporation    string `koanf:"corporation" json:"corporation" default:"theopenlane, Inc."`
	RootDomain     string `koanf:"rootDomain" json:"rootDomain" default:"https://theopenlane.io"`
	ProductDomain  string `koanf:"productDomain" json:"productDomain" default:"https://console.theopenlane.io"`
	DocsDomain     string `koanf:"docsDomain" json:"docsDomain" default:"https://docs.theopenlane.io"`
	FromEmail      string `koanf:"fromEmail" json:"fromEmail" default:"no-reply@mail.theopenlane.io"`
	SupportEmail   string `koanf:"supportEmail" json:"supportEmail" default:"support@theopenlane.io"`
	URLConfig
}

// URLConfig includes urls that are used in the email templates
type URLConfig struct {
	VerifyURL           string `koanf:"verifyURL" json:"verifyURL" default:"https://console.theopenlane.io/verify"`
	InviteURL           string `koanf:"inviteURL" json:"inviteURL" default:"https://console.theopenlane.io/invite"`
	ResetURL            string `koanf:"resetURL" json:"resetURL" default:"https://console.theopenlane.io/password-reset"`
	VerifySubscriberURL string `koanf:"verifySubscriberURL" json:"verifySubscriberURL" default:"https://theopenlane.io/subscribe"`
}

// EmailData includes data fields that are common to all the email builders
type EmailData struct {
	Config
	Subject   string    `json:"subject"`
	Recipient Recipient `json:"recipient"`
}

type Recipient struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type WelcomeData struct {
	EmailData
	Organization string `json:"organization"`
}

type VerifyEmailData struct {
	EmailData
}

type SubscriberEmailData struct {
	EmailData
	OrganizationName string `json:"organization_name"`
}

type InviteData struct {
	EmailData
	InviterName      string `json:"inviter_name"`
	OrganizationName string `json:"organization_name"`
	Role             string `json:"role"`
}

type ResetRequestData struct {
	EmailData
}

type ResetSuccessData struct {
	EmailData
}

// Build creates a new email from pre-rendered templates
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

// Verify creates a new email to verify an email address
func Verify(data VerifyEmailData) (*newman.EmailMessage, error) {
	text, html, err := Render("verify_email", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(VerifyEmailSubject, data.CompanyName)

	return data.Build(text, html)
}

// Welcome creates a new email to welcome a new user
func Welcome(data WelcomeData) (*newman.EmailMessage, error) {
	text, html, err := Render("welcome", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(WelcomeSubject, data.CompanyName)

	return data.Build(text, html)
}

// Invite creates a new email to invite a user to an organization
func Invite(data InviteData) (*newman.EmailMessage, error) {
	text, html, err := Render("invite", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(InviteSubject, data.InviterName, data.CompanyName)

	return data.Build(text, html)
}

// InviteAccepted creates a new email to notify a user that their invite has been accepted
func InviteAccepted(data InviteData) (*newman.EmailMessage, error) {
	text, html, err := Render("invite_joined", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(InviteAcceptedSubject, data.CompanyName)

	return data.Build(text, html)
}

// PasswordResetRequest creates a new email to request a password reset
func PasswordResetRequest(data ResetRequestData) (*newman.EmailMessage, error) {
	text, html, err := Render("password_reset_request", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(PasswordResetRequestSubject, data.CompanyName)

	return data.Build(text, html)
}

// PasswordResetSuccess creates a new email to confirm a password reset
func PasswordResetSuccess(data ResetSuccessData) (*newman.EmailMessage, error) {
	text, html, err := Render("password_reset_success", data)
	if err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(PasswordResetSuccessSubject, data.CompanyName)

	return data.Build(text, html)
}
