package metrics

import (
	"github.com/banzaicloud/pipeline/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

var logger *logrus.Logger
var log *logrus.Entry

func init() {
	logger = config.Logger()
	log = logger.WithFields(logrus.Fields{"action": "Helm"})
}

func StartMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9222", nil))
}
