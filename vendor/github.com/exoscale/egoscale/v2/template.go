package v2

import (
	"context"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	papi "github.com/exoscale/egoscale/v2/internal/public-api"
)

// Template represents a Compute instance template.
type Template struct {
	BootMode        *string
	Build           *string
	Checksum        *string `req-for:"create"`
	CreatedAt       *time.Time
	DefaultUser     *string
	Description     *string
	Family          *string
	ID              *string
	Name            *string `req-for:"create"`
	PasswordEnabled *bool   `req-for:"create"`
	SSHKeyEnabled   *bool   `req-for:"create"`
	Size            *int64
	URL             *string `req-for:"create"`
	Version         *string
	Visibility      *string
}

func templateFromAPI(t *papi.Template) *Template {
	return &Template{
		BootMode:        (*string)(t.BootMode),
		Build:           t.Build,
		Checksum:        t.Checksum,
		CreatedAt:       t.CreatedAt,
		DefaultUser:     t.DefaultUser,
		Description:     t.Description,
		Family:          t.Family,
		ID:              t.Id,
		Name:            t.Name,
		PasswordEnabled: t.PasswordEnabled,
		SSHKeyEnabled:   t.SshKeyEnabled,
		Size:            t.Size,
		URL:             t.Url,
		Version:         t.Version,
		Visibility:      (*string)(t.Visibility),
	}
}

// RegisterTemplate registers a new Template in the specified zone.
func (c *Client) RegisterTemplate(ctx context.Context, zone string, template *Template) (*Template, error) {
	if err := validateOperationParams(template, "create"); err != nil {
		return nil, err
	}

	resp, err := c.RegisterTemplateWithResponse(
		apiv2.WithZone(ctx, zone),
		papi.RegisterTemplateJSONRequestBody{
			BootMode:        (*papi.RegisterTemplateJSONBodyBootMode)(template.BootMode),
			Checksum:        *template.Checksum,
			DefaultUser:     template.DefaultUser,
			Description:     template.Description,
			Name:            *template.Name,
			PasswordEnabled: *template.PasswordEnabled,
			SshKeyEnabled:   *template.SSHKeyEnabled,
			Url:             *template.URL,
		})
	if err != nil {
		return nil, err
	}

	res, err := papi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetTemplate(ctx, zone, *res.(*papi.Reference).Id)
}

// ListTemplates returns the list of existing Templates in the specified zone.
func (c *Client) ListTemplates(ctx context.Context, zone, visibility, family string) ([]*Template, error) {
	list := make([]*Template, 0)

	resp, err := c.ListTemplatesWithResponse(apiv2.WithZone(ctx, zone), &papi.ListTemplatesParams{
		Visibility: (*papi.ListTemplatesParamsVisibility)(&visibility),
		Family: func() *string {
			if family != "" {
				return &family
			}
			return nil
		}(),
	})
	if err != nil {
		return nil, err
	}

	if resp.JSON200.Templates != nil {
		for i := range *resp.JSON200.Templates {
			list = append(list, templateFromAPI(&(*resp.JSON200.Templates)[i]))
		}
	}

	return list, nil
}

// GetTemplate returns the Template corresponding to the specified ID in the specified zone.
func (c *Client) GetTemplate(ctx context.Context, zone, id string) (*Template, error) {
	resp, err := c.GetTemplateWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return templateFromAPI(resp.JSON200), nil
}

// DeleteTemplate deletes the specified Template in the specified zone.
func (c *Client) DeleteTemplate(ctx context.Context, zone, id string) error {
	resp, err := c.DeleteTemplateWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return err
	}

	_, err = papi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
