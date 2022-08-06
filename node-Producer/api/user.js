module.exports = class UserApi {
  constructor(service) {
    this.userService = service;
  }

  CreateUser = async (req, res, next) => {
    const { name, email, password } = req.body;
    const response = await this.userService.CreateUser({
      name,
      email,
      password,
    });
    return res.json(response);
  };
  GetUserById = async (req, res, next) => {
    const id = req.params.id;
    const response = await this.userService.GetUserById({ id });
    return res.json({ user: response });
  };
  DeleteUser = async (req, res, next) => {
    const id = req.params.id;
    const response = await this.userService.DeleteUserWithID({ id });
    return res.json({ message: response });
  };
};
