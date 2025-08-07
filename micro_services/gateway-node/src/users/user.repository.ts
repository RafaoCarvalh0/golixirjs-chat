import { User } from './user.entity';

export interface UserRepository {
    createUser(user: User): Promise<void>;
    deleteUser(userId: string): Promise<void>;
    getUserByUsername(username: string): Promise<User>;
}