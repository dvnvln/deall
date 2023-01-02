// Create user
dbAdmin = db.getSiblingDB("admin");
dbAdmin.createUser({
  user: "intv",
  pwd: "pwd",
  roles: [{ role: "userAdminAnyDatabase", db: "admin" }],
  mechanisms: ["SCRAM-SHA-1"],
});

// Authenticate user
dbAdmin.auth({
  user: "intv",
  pwd: "pwd",
  mechanisms: ["SCRAM-SHA-1"],
  digestPassword: true,
});

// Create DB and collection
db = new Mongo().getDB("dealls");
db.createCollection("user", { capped: false });
db.user.createIndex( { "username": 1 }, { unique: true } )