package handlers

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateProjectParams struct {
	Name     string `json:"name"`
	CpesData []struct {
		CpeId string `json:"cpe_id"`
	} `json:"cpes_data,omitempty"`
}

type RegisterCPEParams struct {
	ProjectName string `json:"project_name"`
	CPEData     []CPE  `json:"cpe_data"`
}

type CPE struct {
	Part     string `json:"part"`
	Vendor   string `json:"vendor"`
	Product  string `json:"product"`
	Version  string `json:"version"`
	Update   string `json:"update,omitempty"`
	Edition  string `json:"edition,omitempty"`
	Language string `json:"language,omitempty"`
}

type ResolveProjectCVEParams struct {
	ProjectName string `json:"project_name"`
	CVE         string `json:"cve"`
	CPE         string `json:"cpe"`
}
