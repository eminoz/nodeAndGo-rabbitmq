const dotenv = require("dotenv");
dotenv.config({
    path: "./pkg/config/env/config.env",
})

module.exports = {

    PORT: process.env.PORT,
    DB_URL: process.env.MONGODB_URI,
    APP_SECRET: process.env.APP_SECRET
}