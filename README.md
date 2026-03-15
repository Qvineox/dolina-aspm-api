# Dolina ASPM API

## Features

## Project structure

```
internal/
└── entities/
    └── model1.go
    └── model2.go
    └── model3.go
└── repo/
    └── model1_repository.go  // Impl imports models
    └── model2_repository.go  // Impl imports models
    └── model3_repository.go  // Impl imports models
api/
└── repo/
    └── vuln_repository.go  // Interface only
```

```mermaid
sequenceDiagram
    participant W as Worker
    participant C as Control Center
    Note over W, C: Worker initiates connection
    W ->>+ C: Register(WorkerRegistration)
    C ->>+ W: WorkerIdentity

    loop WorkerControlService
        W ->> C: Heartbeat(WorkerStatus)
        C ->> W: WorkerSignal
    end

    Note over W, C: Job orchestration
    loop WorkerQueueService
        Note over W: Received new job assignment signal
        W ->> C: GetJob
        C ->> W: WorkerJob
    end

    Note over W: Report processing
    Note over C: ReportsService
    W ->> C: GetReportByUUID(common.v1.UUID)
    C ->> W: reports.v1.Report
    W ->> C: GetScannerReports(common.v1.UUIDs)
    C ->> W: reports.v1.ScannerReports
    Note over C: RulesetsService
    W ->> C: GetScannerRulesets()
    C ->> W: rules.v1.Rulesets
    Note over C: AnalysisService
    W ->> C: CreateReportAnalysis(stream AnalysisProcessStream)
    C ->> W: analysis.v1.Analysis
    W ->> C: StreamAnalysisProcess(stream AnalysisProcessStream)
```