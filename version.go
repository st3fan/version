// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/

package version

import "fmt"

//
// These are filled in by the linker.
//
// For Gitlab CI:
//
//  go build -ldflags "-X version.sha=${CI_COMMIT_SHA} -X version.ref=${CI_COMMIT_REF_NAME} -X version.tag=${CI_COMMIT_TAG}"
//
// For Travis CI:
//
//  go build -ldflags "-X version.sha=${TRAVIS_COMMIT} -X version.ref=${TRAVIS_BRANCH} -X version.tag=${TRAVIS_TAG}"
//

var (
	sha = ""
	ref = ""
	tag = ""
)

func Tag() string {
	if tag != "" {
		return tag
	}
	return "unknown"
}

func Ref() string {
	if ref != "" {
		return ref
	}
	return "unknown"
}

func Sha() string {
	if sha != "" {
		return sha
	}
	return "unknown"
}

func Banner(name string) string {
	return fmt.Sprintf("This is %s <%s> (%s / %s)", name, Tag(), Ref(), Sha())
}
