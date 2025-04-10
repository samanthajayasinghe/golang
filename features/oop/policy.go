package main

type PolicyDocument struct {
	Version   string            `json:"Version"`
	Statement []PolicyStatement `json:"Statement"`
}

func NewPolicyDocument(version string, statements []PolicyStatement) PolicyDocument {
	return PolicyDocument{
		Version:   version,
		Statement: statements,
	}
}
