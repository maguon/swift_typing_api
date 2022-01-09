package util

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"swift_typing_api/common"
	"swift_typing_api/conf"
	"time"
)

func SendCaptchaSms(phone string, captcha string) bool {
	smsConfig := conf.Config.Sms
	timeStampStr := time.Now().Format("20060102150405")
	originSignStr := smsConfig.AccountSID + smsConfig.AccountToken + timeStampStr
	signature := fmt.Sprintf("%x", md5.Sum([]byte(originSignStr)))

	originAuthStr := smsConfig.AccountSID + ":" + timeStampStr
	auth := base64.StdEncoding.EncodeToString([]byte(originAuthStr))
	url := "/2013-12-26/" + smsConfig.Type + "/" +
		smsConfig.AccountSID + "/" + smsConfig.Action + "?sig="

	url = url + signature
	url = fmt.Sprintf("https://%s:%d%s", smsConfig.Server, smsConfig.Port, url)

	jsonStr := fmt.Sprintf(`{"to":"%s","appId":"%s","templateId":"%d","datas":[%s,15]}`, phone, smsConfig.AppSID, smsConfig.Id, captcha)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return false
	}

	req.Header.Set("Authorization", auth)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var smsRes struct {
		StatusCode string `json:"statusCode"`
	}
	json.Unmarshal(body, &smsRes)
	if smsRes.StatusCode == "000000" {
		return true
	} else {
		common.GetLogger().Error("send sms error", string(body))
		return false
	}
}
