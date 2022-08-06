module.exports = class UserRepository {
  constructor(database) {
    this.User = database;
  }

  async CreateUser({ name, email, password }) {
    const user = new this.User({ name, email, password });
    const responseUser = await user.save();
    return responseUser;
  }
  GetUserByEmail = async ({ email }) => {
    return await this.User.find({ email: email });
  };
  GetUserById = async ({ id }) => {
    return await this.User.find({ _id: id }).exec();
  };
  DeleteUserById = async ({ id }) => {
    return await this.User.deleteOne({ _id: id });
  };
};
