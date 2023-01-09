package main

func main() {

	//
	//url := "http://9.136.171.180:9271/api/v1/ticket/get_apply_list"
	//method := "GET"
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, nil)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Add("staffname", "v_ljtli")
	//req.Header.Add("w-seq-id", "c567f0d4-bd26-43")
	//req.Header.Add("host", "tmoss.oa.com")
	//req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	//req.Header.Add("content-length", "1462")
	//req.Header.Add("accept", "application/json, text/plain, */*")
	//req.Header.Add("accept-encoding", "gzip")
	//req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	//req.Header.Add("content-type", "application/json")
	//req.Header.Add("cookie", "x-host-key-idcback=3804246a9f7341e3fd6277a3b054fe795f2d7be8_s; x-host-key-oaback=3804246a9f7341e3fd6277a3b054fe795f2d7be8_s; x-host-key-front=3804246a9f7341e3fd6277a3b054fe795f2d7be8_s; x_host_key=3804246a9f7341e3fd6277a3b054fe795f2d7be8_s; x-host-key-ngn=3804246a9f7341e3fd6277a3b054fe795f2d7be8_s; wujiMossLogin=Ne3ujBhFf8Wk6WkQQlV-eg.k1h2hh1JDP87yNU1nWMYs4c36fQN0LuZ7eUOpgUYwUNT9JJ_U1E2KLSjEC6iVDCQVRJPmMW1YMAsvG0Qju32Lu-DRFeSa3cSY-kxbp8iE1pwFgBWcdXVNyf7sQI2kqpodE86AUiJyP8BmpzhhNtIy7rAyypDMerv2ebp4ooC53tBDR6XBVwlxTOuzKoMppsv.1670486045631.86400000.abb8a8ika5QpOJwZaSRnDMT-Hw9QkKpl0-FmUyC35S4; x_host_key_access_https=3804246a9f7341e3fd6277a3b054fe795f2d7be8_s; x-client-ssid=184f45d98e0-eb29945978ff3739e3f428825c0dc0e449d55c5b; x-tofapi-host-key=184f45d98ff-a94e883459a7d4ac34615fe9a3841c140971d4fe; RIO_TCOA_TICKET=tof:TOF4TeyJ2IjoiNCIsInRpZCI6IkhhN2d5WE1PM1lWTDF1T2ZYZkdmTWNFMGhCRnlGbWhLIiwiaXNzIjoiMTAuOTkuMTUuNDkiLCJpYXQiOiIyMDIyLTEyLTA5VDA4OjUxOjE4LjQ5Njc1MzI1NyswODowMCIsImF1ZCI6IjEwLjk2LjIwLjQ2IiwiaGFzaCI6IjAyNkJCNUNENkI4MTlGNkRERTA1NDFCMDQ3MERBMUM0NTQ4RUQ5NTBGODRFNzU2Mjk4QTE5RDE2REM4QjgwNjYiLCJuaCI6IjEzMTYzNEQwRDQ4RTE3NzhDQjNBRjEzMjA3QUNBNDY0RkZEQjEzQjFFQ0UwMkEwMzkwQUNFNzUwQUUxODI2M0IifQ; RIO_TCOA_TICKET_HTTPS=tof:TOF4TeyJ2IjoiNCIsInRpZCI6IkhhN2d5WE1PM1lWTDF1T2ZYZkdmTWNFMGhCRnlGbWhLIiwiaXNzIjoiMTAuOTkuMTUuNDkiLCJpYXQiOiIyMDIyLTEyLTA5VDA4OjUxOjE4LjQ5Njc1MzI1NyswODowMCIsImF1ZCI6IjEwLjk2LjIwLjQ2IiwiaGFzaCI6IjAyNkJCNUNENkI4MTlGNkRERTA1NDFCMDQ3MERBMUM0NTQ4RUQ5NTBGODRFNzU2Mjk4QTE5RDE2REM4QjgwNjYiLCJuaCI6IjEzMTYzNEQwRDQ4RTE3NzhDQjNBRjEzMjA3QUNBNDY0RkZEQjEzQjFFQ0UwMkEwMzkwQUNFNzUwQUUxODI2M0IifQ")
	//req.Header.Add("origin", "https://tmoss.woa.com")
	//req.Header.Add("protocol", "https:")
	//req.Header.Add("qauth-seq", "88ab168b-0f35-4afe-864e-7e0534e58d02")
	//req.Header.Add("qauth-sign", "d842c6deafd44045814504ee559eb8c711f26f967d6272763f47cca435d8e628")
	//req.Header.Add("qauth-timestamp", "1670550793")
	//req.Header.Add("qauth-user-domain", "OA")
	//req.Header.Add("qauth-user-id", "1111")
	//req.Header.Add("qauth-user-name", "v_ljtli(é»é§æ¶)")
	//req.Header.Add("qauth-user-openid", "300505")
	//req.Header.Add("referer", "https://tmoss.woa.com/wuji/xy/project/config/collection/list?projectid=ingame_sy&editingid=sop_test")
	//req.Header.Add("sec-ch-ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\"")
	//req.Header.Add("sec-ch-ua-mobile", "?0")
	//req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	//req.Header.Add("sec-fetch-dest", "empty")
	//req.Header.Add("sec-fetch-mode", "cors")
	//req.Header.Add("sec-fetch-site", "same-origin")
	//req.Header.Add("server-ip", "tmoss.woa.com:80")
	//req.Header.Add("signature", "91284d85140a224dc673c662696237d67206f3f4b45b5f40672cc4b027dc6902")
	//req.Header.Add("staffid", "300505")
	//req.Header.Add("timestamp", "1670550793")
	//req.Header.Add("x-client-ip", "10.96.20.46")
	//req.Header.Add("x-client-ip-port", "10.96.20.46:65284")
	//req.Header.Add("x-client-seq", "20221209_095018%3A135320")
	//req.Header.Add("x-ext-data", "")
	//req.Header.Add("x-forwarded-for", "10.96.20.46, 11.176.16.222, 9.136.158.208")
	//req.Header.Add("x-forwarded-proto", "https")
	//req.Header.Add("x-ngn-connect-url", "9.136.158.208:443")
	//req.Header.Add("x-ngn-network", "intranet")
	//req.Header.Add("x-ngn-platform", "pc")
	//req.Header.Add("x-proxy-by", "SmartGate-IDC")
	//req.Header.Add("x-real-ip", "10.96.20.46")
	//req.Header.Add("x-rio-seq", "9011630a:0184f4962b48:0b5aa3")
	//req.Header.Add("x-sg-chcode", "sg2-ngnpc-s")
	//req.Header.Add("x-sg-ip-chain", "10.99.17.144,10.99.3.239,10.99.3.239,10.99.17.206,11.176.16.222")
	//req.Header.Add("x-sg-seq", "9011630a:0184f4962b48:0b5aa3")
	//req.Header.Add("x-sg-sname", "sg2-ngnpc-s.sgw.woa.com")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
}
