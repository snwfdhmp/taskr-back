package main

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/rjz/githubhook.v0"
)

var (
	log = logrus.New()
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/webhook", func(rw http.ResponseWriter, req *http.Request) {
		hook, err := githubhook.Parse(appSecret, req)
		if err != nil {
			log.Println("fatal:", err)
			return
		}
		log.Println(hook)
	})

	r.HandleFunc("/callback", func(rw http.ResponseWriter, req *http.Request) {
		log.Infoln("Trigger /callback")
		print(req)
	})

	r.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		log.Infoln("Trigger /")
		print(req)
	})

	panic(http.ListenAndServe(":9876", r))
}

func print(req *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	log.Println(buf.String())
	log.Println(req.Header.Get("x-hub-signature"))
	log.Println(req.Header.Get("x-github-event"))
	log.Println(req.Header.Get("x-github-delivery"))
}

const (
	appSecret = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAq6TyY9yHbly0xlRbWo+pPaBpT597gNBBCjbqOcZLCqbHI+ML
ULrwyUwGNJub5L72hPt2hTcrTTXMsndzMnhKyEBMKg032J/LtH+1z9gieF8go6WD
lExz/xw7h7zvnCZL7MBqsmB5VkB5mvu8g0TTqzGHpXi82LmXt7YKfF1MFkKEind+
W4730VAiLdKevMrBZ8Ylwy9z8MTFKin6h51dDMiFZ9VtDUG3qwjgl0tKQs5hp7aY
JuzZWZ8ENDYFkU036LpGlA8tyJ0I48fJDYS/lE8fzY0NnLVrvCwG1O/HiMvSh3NM
vUVg4PRxqcNHVHvSmYBnpYkU0+X04s1iF6ItrwIDAQABAoIBAD5J6qiws/kp7XR3
0nsn3Uv+9ZiukJwrdx7k1NVIj5z67xOn5khSvuTeEPZwbf9yCXYy50zqu20WlAVD
8esj2keXhcxQ5a5YNw12tx0JG2FBbE/W1cwwfnva3AOjXfT1tOHi5hV7iyzw0vCa
YEvm40WiAM7c5PNlTpidmGqPVPhSbKewabhl7xSoGSRxQoZKdHKzwO5fa1ETacYn
57tLobSSh0NpKMYA/TEv8w9p2zyfFfFRx6iNHmv+s0en1ywXw7HdVzWnhtKHaigd
8wQA4OAP+lAQHse+En4NIEKe16J61E4hNvspzdHGZnHPWfsk/XLa3Xqa8FCzssSI
9PYZjGECgYEA4+8RBjs5uEmwuWUixC7ZQCvl/QNix+mmwJ/RqST+tbDiJGnHDoSc
TLu+j7y3OjWNO/d7NNGApOj6vaQEJX4oE48EH3mSh8t/a+f6Ly3crvjBWyGXUtV3
WJrzm0hpLcboLaNarO9gJ68XjXJAgwzTyiRjJQrmgNNzZ0+P3IN16l8CgYEAwMeF
aUupatv3p5jFprLJ6ux/8Z1cmcRJWu0TZbSa6yuAo6/O4ZrqS1yu5xU7dRESfXFf
jiY21lcIaPuQeJnr9IevJQPq65wQop5DTctV0YrDdzwxVN3F6eyDY1E0TqHL+fWS
QZJ9wpdEnycek+a0lpY+E3PkFrD0i3D+BQjcHrECgYEAiYZW6UGWw+CCZnnlDiWO
Iy/ZZl1CnogVeqdzIvEVVZ7R81J3LgVJIOsuQR+GFfL1sjbR+b5mNhn0H7P0ZILH
v3VM8Yiypohb60leJFUxj8i7MTcFTI9LoQRs80YvX9VhhRfIwAr2Izt7pHZJAkkd
CBwITfNTvGrYeH8Ct/QBk1UCgYEAr58UcvByzt7FpvEOXubR/CqoBD0PdyXSjfsc
iLMXztY6wuQO3Ih8mVs+G1bEa8m5xn3aR22Y13tCNiG3Y6cVxhc1/8qp4Sq0HPGA
7wYmkFaR67XyWxYOjyTgLUJ1HI1ZfcWyD/qchwG3iiXhPWwXaHhG5QHW8LUUk1PE
lmGT5oECgYACBm122fP/14zbc7E7UR+oeS7xcDy7keI4eY6oa/5A73DgH6FfKPEm
93bg0LIVJgj2OMiwaXc9Vzp/eApwcOK6VZjLGEIzFU8limsAS49aZAfCcWFbZCJj
VFN9+HzThCaM53wNMeYl55uaTJzen/feVcPhsn+3kHuSxkjiInIn2A==
-----END RSA PRIVATE KEY-----
`
)
