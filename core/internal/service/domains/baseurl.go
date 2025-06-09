package domains

import (
	v1 "billionmail-core/api/domains/v1"
	"billionmail-core/internal/consts"
	"billionmail-core/internal/service/public"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sync"
)

var (
	baseurl    = ""
	baseurlMap = make(map[string]string)
)

// GetBaseURL get baseurl of console panel
func GetBaseURL() string {
	return baseurl
}

// GetBaseURLBySender get baseurl by sender email address
func GetBaseURLBySender(sender string) string {
	err := g.Validator().Data(sender).Rules("email").Run(context.Background())

	if err != nil {
		g.Log().Warning(context.Background(), "GetBaseURLBySender --> Invalid email address", sender, err)
		return GetBaseURL()
	}

	if s, ok := baseurlMap[sender]; ok {
		return s
	}

	return GetBaseURL()
}

func UpdateBaseURL(ctx context.Context, domain ...string) {
	g.Log().Debug(context.Background(), "UpdateBaseURL --> Starting")
	defer func() {
		g.Log().Debug(context.Background(), "UpdateBaseURL --> Ending")
	}()

	var domains []string

	if len(domain) > 0 {
		domains = domain
	} else {
		ds, err := All(ctx)

		if err != nil {
			return
		}

		for _, d := range ds {
			domains = append(domains, d.Domain)
		}
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		baseurl = buildBaseURL("")
	}()

	for _, d := range domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			g.Log().Debug(ctx, "UpdateBaseURL --> Updating base URL for domain:", domain)
			baseurlMap[domain] = buildBaseURL(domain)
		}(d)
	}

	wg.Wait()
}

func buildBaseURL(hostname string) (s string) {

	s := "https://mail2.psychiatr.ru"

	return
}
