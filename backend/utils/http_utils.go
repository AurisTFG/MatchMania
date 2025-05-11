package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpRequest[T any](client *http.Client, req *http.Request, body any) (*T, error) {
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encoding request body: %w", err)
		}
		req.Body = io.NopCloser(bytes.NewReader(jsonBody))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad response: %s\n%s", resp.Status, string(respBytes))
	}

	var decoded T
	if err = json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return &decoded, nil
}

func GetBasicAuthHeader(email, password string) string {
	if email == "" || password == "" {
		panic("Email or password is empty")
	}

	credentials := email + ":" + password
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))
	basicAuthHeader := "Basic " + encodedCredentials

	return basicAuthHeader
}

func GetUbisoftAuthHeader(ubisoftTicket string) string {
	if ubisoftTicket == "" {
		panic("Ticket is empty")
	}

	return "ubi_v1 t=" + ubisoftTicket
}

func GetNadeoAuthHeader(nadeoToken string) string {
	if nadeoToken == "" {
		panic("Token is empty")
	}

	return "nadeo_v1 t=" + nadeoToken
}
