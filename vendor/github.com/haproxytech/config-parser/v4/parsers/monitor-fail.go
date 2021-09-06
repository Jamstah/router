/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package parsers

import (
	"fmt"
	"strings"

	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"
)

type MonitorFail struct {
	data        *types.MonitorFail
	preComments []string
}

func (p *MonitorFail) Parse(line string, parts, previousParts []string, comment string) (changeState string, err error) {
	if len(parts) > 3 && parts[0] == "monitor" && parts[1] == "fail" {
		if op := parts[2]; op == "if" || op == "unless" {
			p.data = &types.MonitorFail{
				Condition: op,
				ACLList:   parts[3:],
			}
			return "", nil
		}
		return "", &errors.ParseError{Parser: "monitor fail", Line: line, Message: "Unrecognized operator"}
	}
	return "", &errors.ParseError{Parser: "monitor fail", Line: line}
}

func (p *MonitorFail) Result() ([]common.ReturnResultLine, error) {
	if p.data == nil {
		return nil, errors.ErrFetch
	}
	return []common.ReturnResultLine{
		{
			Data: fmt.Sprintf("monitor fail %s %s", p.data.Condition, strings.Join(p.data.ACLList, " ")),
		},
	}, nil
}
