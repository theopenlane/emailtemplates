package emailtemplates

import (
	"errors"
	"fmt"
	"net/mail"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

var (
	ErrInvalidSenderEmail = errors.New("please provide a valid sender email ( from email )")
	templateLoadOnce      sync.Once
	templateLoadErr       error
)

// ensureCustomTemplatesLoaded ensures templates are loaded only once
// Also this makes sure if we have a custom template path, we should load them
func ensureCustomTemplatesLoaded(templatePath string) error {
	templateLoadOnce.Do(func() {
		if templatePath != defaultTemplatesDir && templatePath != "" {
			templateFiles, err := os.ReadDir(templatePath)
			if err != nil {
				templateLoadErr = fmt.Errorf("could not read template files from %q: %w", templatePath, err)
				return
			}

			for _, file := range templateFiles {
				if file.IsDir() {
					continue
				}

				pattern := filepath.Join(templatePath, file.Name())

				tmpl, err := template.New(file.Name()).
					Funcs(fm).
					ParseFiles(pattern)

				if err != nil {
					templateLoadErr = fmt.Errorf("could not parse template %q: %w", file.Name(), err)
					return
				}

				templates[file.Name()] = tmpl
			}
		}
	})

	return templateLoadErr
}

// validate checks if all required fields are set and valid
func (c *Config) validate() error {
	if c.TemplatesPath != defaultTemplatesDir && c.TemplatesPath != "" {
		if err := ensureCustomTemplatesLoaded(c.TemplatesPath); err != nil {
			return err
		}
	}

	if len(c.CompanyAddress) == 0 {
		return newMissingRequiredFieldError("company address")
	}

	if len(c.CompanyName) == 0 {
		return newMissingRequiredFieldError("company name")
	}

	if c.FromEmail == "" {
		return newMissingRequiredFieldError("sender email")
	}

	if _, err := mail.ParseAddress(c.FromEmail); err != nil {
		return ErrInvalidSenderEmail
	}

	return nil
}

// New is a function that creates a new config for the email templates
func New(options ...Option) (*Config, error) {
	// initialize the resendEmailSender
	c := &Config{
		TemplatesPath: defaultTemplatesDir,
	}

	// apply the options
	for _, option := range options {
		option(c)
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// Option is a function that sets a field on an EmailMessage
type Option func(*Config)

// WithCompanyName sets the company name for the email
func WithCompanyName(name string) Option {
	return func(t *Config) {
		t.CompanyName = name
	}
}

// WithCompanyAddress sets company address for the footer of the email
func WithCompanyAddress(address string) Option {
	return func(t *Config) {
		t.CompanyAddress = address
	}
}

// WithCorporation sets the corporation used in the footer of the email
func WithCorporation(corp string) Option {
	return func(t *Config) {
		t.Corporation = corp
	}
}

// WithRootDomain sets the root domain for the email
func WithRootDomain(domain string) Option {
	return func(t *Config) {
		t.URLS.Root = domain
	}
}

// WithProductDomain sets the product domain for the email
func WithProductDomain(domain string) Option {
	return func(t *Config) {
		t.URLS.Product = domain
	}
}

// WithDocsDomain sets the docs domain for the email
func WithDocsDomain(domain string) Option {
	return func(t *Config) {
		t.URLS.Docs = domain
	}
}

// WithFromEmail sets the from email for the email
func WithFromEmail(email string) Option {
	return func(t *Config) {
		t.FromEmail = email
	}
}

// WithSupportEmail sets the support email for the email
func WithSupportEmail(email string) Option {
	return func(t *Config) {
		t.SupportEmail = email
	}
}

// WithVerifyURL sets the verify URL for the email
func WithVerifyURL(url string) Option {
	return func(t *Config) {
		t.URLS.Verify = url
	}
}

// WithInviteURL sets the invite URL for the email
func WithInviteURL(url string) Option {
	return func(t *Config) {
		t.URLS.Invite = url
	}
}

// WithResetURL sets the reset URL for the email
func WithResetURL(url string) Option {
	return func(t *Config) {
		t.URLS.PasswordReset = url
	}
}

// WithVerifySubscriberURL sets the verify subscriber URL for the email
func WithVerifySubscriberURL(url string) Option {
	return func(t *Config) {
		t.URLS.VerifySubscriber = url
	}
}

// WithLogoURL sets the logo URL for the email, this field is optional and
// omitted from the email if not provided
func WithLogoURL(url string) Option {
	return func(t *Config) {
		t.LogoURL = url
	}
}

// WithTemplatesPath allows you configure the path to your templates
// else we will use the default templates
func WithTemplatesPath(p string) Option {
	return func(c *Config) {
		c.TemplatesPath = p
	}
}
