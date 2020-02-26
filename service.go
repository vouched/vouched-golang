package client

import (
	"context"
)

// APIClient service structure
type APIClient struct {
	Key string
}

// JobResponse struct for job
type JobResponse struct {
	ID      string
	Status  string
	Review  string
	ReviewPass	string
	ReviewedAt  string
	Logs struct {
		ReviewedEntities  string
		CreatedAt  string
		FirstName  string
		LastName  string
	}
	Request struct {
		Type        string
		CallbackURL string
		Properties  []struct {
			Name  string
			Value string
		}
		Parameters struct {
			IDPhoto   string
			UserPhoto string
			FirstName string
			LastName  string
			Dob       string
		}
	}
	Errors []struct {
		Type       string
		Message    string
		Suggestion string
	}
	submitted string
	Result    struct {
		Success     bool
		Type        string
		Country     string
		State       string
		ID          string
		FirstName   string
		LastName    string
		MiddleName  string
		BirthDate   string
		ExpireDate  string
		Confidences struct {
			ID        float64
			BackID    float64
			Selfie    float64
			IDMatch   float64
			FaceMatch float64
		}
	}
}

// JobsResponse struct for jobs
type JobsResponse struct {
	Jobs       []JobResponse `json:"items"`
	TotalPages int
	PageSize   int
	Page       int
	Total      int
}

// New create a new client
func New(key string) *APIClient {
	return &APIClient{Key: key}
}
func setHeader(key string, req *Request) {
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("X-Api-Key", key)
}

func (c APIClient) toResponse(query string, resp interface{}, params map[string]interface{}) error {
	config := NewConfig()
	req := NewRequest(query)
	setHeader(c.Key, req)
	ctx := context.Background()

	for k, v := range params {
		req.Var(k, v)
	}

	client := NewClient(config.URL)
	if err := client.Run(ctx, req, &resp); err != nil {
		return err
	}
	return nil
}

// RemoveJob remove a job
func (c APIClient) RemoveJob(id string) (*JobResponse, error) {
	var resp struct {
		RemoveJob JobResponse
	}
	params := map[string]interface{}{
		"id": id}

	err := c.toResponse(RemoveMutation, &resp, params)
	if err != nil {
		return nil, err
	}
	return &resp.RemoveJob, nil
}

// Jobs query jobs
func (c APIClient) Jobs(params map[string]interface{}) (*JobsResponse, error) {
	var resp struct {
		Jobs JobsResponse
	}
	err := c.toResponse(JobsQuery, &resp, params)
	if err != nil {
		return nil, err
	}
	return &resp.Jobs, nil
}

// UpdateSecretClientKeyResponse updateSecretClientKey response
type UpdateSecretClientKeyResponse struct {
	SecretClientKey string
}

// UpdateSecretClientKey update the secret key
func (c APIClient) UpdateSecretClientKey(params map[string]interface{}) (*UpdateSecretClientKeyResponse, error) {
	var resp struct {
		UpdateSecretClientKey UpdateSecretClientKeyResponse
	}
	err := c.toResponse(UpdateSecretClientKeyMutation, &resp, params)
	if err != nil {
		return nil, err
	}
	return &resp.UpdateSecretClientKey, nil
}

// Submit submit a job
func (c APIClient) Submit(params map[string]interface{}) (*JobResponse, error) {
	queryParams := map[string]interface{}{
		"type":        params["type"],
		"properties":  params["properties"],
		"callbackURL": params["callbackURL"],
		"params": map[string]interface{}{
			"userPhoto": params["userPhoto"],
			"idPhoto":   params["idPhoto"],
			"firstName": params["firstName"],
			"dob":       params["dob"],
			"lastName":  params["lastName"]}}

	var resp struct {
		SubmitJob JobResponse
	}
	err := c.toResponse(SubmitMutation, &resp, queryParams)
	if err != nil {
		return nil, err
	}
	return &resp.SubmitJob, nil
}
