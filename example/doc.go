// +k8s:deepcopy-gen=package

package example

type Person struct {
	Name      string
	Age       int
	Addresses map[string]string
}

type PublicData struct {
	UserId          int64
	Name            string
	Avatar          string
	Level           int32
	Skins           map[int32]int32
	IsUnderGuardian bool
	AvatarFrame     int32
	Language        int32
}
