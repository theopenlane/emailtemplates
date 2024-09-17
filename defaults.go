package emailtemplates

// NewVerifyTemplate returns a new VerifyEmailData with default values for use at Openlane
func (c Config) NewVerifyTemplate(r Recipient) VerifyEmailData {
	data := VerifyEmailData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r

	return data
}

// NewWelcomeTemplate returns a new WelcomeData with default values for use at Openlane
func (c Config) NewWelcomeTemplate(r Recipient, org string) WelcomeData {
	data := WelcomeData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r
	data.Organization = org

	return data
}

// NewInviteTemplate returns a new InviteData with default values for use at Openlane
func (c Config) NewInviteTemplate(r Recipient, inviterName string, org string, role string) InviteData {
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

// NewPasswordResetRequestTemplate returns a new ResetRequestData with default values for use at Openlane
func (c Config) NewPasswordResetRequestTemplate() ResetRequestData {
	data := ResetRequestData{
		EmailData: EmailData{
			Config: c,
		},
	}

	return data
}

// NewPasswordResetSuccessTemplate returns a new ResetSuccessData with default values for use at Openlane
func (c Config) NewPasswordResetSuccessTemplate(r Recipient) ResetSuccessData {
	data := ResetSuccessData{
		EmailData: EmailData{
			Config: c,
		},
	}

	data.Recipient = r

	return data
}

// NewSubscribedTemplate returns a new SubscriberEmailData with default values for use at Openlane
func (c Config) NewSubscribedTemplate() SubscriberEmailData {
	data := SubscriberEmailData{
		EmailData: EmailData{
			Config: c,
		},
	}

	return data
}
