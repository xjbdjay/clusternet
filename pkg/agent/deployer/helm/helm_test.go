package helm

import (
	"fmt"
	"os"
	"testing"

	appsapi "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	"github.com/clusternet/clusternet/pkg/utils"
	"github.com/stretchr/testify/assert"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/registry"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog/v2"
)

var (
	username = "AWS"
	password = "eyJwYXlsb2FkIjoiakVlaUxSamlvelJBRmIzcHIyMFdqNldtUHFGN0JhVHRwdjU2TmhpcldlVERnc1B0OWlBcXRuQmtvcFoxNit6SkdQdHBGUTFEQ3dGOUh0NkZzdnhHUGhieFRJY3BNNDAyQUVZbVI2emlseEtkMHBYdE1uNWdtalltYzAwOUVGd2lGQ2NzT3ZlbUI3WVdnbDlvTGhuTTFxdkpEU2pRUTJTZlZ4RUJBajI3dE9hL1RlNTlzRW5pbFlhWUFoVTZJQjVrNG03TWtMZk9uUlVRd1UwRzhRRDVlenB5R3JER1J2aXJGOFcrUElYaXJQWk9vdWlKOG5SZGNQTGx5YTg3L295aDZzWkREZkNaWHR5c3VpRWJIZC85VzhqSW5RYldOaGJiRWN2ajB4cnVGWDlleWwxUlVsNmZzNE5vb2pOOVMzdTI1WWFWNWlOeUtMMkxMQkRFN0NicnVCK1NnQ3FGZ1pKYVhqUG9NL0F4TlM2NjJ5eUk1U3BJTTdORklrLzhubVVIUnBPaURuaVBZUS8zMmZSNGQxL05NaGNMUHFKVDhWSnFaN01yblZVNnlmSGJHZkJhWDRGNHplem9qNmRwcWhzak93dVFwRHRPNTlVZit0bm9kMzdsOGxCNHhhMFdPRTJvc0I0OTk4WW1EeStNR3lEeG94UEhSVjdSQko1U2ljdjIzdVJERVlta2ZPYkZSSlkzbVpQMmVpWnRYWmY1K2NuSWhoeG02NWhmM09ZNEd6NGIvQkhSTW9JWXh1MUNOdjFBdVhqZG9GUTdncVZXTkpiTHVNTllEV3BYRk1maTNZWUUvdXp4cUxsM2pmZ0VUdUdPbjY2RGlhK2s0U2NCdXF0ZGExZzNuZmU2dTFVa3hNRFpmSHBMa3FUYWh2cWdjU0VxMEJlQ2k4Q3R6V245TUFNcmJrVDRmbmlvZ205aHo4d2haTGxEeWpYQmxLRVY1VGdleHVjRXU0TUkwTmw4dWxzekM5cGp1YmZsdXo1VXRJd1hvcE14SHh0QUhURFhFTE1UT1FKUnQwN1JFc1RTMmFZU1ptTm1oV3hhZkpIWXkzUG5VbXRic1VMUmdHeDNFbm8xd2MxTnNZUHRGS3B6QkZ4elBFb3RZNHMrRTUwejRBTFk0K0ZmQ0J6ZVVUaDdnYmhDZHhQdjFwOTRTN2lsNkJ3dUhzSWhMdEVJRkh5c1pHTlFnQ1g3SHoyTDgycm80b053SjFTeTgyZDlqMmF0V09yVTlBbVdDOEFQbTBtTVBpTEVNV1E5NXNYWXJYMlhHWExxeVFGUVcwRHdXa2xEMGJIeUhpZWw4QVFJSVZpdWlRbmpaSUFuVXNocHlCRVhVVWlDS1NlZnEzV1psMjZtVlk5QVlCRGpEOVZ4M1ZsaSIsImRhdGFrZXkiOiJBUUVEL3pkU3d2MFJZKzlka0N3K0JDZlFHbXpEdk9HQ0NtR0xmNFlKK2R1Z21NVnp4UUFBQU0wQkFJVHd4KytuK1g0QUFQWk1tUDc3ZmdBQWNEY2ZBQUFBQUFDVDVjb0ZBQUFBQUVBb05weno4ZVZWZFZDcFlaZlBIZmhQK2s0bW9qeWwzaTkxdk9Ha2hJWE5rbW5INEJRQUFBQmc1RDBHQndBQUFKamtQUVlIQUFBQUt3QUFBQUFBQUFBb0FBQUF1S09LK2c4M3VBY0V2TXQ1b09oRFBvdzJXZkFqMjVqTWFzeG9vZENta2dFQU1CdnFURFB6U3NTRzBLNDE3TG5CaThXeG5hNUQ2dXdJazBIbkhYcFIvOTRZaWZpYlNkVlNINFVBSytMNkx4TmRFd0FRcGZQZ3Nud25GVGQzRGxONHAzRGFSd0FBIiwidmVyc2lvbiI6IjIiLCJ0eXBlIjoiREFUQV9LRVkiLCJleHBpcmF0aW9uIjoxNjgyMjg0MDk0fQ=="
)

func TestName(t *testing.T) {
	config, err := clientcmd.LoadFromFile("/Users/zhangfangjie/.kube/config")
	assert.Nil(t, err)
	targetNS := "abc"
	deployCtx, err := utils.NewDeployContext(config, &clientcmd.ConfigOverrides{Context: clientcmdapi.Context{
		Namespace: targetNS,
	}})
	Settings := cli.New()

	assert.Nil(t, err)

	registryClient, err := registry.NewClient(
		registry.ClientOptDebug(Settings.Debug),
		registry.ClientOptWriter(os.Stdout),
		registry.ClientOptCredentialsFile(Settings.RegistryConfig),
	)

	assert.Nil(t, err)

	cfg := new(action.Configuration)
	err = cfg.Init(deployCtx, targetNS, "secret", klog.V(5).Infof)
	assert.Nil(t, err)

	cfg.Releases.MaxHistory = 5
	cfg.RegistryClient = registryClient

	var (
		chart *chart.Chart
	)
	username := "xjbdjay"
	password := "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"
	chart, err = utils.LocateAuthHelmChart(cfg, "oci://ghcr.io/xjbdjay", username, password, "metadata", "0.1.0")
	assert.Nil(t, err)
	fmt.Println(chart)
	createNamespace := true
	hr := &appsapi.HelmRelease{
		Spec: appsapi.HelmReleaseSpec{
			ReleaseName:     &username,
			TargetNamespace: targetNS,
			HelmOptions: appsapi.HelmOptions{
				TimeoutSeconds:  10,
				CreateNamespace: &createNamespace,
			},
		},
	}
	r, err := utils.UpgradeRelease(cfg, hr, chart, nil)
	fmt.Println(r)
	assert.Nil(t, err)
}

func TestEcr(t *testing.T) {
	config, err := clientcmd.LoadFromFile("/Users/zhangfangjie/.kube/config")
	assert.Nil(t, err)
	targetNS := "abc"
	deployCtx, err := utils.NewDeployContext(config, &clientcmd.ConfigOverrides{Context: clientcmdapi.Context{
		Namespace: targetNS,
	}})
	Settings := cli.New()

	assert.Nil(t, err)

	registryClient, err := registry.NewClient(
		registry.ClientOptDebug(Settings.Debug),
		registry.ClientOptWriter(os.Stdout),
		registry.ClientOptCredentialsFile(Settings.RegistryConfig),
	)

	assert.Nil(t, err)

	cfg := new(action.Configuration)
	err = cfg.Init(deployCtx, targetNS, "secret", klog.V(5).Infof)
	assert.Nil(t, err)

	cfg.Releases.MaxHistory = 5
	cfg.RegistryClient = registryClient

	var (
		chart *chart.Chart
	)

	// client, err := registry.NewClient(registry.ClientOptDebug(true), registry.ClientOptWriter(os.Stdout))
	// assert.Nil(t, err)
	// err = client.Login("https://635304352795.dkr.ecr.cn-north-1.amazonaws.com.cn", registry.LoginOptBasicAuth(username, password))
	assert.Nil(t, err)
	releaseName := "tst"
	chart, err = utils.LocateAuthHelmChart(cfg, "oci://635304352795.dkr.ecr.cn-north-1.amazonaws.com.cn", username, password, "metadata", "0.1.0")
	assert.Nil(t, err)
	fmt.Println(chart)
	createNamespace := true
	hr := &appsapi.HelmRelease{
		Spec: appsapi.HelmReleaseSpec{
			ReleaseName:     &releaseName,
			TargetNamespace: targetNS,
			HelmOptions: appsapi.HelmOptions{
				TimeoutSeconds:  10,
				CreateNamespace: &createNamespace,
			},
		},
	}
	cfg.Releases.Last(username)
	r, err := utils.InstallRelease(cfg, hr, chart, nil)
	fmt.Println(r)
	assert.Nil(t, err)
}

func TestLogin(t *testing.T) {
	client, err := registry.NewClient(registry.ClientOptDebug(true), registry.ClientOptWriter(os.Stdout))
	assert.Nil(t, err)
	err = client.Login("https://635304352795.dkr.ecr.cn-north-1.amazonaws.com.cn", registry.LoginOptBasicAuth(username, password))
	assert.Nil(t, err)
}

func TestFind(t *testing.T) {
	res, err := utils.FindOCIChart("oci://635304352795.dkr.ecr.cn-north-1.amazonaws.com.cn", username, password, "metadata", "0.1.0")
	assert.Nil(t, err)
	assert.True(t, res)
}
