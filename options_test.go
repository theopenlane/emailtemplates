package emailtemplates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	t.Run("WithCompanyName", func(t *testing.T) {
		cfg := &Config{}
		opt := WithCompanyName("Test Company")
		opt(cfg)
		assert.Equal(t, "Test Company", cfg.CompanyName)
	})

	t.Run("WithCompanyAddress", func(t *testing.T) {
		cfg := &Config{}
		opt := WithCompanyAddress("123 Test St")
		opt(cfg)
		assert.Equal(t, "123 Test St", cfg.CompanyAddress)
	})

	t.Run("WithCorporation", func(t *testing.T) {
		cfg := &Config{}
		opt := WithCorporation("Test Corp")
		opt(cfg)
		assert.Equal(t, "Test Corp", cfg.Corporation)
	})

	t.Run("WithRootDomain", func(t *testing.T) {
		cfg := &Config{}
		opt := WithRootDomain("https://example.com")
		opt(cfg)
		assert.Equal(t, "https://example.com", cfg.URLS.Root)
	})

	t.Run("WithProductDomain", func(t *testing.T) {
		cfg := &Config{}
		opt := WithProductDomain("https://product.example.com")
		opt(cfg)
		assert.Equal(t, "https://product.example.com", cfg.URLS.Product)
	})

	t.Run("WithDocsDomain", func(t *testing.T) {
		cfg := &Config{}
		opt := WithDocsDomain("https://docs.theopenlane.io")
		opt(cfg)
		assert.Equal(t, "https://docs.theopenlane.io", cfg.URLS.Docs)
	})

	t.Run("WithFromEmail", func(t *testing.T) {
		cfg := &Config{}
		opt := WithFromEmail("test@test.com")
		opt(cfg)
		assert.Equal(t, "test@test.com", cfg.FromEmail)
	})

	t.Run("WithSupportEmail", func(t *testing.T) {
		cfg := &Config{}
		opt := WithSupportEmail("support@theopenlane.io")
		opt(cfg)
		assert.Equal(t, "support@theopenlane.io", cfg.SupportEmail)
	})

	t.Run("WithVerifyURL", func(t *testing.T) {
		cfg := &Config{}
		opt := WithVerifyURL("https://example.com/verify")
		opt(cfg)
		assert.Equal(t, "https://example.com/verify", cfg.URLS.Verify)
	})

	t.Run("WithInviteURL", func(t *testing.T) {
		cfg := &Config{}
		opt := WithInviteURL("https://example.com/invite")
		opt(cfg)
		assert.Equal(t, "https://example.com/invite", cfg.URLS.Invite)
	})

	t.Run("WithResetURL", func(t *testing.T) {
		cfg := &Config{}
		opt := WithResetURL("https://example.com/reset")
		opt(cfg)
		assert.Equal(t, "https://example.com/reset", cfg.URLS.PasswordReset)
	})

	t.Run("WithVerifySubscriberURL", func(t *testing.T) {
		cfg := &Config{}
		opt := WithVerifySubscriberURL("https://example.com/verify-subscriber")
		opt(cfg)
		assert.Equal(t, "https://example.com/verify-subscriber", cfg.URLS.VerifySubscriber)
	})

	t.Run("WithLogoURL", func(t *testing.T) {
		cfg := &Config{}
		opt := WithLogoURL("https://example.com/logo.png")
		opt(cfg)
		assert.Equal(t, "https://example.com/logo.png", cfg.LogoURL)
	})

	t.Run("WithTemplatesPath", func(t *testing.T) {
		cfg := &Config{}
		opt := WithTemplatesPath("./custom/templates")
		opt(cfg)
		assert.Equal(t, "./custom/templates", cfg.templatesPath)
	})
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		options []Option
		wantErr bool
		errMsg  string
	}{
		{
			name:    "missing templates path",
			options: []Option{WithTemplatesPath("")},
			wantErr: true,
			errMsg:  "please provide your templates path",
		},
		{
			name: "missing company address",
			options: []Option{
				WithTemplatesPath("./templates"),
				WithCompanyName("Test Company"),
				WithFromEmail("test@example.com"),
			},
			wantErr: true,
			errMsg:  "please provide your company's address",
		},
		{
			name: "missing company name",
			options: []Option{
				WithTemplatesPath("./templates"),
				WithCompanyAddress("123 Test St"),
				WithFromEmail("test@example.com"),
			},
			wantErr: true,
			errMsg:  "please provide your company's name",
		},
		{
			name: "missing from email",
			options: []Option{
				WithTemplatesPath("./templates"),
				WithCompanyAddress("123 Test St"),
				WithCompanyName("Test Company"),
			},
			wantErr: true,
			errMsg:  "please provide your sender email",
		},
		{
			name: "invalid from email",
			options: []Option{
				WithTemplatesPath("./templates"),
				WithCompanyAddress("123 Test St"),
				WithCompanyName("Test Company"),
				WithFromEmail("invalid-email"),
			},
			wantErr: true,
			errMsg:  "please provide a valid sender email ( from email )",
		},
		{
			name: "valid configuration",
			options: []Option{
				WithTemplatesPath("./templates"),
				WithCompanyAddress("123 Test St"),
				WithCompanyName("Test Company"),
				WithFromEmail("test@example.com"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := New(tt.options...)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.errMsg, err.Error())
				assert.Nil(t, cfg)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, cfg)
		})
	}
}
