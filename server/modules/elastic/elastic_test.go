// Copyright 2019 Jason Ertel (jertel). All rights reserved.
// Copyright 2020-2021 Security Onion Solutions, LLC. All rights reserved.
//
// This program is distributed under the terms of version 2 of the
// GNU General Public License.  See LICENSE for further details.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

package elastic

import (
	"testing"
	"time"

	"github.com/security-onion-solutions/securityonion-soc/module"
	"github.com/stretchr/testify/assert"
)

func TestElasticInit(tester *testing.T) {
	elastic := NewElastic(nil)
	cfg := make(module.ModuleConfig)
	err := elastic.Init(cfg)
	if assert.Nil(tester, err) {
		if assert.Len(tester, elastic.store.hostUrls, 1) {
			assert.Equal(tester, "elasticsearch", elastic.store.hostUrls[0])
		}
		assert.Len(tester, elastic.store.esRemoteClients, 0)
		assert.Len(tester, elastic.store.esAllClients, 1)
		assert.Equal(tester, DEFAULT_TIME_SHIFT_MS, elastic.store.timeShiftMs)
		assert.Equal(tester, DEFAULT_DURATION_MS, elastic.store.defaultDurationMs)
		assert.Equal(tester, DEFAULT_ES_SEARCH_OFFSET_MS, elastic.store.esSearchOffsetMs)
		expectedTimeout := time.Duration(DEFAULT_TIMEOUT_MS) * time.Millisecond
		assert.Equal(tester, expectedTimeout, elastic.store.timeoutMs)
		expectedCache := time.Duration(DEFAULT_CACHE_MS) * time.Millisecond
		assert.Equal(tester, expectedCache, elastic.store.cacheMs)
		assert.Equal(tester, DEFAULT_INDEX, elastic.store.index)
		assert.Equal(tester, DEFAULT_INTERVALS, elastic.store.intervals)
	}
}
