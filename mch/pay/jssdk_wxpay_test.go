package pay

import (
	"fmt"
	"github.com/kprince/gowechat/wxcontext"
	"testing"
)

func TestPay_CheckRefundNotifyData(t *testing.T) {
	str := "<xml><return_code>SUCCESS</return_code><appid><![CDATA[wx2469385b57e0731c]]></appid><mch_id><![CDATA[1579040431]]></mch_id><nonce_str><![CDATA[0d52d12be1e5ff44d511fbc7dc78778a]]></nonce_str><req_info><![CDATA[OpfXkyWatdXZK4x0rSSebzZM22knPNKzR6Fz0AuIHob5SGojgNtH2KXo81vn8kYGpy3fNAAq14j4+9TIBVK44Xslrql3VTz3ItXDFlzYRpWICxfr1O7gQ9ZOFrndFjbU9NlnHX6zJY1C0AFhOdmj8EGHXBUeetxMLaH57zo0k6pTAMh3A1GmQyFc16uYGc9HnXZLV1psAv9W+Jg2OBXUsTToPiBlTZotapW4LUOepgdMIoAFVbA8N9qXooGdBu1YdQiOI9WknaUPmsDEoo3C3FcUb1Mm78IRAM7HNekDSAow8fw2yajAA4bgoST9+rp5tecIp2Op3HxgZIF0oS05zauF5Huz8BNICEe1sPKBcpF+V8x64s8onUQ9P7P3I67uBMQqEFLTzqrMUx3QOX4px3OBIeHDya23DB7RIcTUDHFKBY1xhOtAiLKyhuxVONV9o4X4Vp4ljjoPnxZXtclKcM41OjLUhmsmp4CjIb5achwBu7dK3s+Fv4FooOWm0adHXe1pb4YIgzvb1Nt+I9jRZwWnPVqGPCfbdFmx82PtVda84VPfDh60a+due/LZZz81xTlXXZpHf3Uo+LWFw6WwyceQN4+Xg0t38x8gnnB3FfCJxGVObCXXjIZPklCdw5XexFJ/G4yBHu5y7K6yeZ0UrpEh9oeipQBrvXbAfpY1EECLxWGR1+8bESyiG9eKv+uiuCtyQdSWiNzrO5nT236+0+BFOvMDdHwNvvQ3njlKDTGKltcSiFA4UfeDlVpehLNS/TAAvvKpQiFGCKxBefJQE/wk1176LMacVyzB1i9D1ADfx04n/NtUHR0T+cfhCUgKG0g5SJU8sU5nO6Lng/7G2JGzRLDNJo6SsJ/pSa7yAw3ws4kxlvDM3sZMlOH8y13FQR98ONv8U9ILhpyw+/NX5sjE/cB+D5mGy156eJASLTtqFwzTGZcUEJtKyqpScFTVZpES2S3JwKrfu5FU218qVgD+iiHtUQYTpDzs2DW4Tj8/x/gVzxh3j7yFwNu4+PYSlANxVu3WBq4PCUHxzDMJuaffEw6WCKFpiKV0D+wmIelNCoTj1IRgyev24qTllqW0]]></req_info></xml>"
	pay := Pay{}
	pay.Context = &wxcontext.Context{Config: &wxcontext.Config{MchAPIKey: "1240291536688648192adcedsasdasda",AppID: "wx2469385b57e0731c",MchID: "1579040431"}}
	succ, refundNo, err := pay.CheckRefundNotifyData([]byte(str))
	if err != nil || !succ || refundNo == "" {
		t.Fail()
	}
}

func TestDecodeReqInfo(t *testing.T) {
	str := "OpfXkyWatdXZK4x0rSSebzZM22knPNKzR6Fz0AuIHoa4sj7wIwa970D0Cy9+/WjTpy3fNAAq14j4+9TIBVK44Xslrql3VTz3ItXDFlzYRpXbtcO4As4RD/r7kGVfygoZhn0cPLtoGau43UINhmGFxUGHXBUeetxMLaH57zo0k6pTAMh3A1GmQyFc16uYGc9HnXZLV1psAv9W+Jg2OBXUsTToPiBlTZotapW4LUOepgdMIoAFVbA8N9qXooGdBu1YdQiOI9WknaUPmsDEoo3C3NdRaYT764gaXgiAdpvhSp4w8fw2yajAA4bgoST9+rp5tecIp2Op3HxgZIF0oS05zdA6UcAGddER1ah7C8libW1ks9FMHnaoXALjPB9MuhwWFYEhnq9CFmGHiC4MVLPXmnOBIeHDya23DB7RIcTUDHG18ozmI0VHEKvEuYptV6MXzXXMDv+jqkgQiP9Z99ESIdKPCOa/u5DbVcIrDdIPejz0hhEggdyPv5cw8AembC7Re1TnuGXaEyhexMD/ibcDqhg0iWs1bhK3Y2otfkLBz2diDS2nm8qTSqqVfsNDwhEWVoAHTBcC5m7xVKQNCV8nOcYgbiJIMkoYCKdXn0AiM/Ebo2DLTMqgx4+8KG1b16YV0La4zAh8TkUITUu6j/bmc6euDlU8lj3JWxZMQJ084sXXUWmE++uIGl4IgHab4UqeQJXPvlBRZgQnh5HGGF3AjG9umWELPOLpg8Bx3Uj/PtH/EdKaTHQVD2O6LVf1I0pCxKLO70QSTtgfNl6amYPL2KYYPOgRdqjqwzz+UY2wII3SN8qe6zpsEkJVDk/HmXZbBn1ZgzkfcfP4y0usROYy//+gD3yBvIhX82CvcZ4LuGDQbB1TCZOAv/P/ksLHpiH3zQkPdd4rWS4osPLNPg4R1sSizu9EEk7YHzZempmDy9h9JGa5+DTEbuyZD295pV6agHByGoSVgdITA/XnNV58wdShhEoawsmi1WfFrb9ZkHya32t7lN1VwYVEcA6cCC0l8iuofi/gy7vdre1gktwwpto7kcYMz7FPe7xjPQEu/OCxqYJiTGxtM5Kv3vSs69Ct"
	key := "1240291536688648192adcedsasdasda"
	bts, err := DecodeReqInfo(str, key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bts))

	//b, err := base64.StdEncoding.DecodeString(str)
	//if err != nil {
	//	fmt.Println(b, err)
	//	return
	//}
	//gocrypto.SetAesKey(strings.ToLower(gocrypto.Md5(key)))
	//plaintext, err := gocrypto.AesECBDecrypt(b)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(plaintext))
}
