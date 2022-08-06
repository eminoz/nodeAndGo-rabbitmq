//RabbitMQ
const amqp = require("amqplib");

const rabbitUrl = "amqp://eminoz:eminoz@localhost:5672";
module.exports = class RabbitMQ {
  sendRabbitMQ = async (queueName, data) => {
    try {
      const connection = await amqp.connect(rabbitUrl);
      const channel = await connection.createChannel();
      await channel.assertQueue(queueName);
      await channel.sendToQueue(queueName, Buffer.from(JSON.stringify(data)));
      setTimeout(() => {
        connection.close();
      }, 2000);
    } catch (error) {
      console.log(error);
    }
  };
};
