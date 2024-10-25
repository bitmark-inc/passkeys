package passkey

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
)

type InitRegistrationRequest struct {
	UserId      string  `json:"user_id"`
	Username    string  `json:"username"`
	DisplayName *string `json:"display_name"`
	Icon        *string `json:"icon"`
}

type InitLoginRequest struct {
	UserId *string `json:"user_id"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (c *Client) InitRegistration(
	ctx context.Context,
	req InitRegistrationRequest) (*protocol.CredentialCreation, error) {
	var resp protocol.CredentialCreation
	if err := c.requester.do(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/%s/registration/initialize", c.tenantID),
		req,
		&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) FinalizeRegistration(
	ctx context.Context,
	req protocol.CredentialCreationResponse) (string, error) {
	var resp TokenResponse
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
	req InitLoginRequest) (*protocol.CredentialAssertion, error) {
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
	var resp TokenResponse
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
