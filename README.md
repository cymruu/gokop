# Example usage:
	package main

	client := v1.CreateWykopV1API("apikey", "secret", "connectionkey")
	entry, err := client.Index(int64(38276843))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(entry)


# Project