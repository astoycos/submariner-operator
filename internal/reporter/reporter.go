/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

package reporter

import (
	"fmt"

	"github.com/submariner-io/submariner-operator/internal/cli"
	"github.com/submariner-io/submariner-operator/pkg/eventreporter"
)

type cliReporter struct {
	status *cli.Status
}

func NewCLIReporter() eventreporter.Reporter {
	return &cliReporter{status: cli.NewStatus()}
}

func (r *cliReporter) Started(message string, args ...interface{}) {
	r.status.Start(fmt.Sprintf(message, args...))
}

func (r *cliReporter) Succeeded(message string, args ...interface{}) {
	if message != "" {
		r.status.QueueSuccessMessage(fmt.Sprintf(message, args...))
	}
}

func (r *cliReporter) Warned(message string) {
	if message != "" {
		r.status.QueueWarningMessage(message)
	}
}

func (r *cliReporter) Failed(message string) {
	if message != "" {
		r.status.QueueFailureMessage(message)
	}
}

func (r *cliReporter) EndedWith(err error) {
	r.status.End(cli.CheckForError(err))
}
