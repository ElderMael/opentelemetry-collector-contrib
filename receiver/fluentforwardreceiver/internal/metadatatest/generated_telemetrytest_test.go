// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver/internal/metadata"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := SetupTelemetry()
	tb, err := metadata.NewTelemetryBuilder(
		testTel.NewTelemetrySettings(),
	)
	require.NoError(t, err)
	require.NotNil(t, tb)
	tb.FluentClosedConnections.Add(context.Background(), 1)
	tb.FluentEventsParsed.Add(context.Background(), 1)
	tb.FluentOpenedConnections.Add(context.Background(), 1)
	tb.FluentParseFailures.Add(context.Background(), 1)
	tb.FluentRecordsGenerated.Add(context.Background(), 1)

	testTel.AssertMetrics(t, []metricdata.Metrics{
		{
			Name:        "otelcol_fluent_closed_connections",
			Description: "Number of connections closed to the fluentforward receiver",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: false,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_fluent_events_parsed",
			Description: "Number of Fluent events parsed successfully",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: false,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_fluent_opened_connections",
			Description: "Number of connections opened to the fluentforward receiver",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: false,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_fluent_parse_failures",
			Description: "Number of times Fluent messages failed to be decoded",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: false,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_fluent_records_generated",
			Description: "Number of log records generated from Fluent forward input",
			Unit:        "1",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: false,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
	}, metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreValue())
	require.NoError(t, testTel.Shutdown(context.Background()))
}