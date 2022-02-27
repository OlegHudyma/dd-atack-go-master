package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"math/rand"
)

const (
	siteFilePath       = "./sites.json"
	proxyFilePath      = "./proxy.json"
	responseProxyCount = 15
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

	siteToAttackIndex := getRandIntInRange(randGen, len(needAttackSites))

	siteToAttack := needAttackSites[siteToAttackIndex]

	randProxies, err := getNRandProxyFromSlice(randGen, loadedProxies, responseProxyCount)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Site: siteResponse{
			ID:   siteToAttack.ID,
			URL:  siteToAttack.URL,
			Page: siteToAttack.Page,
		},
		Proxy: randProxies,
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}

func getNRandProxyFromSlice(randGenerator *rand.Rand, proxies []proxy, count int) ([]proxy, error) {
	randProxies := make([]proxy, count)

	proxiesLength := len(proxies)

	for i := 0; i < count; i++ {
		randIndex := getRandIntInRange(randGenerator, proxiesLength)

		randProxies[i] = proxies[randIndex]

		// remove randProxy from loadedProxies slice
		proxies = append(proxies[:randIndex], proxies[randIndex+1:]...)
		proxiesLength--
	}

	return randProxies, nil
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
