package info

type CoreList []string
type CoreInfo struct {
	name             string
	supportMcVersion []string
	historyVersion   []string
	versionData      map[string]CoreVersionInfo
}
type CoreVersionInfo struct {
	targetMcVersion    string
	requestJavaVersion int
	downloadUrl        string
	sha256             string
	size               int
}
