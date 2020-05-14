package model

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/aquasecurity/starboard/pkg/kube"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	security "github.com/aquasecurity/starboard/pkg/apis/aquasecurity/v1alpha1"
	"github.com/vmware-tanzu/octant/pkg/plugin/service"
	"github.com/vmware-tanzu/octant/pkg/store"
)

const (
	WorkloadKindPod        = "Pod"
	WorkloadKindDeployment = "Deployment"
	WorkloadKindDaemonSet  = "DaemonSet"
	KindNamespace          = "Namespace"
	KindNode               = "Node"
)

const (
	labelWorkloadKind  = "starboard.workload.kind"
	labelWorkloadName  = "starboard.workload.name"
	labelContainerName = "starboard.container.name"
)

const (
	aquaSecurityAPIVersion = "aquasecurity.github.com/v1alpha1"
	vulnerabilitiesKind    = "Vulnerability"
)

type Workload struct {
	Kind string
	Name string
}

type Repository struct {
	client service.Dashboard
}

func NewRepository(client service.Dashboard) *Repository {
	return &Repository{
		client: client,
	}
}

type ContainerImageScanReport struct {
	Name   string
	Report security.Vulnerability
}

func (r *Repository) GetVulnerabilitiesSummary(ctx context.Context, options Workload) (vs security.VulnerabilitySummary, err error) {
	containerReports, err := r.GetVulnerabilitiesForWorkload(ctx, options)
	if err != nil {
		return vs, err
	}
	for _, cr := range containerReports {
		for _, v := range cr.Report.Report.Vulnerabilities {
			switch v.Severity {
			case security.SeverityCritical:
				vs.CriticalCount++
			case security.SeverityHigh:
				vs.HighCount++
			case security.SeverityMedium:
				vs.MediumCount++
			case security.SeverityLow:
				vs.LowCount++
			default:
				vs.UnknownCount++
			}
		}
	}
	return
}

func (r *Repository) GetVulnerabilitiesForNamespace(ctx context.Context, namespace string) (report ContainerImageScanReport, err error) {
	unstructuredList, err := r.client.List(ctx, store.Key{
		APIVersion: aquaSecurityAPIVersion,
		Kind:       vulnerabilitiesKind,
		Namespace:  namespace,
	})
	if err != nil {
		return
	}
	var reportList security.VulnerabilityList
	err = r.structure(unstructuredList, &reportList)
	if err != nil {
		return
	}

	var vulnerabilities []security.VulnerabilityItem

	for _, i := range reportList.Items {
		if _, containerNameSpecified := i.Labels[labelContainerName]; !containerNameSpecified {
			continue
		}
		vulnerabilities = append(vulnerabilities, i.Report.Vulnerabilities...)
	}

	sort.SliceStable(vulnerabilities, func(i, j int) bool {
		return strings.Compare(vulnerabilities[i].VulnerabilityID, vulnerabilities[j].VulnerabilityID) < 0
	})

	report = ContainerImageScanReport{
		Name: fmt.Sprintf("Namespace %s", namespace),
		Report: security.Vulnerability{
			Report: security.VulnerabilityReport{
				Vulnerabilities: vulnerabilities,
			},
		},
	}

	return
}

func (r *Repository) GetVulnerabilitiesForWorkload(ctx context.Context, options Workload) (reports []ContainerImageScanReport, err error) {
	unstructuredList, err := r.client.List(ctx, store.Key{
		APIVersion: aquaSecurityAPIVersion,
		Kind:       vulnerabilitiesKind,
		// TODO Report bug to Octant? Apparently the label selector doesn't work and I have to do filtering manually in a loop :(
		//Selector: &labels.Set{
		//	labelWorkloadKind: options.Kind,
		//	labelWorkloadName: options.Name,
		//},
	})
	if err != nil {
		err = fmt.Errorf("listing vulnerabilities: %w", err)
		return
	}
	var reportList security.VulnerabilityList
	err = r.structure(unstructuredList, &reportList)
	if err != nil {
		err = fmt.Errorf("unmarshalling JSON to VulnerabilityList: %w", err)
		return
	}
	for _, item := range reportList.Items {
		containerName, containerNameSpecified := item.Labels[labelContainerName]
		if item.Labels[labelWorkloadKind] == options.Kind &&
			item.Labels[labelWorkloadName] == options.Name &&
			containerNameSpecified {
			reports = append(reports, ContainerImageScanReport{
				Name:   fmt.Sprintf("Container %s", containerName),
				Report: item,
			})
		}
	}

	sort.SliceStable(reports, func(i, j int) bool {
		return strings.Compare(reports[i].Name, reports[j].Name) < 0
	})

	return
}

func (r *Repository) GetCISKubeBenchReport(ctx context.Context, node string) (report *security.CISKubeBenchReport, err error) {
	unstructuredList, err := r.client.List(ctx, store.Key{
		APIVersion: aquaSecurityAPIVersion,
		Kind:       security.CISKubeBenchReportKind,
		Name:       node,
	})
	if err != nil {
		err = fmt.Errorf("listing CIS Kubernetes Benchmarks: %w", err)
		return
	}
	var reportList security.CISKubeBenchReportList
	err = r.structure(unstructuredList, &reportList)
	if err != nil {
		err = fmt.Errorf("unmarshalling JSON to CISKubernetesBenchmarkList: %w", err)
		return
	}

	for _, r := range reportList.Items {
		if r.Labels[kube.LabelResourceKind] == "Node" &&
			r.Labels[kube.LabelResourceName] == node &&
			r.Labels[kube.LabelHistoryLatest] == "true" {
			report = &r
			return
		}
	}

	return
}

func (r *Repository) GetKubeHunterReport(ctx context.Context) (report security.KubeHunterReport, err error) {
	unstructuredList, err := r.client.List(ctx, store.Key{
		APIVersion: aquaSecurityAPIVersion,
		Kind:       security.KubeHunterReportKind,
	})
	if err != nil {
		return
	}
	var reportList security.KubeHunterReportList
	err = r.structure(unstructuredList, &reportList)
	if err != nil {
		return
	}

	for _, r := range reportList.Items {
		if r.Name == "cluster" {
			report = r
			return
		}
	}

	return
}

func (r *Repository) structure(ul *unstructured.UnstructuredList, v interface{}) (err error) {
	b, err := ul.MarshalJSON()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, v)
	return
}