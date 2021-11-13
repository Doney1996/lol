package rpc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Post  获取lol战绩
func Post(uid string, tgpTicket string) Res {
	var url = "https://www.wegame.com.cn/api/v1/wegame.pallas.game.LolBattle/GetBattleList"
	var jsonReq = `{"account_type":2,"id":"&&uid&&","area":8,"offset":0,"count":1,"filter":"","from_src":"lol_helper"}`
	jsonReq = strings.Replace(jsonReq, "&&uid&&", uid, -1)

	client := &http.Client{}

	request, err := http.NewRequest("POST", url, strings.NewReader(jsonReq))
	//增加header选项
	request.Header.Add("Host", "www.wegame.com.cn")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Accept", "image/sharpp,application/json, text/plain, */*")
	request.Header.Add("Content-Type", "application/json;charset=UTF-8")
	request.Header.Add("Origin", "https://www.wegame.com.cn")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.124 Safari/537.36 qblink wegame.exe WeGame/5.0.5.10130 QBCore/3.70.101.400 QQBrowser/9.0.2524.400")
	request.Header.Add("trpc-caller", "wegame.pallas.web.LolBattle")
	request.Header.Add("Referer", "https://www.wegame.com.cn/helper/lol/record/profile.html")
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.5;q=0.7")
	request.Header.Add("Cookie", "pgv_pvid=9213159304; ts_uid=3968525854; PTTuserFirstTime=1628208000000; weekloop=43-44-45-46; isHostDate=18944; puin=995068823; pt2gguin=o0995068823; uin=o0995068823; tgp_id=11513523; geoid=45; lcid=2052; tgp_env=online; tgp_user_type=0; colorMode=1; pkey=0001618EB31100702EA7C8F54DFC4AE99D83193F6C506112FB5D628B974E56DE65F1DA58D9F2271736C03199C6B6EF72471E77DCB734A002CEC189F6347C4F1F20121181ED8089FD8CBBF50D3E0B8D488B3747EADC645EEAD8302C356F513A5B5CEE90C8F812EE3129CBFB81AAABD85E4F1630C39BBDE455; tgp_biz_ticket=0100000000000000001B0D35EEF8B7218F03CD420BA6CDF659A6AB21D93BF7C9711E13095022B236EE0920AC21E20469DF286E633952F1A68BB8AED468A74EA745E6D6AE9426B8F67D; ssr=0; colorMode=1; BGTheme=[object Object]; pgv_info=ssid=s4279457552; language=zh_CN; ts_last=www.wegame.com.cn/helper/lol/record/profile.html; tgp_ticket=6ACDCB8422BA13DA888303110E49C1777053435D85035C94F3CC7D5E9ECB008C3785E2922E11A94B443FE324F9A199CB86A7899EF0417A2D97DA29F8604666A38D4B5107CB1348242396E3BBEE40BA13A3B4D78DA8B07E0C9290B3E18BE0D5957A96C835F5CC091645195776155AC76F9F7B44C36CB13B7C2428A7C0FF8558B5; region=CN")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()
	dataBts, _ := ioutil.ReadAll(response.Body)
	var result = Res{}
	err = json.Unmarshal(dataBts, &result)
	return result
}

type Res struct {
	Battles []Info `json:"battles""`
}

type Info struct {
	GameId                     string `json:"game_id"`
	GameStartTime              string `json:"game_start_time"`
	GameTimePlayed             int    `json:"game_time_played"`
	MapId                      int    `json:"map_id"`
	GameQueueId                int    `json:"game_queue_id"`
	WasMvp                     int    `json:"was_mvp"`
	WasSvp                     int    `json:"was_svp"`
	WasEarlySurrender          int    `json:"was_early_surrender"`
	PlayGt25Mask               int    `json:"play_gt25_mask"`
	GameServerVersion          string `json:"game_server_version"`
	ChampionId                 int    `json:"champion_id"`
	Position                   string `json:"position"`
	SkinIndex                  int    `json:"skin_index"`
	GameScore                  int    `json:"game_score"`
	TeamId                     string `json:"team_id"`
	Win                        string `json:"win"`
	Kills                      int    `json:"kills"`
	Deaths                     int    `json:"deaths"`
	Assists                    int    `json:"assists"`
	GoldEarned                 int    `json:"gold_earned"`
	WasSurrender               int    `json:"was_surrender"`
	WasAfk                     int    `json:"was_afk"`
	MostKills                  int    `json:"most_kills"`
	MostAssists                int    `json:"most_assists"`
	MostMinionsKilled          int    `json:"most_minions_killed"`
	MostGoldEarned             int    `json:"most_gold_earned"`
	MostDamageDealtToChampions int    `json:"most_damage_dealt_to_champions"`
	MostTotalDamageTaken       int    `json:"most_total_damage_taken"`
	MostTurretsKilled          int    `json:"most_turrets_killed"`
	DoubleKills                int    `json:"double_kills"`
	TripleKills                int    `json:"triple_kills"`
	QuadraKills                int    `json:"quadra_kills"`
	PentaKills                 int    `json:"penta_kills"`
	UnrealKills                int    `json:"unreal_kills"`
	GameLevel                  string `json:"game_level"`
	WinWithLessTeammate        int    `json:"win_with_less_teammate"`
	TeamMadeSize               int    `json:"team_made_size"`
	BattleType                 int    `json:"battle_type"`
	SubTotal                   int
}
