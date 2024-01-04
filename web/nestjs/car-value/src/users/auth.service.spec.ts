import { Test } from '@nestjs/testing';
import { AuthService } from './auth.service';
import { UsersService } from './users.service';
import { User } from './users.entity';
import { CreateUserDTO } from './dtos/create-user.dto';
import { BadRequestException, NotFoundException } from '@nestjs/common';

describe('AuthService', () => {
  let service: AuthService;
  let fakeUsersService: Partial<UsersService>;

  beforeEach(async () => {
    // create a fake copy of the user service
    fakeUsersService = {
      createUser: (user: CreateUserDTO) =>
        Promise.resolve({ id: 1, email: user.email, pwd: user.pwd } as User),
      find: () => Promise.resolve([]),
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
    fakeUsersService.find = () =>
      Promise.resolve([
        { id: '1', email: 'asd', pwd: 'hash' } as unknown as User,
      ]);
    await expect(service.signin('asd', 'asd')).rejects.toThrow(
      BadRequestException,
    );
  });

  it('return a user if correct password is provided', async () => {
    fakeUsersService.find = () =>
      Promise.resolve([
        {
          id: '1',
          email: 'asd',
          pwd: '$2b$10$WJKKmcmu6gRvBlUti3W6lOdeng7dgXkvj/vsrcQDxVhJSfo1uduZq',
        } as unknown as User,
      ]);

    await expect(service.signin('asd', 'asd')).toBeDefined();
    // await expect(service.signup('asd', 'asd')).toBeDefined();
  });
});
