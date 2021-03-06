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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"

	"github.com/statiko-dev/stkcli/utils"
)

func init() {
	var (
		domain         string
		aliases        []string
		tlsCertificate string
	)

	c := &cobra.Command{
		Use:   "set",
		Short: "Updates the configuration for a site",
		Long: `Updates a site configured in the node.

When creating a site, you must specify the name of a TLS certificate stored in the node or cluster. Alternatively, you can pass one of the following values:

  - ` + "`" + `selfsigned` + "`" + ` for generating a self-signed certificate for your site
  - ` + "`" + `acme` + "`" + ` for requesting a certificate from an ACME provider, such as Let's Encrypt
  - ` + "`" + `akv:[name]:[version]` + "`" + ` for requesting a certificate stored in the Azure Key Vault instance associated with the cluster; the version is optional.

If you omit the ` + "`" + `--certificate` + "`" + ` option, it will default to a self-signed certificate.

The ` + "`" + `--alias` + "`" + ` parameter is used to replace the list of aliases configured for the domain. You can use this parameter multiple time to add more than one alias. Note that using the ` + "`" + `--alias` + "`" + ` flag will replace the entire list of aliases with the new one.
`,
		DisableAutoGenTag: true,

		Run: func(cmd *cobra.Command, args []string) {
			baseURL, client := getURLClient()
			auth := nodeStore.GetAuthToken(optAddress)

			// Request body
			tlsConfig := &siteTLSConfiguration{}
			if tlsCertificate == "" || tlsCertificate == "selfsigned" {
				tlsConfig.Type = TLSCertificateSelfSigned
			} else if tlsCertificate == "acme" || tlsCertificate == "letsencrypt" {
				tlsConfig.Type = TLSCertificateACME
			} else if strings.HasPrefix(tlsCertificate, "akv:") {
				tlsConfig.Type = TLSCertificateAzureKeyVault
				tlsConfig.Certificate = tlsCertificate[4:]
				// Check if there's a version
				i := strings.Index(tlsConfig.Certificate, ":")
				// Start from 1 because the certificate name must be 1 character at least
				if i > 0 {
					tlsConfig.Version = tlsConfig.Certificate[(i + 1):]
					tlsConfig.Certificate = tlsConfig.Certificate[0:i]
				}
			} else {
				tlsConfig.Type = TLSCertificateImported
				tlsConfig.Certificate = tlsCertificate
			}
			reqBody := &siteSetRequestModel{
				Aliases: aliases,
				TLS:     tlsConfig,
			}
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(reqBody)
			if err != nil {
				utils.ExitWithError(utils.ErrorApp, "Error while encoding to JSON", err)
				return
			}

			// Invoke the /site/:domain endpoint and edit the site
			var r siteGetResponseModel
			err = utils.RequestJSON(utils.RequestOpts{
				Authorization:   auth,
				Body:            buf,
				BodyContentType: "application/json",
				Client:          client,
				Method:          utils.RequestPATCH,
				StatusCode:      http.StatusOK,
				Target:          &r,
				URL:             baseURL + "/site/" + domain,
			})
			if err != nil {
				utils.ExitWithError(utils.ErrorNode, "Request failed", err)
				return
			}

			// Print the response
			fmt.Println(siteGetResponseModelFormat(&r))
		},
	}
	siteCmd.AddCommand(c)

	// Flags
	c.Flags().StringVarP(&domain, "domain", "d", "", "primary domain name")
	c.MarkFlagRequired("domain")
	c.Flags().StringArrayVarP(&aliases, "alias", "a", []string{}, "alias domain (can be used multiple times)")
	c.Flags().StringVarP(&tlsCertificate, "certificate", "c", "", "name of the TLS certificate")

	// Add shared flags
	addSharedFlags(c)
}
