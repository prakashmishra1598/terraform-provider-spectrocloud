package client

import (
	"strings"

	"github.com/spectrocloud/terraform-provider-spectrocloud/pkg/client/herr"

	hapitransport "github.com/spectrocloud/hapi/apiutil/transport"
	"github.com/spectrocloud/hapi/models"

	clusterC "github.com/spectrocloud/hapi/spectrocluster/client/v1"
)

func (h *V1Client) DeleteCluster(uid string) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	cluster, err := h.GetCluster(uid)
	if err != nil || cluster == nil {
		return err
	}

	var params *clusterC.V1SpectroClustersDeleteParams
	switch cluster.Metadata.Annotations["scope"] {
	case "project":
		params = clusterC.NewV1SpectroClustersDeleteParamsWithContext(h.Ctx).WithUID(uid)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersDeleteParams().WithUID(uid)
		break
	default:
		break
	}

	_, err = client.V1SpectroClustersDelete(params)
	return err
}

func (h *V1Client) GetCluster(uid string) (*models.V1SpectroCluster, error) {
	cluster, err := h.GetClusterWithoutStatus(uid)
	if err != nil {
		return nil, err
	}

	if cluster == nil || cluster.Status.State == "Deleted" {
		return nil, nil
	}

	return cluster, nil
}

func (h *V1Client) GetClusterWithoutStatus(uid string) (*models.V1SpectroCluster, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	params := clusterC.NewV1SpectroClustersGetParamsWithContext(h.Ctx).WithUID(uid)
	success, err := client.V1SpectroClustersGet(params)
	// handle tenant context here cluster may be a tenant cluster
	if e, ok := err.(*hapitransport.TransportError); ok && e.HttpCode == 404 {
		params := clusterC.NewV1SpectroClustersGetParams().WithUID(uid)
		success, err = client.V1SpectroClustersGet(params)
		if e, ok := err.(*hapitransport.TransportError); ok && e.HttpCode == 404 {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		//return nil, nil
	}
	if e, ok := err.(*hapitransport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// special check if the cluster is marked deleted
	cluster := success.Payload
	return cluster, nil
}

func (h *V1Client) GetClusterByName(name string, ClusterContext string) (*models.V1SpectroCluster, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil, err
	}

	var params *clusterC.V1SpectroClustersListParams
	switch ClusterContext {
	case "project":
		params = clusterC.NewV1SpectroClustersListParamsWithContext(h.Ctx)
		break
	case "tenant":
		params = clusterC.NewV1SpectroClustersListParams()
		break
	default:
		break
	}
	success, err := client.V1SpectroClustersList(params)
	if e, ok := err.(*hapitransport.TransportError); ok && e.HttpCode == 404 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	for _, cluster := range success.Payload.Items {
		if cluster.Metadata.Name == name && cluster.Status.State != "Deleted" {
			return cluster, nil
		}
	}

	return nil, nil
}

func (h *V1Client) GetClusterKubeConfig(uid string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	builder := new(strings.Builder)
	params := clusterC.NewV1SpectroClustersUIDKubeConfigParamsWithContext(h.Ctx).WithUID(uid)
	_, err = client.V1SpectroClustersUIDKubeConfig(params, builder)
	if err != nil {
		if herr.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) GetClusterImportManifest(uid string) (string, error) {
	client, err := h.GetClusterClient()
	if err != nil {
		return "", err
	}

	builder := new(strings.Builder)
	params := clusterC.NewV1SpectroClustersUIDImportManifestParamsWithContext(h.Ctx).WithUID(uid)
	_, err = client.V1SpectroClustersUIDImportManifest(params, builder)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

func (h *V1Client) UpdateClusterProfileValues(uid string, profiles *models.V1SpectroClusterProfiles) error {
	client, err := h.GetClusterClient()
	if err != nil {
		return nil
	}

	resolveNotification := true
	params := clusterC.NewV1SpectroClustersUpdateProfilesParamsWithContext(h.Ctx).WithUID(uid).
		WithBody(profiles).WithResolveNotification(&resolveNotification)
	_, err = client.V1SpectroClustersUpdateProfiles(params)
	return err
}
