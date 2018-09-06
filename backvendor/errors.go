// Copyright (C) 2018 Tim Waugh
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package backvendor

import "errors"

// ErrorNeedImportPath indicates the import path for the project
// cannot be determined automatically and must be provided.
var ErrorNeedImportPath = errors.New("unable to determine import path")

// ErrorVersionNotFound indicates a vendored project does not match any semantic
// tag in the upstream revision control system.
var ErrorVersionNotFound = errors.New("version not found")

// ErrorUnknownVCS indicates the upstream version control system is not one of
// those for which support is implemented in backvendor.
var ErrorUnknownVCS = errors.New("unknown VCS")
