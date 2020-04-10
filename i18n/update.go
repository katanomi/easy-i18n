package i18n

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Update messages
func Update(srcFile string, destFile string) error {
	if len(srcFile) == 0 {
		return fmt.Errorf(Sprintf("srcFile cannot be empty"))
	}

	if len(destFile) == 0 {
		return fmt.Errorf(Sprintf("destFile cannot be empty"))
	}

	srcMessages, err := unmarshal(srcFile)
	if err != nil {
		return err
	}
	dstMessages, err := unmarshal(destFile)
	if err != nil {
		return err
	}

	result := *dstMessages
	for key, value := range *srcMessages {
		if _, ok := result[key]; !ok {
			result[key] = value
		}
	}

	var content []byte
	of := strings.ToLower(destFile)
	if strings.HasSuffix(of, ".json") {
		content, err = marshal(result, "json")
	}
	if strings.HasSuffix(of, ".toml") {
		content, err = marshal(result, "toml")
	}
	if strings.HasSuffix(of, ".yaml") {
		content, err = marshal(result, "yaml")
	}
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destFile, content, 0664)
	if err != nil {
		return nil
	}

	return nil
}