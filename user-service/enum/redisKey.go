package enum

type (
	RedisKey string
)

const (
	REDIS_KEY_OTP                      RedisKey = "user:otp:%s:value:user_id"
	REDIS_KEY_LOGIN_WITH_FAILURE_COUNT RedisKey = "user:login_with_failure_count:%s"
)

func (r RedisKey) String() string {
	switch r {
	case
		REDIS_KEY_OTP,
		REDIS_KEY_LOGIN_WITH_FAILURE_COUNT:
		return string(r)
	default:
		return "unknown"
	}
}
