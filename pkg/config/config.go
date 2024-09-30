// Copyright 2024 Fabian `xx4h` Sylvester
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Hub        ConfigHub        `yaml:"hub"`
	Completion ConfigCompletion `yaml:"completion"`
	Handling   ConfigHandling   `yaml:"handling"`
	Logging    ConfigLogging    `yamn:"logging"`
	Viper      *viper.Viper
}

type ConfigHub struct {
	Type  string `yaml:"type"`
	Url   string `yaml:"url"`
	Token string `yaml:"token"`
}

type ConfigCompletion struct {
	ShortNames bool `yaml:"shortNames"`
	WhatIf     bool `yaml:"whatif"`
}

type ConfigHandling struct {
	Fuzz bool `yaml:"fuzz"`
}

type ConfigLogging struct {
	LogLevel string `yaml:"log_level"`
}

func NewConfig() (*Config, error) {
	v := viper.New()

	v.SetEnvPrefix("HCTL")
	v.AutomaticEnv()
	v.SetConfigType("yaml")
	v.SetConfigName("hctl")
	v.AddConfigPath("$HOME/.config/hctl")
	v.AddConfigPath(".")

	// create empty config and set defaults
	cfg := &Config{}
	cfg.Completion.ShortNames = true
	cfg.Completion.WhatIf = true
	cfg.Handling.Fuzz = true
	cfg.Logging.LogLevel = "error"

	// use defaults for viper as well
	v.SetDefault("completion", &cfg.Completion)
	v.SetDefault("handling", &cfg.Handling)
	v.SetDefault("logging", &cfg.Logging)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug().Msgf("Config File not found! Please run `hctl init` or manually create %s", v.ConfigFileUsed())
		}
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	log_level := v.GetString("log_level")
	if log_level != "" {
		lvl, err := zerolog.ParseLevel(log_level)
		if err != nil {
			return nil, err
		}
		zerolog.SetGlobalLevel(lvl)
	}

	log.Debug().Msgf("Running with the following config: %#v", v)

	cfg.Viper = v

	return cfg, nil
}