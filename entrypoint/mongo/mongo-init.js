// INIT DATABASE NOTIFICATION
db = db.getSiblingDB('notification');
db.createCollection('notification_templates');

// INIT DATABASE COMMERCE
db = db.getSiblingDB('commerce');
db.createCollection('carts');
db.createCollection('wishlists');
