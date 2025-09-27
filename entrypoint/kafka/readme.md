# Naming Convention for RabbitMQ
---
## SNAPSHOT /  CHANGELOG TOPIC
- naming convention for snapshot / changelog :`snapshot-<namespace>-<event_type>`
- naming convention for compensation snapshot / changelog :`compensate-snapshot-<namespace>-<event_type>`
- naming convention for confirm snapshot / changelog :`confirm-snapshot-<namespace>-<event_type>`

### Example : 
- `snapshot-users-users_login`
- `snapshot-users-users_created`
- `snapshot-users-users_updated`
---

## SINK CONNECTOR
- naming convention for sink/connector :`<connector_type>-<provider>-<namespace>-<table or collection>`
- naming convention for sink/connector error handler :`dlq-<connector_type>-<provider>-<namespace>-<table or collection>`

### Example :
- `sink-pg-users-users`
- `source-mongo-notification-notification_templates`
---


## EXAMPLE FLOW
```mermaid
sequenceDiagram
    participant API as API Gateway
    participant US as User Service
    participant PS as Product Service
    participant PayS as Payment Service
    participant K as Kafka
    participant ES as Event Store (via mongo sink connector)

    Note over API, ES: Order Creation Saga - Happy Path
    
    API->>K: PublishCommand(CreateOrder)
    K->>US: CreateOrder Command
    
    US->>ES: StoreEvent(UserValidationRequested)
    US->>US: Validate User & Credit
    
    alt User Valid
        US->>ES: StoreEvent(UserValidated)
        US->>K: PublishEvent(UserValidated)
        K->>PS: UserValidated Event
        
        PS->>ES: StoreEvent(ProductReservationRequested)
        PS->>PS: Reserve Product Inventory
        
        alt Product Available
            PS->>ES: StoreEvent(ProductReserved)
            PS->>K: PublishEvent(ProductReserved)
            K->>PayS: ProductReserved Event
            
            PayS->>ES: StoreEvent(PaymentRequested)
            PayS->>PayS: Process Payment
            
            alt Payment Success
                PayS->>ES: StoreEvent(PaymentCompleted)
                PayS->>K: PublishEvent(PaymentCompleted)
                
                K->>US: PaymentCompleted Event
                K->>PS: PaymentCompleted Event
                
                US->>ES: StoreEvent(OrderConfirmed)
                PS->>ES: StoreEvent(ProductSold)
                
                Note over US, PS: Saga Completed Successfully
                
            else Payment Failed
                Note over PayS, ES: Compensation Flow - Payment Failed
                PayS->>ES: StoreEvent(PaymentFailed)
                PayS->>K: PublishEvent(PaymentFailed)
                
                K->>PS: PaymentFailed Event
                PS->>ES: StoreEvent(ProductReservationCancelled)
                PS->>PS: Release Product Inventory
                PS->>K: PublishEvent(ProductUnreserved)
                
                K->>US: ProductUnreserved Event
                US->>ES: StoreEvent(OrderCancelled)
            end
            
        else Product Unavailable
            Note over PS, ES: Compensation Flow - Product Unavailable
            PS->>ES: StoreEvent(ProductReservationFailed)
            PS->>K: PublishEvent(ProductReservationFailed)
            
            K->>US: ProductReservationFailed Event
            US->>ES: StoreEvent(OrderCancelled)
        end
        
    else User Invalid
        Note over US, ES: Compensation Flow - User Invalid
        US->>ES: StoreEvent(UserValidationFailed)
        US->>K: PublishEvent(UserValidationFailed)
        
        Note over US: Saga Terminated - No further steps needed
    end
```