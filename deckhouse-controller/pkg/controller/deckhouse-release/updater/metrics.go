/*
Copyright 2024 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package d8updater

import (
	metricstorage "github.com/flant/shell-operator/pkg/metric_storage"

	"github.com/deckhouse/deckhouse/go_lib/updater"
)

const d8ReleaseBlockedMetricName = "d8_release_info"

func NewMetricsUpdater(metricStorage *metricstorage.MetricStorage) *MetricsUpdater {
	return &MetricsUpdater{
		metricStorage: metricStorage,
	}
}

type MetricsUpdater struct {
	metricStorage *metricstorage.MetricStorage
}

func (mu *MetricsUpdater) UpdateReleaseMetric(name string, metricLabels updater.MetricLabels) {
	mu.PurgeReleaseMetric(name)
	mu.metricStorage.Grouped().GaugeSet(name, d8ReleaseBlockedMetricName, 1, metricLabels)
}

func (mu *MetricsUpdater) PurgeReleaseMetric(name string) {
	mu.metricStorage.Grouped().ExpireGroupMetricByName(name, d8ReleaseBlockedMetricName)
}
