import { Injectable, Inject } from '@nestjs/common';
import { USER_REPOSITORY } from './user.repository.token';
import { User } from './user.entity';

import type { UserRepository } from './user.repository';

@Injectable()
export class UserService {
    constructor(
        @Inject(USER_REPOSITORY)
        private userRepository: UserRepository
    ) { }

    async joinUserToChat(username: string): Promise<string> {
        const existingUser = await this.userRepository.getUserByUsername(username);
        if (existingUser) {
            throw `username already taken`;
        }

        const newUser = User.new(username)
        this.userRepository.createUser(newUser)

        // TODO: make user join a chat using the chat ms
        return `user ${username} joined chat`;
    }

    async sendMessageToChat(username: string): Promise<string> {
        // TODO: get user from cache and send message to roomId trough a chat micro service
        return `user ${username} message sent`;
    }

    async leaveChat(username: string): Promise<string> {
        // TODO: close chat socket
        const existingUser = await this.userRepository.getUserByUsername(username);
        if (existingUser) {
            this.userRepository.deleteUser(existingUser.id)
        }
        return `user ${username} has left the chat`;
    }

}


