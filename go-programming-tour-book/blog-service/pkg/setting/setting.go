package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig() //基本上解析成为map[string]interface{} 然后再自定义struct将配置统一处理为对象
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	//viper内部是有fsnotify库的
	s.WatchSettingChange()

	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
