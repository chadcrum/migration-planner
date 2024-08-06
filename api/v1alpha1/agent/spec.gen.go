// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/kubev2v/migration-planner/api/v1alpha1"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RXQVPjOBP9Ky5939GJA8spN2CZ3dQuDAUFc6A4CKsTa7BbWqlNiqX837ckxbETKyEw",
	"M7dEanW/fv3ULb+xXFVaISBZNn1jNi+g4v7nhTHKuB/aKA2GJPjlCqzlC3A/BdjcSE1SIZsG+6TdThm9",
	"amBTZslIXLCmSZmBf2ppQLDpw9rNY5OyGc4NH0YSnLglZcI/SVDZodHcAJxzzXNJr3+cuZVVXIkECzCs",
	"SRkp4uW7Rn7l7R3YfnfoMd3G8bjOXz19h5y6CIwbw1/d/0JZstdgzsvaEpiNJHfh604j0FKZ533cIK9i",
	"KXW5AtaVy8oSR8GNYCkT0pk91QSil8R+NnycQzIOtIV07Z5a/emYie1vx++Mt50P6e1Rlva1FQM+wxdA",
	"UuZ1SKpsxfp/A3M2Zf/LuhuUra5PFhTdpOwl1GWf7f2lHSTmjqWrUDF8t6o2OQzB5QY4gTglfzWUqTix",
	"qUsWRiSryK1M3REBSJKXd6aMqkWK+HKfo/1ktIZNuluTljjVtq9KVDTKFSLkTowpW3JJEhejuTKjDrUj",
	"CnyrStmCUwHO4UiidJujDmTKaj0iNXJkRJTdApjhXEXx1Vp8jNqtkkrR3pR1rhsx+4RuVyXtFbaPZLc0",
	"br3jO28Zlcl7Nf9UcbsafozdLari9GyCjqV+f3kD1qd/ZoA/C7XEYe6FtKQWhlfxgfPBLlxJvOdlDXFr",
	"S6APaGNrJ6sToT3FO6prcns65xdlLuXCcOJPJRxq901S8Y0blLiw+89cKdrvfiuzjuwWeiz+VywlQnxr",
	"Pt+514e9x8WG2cDdRkYDL5u7q8Nx6UXmb67r8/b5sn8CDHXb+FH8HF4rnzhcedhSYb+wXAjp1nh5vQF0",
	"p867BLHPxA1wq/BHXaofdWB49Wl+3rtIB92iw69Q7OHChqHSTjNtemsdeMJ2FCJW76FMG9/YQwcuZQ5o",
	"oXsmslPN8wKS4/HEzRg3F1hBpO00y5bL5Zj77bEyi2x11mZ/z84vrm4vRsfjybigqvTESHKcscsWUHJd",
	"ckQwyen1LBklfAFICaDQSqIbZy9gbPiCqFHAXCIIrw4NyLVkU/bbeDI+YinTnApPbsa1zF6OslBWm71J",
	"0WTd7NE1Db9ODOiS55AEs0TNEyogsRpyOZcgkuCL+bgB90ywKbsJx/pD1SMxvILwiH3YjjT7fe299Snd",
	"uoPfPgGm4TnQSYJMDenq6ys2IB+DMVg6U8IP5VwhAfpMudalzD3o7LtV2H3IvXc5Io+FZlOrDphfsFq5",
	"mjuPx5PJT0YQom7y+PUvJ4OTnxgrfNFGQp1xkdwEekPMo18f8w55TYUy8t+g95PJya8PeqXoi6pR+IZE",
	"3PWqB7aS6aO3t2BeWl2HJpCx5rH5LwAA//9h+0QIKBAAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../openapi.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
