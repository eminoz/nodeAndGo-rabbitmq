const User = require("../model/user");
const UserRepository = require("../repository/user");
const UserService = require("../service/user");
const express = require("express");
const UserApi = require("../api/user");
const UserRepo = new UserRepository(User);
const UserSer = new UserService(UserRepo);
const UserController = new UserApi(UserSer);

const router = express.Router();
router.post("/createuser", UserController.CreateUser);
router.get("/getuserbyid/:id", UserController.GetUserById);
router.delete("/deleteuserbyid/:id", UserController.DeleteUser);
module.exports = router;