pub(crate) const SHIPPING_PROVIDERS_SCHEMA: &str = r#"
    {
      "$schema": "http://json-schema.org/draft-07/schema#",
      "title": "ShippingProvider",
      "type": "object",
      "properties": {
        "id": { "type": "string" },
        "name": { "type": "string" },
        "created_at": { "type": "string", "format": "date-time" },
        "updated_at": { "type": "string", "format": "date-time" },
        "discarded_at": {
          "type": ["string", "null"],
          "format": "date-time"
        }
      },
      "required": ["id", "name", "created_at", "updated_at"]
    }
    "#;
