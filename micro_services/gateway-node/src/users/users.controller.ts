import { Body, Controller, Post } from '@nestjs/common';
import { UserService } from './users.service';
import { UserDto } from './dto/user-dto.dto';


@Controller('chat')
export class UserController {
    constructor(private readonly userService: UserService) { }

    @Post('join')
    joinChat(@Body() dto: UserDto): string {
        return this.userService.joinUserToChat(dto.username)
    }

    @Post('send-message')
    sendMessage(@Body() dto: UserDto): string {
        return this.userService.sendMessageToChat(dto.username)
    }

    @Post('terminate')
    terminateChat(@Body() dto: UserDto): string {
        return this.userService.leaveChat(dto.username)
    }
}