# Email Templates

This repository contains some of the common emails which would be sent to users
of the Openlane platform.

## Example Usage

```go
    // setup config using options
   	config, err := emailtemplates.New(
		emailtemplates.WithCompanyName("Avengers"),
		emailtemplates.WithCompanyAddress("1337 Main St. &middot;Metropolis, NY 10010"),
		emailtemplates.WithCorporation("avengers, Inc."),
		emailtemplates.WithSupportEmail("support@avengers.com"),
		emailtemplates.WithFromEmail("no-reply@mail.avengers.com"),
		emailtemplates.WithRootDomain("https://www.avengers.com"),
		emailtemplates.WithDocsDomain("https://docs.avengers.com"),
		emailtemplates.WithProductDomain("https://console.avengers.com"),
	)
    if err != nil {
        return err
    }

    // create email message
	email, err := config.NewWelcomeEmail(
		emailtemplates.Recipient{
			Email:     "ironman@example.com",
			FirstName: "Tony",
			LastName:  "Stark",
		}, "Avengers",
	)
    if err != nil {
        return err
    }

    // we recommend the use of the https://github.com/theopenlane/newman package to send the email
    // which supports several providers including Resend, Mailgun, etc.
    // for brevity, this won't show how to create the client and assumes it was created beforehand
    if err := newmanClient.SendEmailWithContext(ctx, email); err != nil {
        return err
    }
```

## Variables

### Required Variables For All Templates

| Variable          | Example                                      |
| ----------------- | -------------------------------------------- |
| `.CompanyName`    | `Openlane`                                   |
| `.CompanyAddress` | `1337 Main St. &middot;Metropolis, NY 10010` |
| `.Corporation`    | `theopenlane, Inc.`                          |
| `.SupportEmail`   | `support@theopenlane.io`                     |
| `.FromEmail`      | `no-reply@mail.theopenlane.io`               |
| `.URLS.Root`      | `https://theopenlane.io`                     |
| `.URLS.Product`   | `https://console.theopenlane.io`             |
| `.URLS.Docs`      | `https://docs.theopenlane.io`                |

### Additional Variables for Specific Templates

| Template Name           | Variable                 | Example Value                                   |
| ----------------------- | ------------------------ | ----------------------------------------------- |
| Email Verification      | `.URLS.Verify`           | `https://console.theopenlane.io/verify`         |
| Subscriber Verification | `.URLS.VerifySubscriber` | `https://theopenlane.io/verify`                 |
| Password Reset          | `.URLS.PasswordReset`    | `https://console.theopenlane.io/password-reset` |
| Invite Acceptance       | `.URLS.Invite`           | `https://console.theopenlane.io/invite`         |

### Optional Variables

| Variable   | Example                                  |
| ---------- | ---------------------------------------- |
| `.LogoURL` | `http://api.example.com/assets/logo.png` |

## Editing

These are the actual emails, language, format, that will be sent to users of
Openlane platform so please exercise care with their updates. If you're
uncertain, feel free to reach out to @matoszz for assistance.

## Contributing

See the [contributing](.github/CONTRIBUTING.md) guide for more information
