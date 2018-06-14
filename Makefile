.PHONY: run modify-scopes

PORT ?= 8080
CLIENT_ID ?=
CLIENT_SECRET ?=
SCOPES ?= "offline accounts:read"
TYPEFORM_URL ?= "https://api.typeform.com/oauth"

run:
	HTTP_PORT=$(PORT) \
	TYPEFORM_URL=$(TYPEFORM_URL) \
	CLIENT_ID=$(CLIENT_ID) \
	CLIENT_SECRET=$(CLIENT_SECRET) \
	REDIRECT_URI=http://localhost:$(PORT)/callback \
	SCOPES=$(SCOPES) \
	go run *.go

modify-scopes:
	curl -X PUT http://localhost:8080/modify-scopes?scopes=$(SCOPES)

