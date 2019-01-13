package gokop

type IWykopClient interface {
	APIKey() string
	Secret() string

	Userkey() string

	APIVersion() APIVersionT
	APIURL() string

	Useragent() string

	SendRequest(req IWykopRequest) ([]byte, error)
}
