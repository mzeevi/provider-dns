package a

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("dns_a_record_set", func(r *config.Resource) {
		r.ShortGroup = "a"
	})
}
