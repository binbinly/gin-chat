package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// see: https://github.com/chenjiandongx/ginprom

var (
	labels = []string{"status", "endpoint", "method"}

	uptime        *prometheus.CounterVec
	reqCount      *prometheus.CounterVec
	reqDuration   *prometheus.HistogramVec
	reqSizeBytes  *prometheus.SummaryVec
	respSizeBytes *prometheus.SummaryVec
)

// Init registers the prometheus metrics
func Init(namespace string) {

	// 应用运行时长
	uptime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "uptime",
			Help:      "HTTP service uptime.",
		}, nil,
	)

	// QPS
	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, labels,
	)

	// 接口响应时间
	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
		}, labels,
	)

	// 请求大小
	reqSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "http_request_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)

	// 响应大小
	respSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "http_response_size_bytes",
			Help:      "HTTP response sizes in bytes.",
		}, labels,
	)
	prometheus.MustRegister(uptime, reqCount, reqDuration, reqSizeBytes, respSizeBytes)
	go recordUptime()
}

// recordUptime increases service uptime per second.
func recordUptime() {
	for range time.Tick(time.Second) {
		uptime.WithLabelValues().Inc()
	}
}

// calcRequestSize returns the size of request object.
func calcRequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size = len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL.
	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}

// PromOpts represents the Prometheus middleware Options.
// It is used for filtering labels by regex.
type PromOpts struct {
	Namespace            string
	ExcludeRegexStatus   string
	ExcludeRegexEndpoint string
	ExcludeRegexMethod   string
	EndpointLabelFn      func(c *gin.Context) string
}

// PromOpt 设置 PromOpts 选项
type PromOpt func(opts *PromOpts)

// WithNamespace 设置 Namespace
func WithNamespace(namespace string) PromOpt {
	return func(opts *PromOpts) {
		opts.Namespace = namespace
	}
}

// WithEndpointLabel 设置 EndpointLabel
func WithEndpointLabel(fn func(c *gin.Context) string) PromOpt {
	return func(opts *PromOpts) {
		opts.EndpointLabelFn = fn
	}
}

// checkLabel returns the match result of labels.
// Return true if regex-pattern compiles failed.
func (po *PromOpts) checkLabel(label, pattern string) bool {
	if pattern == "" {
		return true
	}

	matched, err := regexp.MatchString(pattern, label)
	if err != nil {
		return true
	}
	return !matched
}

// Prom returns a gin.HandlerFunc for exporting some Web metrics
func Prom(opts ...PromOpt) gin.HandlerFunc {
	promOpts := &PromOpts{
		Namespace: "app",
	}
	for _, opt := range opts {
		opt(promOpts)
	}

	Init(promOpts.Namespace)

	// make sure EndpointLabelMappingFn is callable
	if promOpts.EndpointLabelFn == nil {
		promOpts.EndpointLabelFn = func(c *gin.Context) string {
			return c.Request.URL.Path
		}
	}

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		endpoint := promOpts.EndpointLabelFn(c)
		method := c.Request.Method

		lvs := []string{status, endpoint, method}

		isOk := promOpts.checkLabel(status, promOpts.ExcludeRegexStatus) &&
			promOpts.checkLabel(endpoint, promOpts.ExcludeRegexEndpoint) &&
			promOpts.checkLabel(method, promOpts.ExcludeRegexMethod)

		if !isOk {
			return
		}
		// no response content will return -1
		respSize := c.Writer.Size()
		if respSize < 0 {
			respSize = 0
		}

		reqCount.WithLabelValues(lvs...).Inc()
		reqDuration.WithLabelValues(lvs...).Observe(time.Since(start).Seconds())
		reqSizeBytes.WithLabelValues(lvs...).Observe(calcRequestSize(c.Request))
		respSizeBytes.WithLabelValues(lvs...).Observe(float64(respSize))
	}
}

// PromHandler wrappers the standard http.Handler to gin.HandlerFunc
func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
