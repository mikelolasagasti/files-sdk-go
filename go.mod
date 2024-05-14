module github.com/Files-com/files-sdk-go/v3

go 1.21

require (
	github.com/appscode/go-querystring v0.0.0-20170504095604-0126cfb3f1dc
	github.com/bradfitz/iter v0.0.0-20191230175014-e8f45d346db8
	github.com/chilts/sid v0.0.0-20190607042430-660e94789ec9
	github.com/dnaeon/go-vcr v1.2.0
	github.com/fatih/structs v1.1.0
	github.com/gin-gonic/gin v1.10.0
	github.com/hashicorp/go-retryablehttp v0.7.6
	github.com/itchyny/timefmt-go v0.1.5
	github.com/lpar/date v1.0.0
	github.com/panjf2000/ants/v2 v2.9.1
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06
	github.com/samber/lo v1.39.0
	github.com/snabb/httpreaderat v1.0.1
	github.com/stretchr/testify v1.9.0
	github.com/tunabay/go-infounit v1.1.3
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842
	golang.org/x/text v0.15.0
	moul.io/http2curl/v2 v2.3.0
)

require (
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-44487
replace golang.org/x/net => golang.org/x/net v0.25.0
