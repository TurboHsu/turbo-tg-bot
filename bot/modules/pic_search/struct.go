package picsearch

type saucenaoUserStat struct {
	ShortLimit     int
	LongLimit      string
	ShortRemaining int
	LongRemaining  int
}

type saucenaoResponse struct {
	Header struct {
		UserID            string  `json:"user_id"`
		AccountType       string  `json:"account_type"`
		ShortLimit        string  `json:"short_limit"`
		LongLimit         string  `json:"long_limit"`
		LongRemaining     int     `json:"long_remaining"`
		ShortRemaining    int     `json:"short_remaining"`
		Status            int     `json:"status"`
		ResultsRequested  int     `json:"results_requested"`
		SearchDepth       string  `json:"search_depth"`
		MinimumSimilarity float64 `json:"minimum_similarity"`
		QueryImageDisplay string  `json:"query_image_display"`
		QueryImage        string  `json:"query_image"`
		ResultsReturned   int     `json:"results_returned"`
	} `json:"header"`
	Results []struct {
		Header struct {
			Similarity string `json:"similarity"`
			Thumbnail  string `json:"thumbnail"`
			IndexID    int    `json:"index_id"`
			IndexName  string `json:"index_name"`
			Dupes      int    `json:"dupes"`
			Hidden     int    `json:"hidden"`
		} `json:"header"`
		Data struct {
			ExtUrls    []string `json:"ext_urls"`
			Title      string   `json:"title"`
			PixivID    int      `json:"pixiv_id"`
			MemberName string   `json:"member_name"`
			MemberID   int      `json:"member_id"`
		} `json:"data"`
	} `json:"results"`
}
