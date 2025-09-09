-- +goose Up

INSERT INTO products (id, name, description, uom, image, price, stock, created_at, updated_at, discarded_at)
VALUES
-- 1
(gen_random_uuid(), 'Mountain Boots', 'Durable boots for rugged mountain terrain.', 'pair', 'https://picsum.photos/200?1', 129.99, 50, NOW(), NOW(), NULL),
-- 2
(gen_random_uuid(), 'Camping Tent', 'Lightweight and waterproof tent for two.', 'unit', 'https://picsum.photos/200?2', 199.49, 25, NOW(), NOW(), NULL),
-- 3
(gen_random_uuid(), 'Hiking Backpack', '35L backpack with hydration system.', 'unit', 'https://picsum.photos/200?3', 89.95, 100, NOW(), NOW(), NULL),
-- 4
(gen_random_uuid(), 'Thermal Jacket', 'Insulated jacket for sub-zero conditions.', 'unit', 'https://picsum.photos/200?4', 149.99, 70, NOW(), NOW(), NULL),
-- 5
(gen_random_uuid(), 'Sleeping Bag', 'Comfortable sleeping bag rated to -10°C.', 'unit', 'https://picsum.photos/200?5', 79.95, 80, NOW(), NOW(), NULL),
-- 6
(gen_random_uuid(), 'Water Bottle', '1L BPA-free reusable water bottle.', 'pcs', 'https://picsum.photos/200?6', 12.99, 200, NOW(), NOW(), NULL),
-- 7
(gen_random_uuid(), 'Trekking Poles', 'Adjustable poles with anti-shock system.', 'pair', 'https://picsum.photos/200?7', 39.95, 90, NOW(), NOW(), NULL),
-- 8
(gen_random_uuid(), 'Camping Stove', 'Portable gas stove with piezo ignition.', 'unit', 'https://picsum.photos/200?8', 54.90, 35, NOW(), NOW(), NULL),
-- 9
(gen_random_uuid(), 'Headlamp', 'USB rechargeable LED headlamp.', 'unit', 'https://picsum.photos/200?9', 24.50, 150, NOW(), NOW(), NULL),
-- 10
(gen_random_uuid(), 'Rain Poncho', 'Lightweight poncho for emergency rain.', 'pcs', 'https://picsum.photos/200?10', 9.99, 300, NOW(), NOW(), NULL),
-- 11
(gen_random_uuid(), 'Climbing Harness', 'Secure and comfortable climbing harness.', 'unit', 'https://picsum.photos/200?11', 69.95, 45, NOW(), NOW(), NULL),
-- 12
(gen_random_uuid(), 'Trail Shoes', 'Lightweight trail running shoes.', 'pair', 'https://picsum.photos/200?12', 99.00, 60, NOW(), NOW(), NULL),
-- 13
(gen_random_uuid(), 'Solar Charger', '10,000mAh solar power bank.', 'unit', 'https://picsum.photos/200?13', 39.99, 85, NOW(), NOW(), NULL),
-- 14
(gen_random_uuid(), 'Cooking Set', 'Compact camping cookware set.', 'set', 'https://picsum.photos/200?14', 29.99, 120, NOW(), NOW(), NULL),
-- 15
(gen_random_uuid(), 'Compass', 'Classic magnetic navigation compass.', 'unit', 'https://picsum.photos/200?15', 7.49, 200, NOW(), NOW(), NULL),
-- 16
(gen_random_uuid(), 'First Aid Kit', 'Essential first aid supplies for travel.', 'kit', 'https://picsum.photos/200?16', 25.00, 75, NOW(), NOW(), NULL),
-- 17
(gen_random_uuid(), 'Rope 20m', '20 meter nylon climbing rope.', 'unit', 'https://picsum.photos/200?17', 34.99, 50, NOW(), NOW(), NULL),
-- 18
(gen_random_uuid(), 'Portable Chair', 'Foldable camping chair with cup holder.', 'unit', 'https://picsum.photos/200?18', 22.99, 60, NOW(), NOW(), NULL),
-- 19
(gen_random_uuid(), 'Multi-tool Knife', 'Stainless steel multi-tool for outdoors.', 'unit', 'https://picsum.photos/200?19', 18.50, 140, NOW(), NOW(), NULL),
-- 20
(gen_random_uuid(), 'Camp Lantern', 'LED lantern with adjustable brightness.', 'unit', 'https://picsum.photos/200?20', 29.00, 90, NOW(), NOW(), NULL),
-- 21
(gen_random_uuid(), 'Dry Bag 10L', 'Waterproof dry bag with roll-top.', 'unit', 'https://picsum.photos/200?21', 14.95, 180, NOW(), NOW(), NULL),
-- 22
(gen_random_uuid(), 'Fire Starter', 'Magnesium fire starter stick.', 'unit', 'https://picsum.photos/200?22', 5.99, 250, NOW(), NOW(), NULL),
-- 23
(gen_random_uuid(), 'Mosquito Net', 'Portable mosquito net for camping.', 'unit', 'https://picsum.photos/200?23', 12.00, 130, NOW(), NOW(), NULL),
-- 24
(gen_random_uuid(), 'Climbing Helmet', 'Protective helmet for rock climbing.', 'unit', 'https://picsum.photos/200?24', 59.99, 40, NOW(), NOW(), NULL),
-- 25
(gen_random_uuid(), 'Binoculars', '10x50 field binoculars.', 'unit', 'https://picsum.photos/200?25', 89.95, 30, NOW(), NOW(), NULL),
-- 26
(gen_random_uuid(), 'Hammock', 'Lightweight nylon hammock with straps.', 'unit', 'https://picsum.photos/200?26', 27.99, 110, NOW(), NOW(), NULL),
-- 27
(gen_random_uuid(), 'Sunscreen SPF50', 'Water-resistant sunscreen 100ml.', 'tube', 'https://picsum.photos/200?27', 6.49, 300, NOW(), NOW(), NULL),
-- 28
(gen_random_uuid(), 'Insect Repellent', 'DEET-based bug spray 100ml.', 'bottle', 'https://picsum.photos/200?28', 7.99, 280, NOW(), NOW(), NULL),
-- 29
(gen_random_uuid(), 'Windbreaker Jacket', 'Packable wind-resistant jacket.', 'unit', 'https://picsum.photos/200?29', 44.95, 65, NOW(), NOW(), NULL),
-- 30
(gen_random_uuid(), 'Camp Kettle', 'Aluminum kettle for campfire use.', 'unit', 'https://picsum.photos/200?30', 21.99, 70, NOW(), NOW(), NULL),
-- 31-50
(gen_random_uuid(), 'Survival Blanket', 'Emergency thermal survival blanket.', 'pcs', 'https://picsum.photos/200?31', 3.49, 500, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Trekking Hat', 'UV-protective wide-brim hat.', 'unit', 'https://picsum.photos/200?32', 15.00, 150, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Snow Goggles', 'Anti-fog UV-protective snow goggles.', 'unit', 'https://picsum.photos/200?33', 49.90, 60, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Gaiters', 'Waterproof gaiters for hiking boots.', 'pair', 'https://picsum.photos/200?34', 17.99, 95, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Foam Mat', 'Roll-up foam sleeping mat.', 'unit', 'https://picsum.photos/200?35', 13.50, 80, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Camp Grill', 'Foldable stainless steel grill.', 'unit', 'https://picsum.photos/200?36', 32.50, 40, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Hydration Pack', '2L bladder hydration backpack.', 'unit', 'https://picsum.photos/200?37', 42.00, 60, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Ski Gloves', 'Insulated waterproof gloves.', 'pair', 'https://picsum.photos/200?38', 35.75, 75, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Camping Table', 'Collapsible aluminum table.', 'unit', 'https://picsum.photos/200?39', 49.99, 20, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Lantern Fuel', 'Fuel canister for lanterns.', 'can', 'https://picsum.photos/200?40', 8.25, 90, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Winter Socks', 'Thermal wool socks.', 'pair', 'https://picsum.photos/200?41', 9.99, 150, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Navigation Watch', 'GPS-enabled trekking watch.', 'unit', 'https://picsum.photos/200?42', 129.00, 30, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Snowshoes', 'Durable lightweight snowshoes.', 'pair', 'https://picsum.photos/200?43', 89.95, 25, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Portable Fan', 'USB-powered mini fan.', 'unit', 'https://picsum.photos/200?44', 14.99, 85, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Mini Tripod', 'Flexible leg tripod for cameras.', 'unit', 'https://picsum.photos/200?45', 10.50, 100, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Patch Kit', 'Tire and mat patch repair kit.', 'kit', 'https://picsum.photos/200?46', 5.99, 180, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Trekking Shorts', 'Quick-dry breathable shorts.', 'unit', 'https://picsum.photos/200?47', 27.99, 95, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Ice Axe', 'Lightweight ice axe for mountaineering.', 'unit', 'https://picsum.photos/200?48', 74.90, 20, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Gas Canister', 'Isobutane fuel 230g.', 'can', 'https://picsum.photos/200?49', 6.99, 100, NOW(), NOW(), NULL),
(gen_random_uuid(), 'Thermos Flask', 'Stainless steel 1L thermos.', 'unit', 'https://picsum.photos/200?50', 19.90, 120, NOW(), NOW(), NULL);
