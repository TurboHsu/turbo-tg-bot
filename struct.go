package main

type configStruct struct {
	BotToken       string `json:"tg_bot_token"`
	SaucenaoAPIKey string `json:"saucenao_api_key"`
	Debug          bool   `json:"debug"`
	Silent         bool   `json:"silent"`
	GlotAPIKey     string `json:"glot_api_key"`
	Glot           []struct {
		Name string `json:"name"`
		Ext  string `json:"ext"`
		File string `json:"file"`
	} `json:"glot"`
}

type glotRequest struct {
	Stdin string            `json:"stdin"`
	Files []glotRequestFile `json:"files"`
}

type glotRequestFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type glotResponse struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

type saucenaoJSONStruct struct {
	Header struct {
		UserID           string `json:"user_id"`
		AccountType      string `json:"account_type"`
		ShortLimit       string `json:"short_limit"`
		LongLimit        string `json:"long_limit"`
		LongRemaining    int    `json:"long_remaining"`
		ShortRemaining   int    `json:"short_remaining"`
		Status           int    `json:"status"`
		ResultsRequested int    `json:"results_requested"`
		Index            struct {
			Num0 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"0"`
			Num2 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"2"`
			Num5 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"5"`
			Num6 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"6"`
			Num8 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"8"`
			Num9 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"9"`
			Num10 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"10"`
			Num11 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"11"`
			Num12 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"12"`
			Num16 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"16"`
			Num18 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"18"`
			Num19 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"19"`
			Num20 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"20"`
			Num21 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"21"`
			Num22 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"22"`
			Num23 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"23"`
			Num24 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"24"`
			Num25 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"25"`
			Num26 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"26"`
			Num27 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"27"`
			Num28 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"28"`
			Num29 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"29"`
			Num30 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"30"`
			Num31 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"31"`
			Num32 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"32"`
			Num33 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"33"`
			Num34 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"34"`
			Num35 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"35"`
			Num36 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"36"`
			Num37 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"37"`
			Num38 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"38"`
			Num39 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"39"`
			Num40 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"40"`
			Num41 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"41"`
			Num42 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"42"`
			Num51 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"51"`
			Num52 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"52"`
			Num53 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"53"`
			Num211 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"211"`
			Num341 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"341"`
			Num371 struct {
				Status   int `json:"status"`
				ParentID int `json:"parent_id"`
				ID       int `json:"id"`
				Results  int `json:"results"`
			} `json:"371"`
		} `json:"index"`
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
			YandereID  int      `json:"yandere_id"`
			GelbooruID int      `json:"gelbooru_id"`
			KonachanID int      `json:"konachan_id"`
			Creator    string   `json:"creator"`
			Material   string   `json:"material"`
			Characters string   `json:"characters"`
			Source     string   `json:"source"`
		} `json:"data"`
	} `json:"results"`
}
