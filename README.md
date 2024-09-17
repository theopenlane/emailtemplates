# Email Templates

This repository contains some of the common emails which would be sent to users
of the platform.

## Required Variables For All Templates

| Variable          | Example                                      |
| ----------------- | -------------------------------------------- |
| `.CompanyName`    | `Openlane`                                   |
| `.CompanyAddress` | `1337 Main St. &middot;Metropolis, NY 10010` |
| `.Corporation`    | `theopenlane, Inc.`                          |
| `.SupportEmail`   | `support@theopenlane.io`                     |
| `.URLS.Root`      | `https://theopenlane.io`                     |
| `.URLS.Product`   | `https://console.theopenlane.io`             |
| `.URLS.Docs`      | `https://docs.theopenlane.io`                |

## Additional Variables for Specific Templates

| Template Name           | Variable                 | Example Value                                   |
| ----------------------- | ------------------------ | ----------------------------------------------- |
| Email Verification      | `.URLS.Verify`           | `https://console.theopenlane.io/verify`         |
| Subscriber Verification | `.URLS.VerifySubscriber` | `https://theopenlane.io/verify`                 |
| Password Reset          | `.URLS.PasswordReset`    | `https://console.theopenlane.io/password-reset` |
| Invite Acceptance       | `.URLS.Invite`           | `https://console.theopenlane.io/invite`         |

## Editing

These are the actual emails, language, format, that will be sent to users of
Openlane platform so please exercise care with their updates. If you're
uncertain, feel free to reach out to @matoszz for assistance.
