package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v2"
)

var (
	ErrNoEnv = errors.New("no env variable")
)

// OverrideConfigWithEnv is an interface to implement methode override configuration with environement variable
type OverrideConfigWithEnv interface {
	OverrideWithEnv(prefix string) error
}

// Parse will decode the yaml file into the given interface{}
func Parse(file string, into interface{}) error {
	path := filepath.Clean(file)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("can't read file %s : %w", path, err)
	}

	err = yaml.UnmarshalStrict(bytes, into)
	if err != nil {
		return fmt.Errorf("can't unmarshal: %w", err)
	}
	return nil
}

// FindUnsetEnvVar using key as env name to check and return the list of unset environement variable
func FindUnsetEnvVar(envs map[string]interface{}) []string {
	envNotSet := []string{}

	for key := range envs {
		_, present := os.LookupEnv(key)
		if !present {
			envNotSet = append(envNotSet, key)
		}
	}
	return envNotSet
}

// ParseEnvVarMap using key as env variable name and value as dest in the ParseEnvVar function.
func ParseEnvVarMap(data map[string]interface{}) error {
	var e error
	for name, dest := range data {
		err := ParseEnvVar(name, dest)
		if err != nil && !errors.Is(err, ErrNoEnv) {
			if e == nil {
				e = fmt.Errorf("%w", err)
				continue
			}
			e = fmt.Errorf("%w, %v", err, e)
		}
	}

	return e
}

// ParseEnvVar with the given name and use the dest type to convert the value of the env variable.
// dest MUST be a pointer !
func ParseEnvVar(name string, dest interface{}) error {
	envVar, present := os.LookupEnv(name)
	if !present {
		return fmt.Errorf("the variable %v is not present. Error : %w", name, ErrNoEnv)
	}

	switch d := dest.(type) {
	case *string:
		*d = envVar

	case *int:
		i, err := strconv.Atoi(envVar)
		if err != nil {
			return fmt.Errorf("fail to convert %s=%s to int: %w", name, envVar, err)
		}
		*d = i

	case *int64:
		i, err := strconv.ParseInt(envVar, 10, 64)
		if err != nil {
			return fmt.Errorf("fail to convert %s=%s to int64: %w", name, envVar, err)
		}
		*d = i

	case *bool:
		b, err := strconv.ParseBool(envVar)
		if err != nil {
			return fmt.Errorf("fail to convert %s=%s to bool: %w", name, envVar, err)
		}
		*d = b
	default:
		return fmt.Errorf("unsupported destination type %t", d)
	}

	return nil
}

// SetEnvWithMap set environement variable with key value of map
func SetEnvWithMap(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}

// UnsetEnvWithMap unset environement variable with key map
func UnsetEnvWithMap(envs map[string]string) {
	for key := range envs {
		os.Unsetenv(key)
	}
}
