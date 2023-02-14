package version

const versionPrefix = "golang-sdk-"

//go:generate go run gen.go
func GetSDKVersion() string {
	return versionPrefix + version
}
