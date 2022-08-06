module.exports= class UserService {
    constructor(repo) {
        this.userRepo = repo;
    }

    CreateUser = async ({name, email, password}) => {
        const responseUSer = await this.userRepo.CreateUser({
            name,
            email,
            password,
        });
        return responseUSer;
    };
    GetUserById = async ({id}) => {

        const data = await this.userRepo.GetUserById({id})
        return data.map(res => ({
            id: res._id,
            name: res.name
        }))
    } 
    DeleteUserWithID = async ({id}) => {
        const response = await this.userRepo.DeleteUserById({id})
        return response
    }

}