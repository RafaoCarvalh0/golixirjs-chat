import { v4 as uuidv4 } from 'uuid';

export class User {
    constructor(
        public id: string,
        public username: string,
        public roomId: string
    ) { }

    static new(username: string): User {
        const id = uuidv4();
        const roomId = uuidv4();
        return new User(id, username, roomId);
    }
} 