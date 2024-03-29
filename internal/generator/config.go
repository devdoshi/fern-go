package generator

// Config represents the Fern generator configuration.
type Config struct {
	DryRun     bool
	IRFilepath string
	ImportPath string

	// TODO: Add Workspace, Organization, and Environment, as needed.
}

// NewConfig returns a new *Config for the given values.
func NewConfig(dryRun bool, irFilepath string, importPath string) (*Config, error) {
	return &Config{
		DryRun:     dryRun,
		IRFilepath: irFilepath,
		ImportPath: importPath,
	}, nil
}
