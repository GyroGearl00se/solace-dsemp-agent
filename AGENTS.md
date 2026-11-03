# AGENTS.md - Developer & AI Agent Guide for `solace-dsemp-agent`

This document provides a concise architectural overview, operational mental model, and code navigation guide for AI agents and human contributors working on this codebase.

---

## 🎯 1. Project Purpose & High-Level Concept

**`solace-dsemp-agent`** is a declarative "Configuration-as-Code" (GitOps) reconciliation agent for Solace PubSub+ event brokers. It runs alongside a Solace Edge Broker or in Kubernetes/Docker environments.

### Core Lifecycle:
1. **Target State Ingestion**:
   - **File-based**: Reads a local YAML/JSON file (`SOL_INITIAL_STATE_FILE`, e.g., mounted K8s ConfigMap) and polls it for changes (`pkg/watcher`).
   - **Messaging-based**: Connects to the broker and subscribes to a Solace topic (`SOL_STATE_TOPIC`) listening for JSON/YAML target states (`pkg/messaging`).
2. **Version Gating**:
   - Reads the applied version stored in the metadata queue (`solace-dsemp-agent-metadata`, `Owner` field).
   - Skips reconciliation if the incoming target state version is not strictly newer (unless version check is disabled).
3. **Reconciliation Loop**:
   - Audits the broker state via Solace SEMPv2 REST API (`/SEMP/v2/config`).
   - Computes structural diffs between actual and desired states.
   - **Deletes** unmanaged resources (in reverse dependency order).
   - **Upserts** (Creates/Updates) desired resources (in forward dependency order).
4. **Status Reporting**:
   - Publishes run results and any error collections to an HTTP webhook (`SOL_STATUS_WEBHOOK_URL`) and/or a Solace topic (`SOL_STATUS_MESSAGE_TOPIC`).

---

## 🏗️ 2. Repository Architecture & Layout

```
.
├── main.go                     <- Application entrypoint: CLI configuration, signals, worker orchestration
├── main_test.go                <- Version comparison and smoke tests
├── cmd/
│   └── aesencryptor/           <- Standalone CLI utility for encrypting passwords with AES-CBC
│       └── main.go
├── pkg/
│   ├── controllers/            <- Resource adapters implementing CRUD and SEMP cursor pagination
│   │   ├── types.go            <- ResourceController interface & GenericCRUDHandler
│   │   ├── diff.go             <- Structural pointer-aware deep diffing (NeedsUpdate)
│   │   ├── queue_controller.go
│   │   ├── acl_profile_controller.go
│   │   └── ... (24 controllers)
│   ├── reconciler/             <- Thread-safe reconciliation orchestration & version store
│   │   ├── engine.go           <- Delete & Upsert sequencing pipeline (mutex-protected)
│   │   └── version.go          <- Version inspection & persistence in metadata queue
│   ├── state/                  <- TargetState models & templating resolution
│   │   ├── state.go            <- TargetState struct, YAML/JSON parser
│   │   ├── error.go            <- State error definition
│   │   └── templating          <- Recursive $env{VAR} and $aes{CIPHER} resolution
│   ├── crypto/                 <- AES-128/192/256-CBC cipher with PKCS#7 padding
│   │   └── aes.go
│   ├── messaging/              <- Solace Go Messaging API client for inbound TargetState
│   │   └── consumer.go
│   ├── watcher/                <- File system polling watcher for TargetState files
│   │   └── watcher.go
│   ├── status/                 <- Outbound webhook & status message publishers
│   │   ├── reporter.go
│   │   └── send_status_message.go
│   ├── semputil/               <- SEMP client initialization, TLS cert pool, 404 detector
│   │   └── semp.go
│   └── semplegacy/             <- Legacy SEMP broker version retriever
│       └── semplegacy.go
├── semp_swagger/               <- Generated Go Swagger SDK for Solace SEMPv2
│   └── config/
├── spec/                       <- OpenAPI / Swagger 2.0/3.0 SEMPv2 specifications
├── config.yaml                 <- Sample configuration file
├── targetstate.yaml            <- Sample target state definition
├── Dockerfile                  <- Multi-stage container build
└── Makefile                    <- Build & Swagger generator scripts
```

---

## 🔑 3. Critical Implementation Rules & Gotchas

When modifying or extending this agent, follow these mandatory conventions:

### A. Pointer Comparison in SEMP Models
- Generated Swagger models (`semp_swagger/config`) use pointers for optional fields (`*string`, `*bool`, `*int64`) and multi-level pointers (`**bool`).
- **Never** use raw `reflect.DeepEqual(a, b)` on Swagger structs without pointer dereferencing, as two distinct pointers holding identical values will fail equality checks.
- Always use or delegate to `controllers.NeedsUpdate(brokerObj, stateObj)` in [`pkg/controllers/diff.go`](pkg/controllers/diff.go).
- Missing/zero-value fields in the desired state (`stateField.IsZero()`) are intentionally skipped to preserve broker default configurations.

### B. SEMP Pagination (Cursor Following)
- SEMP v2 endpoints return pages (default page count: 100).
- When writing a `Get` method in any controller, you **must follow the cursor loop**:
  ```go
  var cursor string
  for {
      opts := &swagger.XxxOpts{Count: optional.NewInt32(100)}
      if cursor != "" {
          opts.Cursor = optional.NewString(cursor)
      }
      resp, _, err := client.XxxApi.GetXxx(ctx, msgVpn, opts)
      if err != nil { return nil, err }
      // append items...
      if resp.Meta == nil || resp.Meta.Paging == nil || resp.Meta.Paging.CursorQuery == "" {
          break
      }
      cursor = resp.Meta.Paging.CursorQuery
  }
  ```

### C. Concurrency Safety
- Both the file watcher and the messaging consumer can invoke state reconciliation asynchronously.
- The `reconciler.Engine` uses a central `sync.Mutex` (`e.mutex.Lock()`) around `Reconcile()`. Ensure this lock is never bypassed when mutating broker state.

### D. System & Whitelisted Resources
- Resources prefixed with `#` (e.g. `#acl-profile`, `#client-username`) or equal to `default` are system resources.
- Controllers implement `ShouldManage(obj)` to filter out system resources and whitelisted regex/path patterns (`WhitelistPatterns`) before calculating deletions.

### E. Child Resource Cleanup
- Sub-resources (e.g. `QueueSubscriptions`, `PublishExceptions`, `SubscribeExceptions`, `RDPQueueBindings`) must be reconciled even when the parent struct contains an empty array (`len(...) == 0`). This ensures unmanaged child resources on the broker are correctly pruned.

---

## 🛠️ 4. Common Developer Workflows

### Run all tests
```bash
go test -v ./...
```

### Build application binary
```bash
go build -o solace-dsemp-agent main.go
```

### Build encryption CLI tool
```bash
go build -o aesencryptor cmd/aesencryptor/main.go
```

### Run local dry-run test
```bash
SOL_DRYRUN=true SOL_INITIAL_STATE_ENABLED=true SOL_INITIAL_STATE_FILE=targetstate.yaml ./solace-dsemp-agent
```

### Regenerate SEMP Swagger client
```bash
make swagger_semp
```
*(Requires Docker and a running Solace broker at `localhost:8080`)*

