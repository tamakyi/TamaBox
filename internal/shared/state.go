package shared

import "sync"

type AppState struct {
    mu            sync.Mutex
    Lang          string
    Translation   string
}

func (s *AppState) determineLangBasedOnUser(translation string) string {
    switch translation {
    case "zh_CN":
        return "zh_CN"
    case "zh_HK":
        return "zh_HK"
    case "en":
        return "en"
    case "ja":
        return "ja"
    default:
        return "zh_CN"
    }
}

func (s *AppState) SetTranslation(translation string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.Translation = translation
    s.Lang = determineLangBasedOnUser(translation)
}

func (s *AppState) GetLang() string {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.Lang
}
