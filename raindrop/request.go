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
	client     *Client
	method     string
	path       string
	queryParam url.Values
	body       interface{}
}

// newRequest initialise un RequestBuilder.
func (c *Client) newRequest(method, path string, body interface{}) *RequestBuilder {
	return &RequestBuilder{
		client:     c,
		method:     method,
		path:       path,
		queryParam: url.Values{},
		body:       body,
	}
}

// Param ajoute un paramètre de requête.
func (r *RequestBuilder) Param(key, value string) *RequestBuilder {
	r.queryParam.Add(key, value)
	return r
}

// Do exécute la requête HTTP et décode la réponse JSON dans `out` si fourni.
func (r *RequestBuilder) Do(ctx context.Context, out interface{}) error {
	u := fmt.Sprintf("%s%s", r.client.BaseURL, r.path)
	if len(r.queryParam) > 0 {
		u += "?" + r.queryParam.Encode()
	}

	var body io.Reader
	if r.body != nil {
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		if err := encoder.Encode(r.body); err != nil {
			return fmt.Errorf("failed to encode request body: %w", err)
		}

		// Vérification que le JSON est valide
		var js json.RawMessage
		if err := json.Unmarshal(buf.Bytes(), &js); err != nil {
			return fmt.Errorf("invalid JSON request body: %w", err)
		}

		body = &buf

		if r.client.Debug {
			prettyJSON, err := json.MarshalIndent(r.body, "", "  ")
			if err != nil {
				fmt.Printf("[DEBUG] Failed to pretty-print JSON: %v\n", err)
			} else {
				fmt.Printf("[DEBUG] Request JSON body:\n%s\n", string(prettyJSON))
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, r.method, u, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+r.client.Token)
	req.Header.Set("Content-Type", "application/json")

	if r.client.Debug {
		fmt.Printf("[DEBUG] Sending %s request to %s\n", r.method, u)
	}

	resp, err := r.client.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if r.client.Debug {
		fmt.Printf("[DEBUG] Response status: %s\n", resp.Status)
	}

	if resp.StatusCode != http.StatusOK {
		responseBody, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("request failed: %d, unable to read response body: %w", resp.StatusCode, readErr)
		}
		if r.client.Debug {
			fmt.Printf("[DEBUG] Response body: %s\n", string(responseBody))
			fmt.Printf("[DEBUG] Request headers: %+v\n", req.Header)
		}
		return fmt.Errorf("request failed: status code %d, response: %s", resp.StatusCode, string(responseBody))
	}

	if out == nil {
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("failed to decode response JSON: %w", err)
	}

	return nil
}
