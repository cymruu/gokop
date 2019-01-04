package gokop

type IClient interface {
	APIKey() string
	Secret() string

	Userkey() string

	APIVersion() string
	Useragent() string
}
