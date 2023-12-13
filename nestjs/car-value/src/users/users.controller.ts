import {
  Body,
  Controller,
  Get,
  NotFoundException,
  Param,
  Post,
  Query,
} from '@nestjs/common';
import { CreateUserDTO } from './dtos/create-user.dto';
import { UsersService } from './users.service';

@Controller('auth')
export class UsersController {
  constructor(private userService: UsersService) {}

  @Post('/signup')
  createUser(@Body() body: CreateUserDTO) {
    this.userService.createUser(body);
    console.log(body);
  }

  @Get('/:id')
  getUser(@Param('id') id: number) {
    return JSON.stringify({
      id,
    });
  }
  @Get()
  getUserByEmail(@Query('email') email: string) {
    if (email) {
      return JSON.stringify({
        email,
      });
    }
    throw new NotFoundException();
  }
}
