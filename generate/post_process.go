package generate

import (
	"strconv"

	"github.com/gunk/gunk/config"
)

// parseBoolParam returns true if the given parameter name exists and is a truthful value.
func parseBoolParam(paramName string, gen config.Generator) (bool, error) {
	for _, param := range gen.Params {
		if param.Key == paramName {
			result, err := strconv.ParseBool(param.Value)
			if err != nil {
				return false, err
			}
			return result, nil
		}
	}
	return false, nil
}

// postProcess processes the input file before writing to output file.
func postProcess(input []byte, gen config.Generator) ([]byte, error) {
	const (
		protocGenGoCmd           = "protoc-gen-go"
		jsonTagPostprocParamName = "json_tag_postproc"
	)

	switch gen.Command {
	case protocGenGoCmd:
		jsonTagPostprocEnabled, err := parseBoolParam(jsonTagPostprocParamName, gen)
		if err != nil {
			return nil, err
		}
		if jsonTagPostprocEnabled {
			return jsonTagPostProcessor(input)
		}
	}

	return input, nil
}
