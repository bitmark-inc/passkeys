package passkey

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/gofrs/uuid"
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

type Credential struct {
	ID              string     `json:"id"`
	Name            *string    `json:"name,omitempty"`
	PublicKey       string     `json:"public_key"`
	AttestationType string     `json:"attestation_type"`
	AAGUID          uuid.UUID  `json:"aaguid"`
	LastUsedAt      *time.Time `json:"last_used_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	Transports      []string   `json:"transports"`
	BackupEligible  bool       `json:"backup_eligible"`
	BackupState     bool       `json:"backup_state"`
	IsMFA           bool       `json:"is_mfa"`
}

func (c *Client) GetCredential(ctx context.Context, userID string) (*Credential, error) {
	var resp []Credential
	if err := c.requester.do(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/%s/credentials?user_id=%s", c.tenantID, userID),
		nil,
		&resp); err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, nil
	}
	if len(resp) > 1 {
		return nil, fmt.Errorf("multiple credentials found for user %s", userID)
	}

	return &resp[0], nil
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
