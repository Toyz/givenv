package providers

type ProviderInterface interface {
	Get(secretId string) (error, map[string]string)
}

var providerMap = make(map[string]ProviderInterface)

func InitProvider(provider string) ProviderInterface {
	if method, ok := providerMap[provider]; ok {
		return method
	}

	return nil
}