# conf

This project is a module to parse configuration in YAML file

## Usage

```go
type Conf struct {
	Structs    *structs.Conf    `yaml:"structs"`
	Storage    *storage.Conf    `yaml:"storage"`
	Controller *controller.Conf `yaml:"controller"`
	Routers    *routers.Conf    `yaml:"router"`
}

func LoadConf(path string) error {
	Config = new(Conf)
	Config.Structs = &structs.Config
	Config.Storage = &storage.Config
	Config.Controller = &controller.Config
	Config.Routers = &routers.Config

	err := conf.Parse(path, Config)
	if err != nil {
		return fmt.Errorf("fail to parse config: %w", err)
	}

	return nil
}
```
