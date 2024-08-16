
package version

import "github.com/hashicorp/packer-plugin-sdk/version"

var (
	Version           = "1.1.0"
	VersionPrerelease = ""
	VersionMetadata   = ""
	PluginVersion     = version.NewPluginVersion(Version, VersionPrerelease, VersionMetadata)
)