package logt

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"
)

// TODO: CONTINUE HERE (write both log and trace)
func Info(span trace.Span, args ...any) {
	// write reduce function for array here
	template := ""
	for range args {
		template = template + "%v "
	}
	logStr := fmt.Sprintf(template, args...)
	logrus.Info(logStr)
	span.AddEvent(logStr)
}
