package app

var (
	Release, BuildDate, BuildUser, GitTag, GitHash, Image string
)

type AppVersion struct {
	Release   string `json:"release"`
	BuildDate string `json:"buildDate"`
	BuildUser string `json:"buildUser"`
	GitTag    string `json:"gitTag"`
	GitHash   string `json:"gitHash"`
	Image     string `json:"image"`
}

func GetVersion() *AppVersion {
	return &AppVersion{
		Release:   Release,
		BuildDate: BuildDate,
		BuildUser: BuildUser,
		GitTag:    GitTag,
		GitHash:   GitHash,
		Image:     Image}
}
