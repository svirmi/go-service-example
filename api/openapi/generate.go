// Package openapi contains OpenAPI spec and code generated by go-swagger.
package openapi

//go:generate rm -rf model restapi client
//go:generate swagger generate server --api-package op --model-package model --strict-responders --strict-additional-properties --keep-spec-order --principal github.com/powerman/go-service-goswagger-clean-example/internal/app.Auth --exclude-main
//go:generate swagger generate client --api-package op --model-package model --strict-responders --strict-additional-properties --keep-spec-order --principal github.com/powerman/go-service-goswagger-clean-example/internal/app.Auth
//go:generate find restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;
