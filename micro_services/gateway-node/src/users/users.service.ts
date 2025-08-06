import { Injectable } from '@nestjs/common';


@Injectable()
export class UserService {
    joinUserToChat(username: string): string {
        // TODO: check if username is available (not present in cache) 
        // TODO: add user to cache, maybe an map with username and room_id fields
        return `user ${username} joined chat`;
    }

    sendMessageToChat(username: string): string {
        // TODO: check for user in cache
        // TODO: send message to room_id trough chat service
        return `user ${username} message sent`;
    }

    leaveChat(username: string): string {
        // TODO: close chat socket
        // TODO: remove username from cache to make it available
        return `user ${username} has left the chat`;
    }
}
