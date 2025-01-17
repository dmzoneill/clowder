package database

import (
	crd "github.com/RedHatInsights/clowder/apis/cloud.redhat.com/v1alpha1"
	"github.com/RedHatInsights/clowder/controllers/cloud.redhat.com/providers"
)

type noneDbProvider struct {
	providers.Provider
}

// NewNoneDBProvider returns a new none db provider object.
func NewNoneDBProvider(p *providers.Provider) (providers.ClowderProvider, error) {
	return &noneDbProvider{Provider: *p}, nil
}

func (db *noneDbProvider) EnvProvide() error {
	return nil
}

func (db *noneDbProvider) Provide(app *crd.ClowdApp) error {
	return nil
}
