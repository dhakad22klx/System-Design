# Learning: Payment Gateways and Payment Processors

**Blog:** [Stripe — Payment Gateways 101](https://stripe.com/in/resources/more/payment-gateways-101)

## 1. Payment Gateway

A payment gateway securely receives payment information from a checkout, forwards the payment request to the processor or acquiring infrastructure, and returns the result to the merchant.

Common capabilities include:
- Secure transmission and tokenization of payment details.
- Fraud detection and validation.
- Communication of payment status to the website or app.

## 2. Payment Processor

A payment processor is a company that operates the backend infrastructure to process transactions for merchants.

It:
- Routes payment requests through card networks to issuing banks.
- Receives approval or decline responses and returns them to the merchant.
- Supports clearing and settlement.

**Examples:** Stripe and Adyen.

Here, “business” means a payment services company. The shop selling the product is called the **merchant**.

## 3. Participants in a Card Payment

| Participant | Role |
| --- | --- |
| Customer | Purchases a product or service. |
| Merchant | Sells the product or service and accepts payment. |
| Payment gateway | Securely receives and forwards payment information. |
| Payment processor | Handles backend transaction processing and communication. |
| Card network | Connects the acquiring and issuing sides; examples include Visa and Mastercard. |
| Issuing bank | Issues the customer’s card and authorizes the payment. |
| Acquirer | Provides the merchant-side arrangement for accepting card payments. |

## 4. Example: Buying a ₹1,000 Shirt

Suppose you pay an online shop using an HDFC Visa card.

1. You enter your card details at checkout.
2. The gateway securely forwards the payment request.
3. The processor routes the authorization request through Visa to HDFC.
4. HDFC checks available funds or credit and relevant security conditions.
5. HDFC approves or declines the transaction.
6. The response returns through Visa, the processor, and the gateway to the shop.
7. If approved, the merchant captures the payment.
8. Clearing and settlement follow, and the merchant receives its payout according to its provider’s arrangement.

This is a simplified card-payment example. Other payment methods have different flows.

## 5. Authorization vs. Capture vs. Settlement

| Stage | Meaning |
| --- | --- |
| Authorization | The issuer approves the payment and typically reserves funds or available credit. |
| Capture | The merchant requests completion of the authorized payment. |
| Settlement | Financial obligations are settled and funds move between participating institutions. |

**Payment approval does not mean the money has already reached the merchant’s bank account.**

## 6. Can One Company Perform Multiple Roles?

Yes. Gateway, processor, and acquirer describe different roles, but they do not always represent separate companies.

For example, **Adyen combines gateway, processing, and acquiring capabilities** in one platform.

A merchant can integrate with one provider that performs several roles behind the scenes.

## Key Takeaways

- The gateway securely receives and forwards the payment request.
- The processor handles backend processing and communication with networks and banks.
- For a typical card payment, the issuing bank makes the authorization decision.
- Processors can also reject requests through their own validation or risk controls.
- Authorization, capture, settlement, and merchant payout are distinct events.
- One payment provider can perform multiple roles.

## Additional References

- [Stripe — Issuing Banks](https://stripe.com/resources/more/issuing-banks)
- [Adyen — Payment Gateways](https://www.adyen.com/knowledge-hub/payment-gateway)
- [Adyen — Payment Terminology](https://docs.adyen.com/adyen-glossary)