package emailtemplates

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		data    EmailData
		wantErr bool
	}{
		{
			name: "valid data",
			data: EmailData{
				Subject: "Test Subject",
				Recipient: Recipient{
					Email: "test@example.com",
				},
			},
			wantErr: false,
		},
		{
			name: "missing subject",
			data: EmailData{
				Recipient: Recipient{
					Email: "test@example.com",
				},
			},
			wantErr: true,
		},
		{
			name: "missing recipient email",
			data: EmailData{
				Subject: "Test Subject",
				Recipient: Recipient{
					Email: "",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.data.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	data := VerifyEmailData{
		EmailData: EmailData{
			Subject: "Verify Email",
			Recipient: Recipient{
				Email: "test@example.com",
			},
			Config: Config{
				CompanyName: "Test Company",
			},
		},
	}

	email, err := Verify(data)
	require.NoError(t, err)
	require.NotNil(t, email)

	assert.Equal(t, "Please verify your email address to login to Test Company", email.Subject)
}

func TestWelcome(t *testing.T) {
	data := WelcomeData{
		EmailData: EmailData{
			Subject: "Welcome",
			Recipient: Recipient{
				Email: "test@example.com",
			},
			Config: Config{
				CompanyName: "Test Company",
			},
		},
		Organization: "Test Org",
	}

	email, err := Welcome(data)
	require.NoError(t, err)
	require.NotNil(t, email)

	assert.Equal(t, "Welcome to Test Company!", email.Subject)
}

func TestInvite(t *testing.T) {
	data := InviteData{
		EmailData: EmailData{
			Subject: "Invite",
			Recipient: Recipient{
				Email: "test@example.com",
			},
			Config: Config{
				CompanyName: "Test Company",
			},
		},
		InviterName:      "John Doe",
		OrganizationName: "Test Org",
		Role:             "Admin",
	}

	email, err := Invite(data)
	require.NoError(t, err)
	require.NotNil(t, email)

	assert.Equal(t, "Join Your Teammate John Doe on Test Company!", email.Subject)
}

func TestInviteAccepted(t *testing.T) {
	data := InviteData{
		EmailData: EmailData{
			Subject: "Invite Accepted",
			Recipient: Recipient{
				Email: "test@example.com",
			},
			Config: Config{
				CompanyName: "Test Company",
			},
		},
	}

	email, err := InviteAccepted(data)
	require.NoError(t, err)
	require.NotNil(t, email)

	assert.Equal(t, "You've been added to an Organization on Test Company", email.Subject)
}

func TestPasswordResetRequest(t *testing.T) {
	data := ResetRequestData{
		EmailData: EmailData{
			Subject: "Password Reset Request",
			Recipient: Recipient{
				Email: "test@example.com",
			},
			Config: Config{
				CompanyName: "Test Company",
			},
		},
	}

	email, err := PasswordResetRequest(data)
	require.NoError(t, err)
	require.NotNil(t, email)

	assert.Equal(t, "Test Company Password Reset - Action Required", email.Subject)
}

func TestPasswordResetSuccess(t *testing.T) {
	data := ResetSuccessData{
		EmailData: EmailData{
			Subject: "Password Reset Success",
			Recipient: Recipient{
				Email: "test@example.com",
			},
			Config: Config{
				CompanyName: "Test Company",
			},
		},
	}

	email, err := PasswordResetSuccess(data)
	require.NoError(t, err)
	require.NotNil(t, email)

	assert.Equal(t, "Test Company Password Reset Confirmation", email.Subject)
}
