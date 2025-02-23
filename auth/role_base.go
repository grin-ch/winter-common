package auth

const DefaultVersion = "1.0.0"
const UnknownSex = "Unknown"

type roleBase struct {
	Version  string `json:"version"`
	Uid      int64  `json:"uid"`
	Gid      string `json:"gid"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender,omitempty"`
	Ip       string `json:"ip"`
}

func WithUid(uid int64) Option {
	return func(rb *roleBase) {
		rb.Uid = uid
	}
}

func WithGender(gender string) Option {
	return func(rb *roleBase) {
		rb.Gender = gender
	}
}

func WithVersion(version string) Option {
	return func(rb *roleBase) {
		rb.Version = version
	}
}

type Option func(*roleBase)

func MakeRoleBase(gid, avatar, nickname, ip string, opts ...Option) roleBase {
	rb := roleBase{
		Gid:      gid,
		Version:  DefaultVersion,
		Avatar:   avatar,
		Nickname: nickname,
		Gender:   UnknownSex,
		Ip:       ip,
	}
	for _, fn := range opts {
		fn(&rb)
	}
	return rb
}
