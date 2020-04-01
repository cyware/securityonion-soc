// Copyright 2019 Jason Ertel (jertel). All rights reserved.
//
// This program is distributed under the terms of version 2 of the
// GNU General Public License.  See LICENSE for further details.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package securityonion

import (
  "github.com/security-onion-solutions/securityonion-soc/module"
  "github.com/security-onion-solutions/securityonion-soc/server"
)

const DEFAULT_TIME_SHIFT_MS = 120000
const DEFAULT_TIMEOUT_MS = 60000
const DEFAULT_INDEX = "*:so-*"

type SecurityOnion struct {
  config			module.ModuleConfig
  server			*server.Server
  elastic			*SoElastic
}

func NewSecurityOnion(srv *server.Server) *SecurityOnion {
  return &SecurityOnion {
    server: srv,
    elastic: NewSoElastic(),
  }
}

func (somodule *SecurityOnion) PrerequisiteModules() []string {
  return nil
}

func (somodule *SecurityOnion) Init(cfg module.ModuleConfig) error {
  somodule.config = cfg
  host := module.GetStringDefault(cfg, "elasticsearchHost", "elasticsearch")
  verifyCert := module.GetBoolDefault(cfg, "elasticsearchVerifyCert", true)
  username := module.GetStringDefault(cfg, "elasticsearchUsername", "")
  password := module.GetStringDefault(cfg, "elasticsearchPassword", "")
  timeShiftMs := module.GetIntDefault(cfg, "timeShiftMs", DEFAULT_TIME_SHIFT_MS)
  timeoutMs := module.GetIntDefault(cfg, "timeoutMs", DEFAULT_TIMEOUT_MS)
  index := module.GetStringDefault(cfg, "index", DEFAULT_INDEX)
  return somodule.elastic.Init(host, username, password, verifyCert, timeShiftMs, timeoutMs, index)
}

func (somodule *SecurityOnion) Start() error {
  somodule.server.Host.Register("/securityonion/joblookup", NewSoJobLookupHandler(somodule.server, somodule.elastic))
  return nil
}

func (somodule *SecurityOnion) Stop() error {
  return nil
}

func (somodule *SecurityOnion) IsRunning() bool {
  return false
}