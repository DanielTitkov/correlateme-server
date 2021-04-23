package configs

type (
	DataConfig struct {
		Presets PresetsConfig
	}
	PresetsConfig struct {
		ScalePresetsPath           string `yaml:"scalePresetsPath"`
		IndicatorPresetsPath       string `yaml:"indicatorPresetsPath"`
		DictionaryPresetsPath      string `yaml:"dictionaryPresetsPath"`
		DictionaryEntryPresetsPath string `yaml:"dictionaryEntryPresetsPath"`
	}
)
