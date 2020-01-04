package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

func MapHandler(pathsToUrl map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrl[path]; ok {
			http.Redirect(w, r,dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHanlder(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse the YAML
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	// 2. Convert YAML array to map
	pathsToUrls := buildMap(pathUrls)

	// 3. return a map handler using the map
	return MapHandler(pathsToUrls, fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.Url
	}

	return pathsToUrls
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url string	`yaml:url`
}