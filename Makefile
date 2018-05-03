.PHONY: run

CLIENT_ID ?=
CLIENT_SECRET ?=

run:
	HTTP_PORT=8080 \
	TYPEFORM_URL=https://api.typeform.com/oauth \
	CLIENT_ID=$(CLIENT_ID) \
	CLIENT_SECRET=$(CLIENT_SECRET) \
	REDIRECT_URI=http://localhost:8080/callback \
	go run *.go

