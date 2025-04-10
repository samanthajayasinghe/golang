package main

type PolicyStatement struct {
	Sid       string            `json:"Sid"`        // Statement ID
	Effect    string            `json:"Effect"`     // Allow or Deny
	Action    []string          `json:"Action"`     // allowed or denied action
	Principal map[string]string `json:",omitempty"` // principal that is allowed or denied
	Resource  *string           `json:",omitempty"` // object or objects that the statement covers
}

func NewPolicyStatement(sid string, affect string, action []string) PolicyStatement {
	return PolicyStatement{
		Sid:    sid,
		Effect: affect,
		Action: action,
	}
}

func (ps PolicyStatement) AddResource(resource *string) PolicyStatement {
	ps.Resource = resource
	return ps
}
