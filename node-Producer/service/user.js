const { ErrorResult, SuccessDataResult } = require("clean-response");
module.exports = class UserService {
  constructor(repo, broker) {
    this.userRepo = repo;
    this.broker = broker;
  }

  CreateUser = async ({ name, email, password }) => {
    const user = await this.userRepo.GetUserByEmail({ email });

    // if (user.length !== 0) {
    //   return new ErrorResult(400).dataResult();
    // }
    const responseUSer = await this.userRepo.CreateUser({
      name,
      email,
      password,
    });
    this.broker.sendRabbitMQ("newuser", { name, email });//this send mail to message broker to send verification mail
    this.broker.sendRabbitMQ("mailqueue", { name, email });//this send mail to message broker to send verification mail
    return new SuccessDataResult(200, responseUSer).dataResult();
  };
  GetUserById = async ({ id }) => {
    const data = await this.userRepo.GetUserById({ id });
    return data.map((res) => ({
      id: res._id,
      name: res.name,
    }));
  };
  DeleteUserWithID = async ({ id }) => {
    const response = await this.userRepo.DeleteUserById({ id });
    return response;
  };
};
