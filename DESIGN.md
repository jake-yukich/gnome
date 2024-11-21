## Overview
### **Gene-Editing Variant Risk Explorer**
**Overview:** Build a web application to assess the safety of genetic edits by analyzing off-target effects and variant risks, leveraging gnomAD data and an LLM for interpreting results and providing context.

---

#### **Project Components**

1. **Go Backend**
   - **Core Functionality:**
     - Load and query gnomAD data (e.g., variant frequencies, population metrics) stored in Postgres.
     - Implement an endpoint to assess potential off-target edits by comparing input sequences to nearby variants in the gnomAD dataset.
     - Provide APIs to fetch summarized safety insights for specific variants.
   - **LLM Integration:**
     - Add an endpoint to accept user questions (e.g., "What are the implications of editing near BRCA1 variants in population X?").
     - Use an LLM (e.g., OpenAI or a fine-tuned model) to generate insights based on gnomAD data and predefined safety guidelines.

2. **React + TypeScript Frontend**
   - **Features:**
     - A search bar for entering genes or variants of interest.
     - A dashboard displaying safety metrics for specific edits, including:
       - Variant population frequency.
       - Predicted risk levels based on gnomAD data.
     - A chatbot-like interface for querying the LLM (e.g., "Explain why this variant is flagged as high risk").
   - **Visualization:** Add charts or visual summaries for off-target risk scores, variant distribution, and population statistics.

3. **Python ETL Pipeline**
   - **Tasks:**
     - Extract relevant gnomAD data (e.g., VCF files) and process it into a format suitable for querying in Postgres.
     - Precompute off-target risk metrics or variant impact scores for faster retrieval.

4. **LLM-Driven RAG**
   - **Use Case:** When a user queries the system, retrieve relevant documentation or scientific papers explaining the variant or gene from a knowledge base (e.g., PubMed or gnomAD documentation).
   - **Implementation:**
     - Generate embeddings for gnomAD documentation using Python during ETL.
     - Use `pgvector` in Postgres or a separate vector store to enable fast similarity searches.
     - Combine retrieved documents with the LLM’s response for a comprehensive answer.


## Implementation Plan

### **1. Backend Core (Go)**
   - Set up the Go project structure.
   - Database Setup:
   - Create a basic API for querying gnomAD data from Postgres:
     - Endpoint: `/variants` to fetch variants by gene or region.
     - Endpoint: `/offtarget-risk` to calculate and return off-target risk metrics.
   - Test API responses with mock data.

---

### **2. ETL Pipeline (Python)**
   - Process gnomAD data (e.g., extract VCF files, clean, and normalize).
   - Load processed data into Postgres with appropriate indexing for fast lookups.
   - Optionally precompute risk metrics if computationally expensive.

---

### **3. Frontend Basics (React + TypeScript)**
   - Create a basic React app with:
     - A search bar for gene or variant queries.
     - A table or card layout to display query results.
   - Connect the frontend to the backend APIs (`/variants`).

---

### **4. LLM Integration**
   - Add a `/query-llm` endpoint to the backend:
     - Use the LLM API for natural language interpretation.
     - Combine user input with backend data (e.g., off-target risk scores).
   - Test with simple questions and responses (e.g., "What does this variant do?").

---

### **5. Advanced Frontend Features**
   - Implement a chatbot-like interface for the LLM queries.
   - Add visualizations (e.g., risk scores, population frequency plots).
   - Optimize the UX for searching and exploring data.

---

### **6. RAG (Retrieval-Augmented Generation)**
   - Generate embeddings for gnomAD documentation or related literature during ETL.
   - Store embeddings in Postgres (with `pgvector`) or a vector store.
   - Update the `/query-llm` endpoint to retrieve and augment LLM responses with relevant documents.

---

### **7. Final Touches**
   - Add sample gene-editing scenarios to demonstrate functionality.
   - Write clear documentation for the project and APIs.

---

### Sample monorepo structure:

```plaintext
gene-editor/
├── .github/                    # GitHub Actions and PR templates
│   └── workflows/
├── backend/                    # Go backend service
│   ├── cmd/                   
│   │   └── server/            # Main application entrypoint
│   ├── internal/              # Private application code
│   │   ├── api/              # API handlers
│   │   ├── domain/           # Core business logic
│   │   ├── database/         # Database interactions
│   │   └── llm/              # LLM integration logic
│   ├── pkg/                   # Public packages
│   └── tests/                 # Integration tests
│
├── frontend/                   # React + TypeScript frontend
│   ├── src/
│   │   ├── components/       # Reusable UI components
│   │   ├── pages/           # Page components
│   │   ├── services/        # API client and services
│   │   ├── hooks/           # Custom React hooks
│   │   ├── types/           # TypeScript definitions
│   │   └── utils/           # Helper functions
│   ├── public/              # Static assets
│   └── tests/               # Frontend tests
│
├── etl/                       # Python ETL pipeline
│   ├── src/
│   │   ├── extractors/      # Data extraction logic
│   │   ├── transformers/    # Data transformation logic
│   │   ├── loaders/         # Database loading logic
│   │   └── embeddings/      # Document embedding generation
│   ├── scripts/             # ETL running scripts
│   └── tests/               # ETL tests
│
├── docs/                      # Project documentation
│   ├── api/                 # API documentation
│   ├── architecture/        # Architecture diagrams
│   └── setup/               # Setup guides
│
├── deploy/                    # Deployment configurations
│   ├── docker/              # Dockerfile(s)
│   ├── k8s/                 # Kubernetes manifests
│   └── terraform/           # Infrastructure as code
│
├── scripts/                   # Development scripts
│   ├── setup.sh
│   └── seed-data.sh
│
├── .gitignore
├── docker-compose.yml         # Local development setup
├── README.md
└── Makefile                  # Common commands
```
