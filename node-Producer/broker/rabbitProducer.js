//RabbitMQ
const amqp = require("amqplib");

const rabbitUrl = "amqp://eminoz:eminoz@localhost:5672";

const sendRabbitMQ = async (queueName, data) => {
  try {
    const connection = await amqp.connect(rabbitUrl);
    const channel = await connection.createChannel();
    const assertion = await channel.assertQueue(queueName);
    channel.sendToQueue(queueName, Buffer.from(JSON.stringify(data)));
  } catch (error) {
    console.log(error);
  }
};
module.exports = sendRabbitMQ;
