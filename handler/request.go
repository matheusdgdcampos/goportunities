package handler

import (
	"fmt"
	"goportunities/util"
)

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Role == "" {
		return util.ErroParamIsMissing("role", "string")
	}

	if r.Company == "" {
		return util.ErroParamIsMissing("company", "string")
	}

	if r.Location == "" {
		return util.ErroParamIsMissing("location", "string")
	}

	if r.Link == "" {
		return util.ErroParamIsMissing("link", "string")
	}

	if r.Remote == nil {
		return util.ErroParamIsMissing("remote", "bool")
	}

	if r.Salary <= 0 {
		return util.ErroParamIsMissing("salary", "int64")

	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one valid must be provided")
}
