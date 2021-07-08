module simple-ims

go 1.16

replace github.com/zing-dev/soft-version v0.0.0 => ../soft-version

require (
	github.com/Joker/hpp v1.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elliotchance/pie v1.34.0
	github.com/iris-contrib/middleware/jwt v0.0.0-20201115103636-07e8bced147f
	github.com/judwhite/go-svc v1.1.2
	github.com/kataras/iris/v12 v12.2.0-alpha.0.20201113181155-4d09475c290d
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945 // indirect
	github.com/urfave/cli/v2 v2.3.0
	github.com/zing-dev/soft-version v0.0.0
	gorm.io/driver/sqlite v1.0.8
	gorm.io/gorm v0.2.20
)
