# Coinbase Notification Platform — Key Learnings

Source: [Building a notification platform at Coinbase](https://www.coinbase.com/en-in/blog/building-a-notification-platform-at-coinbase), December 14, 2022.

1. **API-first components:** Single-purpose APIs support direct integrations, debugging CLIs, and internal UIs independently.
2. **Product independence:** Abstract recipient identities, delivery endpoints, and preference models across products.
3. **Priority isolation:** Separate queues let urgent alerts receive capacity independently of bulk announcements.
4. **Centralized delivery:** Workers resolve recipients, enforce preferences and rate limits, render templates, and contact providers.
5. **Feedback matters:** Collect provider delivery outcomes, errors, and user interactions for downstream consumers.
6. **Push and pull differ:** Store in-app content separately; support both campaign write bursts and frequent application reads.
7. **Continuous targeting:** Reevaluate eligibility continuously, keeping campaign logic outside product services.
8. **Workflow orchestration:** Per-user DAGs run on Temporal; branches depend on attributes, experiments, or previous actions.
9. **Experimentation:** Integrate content variants and workflow branching with the experimentation platform.
10. **Controlled self-service:** Rich editing, previews, reviewer approval, and role-based authorization enable non-engineers to publish.
11. **Iterative delivery:** Begin with configuration files; add datastore-backed templates and UI after feedback.
12. **Reliability and economics:** Agree on SLOs with consumers. Direct channel integrations cut reported costs over 50% while handling 10× workload.
