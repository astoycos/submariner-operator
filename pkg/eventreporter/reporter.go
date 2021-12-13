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

package eventreporter

// Reporter is responsible for reporting back on the progress of the specified preparation.
type Reporter interface {
	// Started will report that an operation started.
	Started(message string, args ...interface{})

	// Succeeded will report that the last operation has succeeded with the specified message.
	Succeeded(message string, args ...interface{})

	// Failed will report that the last operation has failed with the specified error.
	Failed(message string)

	// EndedWith ends the Started eventreporter with Success if there are no errors otherwise with a Failure.
	EndedWith(err error)

	// Warned will report that the last operation has thrown a warning.
	Warned(message string)
}
