# Learning from Uber’s Genie: Production RAG

[Original blog](https://www.uber.com/us/en/blog/genie-ubers-gen-ai-on-call-copilot/)

## 1. What Genie Does

Genie answers engineering support questions using internal documentation, reducing repetitive work for on-call engineers.

It uses RAG: retrieve relevant evidence, then give that evidence to an LLM to generate an answer.

## 2. Two Pipelines

**Ingestion:** Fetch documents → chunk content → generate embeddings → build and distribute search indexes.

**Answering:** Embed the question → search for relevant chunks → provide their text to the LLM → return an answer with source URLs.

The LLM receives retrieved text, not the entire document collection.

## 3. Pusher Architecture

The Pusher prepares and distributes searchable index artifacts:

1. A bootstrap/ingest job triggers processing.
2. Spark jobs build and merge indexes.
3. Index artifacts are stored in Terrablob.
4. Leaf servers download and sync the base index and snapshots.
5. Queries are sent to the leaves.

| Component | Purpose |
|---|---|
| Terrablob | Durable storage for embeddings and index artifacts |
| Sia | Uber’s vector database system |
| Leaf | A serving node that searches its downloaded index |

**Does Sia store embeddings?** Its leaves hold searchable indexes representing the vectors. Blob storage and serving nodes can hold copies for different purposes.

**How are requests routed?** The blog says queries reach each leaf, but doesn’t explain routing, result merging, or the partition/replica layout.

Generally, partitioned search fans out to relevant partitions, searches them in parallel, and merges their best results. Replicas provide additional serving copies. This is a general pattern, not a confirmed Genie implementation detail.

## 4. Feedback and Evaluation

Users select Resolved, Helpful, Not Helpful, or Not Relevant.

Feedback flows through Slack → Kafka → Hive → dashboards. A separate evaluation pipeline combines context, responses, feedback, and metadata; an LLM judge produces quality metrics.

The blog describes evaluations supporting tuning, but does not establish automatic retraining or automatic fixes.

A practical improvement loop is:

Collect feedback → evaluate → identify recurring failures → change the system → rerun evaluations.

## 5. Lessons to Remember

- Missing evidence, poor retrieval, and incorrect generation need different fixes.
- Negative feedback identifies a failure; it doesn’t identify its cause.
- Automate evaluation and review representative or ambiguous failures.
- Preserve source URLs and require evidence-backed answers.
- Clear, current documentation is essential.
- Measure retrieval quality and answer quality separately.
- Track latency and cost alongside accuracy.
- Treat an LLM judge’s scores as estimates requiring validation.