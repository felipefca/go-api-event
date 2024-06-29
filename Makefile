swagger: 
 swag init --parseDependency --parseInternal -d cmd,internal/controllers -o api --ot go,json --md api
.PHONY: swagger