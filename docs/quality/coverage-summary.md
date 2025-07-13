# Coverage Summary

> Last updated: Sun Jul 13 21:43:35 UTC 2025

## Overall Coverage: 82.3%

### Coverage by Package

```
github.com/brunseba/cdevents-tools/cmd/generate.go:41:			init				100.0%
github.com/brunseba/cdevents-tools/cmd/generate.go:46:			addCommonGenerateFlags		100.0%
github.com/brunseba/cdevents-tools/cmd/generate.go:62:			parseCustomData			83.3%
github.com/brunseba/cdevents-tools/cmd/generate.go:75:			getDefaultSource		66.7%
github.com/brunseba/cdevents-tools/cmd/generate.go:88:			outputEvent			100.0%
github.com/brunseba/cdevents-tools/cmd/generate.go:93:			outputEventWithCustomData	80.0%
github.com/brunseba/cdevents-tools/cmd/generate_build.go:42:		init				100.0%
github.com/brunseba/cdevents-tools/cmd/generate_pipeline.go:42:		init				100.0%
github.com/brunseba/cdevents-tools/cmd/generate_service.go:41:		init				100.0%
github.com/brunseba/cdevents-tools/cmd/generate_task.go:43:		init				100.0%
github.com/brunseba/cdevents-tools/cmd/root.go:32:			Execute				100.0%
github.com/brunseba/cdevents-tools/cmd/root.go:36:			init				100.0%
github.com/brunseba/cdevents-tools/cmd/root.go:50:			initConfig			69.2%
github.com/brunseba/cdevents-tools/cmd/send.go:32:			init				100.0%
github.com/brunseba/cdevents-tools/cmd/send.go:49:			sendEvent			83.3%
github.com/brunseba/cdevents-tools/cmd/send.go:72:			SendEventWithRetry		100.0%
github.com/brunseba/cdevents-tools/cmd/send_pipeline.go:45:		init				100.0%
github.com/brunseba/cdevents-tools/pkg/events/factory.go:26:		NewEventFactory			100.0%
github.com/brunseba/cdevents-tools/pkg/events/factory.go:33:		CreatePipelineRunEvent		96.2%
github.com/brunseba/cdevents-tools/pkg/events/factory.go:93:		CreateTaskRunEvent		77.8%
github.com/brunseba/cdevents-tools/pkg/events/factory.go:157:		CreateBuildEvent		69.2%
github.com/brunseba/cdevents-tools/pkg/events/factory.go:217:		CreateServiceEvent		75.0%
github.com/brunseba/cdevents-tools/pkg/events/factory.go:272:		CreateTestEvent			69.4%
github.com/brunseba/cdevents-tools/pkg/events/factory.go:357:		applyCustomData			0.0%
github.com/brunseba/cdevents-tools/pkg/events/factory.go:363:		ParseCustomDataFromJSON		100.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:19:		FormatOutput			100.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:24:		FormatOutputWithCustomData	100.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:38:		formatJSON			0.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:43:		formatJSONWithCustomData	80.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:75:		formatYAML			0.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:80:		formatYAMLWithCustomData	78.9%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:116:	formatCloudEvent		0.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:121:	formatCloudEventWithCustomData	78.3%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:164:	FormatMultipleEvents		100.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:178:	formatMultipleJSON		75.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:187:	formatMultipleYAML		75.0%
github.com/brunseba/cdevents-tools/pkg/output/formatters.go:196:	formatMultipleCloudEvents	80.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:24:	NewHTTPTransport		85.7%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:46:	WithHTTPHeaders			100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:53:	Send				77.8%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:76:	NewConsoleTransport		100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:83:	Send				100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:97:	NewFileTransport		100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:105:	Send				100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:119:	NewKafkaTransport		100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:125:	Send				0.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:133:	NewTransportFactory		100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:138:	CreateTransport			100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:165:	NewMultiTransport		100.0%
github.com/brunseba/cdevents-tools/pkg/transport/transport.go:172:	Send				100.0%
total:									(statements)			82.3%
```

### Coverage Threshold
- **Target**: 70%
- **Current**: 82.3%
- **Status**: ✅ PASS

### View Interactive Report
[Open Coverage Report](coverage.html)
