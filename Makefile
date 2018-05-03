.PHONY: run

CLIENT_ID ?=
CLIENT_SECRET ?=
SCOPES ?= accounts:read

run:
	HTTP_PORT=8080 \
	TYPEFORM_URL=https://api.typeform.com/oauth \
	CLIENT_ID=$(CLIENT_ID) \
	CLIENT_SECRET=$(CLIENT_SECRET) \
	REDIRECT_URI=http://localhost:8080/callback \
	go run *.go

modify-scopes:
	curl -X PUT http://localhost:8080/modify-scopes?scopes=$(SCOPES)

