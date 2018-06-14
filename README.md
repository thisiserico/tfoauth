# tfoauth
> test [typeform's oauth flow][typeform-oauth] locally

## How to run

First thing you need to do is head over [typeform's applications page][typeform-application]
and register your application:

```
Name:         tfoauth running locally
Description:  testing the oauth integration with tfoauth
Homepage URL: http://localhost:8080
Callback URL: http://localhost:8080/callback
Logo URL:     http://localhost:8080/logo.png
```

After saving it, you'll get your `client_id` and `client_secret`.
You're ready to test the flow:

```
CLIENT_ID={your client id} CLIENT_SECRET={your client secret} make run
```

Head over to [`localhost:8080`][localhost] and follow further instructions.
If everything was set up correctly, you'll get a response with your access token :)

## Working with scopes

Typeform implements [oauth scopes][typeform-scopes] to grant access to certain resources.
By default, the ones being used are `offline` and `accounts:read`.
If you're interested in other scopes, use another environment variable
when running `make run`.

```
... SCOPES="accounts:read forms:read forms:write responses:read" make run
```

[typeform-oauth]: https://developer.typeform.com/get-started/applications
[typeform-scopes]: https://developer.typeform.com/get-started/scopes
[typeform-application]: https://admin.typeform.com/account#/section/apps
[localhost]: http://localhost:8080

