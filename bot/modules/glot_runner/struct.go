package glotrunner

type request struct {
	Stdin string        `json:"stdin"`
	Files []requestFile `json:"files"`
}

type requestFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type response struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

type glotLang struct {
	Name string
	Ext  string
	File string
}

var glotLangs = []glotLang{
	{Name: "ats", Ext: "dats"},
	{Name: "cobol", Ext: "cob"},
	{Name: "crystal", Ext: "cr"},
	{Name: "d", Ext: "d"},
	{Name: "elixir", Ext: "ex"},
	{Name: "elm", Ext: "elm"},
	{Name: "idris", Ext: "idr"},
	{Name: "mercury", Ext: "m"},
	{Name: "assembly", Ext: "asm"},
	{Name: "nim", Ext: "nim"},
	{Name: "nix", Ext: "nix"},
	{Name: "ocaml", Ext: "ml"},
	{Name: "raku", Ext: "raku"},
	{Name: "bash", Ext: "sh"},
	{Name: "c", Ext: "c"},
	{Name: "clojure", Ext: "clj"},
	{Name: "coffeescript", Ext: "coffe"},
	{Name: "cpp", Ext: "cpp"},
	{Name: "csharp", Ext: "cs"},
	{Name: "erlang", Ext: "erl"},
	{Name: "fsharp", Ext: "fs"},
	{Name: "go", Ext: "go"},
	{Name: "groovy", Ext: "groovy"},
	{Name: "haskell", Ext: "hs"},
	{Name: "java", Ext: "java", File: "Main"},
	{Name: "javascript", Ext: "js"},
	{Name: "julia", Ext: "jl"},
	{Name: "kotlin", Ext: "kt"},
	{Name: "lua", Ext: "lua"},
	{Name: "perl", Ext: "pl"},
	{Name: "php", Ext: "php"},
	{Name: "python", Ext: "py"},
	{Name: "ruby", Ext: "rb"},
	{Name: "rust", Ext: "rs"},
	{Name: "scala", Ext: "scala"},
	{Name: "swift", Ext: "swift"},
	{Name: "typescript", Ext: "ts"},
}
