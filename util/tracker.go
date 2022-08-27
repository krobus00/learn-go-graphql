package util

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
)

func StartTracer(ctx context.Context, tag, trackerName string) *newrelic.Segment {
	trx := newrelic.FromContext(ctx)
	segment := trx.StartSegment(fmt.Sprintf("%s %s", tag, trackerName))
	return segment
}

func StartRepositoryTracer(ctx context.Context) *newrelic.DatastoreSegment {
	txn := newrelic.FromContext(ctx)
	segment := newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product:   newrelic.DatastoreMySQL,
	}
	return &segment
}

func ErrorLogger(logger *zap.Logger, tag, tracing string, err error) {
	logger.Error(fmt.Sprintf("%s %s with error: %v", tag, tracing, err))
}
