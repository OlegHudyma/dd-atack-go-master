package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	siteFilePath       = "./loadedSites.json"
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

	siteToAttackIndex := getRandIntInRange(len(needAttackSites))
	siteToAttack := needAttackSites[siteToAttackIndex]

	randProxies := getNRandProxyFromSlice(loadedProxies, responseProxyCount)

	response := &Response{
		Site:  siteToAttack,
		Proxy: randProxies,
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}

func getNRandProxyFromSlice(proxies []proxy, count int) []proxy {
	randProxies := make([]proxy, count)

	proxiesLength := len(proxies)

	for i := 0; i < count; i++ {
		randIndex := getRandIntInRange(proxiesLength)

		randProxies[i] = proxies[randIndex]

		// remove randProxy from loadedProxies slice
		proxies = append(proxies[:randIndex], proxies[randIndex+1:]...)
		proxiesLength--
	}

	return randProxies
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
