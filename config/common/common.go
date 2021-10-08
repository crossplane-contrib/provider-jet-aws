package common

func ExternalNameAsName(base map[string]interface{}, externalName string) {
	base["name"] = externalName
}
