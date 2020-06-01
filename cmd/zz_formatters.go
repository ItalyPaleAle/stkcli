/*
Copyright © 2020 Alessandro Segala (@ItalyPaleAle)

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

	"github.com/statiko-dev/stkcli/utils"
)

// Fromat siteGetResponseModel
func siteGetResponseModelFormat(m *siteGetResponseModel) (result string) {
	aliases := "\033[2m<nil>\033[0m"
	if len(m.Aliases) > 0 {
		aliases = strings.Join(m.Aliases, ", ")
	}

	app := "\033[2m<nil>\033[0m"
	if m.App != nil && m.App.Name != "" {
		app = m.App.Name
	}

	tlsCert := "\033[2m<nil>\033[0m"
	if m.TLS != nil {
		if m.TLS.Type == TLSCertificateSelfSigned {
			tlsCert = "self-signed"
		} else if m.TLS.Type == TLSCertificateACME {
			tlsCert = "acme"
		} else if m.TLS.Certificate != "" {
			tlsCert = m.TLS.Certificate
			if m.TLS.Version != "" {
				tlsCert += " (" + m.TLS.Version + ")"
			}
		}
	}

	result = fmt.Sprintf(`Domain:        %s
Aliases:       %s
TLSCert:       %s
App:           %s`, m.Domain, aliases, tlsCert, app)
	return
}

// Format siteListResponseModel
func siteListResponseModelFormat(m siteListResponseModel) (result string) {
	result = ""
	l := len(m)
	if l == 0 {
		result += "No site configured"
	}
	for i := 0; i < l; i++ {
		result += siteGetResponseModelFormat(&m[i])
		if i < l-1 {
			result += "\n\n"
		}
	}
	return
}

// Format appListResponseModel
func appListResponseModelFormat(m appListResponseModel) (result string) {
	result = ""
	l := len(m)
	if l == 0 {
		result += "No apps stored"
	}
	for i := 0; i < l; i++ {
		result += fmt.Sprintf(`Name:         %s
Size:         %s
LastModified: %s`, m[i].Name, utils.FormatBytes(m[i].Size), m[i].LastModified.Format(time.RFC3339))
		if i < l-1 {
			result += "\n\n"
		}
	}
	return
}

// Format statusResponseModel
func statusResponseModelFormat(m *statusResponseModel, indent bool) (result string) {
	// Indentation
	prefix := ""
	if indent {
		prefix = "    "
	}

	// Info (Nginx and sync status)
	nginxRunning := "no"
	if m.Nginx.Running {
		nginxRunning = "yes"
	}
	syncRunning := "no"
	if m.Sync.Running {
		syncRunning = "yes"
	}
	syncError := "\033[2m<nil>\033[0m"
	if m.Sync.SyncError != "" {
		syncError = m.Sync.SyncError
	}
	storeHealthy := "no"
	if m.Store.Healthy {
		storeHealthy = "yes"
	}

	result = fmt.Sprintf("%[1]sInfo\n%[1]s----\n"+`
%[1]sNode name:        %[2]s
%[1]sNginx is running: %[3]s
%[1]sSync is running:  %[4]s
%[1]sLast sync time:   %[5]s
%[1]sSync error:       %[6]s
%[1]sStore is healthy: %[7]s

`, prefix, m.NodeName, nginxRunning, syncRunning, m.Sync.LastSync.Format(time.RFC3339), syncError, storeHealthy)

	// Sites
	result += prefix + "Sites\n" + prefix + "-----\n\n"

	l := len(m.Health)
	if l == 0 {
		result += "No site configured"
	}
	for i := 0; i < l; i++ {
		el := m.Health[i]

		healthy := "no"
		if el.Healthy {
			healthy = "yes"
		}

		app := "\033[2m<nil>\033[0m"
		if el.App != nil {
			app = *el.App
		}

		ts := "\033[2m<nil>\033[0m"
		if el.Time != nil {
			ts = el.Time.Format(time.RFC3339)
		}

		err := "\033[2m<nil>\033[0m"
		if el.Error != nil {
			err = *el.Error
		}

		result += fmt.Sprintf(`%[1]sDomain:       %[2]s
%[1]sHealthy:      %[3]s
%[1]sApp:          %[4]s
%[1]sLast check:   %[5]s
%[1]sError:        %[6]s

`, prefix, el.Domain, healthy, app, ts, err)
	}

	return
}

// Format clusterStatusResponseModel
func clusterStatusResponseModelFormat(m clusterStatusResponseModel) (result string) {
	// Return all nodes' status
	if m == nil || len(m) == 0 {
		return "Empty response"
	}
	for name, status := range m {
		result += fmt.Sprintf("Node %s:\n\n%s\n", name, statusResponseModelFormat(status, true))
	}
	return
}
