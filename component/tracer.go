package component

import (
	"go.opentelemetry.io/otel"
	"lesson-5-goland/common"
)

var Tracer = otel.Tracer(common.TRACE_SERVICE)
