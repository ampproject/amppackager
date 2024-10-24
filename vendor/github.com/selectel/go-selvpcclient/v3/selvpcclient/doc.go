/*
Package selvpcclient provides a library to work with the Selectel VPC API.

# Authentication

	To work with the Selectel VPC API you first need to:

	  - create a Selectel account: https://my.selectel.ru/registration
	  - create the service user: https://docs.selectel.ru/control-panel-actions/users-and-roles/add-user/

# Usage example

	ctx := context.Background()

	options := &selvpcclient.ClientOptions{
		Context:    ctx,
		DomainName: "999999",
		Username:   "admin",
		Password:   "m1-sup3r-p@ssw0rd-p3w-p3w",
	}

	client, err := selvpcclient.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}

	result, resp, err := projects.List(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response StatusCode: %d \n", resp.StatusCode)

	for _, project := range result {
		fmt.Printf("Project name: %s, enabled: %t \n", project.Name, project.Enabled)
	}
*/
package selvpcclient
