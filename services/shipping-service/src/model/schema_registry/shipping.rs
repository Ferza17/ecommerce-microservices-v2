pub(crate) const SHIPPING: &str = r#"
    {
      "$schema": "http://json-schema.org/draft-07/schema#",
      "title": "Shippings",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "user_id": {
          "type": "string"
        },
        "payment_id": {
          "type": "string"
        },
        "shipping_provider_id": {
          "type": "string"
        },
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        },
        "discarded_at": {
          "type": ["string", "null"],
          "format": "date-time"
        }
      },
      "required": [
        "id",
        "user_id",
        "payment_id",
        "shipping_provider_id",
        "created_at",
        "updated_at"
      ]
    }
"#;
