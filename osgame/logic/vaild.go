package logic

import "git.woa.com/imkd-ingame/community_admin/services"

// CheckLanguage 检查用户输入的语言
func CheckLanguage(language string) bool {
	for _, v := range services.Cfg.OsGame.Language {
		if v == language {
			return true
		}
	}
	return false
}
