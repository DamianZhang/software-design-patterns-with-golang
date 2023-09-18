package logger

type CompositeExporter struct {
	exporters []Exporter
}

func NewCompositeExporter(exporters ...Exporter) *CompositeExporter {
	return &CompositeExporter{
		exporters: exporters,
	}
}

func (c *CompositeExporter) Export(message Message) {
	for _, exporter := range c.exporters {
		exporter.Export(message)
	}
}
