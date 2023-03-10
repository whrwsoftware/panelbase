// Copyright 2023 panelbase Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package appver

type ver struct {
	Pkg       string
	Service   string
	Version   string
	VersionId int
	VersionV  int
	LogFile   string
}

func Ver(pkg, version string, versionId int, logFile string) *ver {
	return &ver{Pkg: pkg, Version: version, VersionId: versionId, LogFile: logFile}
}

func VerVV(pkg string, service string, version string, versionId int, versionV int, logFile string) *ver {
	return &ver{Pkg: pkg, Service: service, Version: version, VersionId: versionId, VersionV: versionV, LogFile: logFile}
}
