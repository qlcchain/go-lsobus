/*
 * Copyright (c) 2019 QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package version

import (
	"fmt"
	"strings"
	"time"
)

var (
	Version   = ""
	GitRev    = ""
	BuildTime = ""
)

func VersionString() string {
	if len(BuildTime) == 0 {
		BuildTime = time.Now().Format(time.RFC3339)
	}
	var b strings.Builder
	t := BuildTime
	if parse, err := time.Parse(time.RFC3339, BuildTime); err == nil {
		t = parse.Local().Format("2006-01-02 15:04:05-0700")
	}

	b.WriteString(fmt.Sprintf("%-15s%s\n", "build time:", t))
	b.WriteString(fmt.Sprintf("%-15s%s\n", "version:", Version))
	b.WriteString(fmt.Sprintf("%-15s%s\n", "hash:", GitRev))
	return b.String()
}

func ShortVersion() string {
	t := BuildTime
	if parse, err := time.Parse(time.RFC3339, BuildTime); err == nil {
		t = parse.Local().Format("2006-01-02 15:04:05-0700")
	}
	return fmt.Sprintf("go-lsobus %s-%s @ %s", Version, GitRev, t)
}
