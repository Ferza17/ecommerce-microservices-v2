
// Create databases & collections NOTIFICATION
print("Creating database notifications & collections...");
db = db.getSiblingDB('notification');
db.createUser({
    user: "mongo",
    pwd: "1234",
    roles: [{ role: "readWrite", db: "notification" }]
});
db.createCollection('notification_templates')

// Create databases & collections COMMERCE
print("Creating database commerce & collections...");
db = db.getSiblingDB('commerce');
db.createUser({
    user: "mongo",
    pwd: "1234",
    roles: [{ role: "readWrite", db: "commerce" }]
});
db.createCollection('carts');
db.createCollection('wishlists');

// Create databases & collections EVENT
print("Creating database commerce & collections...");
db = db.getSiblingDB('event');
db.createUser({
    user: "mongo",
    pwd: "1234",
    roles: [{ role: "readWrite", db: "event" }]
});
db.createCollection('user_event_stores');
db.createCollection('notification_event_stores');
db.createCollection('payment_event_stores');
db.createCollection('product_event_stores');
db.createCollection('shipping_event_stores');
db.createCollection('commerce_event_stores');
