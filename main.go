package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"math/rand"
)

const (
	siteFilePath       = "./sites.json"
	proxyFilePath      = "./proxy.json"
	responseProxyCount = 20
	responseSitesCount = 20
)

var (
	loadedSites   []site
	loadedProxies []proxy
)

func init() {
	var err error
	loadedSites, err = getSitesFromFile(siteFilePath)
	check(err)

	loadedProxies, err = getProxiesFromFile(proxyFilePath)
	check(err)
}

func HandleRequest() (*Response, error) {
	var needAttackSites []site

	for _, s := range loadedSites {
		if s.Attack == 1 {
			needAttackSites = append(needAttackSites, s)
		}
	}

	var src cryptoSource
	randGen := rand.New(src)

	randSites, err := getNRandSitesFromSlice(randGen, responseSitesCount)
	if err != nil {
		return nil, err
	}

	randProxies, err := getNRandProxyFromSlice(randGen, responseProxyCount)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Sites: randSites,
		Proxy: randProxies,
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}

func getNRandProxyFromSlice(randGenerator *rand.Rand, count int) ([]proxy, error) {
	randProxies := make([]proxy, count)

	proxiesToProceed := make([]proxy, len(loadedProxies))

	_ = copy(proxiesToProceed, loadedProxies)

	proxiesLength := len(proxiesToProceed)

	for i := 0; i < count; i++ {
		randIndex := getRandIntInRange(randGenerator, proxiesLength)

		randProxies[i] = proxiesToProceed[randIndex]

		// remove randProxy from loadedProxies slice
		proxiesToProceed = append(proxiesToProceed[:randIndex], proxiesToProceed[randIndex+1:]...)
		proxiesLength--
	}

	return randProxies, nil
}

func getNRandSitesFromSlice(randGenerator *rand.Rand, count int) ([]site, error) {
	sitesToProceed := make([]site, len(loadedSites))

	_ = copy(sitesToProceed, loadedSites)

	randSites := make([]site, count)

	proxiesLength := len(sitesToProceed)

	for i := 0; i < count; i++ {
		randIndex := getRandIntInRange(randGenerator, proxiesLength)

		randSites[i] = sitesToProceed[randIndex]

		// remove randProxy from loadedProxies slice
		sitesToProceed = append(sitesToProceed[:randIndex], sitesToProceed[randIndex+1:]...)
		proxiesLength--
	}

	return randSites, nil
}

func getSitesFromFile(filename string) ([]site, error) {
	sites := new([]site)
	siteFile, err := readFromFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(siteFile, sites)
	return *sites, err
}

func getProxiesFromFile(filename string) ([]proxy, error) {
	proxies := new([]proxy)
	proxyFile, err := readFromFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(proxyFile, proxies)
	return *proxies, err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
