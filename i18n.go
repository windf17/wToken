package wtoken

// ErrorI18nMessages 定义多语言错误信息映射
var ErrorI18nMessages = map[Language]map[ErrorCode]string{
	LangChinese: {
		ErrCodeSuccess:              "成功",
		ErrCodeInvalidToken:         "无效的token",
		ErrCodeTokenExpired:         "token已过期",
		ErrCodeTokenNotFound:        "token不存在",
		ErrCodeTokenLimitExceeded:   "超出token数量限制",
		ErrCodeInvalidUserID:        "无效的用户ID",
		ErrCodeUserIDNotFound:       "用户ID不存在",
		ErrCodeGroupNotFound:        "用户组不存在",
		ErrCodeInvalidGroupID:       "无效的用户组ID",
		ErrCodeInvalidIP:            "无效的IP地址",
		ErrCodeIPMismatch:           "IP地址不匹配",
		ErrCodeInvalidConfig:        "无效的配置",
		ErrCodeCacheFileLoadFailed:  "加载缓存文件失败",
		ErrCodeCacheFileParseFailed: "缓存文件解析错误",
		ErrCodeAccessDenied:         "无权访问该API",
		ErrCodeInternalError:        "内部错误",
		ErrCodeTypeAssertionError:   "类型断言错误",
		ErrCodeInvalidURL:           "无效的URL",
	},
	LangEnglish: {
		ErrCodeSuccess:              "Success",
		ErrCodeInvalidToken:         "Invalid token",
		ErrCodeTokenExpired:         "Token expired",
		ErrCodeTokenNotFound:        "Token not found",
		ErrCodeTokenLimitExceeded:   "Token limit exceeded",
		ErrCodeInvalidUserID:        "Invalid user ID",
		ErrCodeUserIDNotFound:       "User ID not found",
		ErrCodeGroupNotFound:        "User group not found",
		ErrCodeInvalidGroupID:       "Invalid user group ID",
		ErrCodeInvalidIP:            "Invalid IP address",
		ErrCodeIPMismatch:           "IP address mismatch",
		ErrCodeInvalidConfig:        "Invalid configuration",
		ErrCodeCacheFileLoadFailed:  "Cache file load failed",
		ErrCodeCacheFileParseFailed: "Cache file parse failed",
		ErrCodeAccessDenied:         "API access not allowed",
		ErrCodeInternalError:        "Internal error",
		ErrCodeTypeAssertionError:   "Type assertion error",
		ErrCodeInvalidURL:           "Invalid URL",
	},
}
