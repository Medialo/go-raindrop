package raindrop

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// RequestBuilder construit et exécute les requêtes HTTP.
type RequestBuilder struct {
	client *Client
	method string
	path   string
	query  url.Values
	body   interface{}
}

// newRequest initialise un RequestBuilder.
func (c *Client) newRequest(method, path string, body interface{}) *RequestBuilder {
	return &RequestBuilder{
		client: c,
		method: method,
		path:   path,
		query:  url.Values{},
		body:   body,
	}
}

// Param ajoute un paramètre de requête.
func (r *RequestBuilder) Param(key, value string) *RequestBuilder {
	r.query.Add(key, value)
	return r
}

// Do exécute la requête et décode la réponse dans out.
func (r *RequestBuilder) Do(ctx context.Context, out interface{}) error {
	u := fmt.Sprintf("%s%s", r.client.BaseURL, r.path)
	if len(r.query) > 0 {
		u += "?" + r.query.Encode()
	}
	var body io.Reader
	if r.body != nil {
		buf := &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(r.body); err != nil {
			return err
		}
		body = buf
	}
	req, _ := http.NewRequestWithContext(ctx, r.method, u, body)
	req.Header.Set("Authorization", "Bearer "+r.client.Token)
	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("Sending %s request to %s\n", r.method, u)
	if body != nil {
		fmt.Printf("Request body: %v\n", r.body)
	}

	resp, err := r.client.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Response status: %s\n", resp.Status)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request headers sent: %v\n", req.Header)
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("request failed: %d, failed to read response body: %v", resp.StatusCode, err)
		}
		return fmt.Errorf("request failed: %d, response body: %s", resp.StatusCode, string(responseBody))
	}

	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
