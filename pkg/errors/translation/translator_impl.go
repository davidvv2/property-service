package translation

import (
	"embed"
	"encoding/json"
	"strings"

	"property-service/pkg/infrastructure/log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var _ Translator = (*I18nImpl)(nil)

type I18nImpl struct {
	bundle *i18n.Bundle

	languages embed.FS

	log log.Logger
}

//go:embed locales/*.json
var languages embed.FS

// New :
func New(l log.Logger) Translator {
	l.Info("Initialising translation service")

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	s, _ := languages.ReadDir("locales")

	for i := 0; i < len(s); i++ {
		bundle.LoadMessageFileFS(languages, "locales/"+s[i].Name())
		l.Info("Language Loaded %+v", strings.Split(s[i].Name(), "."))
	}
	l.Info("Initialised translation service")

	return &I18nImpl{
		languages: languages,
		bundle:    bundle,
		log:       l,
	}
}

// GetTranslation implements Translator.
func (ti *I18nImpl) GetTranslation(item string, lang string) string {
	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: item,
	}

	localizer := i18n.NewLocalizer(ti.bundle, lang)

	localizationUsingJSON, _ := localizer.Localize(&localizeConfigWelcome)
	return localizationUsingJSON
}
