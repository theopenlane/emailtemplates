package emailtemplates

import "github.com/theopenlane/newman"

// NewVerifyEmail returns a new VerifyEmailData with default values for use at Openlane
func (c Config) NewVerifyEmail(r Recipient) (*newman.EmailMessage, error) {
	data := VerifyEmailData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r

	return verify(data)
}

// NewWelcomeEmail returns a new WelcomeData with default values for use at Openlane
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

// NewInviteEmail returns a new InviteData with default values for use at Openlane
func (c Config) NewInviteEmail(r Recipient, inviterName string, org string, role string) (*newman.EmailMessage, error) {
	data := InviteData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r
	data.InviterName = inviterName
	data.OrganizationName = org
	data.Role = role

	return invite(data)
}

// NewPasswordResetRequestEmail returns a new ResetRequestData with default values for use at Openlane
func (c Config) NewPasswordResetRequestEmail() (*newman.EmailMessage, error) {
	data := ResetRequestData{
		EmailData: EmailData{
			Config: c,
		},
	}

	return passwordResetRequest(data)
}

// NewPasswordResetSuccessEmail returns a new ResetSuccessData with default values for use at Openlane
func (c Config) NewPasswordResetSuccessEmail(r Recipient) (*newman.EmailMessage, error) {
	data := ResetSuccessData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r

	return passwordResetSuccess(data)
}

// NewSubscribedEmail returns a new SubscriberEmailData with default values for use at Openlane
func (c Config) NewSubscribedEmail() (*newman.EmailMessage, error) {
	data := SubscriberEmailData{
		EmailData: EmailData{
			Config: c,
		},
	}

	return subscribe(data)
}
