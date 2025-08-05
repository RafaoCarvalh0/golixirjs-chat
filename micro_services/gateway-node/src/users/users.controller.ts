import { Controller, Post } from '@nestjs/common';


@Controller('join')
export class UserController {
    @Post()
    joinUserIntoChat(): string {
        return 'user joined chat';
    }
}