db = db.getSiblingDB('admin');

// Ensure root user exists
if (db.system.users.find({ user: 'root' }).count() === 0) {
  db.Create({
    user: 'root',
    pwd: 'password',
    roles: [{ role: 'root', db: 'admin' }]
  });
}

// Create 'billbuddies' database and user
db = db.getSiblingDB('billbuddies');

if (db.system.users.find({ user: 'billbuddies_user' }).count() === 0) {
  db.createUser({
    user: 'billbuddies_user',
    pwd: 'billbuddies_password',
    roles: [{ role: 'readWrite', db: 'billbuddies' }]
  });
}

// Create a dummy collection and document to ensure DB exists
db.createCollection('init');
db.init.insertOne({ createdAt: new Date() });
db.init.deleteMany({});