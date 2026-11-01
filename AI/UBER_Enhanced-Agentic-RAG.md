# Learnings: Uber’s Enhanced Agentic RAG

[Original blog](https://www.uber.com/us/en/blog/enhanced-agentic-rag/)

## 1. Main Idea

RAG quality depends on document preparation, retrieval, and generation. Uber improves all three through enriched content and agents around retrieval.

## 2. Offline Document Processing

“Offline” means preparing documents outside the user’s request path. It does not mean disconnected from the internet; processing can run when documents change.

Uber’s preparation includes:
- Extracting structured content using a custom Google Docs loader.
- Preserving table structure during chunking.
- Adding summaries, FAQs, and keywords.
- Creating chunk embeddings and storing reusable enrichment artifacts.

**Why preserve tables?** A value like “30 days” is ambiguous if its row label and column heading are lost. Parsing and chunking must preserve that meaning.

## 3. Feature Store vs Vector Store

| Store | Role |
|---|---|
| Feature store | Supplies prepared titles, summaries, and FAQs to agents |
| Vector store | Supports similarity search over chunk embeddings |

Document summaries are shared across a document’s chunks; FAQs and keywords align with individual chunks.

**Are features embedded?** The blog confirms chunk embeddings, but does not specify every field included in the embedding input or describe independently embedding every feature.

An agent can read a document catalog directly as text in its prompt. Stored information does not need an embedding unless the system needs vector search over it.

## 4. Answer Generation

1. **Query Optimizer:** Refines ambiguous questions and decomposes complex ones.
2. **Source Identifier:** Selects likely relevant documents using the prepared catalog.
3. **Retrieval:** Combines vector-search and BM25 results.
4. **Post-Processor:** Removes duplicates and orders chunks by their original positions.
5. **Answer Generator:** Uses the questions and assembled evidence.

## 5. Evaluation

SMEs establish reference answers or feedback. New chatbot versions answer the test questions, and an LLM judge evaluates them using references, instructions, and retrieved evidence, producing scores and reasoning.

## 6. Engineering Lessons to Remember

- Fix extraction problems before tuning prompts: missing table context cannot reliably be recovered downstream.
- Evaluate retrieval and generation separately to locate failures.
- Narrowing the document set can improve relevance, but selecting the wrong sources can exclude the answer.
- Combining keyword and semantic retrieval provides complementary evidence.
- Reuse prepared summaries and FAQs to avoid regenerating them per question.
- Treat LLM-judge scores as estimates; check their agreement with expert judgment.
- Compare changes on a consistent test set and track latency and cost alongside quality.
- At million-document scale, design a scalable way to select catalog entries rather than putting the entire catalog into every prompt.