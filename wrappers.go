package mocks

import embedded "go.opentelemetry.io/otel/metric/embedded"

type Float64CounterWrapper struct {
	embedded.Float64Counter
	*MockFloat64Counter
}

type Float64HistogramWrapper struct {
	embedded.Float64Histogram
	*MockFloat64Histogram
}

type Float64ObservableCounterWrapper struct {
	embedded.Float64ObservableCounter
	*MockFloat64ObservableCounter
}

type Float64ObservableGaugeWrapper struct {
	embedded.Float64ObservableGauge
	*MockFloat64ObservableGauge
}

type Float64ObservableUpDownCounterWrapper struct {
	embedded.Float64ObservableUpDownCounter
	*MockFloat64ObservableUpDownCounter
}

type Float64ObserverWrapper struct {
	embedded.Float64Observer
	*MockFloat64Observer
}

type Float64UpDownCounterWrapper struct {
	embedded.Float64UpDownCounter
	*MockFloat64UpDownCounter
}

type Int64CounterWrapper struct {
	embedded.Int64Counter
	*MockInt64Counter
}

type Int64HistogramWrapper struct {
	embedded.Int64Histogram
	*MockInt64Histogram
}

type Int64ObservableCounterWrapper struct {
	embedded.Int64ObservableCounter
	*MockInt64ObservableCounter
}

type Int64ObservableGaugeWrapper struct {
	embedded.Int64ObservableGauge
	*MockInt64ObservableGauge
}

type Int64ObservableUpDownCounterWrapper struct {
	embedded.Int64ObservableUpDownCounter
	*MockInt64ObservableUpDownCounter
}

type Int64ObserverWrapper struct {
	embedded.Int64Observer
	*MockInt64Observer
}

type Int64UpDownCounterWrapper struct {
	embedded.Int64UpDownCounter
	*MockInt64UpDownCounter
}

type MeterWrapper struct {
	embedded.Meter
	*MockMeter
}

type MeterProviderWrapper struct {
	embedded.MeterProvider
	*MockMeterProvider
}

type ObserverWrapper struct {
	embedded.Observer
	*MockObserver
}

type RegistrationWrapper struct {
	embedded.Registration
	*MockRegistration
}
