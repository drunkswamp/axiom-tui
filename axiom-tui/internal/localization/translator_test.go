package localization

import "testing"

func TestLanguageFromStringDefaultsToEnglish(t *testing.T) {
	if got := LanguageFromString(""); got != English {
		t.Fatalf("LanguageFromString empty = %q, want %q", got, English)
	}
}

func TestLanguageFromStringSupportsEnglish(t *testing.T) {
	if got := LanguageFromString("en"); got != English {
		t.Fatalf("LanguageFromString en = %q, want %q", got, English)
	}
}

func TestLanguageFromStringSupportsRussian(t *testing.T) {
	if got := LanguageFromString("ru"); got != Russian {
		t.Fatalf("LanguageFromString ru = %q, want %q", got, Russian)
	}
}

func TestLanguageFromStringFallsBackToEnglish(t *testing.T) {
	if got := LanguageFromString("de"); got != English {
		t.Fatalf("LanguageFromString unknown = %q, want %q", got, English)
	}
}
