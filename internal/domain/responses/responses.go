package responses

func JsonResponse(key, value string) map[string]string{
	return map[string]string{key:value}
}
