# 🏠 1. Header & Description

# Aequitas

![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![Coverage](https://img.shields.io/badge/coverage-92%25-brightgreen)
![Go Version](https://img.shields.io/badge/go-1.20+-blue)
![License](https://img.shields.io/badge/license-MIT-blue)

**Aequitas** is a comprehensive, Go-based ledger server and auditing system designed for high-performance transactional workloads. It serves as a robust benchmarking framework that implements and evaluates various Concurrency Control (CC) strategies—including Optimistic Concurrency Control (OCC), Pessimistic, Serializable Snapshot Isolation (SSI), and Adaptive approaches. By integrating rigorous idempotency management (formally specified via TLA+) and comprehensive telemetry, Aequitas is built for researchers and developers evaluating database scalability and transactional integrity under skewed workloads.

![Aequitas Demonstration](https://via.placeholder.com/800x400.png?text=Aequitas+Ledger+and+Benchmarking+In+Action)

---

# 🛠 2. Prerequisites & Installation

### Requirements

* **Go Runtime:** Version 1.20 or higher (required for the `go.mod` and Go packages).
* **Python 3.8+:** Required for generating data visualizations and heatmaps (`plot_heatmap.py`, `plot_results.py`, etc.).
* **PowerShell:** Required to execute the benchmark suite and sweep scripts (e.g., `run_final_sweep.ps1`, `run_set1.ps1`).
* **Database Engine:** A SQL-compatible relational database capable of executing the provided `schema.sql`.

### Step-by-Step Setup

```bash
# 1. Clone the repository
git clone https://github.com/your-org/aequitas.git
cd aequitas

# 2. Download Go dependencies
go mod download

# 3. Initialize the database schema
psql -U postgres -d aequitas_db -f schema.sql
```

### Environment Variables

Create a `.env` file in the root directory. Below are the required variables to ensure the ledger and auditor connect successfully without exposing sensitive secrets in your codebase:

```env
# Server Configuration
AEQUITAS_LEDGER_PORT=8080
AEQUITAS_AUDITOR_PORT=8081

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=aequitas_user
DB_NAME=aequitas_db
# Note: Provide DB_PASSWORD securely at runtime

# Telemetry
ENABLE_GC_MONITOR=true
```

# 🚀 3. Usage & Examples

### Quick Start

To get the project running immediately, start the main ledger server and the auditor in separate terminal sessions:

```bash
# Start the main Ledger Server
go run cmd/ledger_server/main.go

# Start the Auditor process
go run cmd/auditor/main.go
```

### Code Snippets & Benchmarking

Aequitas comes with a built-in workload generator. You can generate a Zipfian distribution workload and run benchmark sweeps using the provided PowerShell scripts:

```bash
# 1. Run the workload generator
go run cmd/workload_gen/main.go --concurrency=100 --skew=0.8

# 2. Execute a full repetition sweep for benchmarking
pwsh run_repetition_sweep.ps1

# 3. Generate plots from the consolidated results
python plot_timeseries.py
python plot_heatmap.py
```

# 🗂 4. Project Blueprint

### Tech Stack

- **Language:** Go (Core ledger, auditor, workload generation, concurrency control strategies).
- **Data Science/Plotting:** Python (`matplotlib`/`pandas` for performance visualization).
- **Automation:** PowerShell (Benchmarking execution sweeps).
- **Formal Verification:** TLA+ (Used in `Idempotency.tla` for formal proofs).
- **Database:** SQL (`schema.sql`).

### Directory Structure

```text
aequitas/
├── cmd/                        # Application entry points
│   ├── auditor/               # Auditor service
│   ├── ledger_server/         # Main ledger server
│   └── workload_gen/          # Load testing/workload generation
├── pkg/                        # Core Go packages
│   ├── auditor/               # Auditing logic
│   ├── cc/                    # Concurrency Control (adaptive, occ, pessimistic, ssi)
│   ├── idempotency/           # Idempotency manager and TTL cleaner
│   ├── ledger/                # Canonical ledger data models
│   ├── telemetry/             # DB stats, GC monitoring, and metrics
│   └── workload/              # Zipf distribution and DB cleaners
├── plots/                      # Generated visual outputs (latency, throughput, heatmaps)
├── *.py                        # Plotting and extraction scripts (plot_heatmap.py, etc.)
├── *.ps1                       # Execution and benchmarking sweeps (run_set1.ps1, etc.)
├── schema.sql                  # Database schema definitions
├── Idempotency.tla             # Formal TLA+ specification for idempotency
└── go.mod & go.sum             # Go module dependencies
```

# 4. Install Python dependencies for plotting
pip install matplotlib pandas seaborn
