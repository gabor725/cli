package ccv2

import (
	"encoding/json"
	"net/http"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/internal"
)

// Application represents a Cloud Controller Application.
type Application struct {
	GUID string
	Name string
}

// UnmarshalJSON helps unmarshal a Cloud Controller Application response.
func (application *Application) UnmarshalJSON(data []byte) error {
	var ccApp struct {
		Metadata internal.Metadata `json:"metadata"`
		Entity   struct {
			Name string `json:"name"`
		} `json:"entity"`
	}
	if err := json.Unmarshal(data, &ccApp); err != nil {
		return err
	}

	application.GUID = ccApp.Metadata.GUID
	application.Name = ccApp.Entity.Name
	return nil
}

// GetApplications returns back a list of Applications based off of the
// provided queries.
func (client *CloudControllerClient) GetApplications(queries []Query) ([]Application, Warnings, error) {
	request := cloudcontroller.NewRequest(
		internal.AppsRequest,
		nil,
		nil,
		FormatQueryParameters(queries),
	)

	fullAppsList := []Application{}
	fullWarningsList := Warnings{}

	for {
		var apps []Application
		wrapper := PaginatedWrapper{
			Resources: &apps,
		}
		response := cloudcontroller.Response{
			Result: &wrapper,
		}

		err := client.connection.Make(request, &response)
		fullWarningsList = append(fullWarningsList, response.Warnings...)
		if err != nil {
			return nil, fullWarningsList, err
		}
		fullAppsList = append(fullAppsList, apps...)

		if wrapper.NextURL == "" {
			break
		}
		request = cloudcontroller.NewRequestFromURI(
			wrapper.NextURL,
			http.MethodGet,
			nil,
		)
	}

	return fullAppsList, fullWarningsList, nil
}

// GetRouteApplications returns a list of Applications associated with a route
// GUID, filtered by provided queries.
func (client *CloudControllerClient) GetRouteApplications(routeGUID string, queryParams []Query) ([]Application, Warnings, error) {
	request := cloudcontroller.NewRequest(
		internal.AppsFromRouteRequest,
		map[string]string{"route_guid": routeGUID},
		nil,
		FormatQueryParameters(queryParams),
	)

	fullAppsList := []Application{}
	fullWarningsList := Warnings{}

	for {
		var apps []Application
		wrapper := PaginatedWrapper{
			Resources: &apps,
		}
		response := cloudcontroller.Response{
			Result: &wrapper,
		}

		err := client.connection.Make(request, &response)
		fullWarningsList = append(fullWarningsList, response.Warnings...)
		if err != nil {
			return nil, fullWarningsList, err
		}
		fullAppsList = append(fullAppsList, apps...)

		if wrapper.NextURL == "" {
			break
		}
		request = cloudcontroller.NewRequestFromURI(
			wrapper.NextURL,
			http.MethodGet,
			nil,
		)
	}

	return fullAppsList, fullWarningsList, nil
}
