
// Create databases & collections
print("Creating database notifications & collections...");
db = db.getSiblingDB('notification');
db.createUser({
    user: "mongo",
    pwd: "1234",
    roles: [{ role: "readWrite", db: "notification" }]
});
db.createCollection('notification_templates');

print("Creating database commerce & collections...");
db = db.getSiblingDB('commerce');
db.createUser({
    user: "mongo",
    pwd: "1234",
    roles: [{ role: "readWrite", db: "commerce" }]
});
db.createCollection('carts');
db.createCollection('wishlists');

print("Creating database commerce & collections...");
db = db.getSiblingDB('event');
db.createUser({
    user: "mongo",
    pwd: "1234",
    roles: [{ role: "readWrite", db: "commerce" }]
});
db.createCollection('user-event-stores');
db.createCollection('notification-event-stores');
db.createCollection('payment-event-stores');
db.createCollection('product-event-stores');
db.createCollection('shipping-event-stores');
db.createCollection('commerce-event-stores');
