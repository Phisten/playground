import { Test } from '@nestjs/testing';
import { AuthService } from './auth.service';
import { UsersService } from './users.service';
import { User } from './users.entity';
import { CreateUserDTO } from './dtos/create-user.dto';
import { BadRequestException, NotFoundException } from '@nestjs/common';

describe('AuthService', () => {
  let service: AuthService;
  let fakeUsersService: Partial<UsersService>;
  const users: User[] = [];

  beforeEach(async () => {
    // create a fake copy of the user service
    fakeUsersService = {
      find: (email: string) => {
        const filtered = users.filter((user) => user.email == email);
        return Promise.resolve(filtered);
      },
      createUser: (user: CreateUserDTO) => {
        const newUser = {
          id: Math.floor(Math.random() * 99999),
          email: user.email,
          pwd: user.pwd,
        } as User;
        users.push(newUser);
        return Promise.resolve(newUser);
      },
    };

    const module = await Test.createTestingModule({
      providers: [
        AuthService,
        {
          provide: UsersService,
          useValue: fakeUsersService,
        },
      ],
    }).compile();

    service = module.get(AuthService);
  });

  it('can create an instance of auth service', async () => {
    expect(service).toBeDefined();
  });

  it('create a new user with a salted and hashed password', async () => {
    const user = await service.signup('asdf@asd.com', 'asdddd');

    expect(user.pwd).not.toEqual('asdddd');
    // const [salt, hash] = user.pwd.split('.');
    // expect(salt).toBeDefined();
    // expect(hash).toBeDefined();
  });

  it('throws an error if user signs up with email that is in use', async () => {
    fakeUsersService.find = async () =>
      Promise.resolve([
        { id: '1', email: '123', pwd: 'hash' } as unknown as User,
      ]);
    await expect(service.signup('asdf@asd.com', 'asd')).rejects.toThrow(
      BadRequestException,
    );
  });
  it('throws if signin is called with an unused email', async () => {
    await expect(service.signin('asd', 'hash')).rejects.toThrow(
      NotFoundException,
    );
  });
  it('throws if signin is called with an incorrect password', async () => {
    await service.signup('asd2', 'asd1');
    await expect(service.signin('asd2', 'asd')).rejects.toThrow(
      BadRequestException,
    );
  });

  it('return a user if correct password is provided', async () => {
    await service.signup('asd', 'asd');
    await expect(service.signin('asd', 'asd')).toBeDefined();
  });
});
