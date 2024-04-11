package auth

const DefaultVersion = 1
const UnknownSex = "Unknown"

type roleBase struct {
	Version  uint8  `json:"version"`
	Uid      int64  `json:"uid"`
	Gid      string `json:"gid"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex,omitempty"`
	Ip       string `json:"ip"`
}

func WithUid(uid int64) Option {
	return func(rb *roleBase) {
		rb.Uid = uid
	}
}

func WithSex(sex string) Option {
	return func(rb *roleBase) {
		rb.Sex = sex
	}
}

func WithVersion(version uint8) Option {
	return func(rb *roleBase) {
		rb.Version = version
	}
}

type Option func(*roleBase)

func MakeRoleBase(gid, avatar, nickname, ip string, opts ...Option) roleBase {
	rb := roleBase{
		Version:  DefaultVersion,
		Avatar:   avatar,
		Nickname: nickname,
		Sex:      UnknownSex,
		Ip:       ip,
	}
	for _, fn := range opts {
		fn(&rb)
	}
	return rb
}
