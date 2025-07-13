# Complexity Analysis

> Last updated: Sun Jul 13 21:43:35 UTC 2025

## High Complexity Functions (>10)

```
47 v04 NewCDEvent sdk-go/pkg/api/v04/types.go:94:1
40 v03 NewCDEvent sdk-go/pkg/api/v03/types.go:84:1
33 main DataFromSchema sdk-go/tools/generator.go:542:1
22 events (*EventFactory).CreateTestEvent pkg/events/factory.go:272:1
21 v04_test TestExamples sdk-go/pkg/api/v04/conformance_test.go:564:1
19 main main sdk-go/tools/generator.go:218:1
16 api_test TestNewFromJsonString sdk-go/pkg/api/bindings_test.go:508:1
16 api_test TestAsCloudEvent sdk-go/pkg/api/bindings_test.go:241:1
15 v04_test TestNewCDEvent sdk-go/pkg/api/v04/factory_test.go:70:1
15 v03_test TestNewCDEvent sdk-go/pkg/api/v03/factory_test.go:70:1
12 events (*EventFactory).CreateBuildEvent pkg/events/factory.go:157:1
12 events (*EventFactory).CreateTaskRunEvent pkg/events/factory.go:93:1
12 events (*EventFactory).CreatePipelineRunEvent pkg/events/factory.go:33:1
11 events (*EventFactory).CreateServiceEvent pkg/events/factory.go:217:1
```

**Total functions with high complexity:** 14

## All Functions Complexity

<details>
<summary>View all functions complexity</summary>

```
47 v04 NewCDEvent sdk-go/pkg/api/v04/types.go:94:1
40 v03 NewCDEvent sdk-go/pkg/api/v03/types.go:84:1
33 main DataFromSchema sdk-go/tools/generator.go:542:1
22 events (*EventFactory).CreateTestEvent pkg/events/factory.go:272:1
21 v04_test TestExamples sdk-go/pkg/api/v04/conformance_test.go:564:1
19 main main sdk-go/tools/generator.go:218:1
16 api_test TestNewFromJsonString sdk-go/pkg/api/bindings_test.go:508:1
16 api_test TestAsCloudEvent sdk-go/pkg/api/bindings_test.go:241:1
15 v04_test TestNewCDEvent sdk-go/pkg/api/v04/factory_test.go:70:1
15 v03_test TestNewCDEvent sdk-go/pkg/api/v03/factory_test.go:70:1
12 events (*EventFactory).CreateBuildEvent pkg/events/factory.go:157:1
12 events (*EventFactory).CreateTaskRunEvent pkg/events/factory.go:93:1
12 events (*EventFactory).CreatePipelineRunEvent pkg/events/factory.go:33:1
11 events (*EventFactory).CreateServiceEvent pkg/events/factory.go:217:1
10 main typesForSchema sdk-go/tools/generator.go:700:1
10 main getWalkProcessor sdk-go/tools/generator.go:459:1
10 v03_test TestExamples sdk-go/pkg/api/v03/examples_test.go:396:1
10 api (*EmbeddedLinksArray).UnmarshalJSON sdk-go/pkg/api/types.go:277:1
10 api_test TestAsJsonBytes sdk-go/pkg/api/bindings_test.go:355:1
10 api Validate sdk-go/pkg/api/bindings.go:144:1
10 output formatCloudEventWithCustomData pkg/output/formatters.go:121:1
9 main validateStringEnumAnyOf sdk-go/tools/generator.go:513:1
9 main getSchemasWalkProcessor sdk-go/tools/generator.go:416:1
9 main generate sdk-go/tools/generator.go:338:1
9 api CDEventTypeFromString sdk-go/pkg/api/types.go:466:1
9 api_test TestNewFromJsonBytes sdk-go/pkg/api/bindings_test.go:670:1
9 events_test TestParseCustomDataFromJSON pkg/events/factory_test.go:564:1
8 api GetCustomData sdk-go/pkg/api/types.go:694:1
8 output formatYAMLWithCustomData pkg/output/formatters.go:80:1
8 events_test TestCreateTestEvent pkg/events/factory_test.go:507:1
8 events_test TestCreateServiceEvent pkg/events/factory_test.go:425:1
8 events_test TestCreateBuildEvent pkg/events/factory_test.go:373:1
8 events_test TestCreateTaskRunEvent pkg/events/factory_test.go:321:1
8 events_test TestCreatePipelineRunEvent pkg/events/factory_test.go:269:1
7 main TestValidateStringEnumAnyOf sdk-go/tools/generator_test.go:223:1
7 main TestExecuteTemplate_Success sdk-go/tools/generator_test.go:134:1
7 api_test TestCDEventTypeFromString sdk-go/pkg/api/types_test.go:368:1
7 transport (*TransportFactory).CreateTransport pkg/transport/transport.go:138:1
7 output formatJSONWithCustomData pkg/output/formatters.go:43:1
7 events_test TestCreateEventsWithoutOptionalFields pkg/events/factory_test.go:208:1
6 api_test TestParseType sdk-go/pkg/api/bindings_test.go:584:1
6 api NewFromJsonBytesContext sdk-go/pkg/api/bindings.go:188:1
6 transport_test TestTransportFactory_CreateTransport pkg/transport/transport_test.go:50:1
6 output_test TestFormatMultipleEvents pkg/output/formatters_test.go:352:1
6 output_test TestFormatOutputWithoutCustomData pkg/output/formatters_test.go:282:1
6 output_test TestFormatCloudEventWithCustomData pkg/output/formatters_test.go:244:1
6 output_test TestFormatYAMLWithCustomData pkg/output/formatters_test.go:207:1
6 output_test TestFormatJSONWithCustomData pkg/output/formatters_test.go:11:1
6 events_test TestValidateNewPipelineRunStartedEvent pkg/events/factory_test.go:11:1
6 cmd_test TestSendCommands cmd/cmd_test.go:151:1
```

</details>

## Recommendations

- Functions with complexity >10 should be reviewed for refactoring
- Consider breaking down complex functions into smaller, more focused functions
- Use early returns and guard clauses to reduce nesting
- Extract common logic into helper functions
