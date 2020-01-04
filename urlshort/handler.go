package urlshort

import "net/http"

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

func YAMLHanlder(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}