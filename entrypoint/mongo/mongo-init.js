
// Create databases & collections NOTIFICATION
print("Creating database notifications & collections...");
db = db.getSiblingDB('notification');
db.createUser({
    user: "mongo",
    pwd: "1234",
    roles: [{ role: "readWrite", db: "notification" }]
});
db.createCollection('notification_templates');
db.createCollection('notification_user_logs');

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

// Create databases & collections OUTBOX
print("Creating database commerce & collections...");
db = db.getSiblingDB('outbox');
db.createUser({
    user: "mongo",
    pwd: "1234",
    roles: [{ role: "readWrite", db: "outbox" }]
});
db.createCollection('event_envelopes');
