package pay

import (
	"fmt"
	"testing"
)

func TestDecodeReqInfo(t *testing.T) {
	str := "OpfXkyWatdXZK4x0rSSebzZM22knPNKzR6Fz0AuIHoa4sj7wIwa970D0Cy9+/WjTpy3fNAAq14j4+9TIBVK44Xslrql3VTz3ItXDFlzYRpXbtcO4As4RD/r7kGVfygoZhn0cPLtoGau43UINhmGFxUGHXBUeetxMLaH57zo0k6pTAMh3A1GmQyFc16uYGc9HnXZLV1psAv9W+Jg2OBXUsTToPiBlTZotapW4LUOepgdMIoAFVbA8N9qXooGdBu1YdQiOI9WknaUPmsDEoo3C3NdRaYT764gaXgiAdpvhSp4w8fw2yajAA4bgoST9+rp5tecIp2Op3HxgZIF0oS05zdA6UcAGddER1ah7C8libW1ks9FMHnaoXALjPB9MuhwWFYEhnq9CFmGHiC4MVLPXmnOBIeHDya23DB7RIcTUDHG18ozmI0VHEKvEuYptV6MXzXXMDv+jqkgQiP9Z99ESIdKPCOa/u5DbVcIrDdIPejz0hhEggdyPv5cw8AembC7Re1TnuGXaEyhexMD/ibcDqhg0iWs1bhK3Y2otfkLBz2diDS2nm8qTSqqVfsNDwhEWVoAHTBcC5m7xVKQNCV8nOcYgbiJIMkoYCKdXn0AiM/Ebo2DLTMqgx4+8KG1b16YV0La4zAh8TkUITUu6j/bmc6euDlU8lj3JWxZMQJ084sXXUWmE++uIGl4IgHab4UqeQJXPvlBRZgQnh5HGGF3AjG9umWELPOLpg8Bx3Uj/PtH/EdKaTHQVD2O6LVf1I0pCxKLO70QSTtgfNl6amYPL2KYYPOgRdqjqwzz+UY2wII3SN8qe6zpsEkJVDk/HmXZbBn1ZgzkfcfP4y0usROYy//+gD3yBvIhX82CvcZ4LuGDQbB1TCZOAv/P/ksLHpiH3zQkPdd4rWS4osPLNPg4R1sSizu9EEk7YHzZempmDy9h9JGa5+DTEbuyZD295pV6agHByGoSVgdITA/XnNV58wdShhEoawsmi1WfFrb9ZkHya32t7lN1VwYVEcA6cCC0l8iuofi/gy7vdre1gktwwpto7kcYMz7FPe7xjPQEu/OCxqYJiTGxtM5Kv3vSs69Ct"
	key := "1240291536688648192adcedsasdasda"
	bts,err:=DecodeReqInfo(str, key)
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