
// Create databases & collections
print("Creating database notifications & collections...");
db = db.getSiblingDB('notification');
db.createCollection('notification_templates');

print("Creating database commerce & collections...");
db = db.getSiblingDB('commerce');
db.createCollection('carts');
db.createCollection('wishlists');
