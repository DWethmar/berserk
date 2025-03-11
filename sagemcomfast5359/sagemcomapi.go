package sagemcomfast5359

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
)

// Client represents a Sagemcom Fast 5359 router client.
type Client struct {
	Host      string
	client    *http.Client
	Username  string
	contextID string
	cookies   []*http.Cookie
}

// NewClient initializes a new router client.
func NewClient(client *http.Client, host string) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{
		Host:   host,
		client: client,
	}
}

// createContextResponse represents the response from createContext.
type createContextResponse struct {
	Status int `json:"status"`
	Data   struct {
		ContextID string `json:"contextID"`
		Username  string `json:"username"`
		Groups    string `json:"groups"`
	} `json:"data"`
}

// Login authenticates with the router and stores session details.
func (c *Client) Login(ctx context.Context, username, password string) error {
	action := Action{
		Service: "sah.Device.Information",
		Method:  "createContext",
		Parameters: map[string]string{
			"applicationName": "webui",
			"username":        username,
			"password":        password,
		},
	}

	resp, err := DoAction(ctx, ActionOptions{
		Client: c.client,
		Host:   c.Host,
		Action: action,
		Header: http.Header{
			"Authorization": {"X-Sah-Login"},
		},
		Cookies: []*http.Cookie{
			{Name: "28b93ef6/accept-language", Value: "nl,en"},
		},
		logger: slog.Default(),
	})
	if err != nil {
		return fmt.Errorf("error logging in: %w", err)
	}
	defer resp.Body.Close()

	var r createContextResponse
	if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	c.contextID = r.Data.ContextID
	c.Username = r.Data.Username

	// extract session cookie
	setCookies := resp.Header.Get("Set-Cookie")
	re := regexp.MustCompile(`([^;=\s]+)=([^;]+)`)
	matches := re.FindAllStringSubmatch(setCookies, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			cookieName := match[1]
			cookieValue := match[2]
			cookieNameRegex := regexp.MustCompile(`(?:[^/]+/)?sessid`)
			if !cookieNameRegex.MatchString(cookieName) {
				continue
			}
			c.cookies = append(c.cookies, &http.Cookie{
				Name:  cookieName,
				Value: cookieValue,
			})
			break
		}
	}
	return nil
}

// Logout terminates the current session.
func (c *Client) Logout(ctx context.Context) error {
	action := Action{
		Service: "sah.Device.Information",
		Method:  "releaseContext",
		Parameters: map[string]string{
			"applicationName": "webui",
		},
	}
	resp, err := DoAction(ctx, ActionOptions{
		Client:  c.client,
		Host:    c.Host,
		Action:  action,
		Cookies: c.cookies,
		logger:  slog.Default(),
	})
	if err != nil {
		return fmt.Errorf("error logging out: %w", err)
	}
	defer resp.Body.Close()
	return nil
}

// LanDevices retrieves all connected devices on the LAN.
func (c *Client) LanDevices(ctx context.Context) ([]Device, error) {
	action := Action{
		Service: "Devices.Device.lan",
		Method:  "topology",
		Parameters: map[string]string{
			"expression": "not logical",
			"flags":      "no_recurse|no_actions",
		},
	}
	header := http.Header{
		"Authorization": {fmt.Sprintf("X-Sah %s", c.contextID)},
		"X-Context":     {c.contextID},
	}
	resp, err := DoAction(ctx, ActionOptions{
		Client:  c.client,
		Host:    c.Host,
		Action:  action,
		Header:  header,
		Cookies: c.cookies,
		logger:  slog.Default(),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting devices: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting devices: %s", resp.Status)
	}
	var topologyResponse TopologyResponse
	if err = json.NewDecoder(resp.Body).Decode(&topologyResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	if len(topologyResponse.Errors) > 0 {
		errMsg := ""
		for _, e := range topologyResponse.Errors {
			errMsg += fmt.Sprintf("%d: %s\n", e.Error, e.Description)
		}
		return nil, fmt.Errorf("error getting devices: %s", errMsg)
	}
	var devices []Device
	for _, s := range topologyResponse.Status {
		for _, a := range s.AccessPoints {
			devices = append(devices, a.Devices...)
		}
	}
	return devices, nil
}
