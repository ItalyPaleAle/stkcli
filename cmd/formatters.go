/*
Copyright © 2019 Alessandro Segala (@ItalyPaleAle)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"fmt"
	"strings"
	"time"
)

// Fromat siteAddResponseModel
func siteAddResponseModelFormat(m *siteAddResponseModel) (result string) {
	clientCaching := "no"
	if m.ClientCaching {
		clientCaching = "yes"
	}
	aliases := strings.Join(m.Aliases, ", ")
	result = fmt.Sprintf(`ID:            %s
Domain:        %s
Aliases:       %s
TLSCert:       %s
ClientCaching: %s`, m.ID, m.Domain, aliases, m.TLSCertificate, clientCaching)
	return
}

// Fromat siteGetResponseModel
func siteGetResponseModelFormat(m *siteGetResponseModel) (result string) {
	clientCaching := "no"
	if m.ClientCaching {
		clientCaching = "yes"
	}
	aliases := strings.Join(m.Aliases, ", ")
	result = fmt.Sprintf(`ID:            %s
Domain:        %s
Aliases:       %s
TLSCert:       %s
ClientCaching: %s`, m.ID, m.Domain, aliases, m.TLSCertificate, clientCaching)
	return
}

// Format siteListResponseModel
func siteListResponseModelFormat(m siteListResponseModel) (result string) {
	result = ""
	l := len(m)
	for i := 0; i < l; i++ {
		result += siteAddResponseModelFormat(&m[i])
		if i < l-1 {
			result += "\n\n"
		}
	}
	return
}

// Format statusResponseModel
func statusResponseModelFormat(m *statusResponseModel) (result string) {
	// Apps
	result = "Apps\n----\n\n"

	l := len(m.Apps)
	for i := 0; i < l; i++ {
		el := m.Apps[i]

		appName := "\033[2m<nil>\033[0m"
		if el.AppName != nil {
			appName = *el.AppName
		}

		appVersion := "\033[2m<nil>\033[0m"
		if el.AppVersion != nil {
			appVersion = *el.AppVersion
		}

		t := "\033[2m<nil>\033[0m"
		if el.Updated != nil {
			t = el.Updated.Format(time.RFC3339)
		}

		result += fmt.Sprintf(`ID:         %s
Domain:     %s
AppName:    %s
AppVersion: %s
Updated:    %s

`, el.ID, el.Domain, appName, appVersion, t)
	}

	// Health
	result += "Health\n------\n\n"

	l = len(m.Health)
	for i := 0; i < l; i++ {
		el := m.Health[i]

		err := "\033[2m<nil>\033[0m"
		if el.Error != nil {
			err = *el.Error
		}

		result += fmt.Sprintf(`Domain:       %s
StatusCode:   %d
ResponseSize: %d
Error:        %s
Time:         %s

`, el.Domain, el.StatusCode, el.ResponseSize, err, el.Time.Format(time.RFC3339))
	}

	return
}

// Fromat deployResponseModel
func deployResponseModelFormat(m *deployResponseModel) (result string) {
	err := "\033[2m<nil>\033[0m"
	if m.Error != nil {
		err = *m.Error
	}

	t := "\033[2m<nil>\033[0m"
	if m.Time != nil {
		t = m.Time.Format(time.RFC3339)
	}

	result = fmt.Sprintf(`DeploymentID: %s
SiteID:       %s
AppName:      %s
AppVersion:   %s
Status:       %s
Error:        %s
Time: %s`, m.DeploymentID, m.SiteID, m.AppName, m.AppVersion, m.Status, err, t)
	return
}