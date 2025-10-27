package wowio

type Character struct {
	Name string `json:"name"`
	Class string `json:"class"`
	Spec string `json:"active_spec_name"`
	Role string `json:"active_spec_role"`
	Race string `json:"race"`
	Realms string `json:"realm"`
	Faction string `json:"faction"`
	BestRuns []Run `json:"mythic_plus_best_runs"`
	Seasons []Season `json:"mythic_plus_scores_by_season"`
	LastCrawl string `json:"last_crawled_at"`
}

type Run struct {
	Dungeon string `json:"dungeon"`
	ShortName string `json:"short_name"`
	MythicLevel int `json:"mythic_level"`
	Score float32 `json:"score"`
	Chests int `json:"num_keystone_upgrades"`
}

type Season struct {
	Season string `json:"season"`
	Scores map[string]float32 `json:"scores"`
}

func (character *Character) Score() float32 {
	return character.Seasons[0].Scores["all"]
}
