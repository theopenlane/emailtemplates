package emailtemplates

import (
	"fmt"

	"github.com/theopenlane/newman"
)

// Email subject lines
const (
	welcomeSubject              = "Welcome to %s!"
	verifyEmailSubject          = "Please verify your email address to login to %s"
	inviteSubject               = "Join Your Teammate %s on %s!"
	passwordResetRequestSubject = "%s Password Reset - Action Required"
	passwordResetSuccessSubject = "%s Password Reset Confirmation"
	inviteAcceptedSubject       = "You've been added to an Organization on %s"
	subscribedSubject           = "You've been subscribed to %s"
)

// Config includes fields that are common to all the email builders that are configurable
type Config struct {
	CompanyName    string    `koanf:"companyName" json:"companyName" default:""`
	CompanyAddress string    `koanf:"companyAddress" json:"companyAddress" default:""`
	Corporation    string    `koanf:"corporation" json:"corporation" default:""`
	FromEmail      string    `koanf:"fromEmail" json:"fromEmail" default:""`
	SupportEmail   string    `koanf:"supportEmail" json:"supportEmail" default:""`
	LogoURL        string    `koanf:"logoURL" json:"logoURL" default:""`
	URLS           URLConfig `koanf:"urls" json:"urls"`
}

// URLConfig includes urls that are used in the email templates
type URLConfig struct {
	Root             string `koanf:"root" json:"root" default:""`
	Product          string `koanf:"product" json:"product" default:""`
	Docs             string `koanf:"docs" json:"docs" default:""`
	Verify           string `koanf:"verify" json:"verify" default:""`
	Invite           string `koanf:"invite" json:"invite" default:""`
	PasswordReset    string `koanf:"reset" json:"reset" default:""`
	VerifySubscriber string `koanf:"verifySubscriber" json:"verifySubscriber" default:""`
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
