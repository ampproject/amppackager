package apiTypes

import (
	"fmt"
	"strings"
)

type Subaccnt struct {
	Active      bool     `json:"active" mapstructure:"active"`
	Domain      string   `json:"domain" mapstructure:"domain"`
	Ip          string   `json:"ip" mapstructure:"ip"`
	ProjectId   int64    `json:"project_id" mapstructure:"project_id"`
	ProjectName string   `json:"project_name" mapstructure:"project_name"`
	RegionId    int      `json:"region_id" mapstructure:"region_id"`
	Status      string   `json:"status" mapstructure:"status"`
	Type        string   `json:"type" mapstructure:"type"`
	UniqId      string   `json:"uniq_id" mapstructure:"uniq_id"`
	Username    string   `json:"username" mapstructure:"username"`
	Categories  []string `json:"categories" mapstructure:"categories"`
}

func (x Subaccnt) String() string {
	var slice []string

	slice = append(slice, fmt.Sprintf("Domain: %s UniqId: %s\n", x.Domain, x.UniqId))

	if len(x.Categories) > 0 {
		slice = append(slice, fmt.Sprintln("\tCategories:"))
		for _, category := range x.Categories {
			slice = append(slice, fmt.Sprintf("\t\t* %s\n", category))
		}
	}

	if x.Ip != "" && x.Ip != "127.0.0.1" {
		slice = append(slice, fmt.Sprintf("\tIp: %s\n", x.Ip))
	}

	if x.ProjectName != "" && x.ProjectId != 0 {
		slice = append(slice, fmt.Sprintf("\tProjectName: %s (id %d)\n", x.ProjectName, x.ProjectId))
	}
	slice = append(slice, fmt.Sprintf("\tRegionId: %d\n", x.RegionId))
	slice = append(slice, fmt.Sprintf("\tStatus: %s\n", x.Status))
	slice = append(slice, fmt.Sprintf("\tType: %s\n", x.Type))

	return strings.Join(slice[:], "")
}
