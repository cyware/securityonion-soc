// Copyright 2020-2021 Security Onion Solutions, LLC. All rights reserved.
//
// This program is distributed under the terms of version 2 of the
// GNU General Public License.  See LICENSE for further details.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package generichttp

import (
  "io"
  "testing"

  "github.com/security-onion-solutions/securityonion-soc/model"
  "github.com/stretchr/testify/assert"
)

func TestConvertCaseToReader(tester *testing.T) {
  socCase := model.NewCase()
  socCase.Id = "123"
  socCase.Title = "MyTitle"
  socCase.Description = "My \"Description\" is this."
  socCase.Severity = 44

  source := `ID: {{ .Id }}; Title: {{ .Title }}; Desc: {{ .Description | js }}; Sev: {{ .Severity }}; Time: {{ .CreateTime.Format "15:04" }}`

  reader, err1 := convertCaseToReader(source, socCase)
  assert.NoError(tester, err1)

  bytes, err2 := io.ReadAll(reader)
  assert.NoError(tester, err2)

  converted := string(bytes)
  assert.Equal(tester, "ID: 123; Title: MyTitle; Desc: My \\\"Description\\\" is this.; Sev: 44; Time: 00:00", converted)
}
