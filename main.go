package main

import (
	"contabo.com/cli/cntb/cmd"
	_ "contabo.com/cli/cntb/cmd/buckets"
	_ "contabo.com/cli/cntb/cmd/datacenters"
	_ "contabo.com/cli/cntb/cmd/images"
	_ "contabo.com/cli/cntb/cmd/instanceActions"
	_ "contabo.com/cli/cntb/cmd/instances"
	_ "contabo.com/cli/cntb/cmd/objectStorage"
	_ "contabo.com/cli/cntb/cmd/objects"
	_ "contabo.com/cli/cntb/cmd/privateNetwork"
	_ "contabo.com/cli/cntb/cmd/roles"
	_ "contabo.com/cli/cntb/cmd/secrets"
	_ "contabo.com/cli/cntb/cmd/snapshots"
	_ "contabo.com/cli/cntb/cmd/tagAssignment"
	_ "contabo.com/cli/cntb/cmd/tags"
	_ "contabo.com/cli/cntb/cmd/users"
	_ "contabo.com/cli/cntb/cmd/vips"
)

func main() {
	cmd.Execute()
}
