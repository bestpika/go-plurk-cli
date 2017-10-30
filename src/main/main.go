package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	plurgo "github.com/clsung/plurgo/plurkgo"
	"github.com/garyburd/go-oauth/oauth"
)

var (
	c string
)

func init() {
	flag.StringVar(&c, "c", "config.json", "載入設定檔")
	flag.Parse()
}

func main() {
	// 認證
	tok := plurkAuth(&c)
	// 取得使用者資料
	opt := map[string]string{}
	opt["include_plurks"] = "false"
	ans, _ := callAPI(tok, "/APP/Profile/getOwnProfile", opt)
	plurker := plurkerObj{} // 使用者
	json.Unmarshal(ans, &plurker)
	printObjIndent(plurker)
	// TODO: 命令列功能
}

func plurkAuth(credPath *string) *oauth.Credentials {
	plurkOAuth, e := plurgo.ReadCredentials(*credPath)
	if e != nil {
		log.Fatalf("%+v", e)
	}
	tok, auth, e := plurgo.GetAccessToken(plurkOAuth)
	if auth {
		b, e := json.MarshalIndent(plurkOAuth, "", "  ")
		if e != nil {
			log.Fatal(e)
		}
		e = ioutil.WriteFile(c, b, 0700)
		if e != nil {
			log.Fatal(e)
		}
	}
	return tok
}

func callAPI(tok *oauth.Credentials, api string, opt map[string]string) ([]byte, error) {
	ans, e := plurgo.CallAPI(tok, api, opt)
	if e != nil {
		log.Fatal(e)
	}
	return ans, e
}

func printJSONIndent(data []byte, indent string) {
	var jsi bytes.Buffer
	json.Indent(&jsi, []byte(data), "", indent)
	fmt.Printf("\n%s\n\n", jsi.Bytes())
}

func printObjIndent(data interface{}) {
	ans, _ := json.Marshal(data)
	printJSONIndent(ans, "  ")
}
