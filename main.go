package main

import (
	"bufio"
	cr "crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getAgent() string {
	agents := []string{
		"Mozilla/5.0 (iPod; CPU iPhone OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A523 Safari/8536.25",
		"Mozilla/5.0 (iPod; CPU iPhone OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10A523",
		"Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn; U8800 Build/HuaweiU8800) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn) AppleWebKit/530.17 (KHTML, like Gecko) FlyFlow/2.2 Version/4.0 Mobile Safari/530.17",
		"Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn; U8800 Build/HuaweiU8800) UC AppleWebKit/534.31 (KHTML, like Gecko) Mobile Safari/534.31",
		"Mozilla/5.0 (Linux; Android 4.0.3; M031 Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
		"Opera/9.80 (Android 4.0.3; Linux; Opera Mobi/ADR-1210241511) Presto/2.11.355 Version/12.10",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M031 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn) AppleWebKit/530.17 (KHTML, like Gecko) FlyFlow/2.2 Version/4.0 Mobile Safari/530.17",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M031 Build/IML74K) UC AppleWebKit/534.31 (KHTML, like Gecko) Mobile Safari/534.31",
		"MQQBrowser/3.7/Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn; U8800 Build/HuaweiU8800) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn; U8800 Build/HuaweiU8800) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (iPad; U; CPU OS 6 like Mac OS X; zh-cn Model:iPad2,1) UC AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9B176 Safari/7543.48.3",
		"Mozilla/5.0 (iPad; CPU OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A523 Safari/8536.25",
		"MQQBrowser/2.7 Mozilla/5.0 (iPad; CPU OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10A523 Safari/7534.48.3",
		"MQQBrowser/3.7/Adr (Linux; U; 2.3.5; zh-cn; U8800 Build/U8800V100R001C00B528G002;480*800)",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M031 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (iPad; CPU OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10A523",
		"Mozilla/5.0 (iPad; CPU OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10A523",
		"Mozilla/5.0 (iPad; U; CPU  OS 4_1 like Mac OS X; en-us)AppleWebKit/532.9(KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7",
		"Mozilla/5.0 (iPad; CPU OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mobile/10A523",
		"MQQBrowser/3.5/Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M9 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M9 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"MQQBrowser/3.5/Adr (Linux; U; 4.0.3; zh-cn; M9 Build/Flyme 1.0.1;640*960)",
		"MQQBrowser/3.7/Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M9 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"MQQBrowser/3.7/Adr (Linux; U; 4.0.3; zh-cn; M9 Build/Flyme 1.0.1;640*960)",
		"MQQBrowser/4.0/Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M031 Build/IML74K) AppleWebKit/533.1 (KHTML, like Gecko) Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M031 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 4.0.4; zh-cn; HTC S720e Build/IMM76D) UC AppleWebKit/534.31 (KHTML, like Gecko) Mobile Safari/534.31",
		"Mozilla/5.0 (Linux; U; Android 4.0.4; zh-cn; HTC S720e Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn; U8800 Build/HuaweiU8800) AppleWebKit/530.17 (KHTML, like Gecko) FlyFlow/2.3 Version/4.0 Mobile Safari/530.17 baidubrowser/042_1.6.3.2_diordna_008_084/IEWAUH_01_5.3.2_0088U/1001a/BE44DF7FABA8768B2A1B1E93C4BAD478%7C898293140340353/1",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; M031 Build/IML74K) AppleWebKit/530.17 (KHTML, like Gecko) FlyFlow/2.3 Version/4.0 Mobile Safari/530.17 baidubrowser/023_1.41.3.2_diordna_069_046/uzieM_51_3.0.4_130M/1200a/963E77C7DAC3FA587DF3A7798517939D%7C408994110686468/1",
		"Mozilla/5.0 (Linux; U; Android 2.3.5; zh-cn; U8800 Build/HuaweiU8800) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 3.2; zh-cn; GT-P6200 Build/HTJ85B) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3) AppleWebKit/534.31 (KHTML, like Gecko) Chrome/17.0.558.0 Safari/534.31 UCBrowser/2.3.1.257",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16",
		"Mozilla/5.0 (Linux; U; Android 4.1.1; zh-cn; M040 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 4.1.1; zh-cn; M040 Build/JRO03H) AppleWebKit/533.1 (KHTML, like Gecko)Version/4.0 MQQBrowser/4.1 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 4.1.1; zh-CN; M031 Build/JRO03H) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/8.8.3.278 U3/0.8.0 Mobile Safari/534.31",
		"Mozilla/5.0 (Linux; U; Android 4.1.1; zh-cn; M031 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (iPhone 5CGLOBAL; CPU iPhone OS 7_0_6 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/6.0 MQQBrowser/5.0.5 Mobile/11B651 Safari/8536.25",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11B554a Safari/9537.53",
		"Mozilla/5.0 (Linux; U; Android 4.1.1; zh-cn; M040 Build/JRO03H) AppleWebKit/534.24 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.24 T5/2.0 baidubrowser/4.2.4.0 (Baidu; P1 4.1.1)",
		"Mozilla/5.0 (Linux; Android 4.1.1; M040 Build/JRO03H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.64 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 4.1.1; M040 Build/JRO03H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/11B554a Safari/7534.48.3",
		"Mozilla/5.0 (Linux; U; Android 4.1.1; zh-CN; M040 Build/JRO03H) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.4.1.362 U3/0.8.0 Mobile Safari/533.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X; zh-CN) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/11B554a UCBrowser/9.3.1.339 Mobile",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/11B554a",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_4 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) CriOS/31.0.1650.18 Mobile/11B554a Safari/8536.25",
		"Mozilla/5.0 (Linux; Android 4.2.1; M040 Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; U; Android 4.2.1; zh-cn; M040 Build/JOP40D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 Maxthon/4.1.3.2000",
		"Mozilla/5.0 (Linux; U; Android 4.2.1; zh-cn; M040 Build/JOP40D) AppleWebKit/534.24 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.24 T5/2.0 baidubrowser/4.1.3.1 (Baidu; P1 4.2.1)",
		"Mozilla/5.0 (Linux; U; Android 4.2.1; zh-cn; M040 Build/JOP40D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		"Mozilla/5.0 (Linux; U; Android 4.2.1; zh-cn; M040 Build/JOP40D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 baidubrowser/4.2.9.2 (Baidu; P1 4.2.1)",
		"Mozilla/5.0 (Linux; U; Android 4.2.1; zh-cn; M040 Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko)Version/4.0 MQQBrowser/5.0 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone 5CGLOBAL; CPU iPhone OS 7_0_5 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/6.0 MQQBrowser/5.0.4 Mobile/11B601 Safari/8536.25",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_5 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Mobile/11B601 baiduboxapp/0_0.0.1.5_enohpi_6311_046/5.0.7_4C2%255enohPi/1099a/0E12BC204E06E175FD283E21BFE1661EE0A20B6CAFNTCGOKCPB/1",
		"Mozilla/5.0 (Linux; U; Android 4.2.1; zh-cn; M040 Build/JOP40D) AppleWebKit/534.24 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.24 T5/2.0 baiduboxapp/5.1 (Baidu; P1 4.2.1)",
		"Mozilla/5.0 (Linux; U; Android 4.2.1; zh-cn; M040 Build/JOP40D) AppleWebKit/534.24 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.24 T5/2.0 baidubrowser/4.3.16.2 (Baidu; P1 4.2.1)",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_6 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11B651 Safari/9537.53",
		"Mozilla/5.0 (Linux; U; Android 4.2.1; zh-CN; M040 Build/JOP40D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.6.0.378 U3/0.8.0 Mobile Safari/533.1",
		"Mozilla/5.0 (iPad; CPU OS 7_1 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/6.0 MQQBrowser/4.0.2 Mobile/11D167 Safari/7534.48.3",
		"Mozilla/5.0 (iPad; CPU OS 7_1 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D167 Safari/9537.53",
		"Mozilla/5.0 (Linux; Android 4.2.1; M040 Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.117 Mobile Safari/537.36 OPR/20.0.1396.72047",
	}

	k, _ := cr.Int(cr.Reader, big.NewInt(100))
	return agents[int(k.Int64())%len(agents)]
}

type AntiClawer struct {
	ProxyIp string
	Agent   string
	Cookie  string
}

func crawlPost(proxyAddr string, address string, params string, agent string, cookie string) []byte {

	// 加入代理IP
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
		DisableKeepAlives:     true,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	// 创建请求
	req, err := http.NewRequest("POST", address, strings.NewReader(params))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 头部添加
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", agent)
	req.Header.Set("Host", "app.gsxt.gov.cn'")
	req.Header.Set("Cookie", cookie)

	// 网络请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	// 获取数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 返回数据
	return body
}

func paInfo(company string, qiyeurl string, proxyAddr1 string) string {
	proxyAddr := "http://" + proxyAddr1
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	req, err := http.NewRequest("POST", qiyeurl+company+".html?nodeNum=120000&entType=1&start=0&sourceType=I", strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; U; Android 2.2.1; zh-cn; HTC_Wildfire_A3333 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1")
	req.Header.Set("User-Agent", getAgent())
	req.Header.Set("Host", "app.gsxt.gov.cn'")
	//req.Header.Set("Cookie", "JSESSIONID=716E87F50A65B64473EAA8F73EA094CE; tlb_cookie=172.16.12.1108080; SECTOKEN=7064461226094101024; __jsluid=05ca2ab69649b8b5bd68b025854917bb")
	req.Header.Set("Cookie", GetCookie())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(body)
}

func searchCompany(company string, qiyeurl string, proxyAddr1 string) []byte {
	proxyAddr := "http://" + proxyAddr1
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	req, err := http.NewRequest("POST", qiyeurl, strings.NewReader("searchword="+company))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; U; Android 2.2.1; zh-cn; HTC_Wildfire_A3333 Build/FRG83D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1")
	req.Header.Set("User-Agent", getAgent())
	req.Header.Set("Host", "app.gsxt.gov.cn'")
	//req.Header.Set("Cookie", "JSESSIONID=716E87F50A65B64473EAA8F73EA094CE; tlb_cookie=172.16.12.1108080; SECTOKEN=7064461226094101024; __jsluid=05ca2ab69649b8b5bd68b025854917bb")
	req.Header.Set("Cookie", GetCookie())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(string(body))
	return body
}

func RandBigStringRunes(n int) string {
	var letterRunes = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandSmallStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandNumberRunes(n int) string {
	var letterRunes = []rune("1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetCookie() string {

	cookie := "JSESSIONID="
	cookie = cookie + RandBigStringRunes(33) + "; "
	cookie = cookie + "tlb_cookie=172.16.12.1108080; "
	cookie = cookie + "SECTOKEN=" + RandNumberRunes(20) + "; "
	cookie = cookie + "__jsluid=" + RandSmallStringRunes(33)
	return cookie
}

func crawlCompany(company string, ac AntiClawer) {

	params := "searchword=" + company
	// 搜索公司
	byt := crawlPost("http://"+ac.ProxyIp, "http://app.gsxt.gov.cn/gsxt/cn/gov/saic/web/controller/PrimaryInfoIndexAppController/search?page=1", params, ac.Agent, ac.Cookie)
	if byt == nil {
		return
	}

	// JSON解析
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		fmt.Println(err)
		return
	}

	// 获取公司ID
	var companyIdent string
	if v1, ok := dat["data"]; ok {
		fmt.Println("1")
		if v2, ok := v1.(map[string]interface{})["result"]; ok {
			fmt.Println("2")
			if v3, ok := v2.(map[string]interface{})["data"]; ok {
				fmt.Println("3")
				if len(v3.([]interface{})) == 0 {
					return
				}
				if v4, ok := v3.([]interface{})[0].(map[string]interface{})["pripid"]; ok {
					companyIdent = v4.(string)
					fmt.Println(v4.(string))
				}
			}
		}
	}

	// 公司没有找到
	if companyIdent == "" {
		return
	}

	// 创建公司数据文件
	f, err := os.Create(company + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	api := []string{
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-primaryinfoapp-entbaseInfo-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-shareholder-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-KeyPerson-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-branch-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-getNeRecItemPubListByPripId-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-liquidation-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-alter-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-mortreginfo-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-stakqualitinfo-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-trademark-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-getDrRaninsRes-",
		"http://app.gsxt.gov.cn/gsxt/Affiche-query-info-assistInfo-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-licenceinfoDetail-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-punishmentdetail-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-entBusExcep-",
		"http://app.gsxt.gov.cn/gsxt/corp-query-entprise-info-illInfo-",
	}

	for _, a := range api {
		time.Sleep(time.Duration(1) * time.Second) // 拿IP需要延时
		data := crawlPost("http://"+ac.ProxyIp, a+companyIdent+".html?nodeNum=120000&entType=1&start=0&sourceType=I", "", ac.Agent, ac.Cookie)
		fmt.Println(string(data))
		w.WriteString(string(data) + "\n")
		w.Flush()
	}

}

func GetIPS(ipNum string) []string {
	url := "http://piping.mogumiao.com/proxy/api/get_ip_bs?appKey=cde47c2f831a4b57918370609f2010e9&count=" + ipNum + "&expiryDate=0&format=2&newLine=2"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("getip")
		// handle error
	}
	//程序在使用完回复后必须关闭回复的主体。
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))

	ips := string(body)

	t := strings.Split(ips, "\r\n")
	return t
}

func GetIPPool(total int) chan string {

	ipsChan := make(chan string, 10)

	go func(ipsChan chan string) {
		for i := 0; i < total; i++ {
			s := GetIPS("10") // 最多一次10个
			for _, k := range s[0 : len(s)-1] {
				ipsChan <- k
			}
			time.Sleep(time.Duration(2) * time.Second) // 拿IP需要延时
		}
	}(ipsChan)

	return ipsChan
}

func LoadCompany() []string {
	b, err := ioutil.ReadFile("companyall.txt")
	if err != nil {
		return nil
	}
	s := string(b)
	return strings.Split(s, "\r\n")
}

func main() {
	// 加载企业数据
	c := LoadCompany()
	for _, key := range c {
		fmt.Println(key)
	}
	// 获取IP

	i := 0
	pool := GetIPPool(1000)
	for ip := range pool {
		fmt.Println(ip)
		//for t := 0; t < 10; t++ { //1个IP爬10个，省些钱
		ac := AntiClawer{
			ProxyIp: ip,
			Agent:   getAgent(),
			Cookie:  GetCookie(),
		}
		go crawlCompany(c[i], ac)
		time.Sleep(time.Duration(1) * time.Second)
		i = i + 1
		//}
	}
}
