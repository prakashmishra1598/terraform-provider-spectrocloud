package spectrocloud

import (
	"github.com/spectrocloud/hapi/models"
)

func getUpdateStrategy(m map[string]interface{}) string {
	updateStrategy := "RollingUpdateScaleOut"
	if m["update_strategy"] != nil {
		updateStrategy = m["update_strategy"].(string)
	}
	return updateStrategy
}

func flattenUpdateStrategy(updateStrategy *models.V1UpdateStrategy, oi map[string]interface{}) {
	if updateStrategy.Type != "" {
		oi["update_strategy"] = updateStrategy.Type
	} else {
		oi["update_strategy"] = "RollingUpdateScaleOut"
	}
}
