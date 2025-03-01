package auth

const DefaultVersion = "1.0.0"
const UnknownGender = "Unknown"

type roleBase struct {
	Uid      int64  `json:"uid,omitempty"`
	Gid      string `json:"gid,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Version  string `json:"version,omitempty"`
	Ip       string `json:"ip,omitempty"`
}

func WithUid(uid int64) Option {
	return func(rb *roleBase) {
		rb.Uid = uid
	}
}

func WithGid(gid string) Option {
	return func(rb *roleBase) {
		rb.Gid = gid
	}
}

func WithNickname(nickname string) Option {
	return func(rb *roleBase) {
		rb.Nickname = nickname
	}
}

func WithAvatar(avatar string) Option {
	return func(rb *roleBase) {
		rb.Avatar = avatar
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

func WithIp(ip string) Option {
	return func(rb *roleBase) {
		rb.Ip = ip
	}
}

type Option func(*roleBase)

func MakeRoleBase(opts ...Option) roleBase {
	rb := roleBase{
		Version: DefaultVersion,
		Gender:  UnknownGender,
	}
	for _, fn := range opts {
		fn(&rb)
	}
	return rb
}
