import { Body, Controller, Post } from '@nestjs/common';
import { CreateUserDTO } from './dtos/create-user.dto';

@Controller('auth')
export class UsersController {
  @Post('/signup')
  createUser(@Body() body: CreateUserDTO) {
    console.log(body);
  }
}
