const mongoose = require('mongoose');
const config  = require('../config/index');

module.exports = async () => {
    try {
        await mongoose.connect(config.DB_URL, {
            useNewUrlParser: true,
            useUnifiedTopology: true,
        })
        console.log('Db Connected');
    } catch (error) {
        console.log('Error ============')
        console.log(error);
        process.exit(1);
    }
}