package localization

type PersianWeatherTranslator struct{}

func NewPersianWeatherTranslator() *PersianWeatherTranslator {
	return &PersianWeatherTranslator{}
}

func (t *PersianWeatherTranslator) Translate(text string) string {
	translations := map[string]string{
		// انگلیسی
		"clear sky":             "آسمان صاف",
		"few clouds":            "کمی ابری",
		"scattered clouds":      "ابری پراکنده",
		"broken clouds":         "ابری",
		"overcast clouds":       "کاملاً ابری",
		"light rain":            "باران سبک",
		"moderate rain":         "باران",
		"heavy intensity rain":  "باران شدید",
		"shower rain":           "رگبار",
		"thunderstorm":          "رعد و برق",
		"snow":                  "برف",
		"mist":                  "مه",
		"fog":                   "غبار مه",
		"haze":                  "غبار",
		"drizzle":               "نم‌نم باران",
		"smoke":                 "دود",
		"dust":                  "گرد و غبار",

		// فارسی (وقتی lang=fa استفاده می‌شه)
		"آسمان صاف":              "آسمان صاف",
		"غیوم قليلة":             "کمی ابری",
		"غیوم متفرقة":            "ابری پراکنده",
		"غیوم مكسرة":             "ابری",
		"غیوم كثيفة":             "کاملاً ابری",
		"مطر خفیف":              "باران سبک",
		"مطر":                   "باران",
		"مطر غزير":              "باران شدید",
		"زخات مطر":              "رگبار",
		"عاصفة رعدية":            "رعد و برق",
		"ثلج":                   "برف",
		"ضباب":                  "مه",
		"ضباب خفیف":             "غبار مه",

		// گروه‌های اصلی (fallback)
		"clear":                 "آسمان صاف",
		"clouds":                "ابری",
		"rain":                  "باران",
	}

	if val, ok := translations[text]; ok {
		return val
	}
	return "نامشخص"
}