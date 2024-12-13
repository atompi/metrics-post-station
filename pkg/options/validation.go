package options

import (
	"gopkg.in/yaml.v3"
)

func createOptions(optsSource, opts any) error {
	optsYaml, err := yaml.Marshal(optsSource)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(optsYaml, opts)
	return err
}
