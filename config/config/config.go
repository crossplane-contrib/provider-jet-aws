package config

import tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"

// Configure configures resources in config group.
func Configure(p *tjconfig.Provider) {
	p.AddResourceConfigurator("aws_config_configuration_recorder_status", func(r *tjconfig.Resource) {
		r.Kind = "AWSConfigurationRecorderStatus"
	})
}
