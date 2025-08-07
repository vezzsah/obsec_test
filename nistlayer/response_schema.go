package nistlayer

type ShortCVEResponse struct {
	Vulnerabilities []struct {
		Cve struct {
			ID           string `json:"id"`
			Descriptions []struct {
				Lang  string `json:"lang"`
				Value string `json:"value"`
			} `json:"descriptions"`
		} `json:"cve"`
	} `json:"vulnerabilities"`
}

type CVEResponse struct {
	ResultsPerPage  int    `json:"resultsPerPage"`
	StartIndex      int    `json:"startIndex"`
	TotalResults    int    `json:"totalResults"`
	Format          string `json:"format"`
	Version         string `json:"version"`
	Timestamp       string `json:"timestamp"`
	Vulnerabilities []struct {
		Cve struct {
			ID               string   `json:"id"`
			SourceIdentifier string   `json:"sourceIdentifier"`
			Published        string   `json:"published"`
			LastModified     string   `json:"lastModified"`
			VulnStatus       string   `json:"vulnStatus"`
			CveTags          []string `json:"cveTags"`
			Descriptions     []struct {
				Lang  string `json:"lang"`
				Value string `json:"value"`
			} `json:"descriptions"`
			Metrics struct {
				CvssMetricV40 []struct {
					Source   string `json:"source"`
					Type     string `json:"type"`
					CvssData struct {
						Version                           string  `json:"version"`
						VectorString                      string  `json:"vectorString"`
						BaseScore                         float64 `json:"baseScore"`
						BaseSeverity                      string  `json:"baseSeverity"`
						AttackVector                      string  `json:"attackVector"`
						AttackComplexity                  string  `json:"attackComplexity"`
						AttackRequirements                string  `json:"attackRequirements"`
						PrivilegesRequired                string  `json:"privilegesRequired"`
						UserInteraction                   string  `json:"userInteraction"`
						VulnConfidentialityImpact         string  `json:"vulnConfidentialityImpact"`
						VulnIntegrityImpact               string  `json:"vulnIntegrityImpact"`
						VulnAvailabilityImpact            string  `json:"vulnAvailabilityImpact"`
						SubConfidentialityImpact          string  `json:"subConfidentialityImpact"`
						SubIntegrityImpact                string  `json:"subIntegrityImpact"`
						SubAvailabilityImpact             string  `json:"subAvailabilityImpact"`
						ExploitMaturity                   string  `json:"exploitMaturity"`
						ConfidentialityRequirement        string  `json:"confidentialityRequirement"`
						IntegrityRequirement              string  `json:"integrityRequirement"`
						AvailabilityRequirement           string  `json:"availabilityRequirement"`
						ModifiedAttackVector              string  `json:"modifiedAttackVector"`
						ModifiedAttackComplexity          string  `json:"modifiedAttackComplexity"`
						ModifiedAttackRequirements        string  `json:"modifiedAttackRequirements"`
						ModifiedPrivilegesRequired        string  `json:"modifiedPrivilegesRequired"`
						ModifiedUserInteraction           string  `json:"modifiedUserInteraction"`
						ModifiedVulnConfidentialityImpact string  `json:"modifiedVulnConfidentialityImpact"`
						ModifiedVulnIntegrityImpact       string  `json:"modifiedVulnIntegrityImpact"`
						ModifiedVulnAvailabilityImpact    string  `json:"modifiedVulnAvailabilityImpact"`
						ModifiedSubConfidentialityImpact  string  `json:"modifiedSubConfidentialityImpact"`
						ModifiedSubIntegrityImpact        string  `json:"modifiedSubIntegrityImpact"`
						ModifiedSubAvailabilityImpact     string  `json:"modifiedSubAvailabilityImpact"`
						Safety                            string  `json:"Safety"`
						Automatable                       string  `json:"Automatable"`
						Recovery                          string  `json:"Recovery"`
						ValueDensity                      string  `json:"valueDensity"`
						VulnerabilityResponseEffort       string  `json:"vulnerabilityResponseEffort"`
						ProviderUrgency                   string  `json:"providerUrgency"`
					} `json:"cvssData"`
				} `json:"cvssMetricV40"`
				CvssMetricV31 []struct {
					Source   string `json:"source"`
					Type     string `json:"type"`
					CvssData struct {
						Version               string  `json:"version"`
						VectorString          string  `json:"vectorString"`
						BaseScore             float64 `json:"baseScore"`
						BaseSeverity          string  `json:"baseSeverity"`
						AttackVector          string  `json:"attackVector"`
						AttackComplexity      string  `json:"attackComplexity"`
						PrivilegesRequired    string  `json:"privilegesRequired"`
						UserInteraction       string  `json:"userInteraction"`
						Scope                 string  `json:"scope"`
						ConfidentialityImpact string  `json:"confidentialityImpact"`
						IntegrityImpact       string  `json:"integrityImpact"`
						AvailabilityImpact    string  `json:"availabilityImpact"`
					} `json:"cvssData"`
					ExploitabilityScore float64 `json:"exploitabilityScore"`
					ImpactScore         float64 `json:"impactScore"`
				} `json:"cvssMetricV31"`
			} `json:"metrics"`
			Weaknesses []struct {
				Source      string `json:"source"`
				Type        string `json:"type"`
				Description []struct {
					Lang  string `json:"lang"`
					Value string `json:"value"`
				} `json:"description"`
			} `json:"weaknesses"`
			Configurations []struct {
				Operator string `json:"operator"`
				Nodes    []struct {
					Operator string `json:"operator"`
					Negate   bool   `json:"negate"`
					CpeMatch []struct {
						Vulnerable      bool   `json:"vulnerable"`
						Criteria        string `json:"criteria"`
						MatchCriteriaID string `json:"matchCriteriaId"`
					} `json:"cpeMatch"`
				} `json:"nodes"`
			} `json:"configurations"`
			References []struct {
				URL    string   `json:"url"`
				Source string   `json:"source"`
				Tags   []string `json:"tags,omitempty"`
			} `json:"references"`
		} `json:"cve"`
	} `json:"vulnerabilities"`
}
