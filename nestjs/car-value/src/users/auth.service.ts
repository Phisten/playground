import {
  BadRequestException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { UsersService } from './users.service';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AuthService {
  constructor(private usersService: UsersService) {}

  async signup(email: string, password: string) {
    // check user email
    const users = await this.usersService.find(email);
    if (users.length > 0) {
      throw new BadRequestException('email is use');
    }

    // get salt
    const salt = await bcrypt.genSalt();
    // hash password
    const hash = await bcrypt.hash(password, salt);
    console.log('signup: ', { salt, hash });
    // const isMatch = await bcrypt.compare(password, hash);

    // create user
    const createdUser = this.usersService.createUser({
      email,
      pwd: hash,
    });

    // return user
    return createdUser;
  }

  async signin(email: string, password: string) {
    // check user email
    const [user] = await this.usersService.find(email);

    if (!user) {
      throw new NotFoundException('email not found');
    }

    console.log('signin', { users: user });

    // hash password
    const isMatch = await bcrypt.compare(password, user?.[0].pwd);

    // return user
    return isMatch;
  }
}
