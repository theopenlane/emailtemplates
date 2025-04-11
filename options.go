package emailtemplates

import (
	"errors"
	"fmt"
	"net/mail"
	"os"
	"path/filepath"
	"sync"

	"github.com/rs/zerolog/log"
)

var (
	ErrInvalidSenderEmail = errors.New("please provide a valid sender email ( from email )")
	templateLoadOnce      sync.Once
)

// New is a function that creates a new config for the email templates
func New(options ...Option) (*Config, error) {
	// initialize the email config
	c := &Config{}

	// apply the options
	for _, option := range options {
		option(c)
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// Validate the config and ensures that all required fields are set
// and that the templates are loaded correctly
// This function is called when the config is created outside the New function
func (c *Config) Validate() error {
	return c.validate()
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

// WithVerifyBillingURL sets the verify billing URL for the email
func WithVerifyBillingURL(url string) Option {
	return func(t *Config) {
		t.URLS.VerifyBilling = url
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

// ensureCustomTemplatesLoaded ensures templates are loaded only once
// Also this makes sure if we have a custom template path, we should load them
// this will include the partials directory as well if it exists
func ensureCustomTemplatesLoaded(templatePath string) (err error) {
	templateLoadOnce.Do(func() {
		partials, err = getPartials(templatePath)
		if err != nil {
			log.Fatal().Err(err).Msgf("could not load partials from %q", templatePath)
			return
		}

		err = loadTemplatesFromDir(templatePath, partials)
		if err != nil {
			log.Fatal().Err(err).Msgf("could not load templates from %q", templatePath)
			return
		}
	})

	return
}

func getPartials(path string) ([]string, error) {
	partials := []string{}

	templateFiles, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range templateFiles {
		if file.Name() == defaultPartialsDir {
			partialPath := filepath.Join(path, file.Name())
			if err := loadTemplatesFromDir(partialPath, partials); err != nil {
				return nil, err
			}

			templateFiles, err := os.ReadDir(partialPath)
			if err != nil {
				return nil, fmt.Errorf("could not read template files from %q: %w", path, err)
			}

			for _, file := range templateFiles {
				if file.IsDir() {
					continue
				}

				partials = append(partials, filepath.Join(defaultPartialsDir, file.Name()))
			}
		}
	}

	return partials, nil
}

// loadTemplatesFromDir loads templates from the specified directory
// and recursively loads partials
func loadTemplatesFromDir(path string, partials []string) error {
	templateFiles, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("could not read template files from %q: %w", path, err)
	}

	for _, file := range templateFiles {
		if file.IsDir() {
			continue
		}

		key := file.Name()
		// path := filepath.Join(path, key)

		templates[key], err = parseCustomTemplate(file, path, partials)
		if err != nil {
			return err
		}
	}

	return nil
}

// validate checks if all required fields are set and valid
func (c *Config) validate() error {
	if c.TemplatesPath != "" {
		if err := ensureCustomTemplatesLoaded(c.TemplatesPath); err != nil {
			return err
		}
	}

	if c.CompanyAddress == "" {
		return newMissingRequiredFieldError("company address")
	}

	if c.CompanyName == "" {
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
