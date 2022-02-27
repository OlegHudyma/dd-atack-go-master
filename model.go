package main

type proxy struct {
	ID   int    `json:"id"`
	IP   string `json:"ip"`
	Auth string `json:"auth"`
}

type site struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
	//NeedParseURL int    `json:"need_parse_url"`
	Page string `json:"page"`
	//PageTime     string `json:"page_time"`
	Attack int `json:"atack"`
}

type siteResponse struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Page string `json:"page"`
}

type Response struct {
	Site  siteResponse `json:"site"`
	Proxy []proxy      `json:"proxy"`
}
