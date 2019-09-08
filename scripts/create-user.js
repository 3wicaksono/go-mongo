use gomongo;
db.createUser({
    user: "gomongo",
    pwd: "gomongo",
    roles: ["readWrite"]
});