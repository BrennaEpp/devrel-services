// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package leif

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dayReg = regexp.MustCompile(`[0-9]+d`)

type stringOrArray []string

func (soa *stringOrArray) UnmarshalJSON(data []byte) error {
	var tempInterface interface{}
	var slice []string

	err := json.Unmarshal(data, &tempInterface)
	if err != nil {
		return err
	}

	str, isString := tempInterface.(string)

	if isString {
		slice = append(slice, str)
	} else {
		err = json.Unmarshal(data, &slice)
	}

	*soa = slice

	return err
}

type duration time.Duration

// A duration is marshalled into the time.Duration int64 representing nanoseconds format
func (d duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d))
}

// Unmarshalling a duration converts it from a string in the time.Duration format (plus days) or an int representing seconds
// It is unmarshalled into the time.Duration representation (an int64 representing nanoseconds)
// Unmarshalling "2d" and then remarshalling it would return 172800000000000
func (d *duration) UnmarshalJSON(data []byte) error {
	var tempInterface interface{}

	err := json.Unmarshal(data, &tempInterface)
	if err != nil {
		return err
	}

	str, isString := tempInterface.(string)

	if isString {
		dur, err := parseDurationWithDays(str)
		*d = duration(dur)
		return err
	}

	value, isNumber := tempInterface.(float64)

	if isNumber {
		*d = duration(time.Duration(int64(value) * int64(time.Second.Nanoseconds())))
		return err
	}

	return errors.New("Invalid duration format")
}

// sloRuleJSON is the intermediary struct between the JSON representation of the SLO representation and the structured leif representation
// The JSON representation of an SLO according to the schema can be marshalled into this struct
// This struct can be marshalled into a JSON representation that corresponds to leif's SLORule struct
type sloRuleJSON struct {
	AppliesToJSON          AppliesToJSON          `json:"appliesTo"`
	ComplianceSettingsJSON ComplianceSettingsJSON `json:"complianceSettings"`
}

// Returns a new SLORuleJSON with the defaults applied,
// except for the responders default which requires knowing if it was partially assigned
func newSLORuleJSON() *sloRuleJSON {
	return &sloRuleJSON{
		AppliesToJSON:          AppliesToJSON{Issues: true, PRs: false},
		ComplianceSettingsJSON: ComplianceSettingsJSON{RequiresAssignee: false},
	}
}

// Applies the default value to the Responders field in the ComplianceSettings
// Only applies it if none of the values in the Responders fiel have been initialized
func (rule *sloRuleJSON) applyResponderDefault() {
	if rule.ComplianceSettingsJSON.RespondersJSON.OwnersRaw == nil &&
		len(rule.ComplianceSettingsJSON.RespondersJSON.Contributors) < 1 &&
		rule.ComplianceSettingsJSON.RespondersJSON.Users == nil {
		rule.ComplianceSettingsJSON.RespondersJSON.Contributors = "WRITE"
	}
}

func parseSLORule(rawRule *json.RawMessage) (*SLORule, error) {
	jsonRule := newSLORuleJSON()

	err := json.Unmarshal(*rawRule, &jsonRule)
	if err != nil {
		return nil, err
	}

	jsonRule.applyResponderDefault()

	marshaled, err := json.Marshal(jsonRule)
	if err != nil {
		return nil, err
	}

	var parsedRule *SLORule
	err = json.Unmarshal(marshaled, &parsedRule)

	return parsedRule, err
}

func unmarshalSLOs(data []byte) ([]*SLORule, error) {
	var sloRules []*SLORule
	var rawSLORules []*json.RawMessage

	if len(data) == 0 {
		return sloRules, nil
	}

	err := json.Unmarshal(data, &rawSLORules)
	if err != nil {
		return sloRules, err
	}

	for _, rawRule := range rawSLORules {
		rule, err := parseSLORule(rawRule)
		if err != nil {
			return sloRules, err
		}
		sloRules = append(sloRules, rule)
	}

	return sloRules, err
}

// AppliesToJSON is the intermediary struct between the JSON representation and the structured leif representation
// that stores structured data on which issues and/or pull requests a SLO applies to
type AppliesToJSON struct {
	GitHubLabelsRaw         stringOrArray `json:"gitHubLabels"`
	ExcludedGitHubLabelsRaw stringOrArray `json:"excludedGitHubLabels"`
	Issues                  bool          `json:"issues"`
	PRs                     bool          `json:"prs"`
}

// ComplianceSettingsJSON is the intermediary struct between the JSON representation and the structured leif representation
// that stores data on the requirements for an issue or pull request to be considered compliant with the SLO
type ComplianceSettingsJSON struct {
	ResponseTime     duration       `json:"responseTime"`
	ResolutionTime   duration       `json:"resolutionTime"`
	RequiresAssignee bool           `json:"requiresAssignee"`
	RespondersJSON   RespondersJSON `json:"responders"`
}

// RespondersJSON is the intermediary struct between the JSON representation and the structured leif representation
// that stores structured data on the responders to the issue or pull request the SLO applies to
type RespondersJSON struct {
	OwnersRaw    stringOrArray `json:"owners"`
	Contributors string        `json:"contributors"`
	Users        []string      `json:"users"`
}

func parseDurationWithDays(duration string) (time.Duration, error) {
	if strings.Contains(duration, ".") {
		return 0, errors.New("Duration should not contain fractions")
	}
	str := dayReg.ReplaceAllStringFunc(duration, func(s string) string {
		days, _ := strconv.Atoi(s[:len(s)-1])
		s = strconv.Itoa(days*24) + "h"
		return s
	})
	return time.ParseDuration(str)
}
