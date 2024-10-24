package passkey

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/teamhanko/passkey-server/api/dto/request"
	"github.com/teamhanko/passkey-server/api/dto/response"
)

func (c *Client) InitRegistration(
	ctx context.Context,
	req request.InitRegistrationDto) (*protocol.CredentialCreation, error) {
	var resp protocol.CredentialCreation
	if err := c.requester.do(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/%s/registration/initialize", c.tenantID),
		req,
		&resp); err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *Client) FinalizeRegistration(
	ctx context.Context,
	req protocol.CredentialCreationResponse) (string, error) {
	var resp response.TokenDto
	if err := c.requester.do(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/%s/registration/finalize", c.tenantID),
		req,
		&resp); err != nil {
		return "", err
	}
	return resp.Token, nil
}

func (c *Client) InitLogin(
	ctx context.Context,
	req request.InitLoginDto) (*protocol.CredentialAssertion, error) {
	var resp protocol.CredentialAssertion
	if err := c.requester.do(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/%s/login/initialize", c.tenantID),
		req,
		&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) FinalizeLogin(
	ctx context.Context,
	req protocol.CredentialAssertionResponse) (string, error) {
	var resp response.TokenDto
	if err := c.requester.do(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/%s/login/finalize", c.tenantID),
		req,
		&resp); err != nil {
		return "", err
	}
	return resp.Token, nil
}
