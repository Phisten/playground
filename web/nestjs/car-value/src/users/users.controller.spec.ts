import { Test, TestingModule } from '@nestjs/testing';
import { UsersController } from './users.controller';
import { UsersService } from './users.service';
import { AuthService } from './auth.service';
import { User } from './users.entity';
import { NotFoundException } from '@nestjs/common';

describe('UsersController', () => {
  let controller: UsersController;
  let fakeUsersService: Partial<UsersService>;
  let fakeAuthService: Partial<AuthService>;

  beforeEach(async () => {
    fakeUsersService = {
      findOne: (id: number) => {
        return Promise.resolve({
          id,
          email: 'asdf@asdf.com',
          password: 'asdf',
        } as unknown as User);
      },
      find: (email: string) => {
        return Promise.resolve([
          { id: 1, email, password: 'asdf' } as unknown as User,
        ]);
      },
      // remove: () => {},
      // update: () => {},
    };
    fakeAuthService = {
      signup: (email: string, password: string) =>
        Promise.resolve({ id: 1, email, password } as unknown as User),
      signin: (email: string, password: string) =>
        Promise.resolve({ id: 1, email, password } as unknown as User),
    };

    const module: TestingModule = await Test.createTestingModule({
      providers: [
        {
          provide: UsersService,
          useValue: fakeUsersService,
        },
        {
          provide: AuthService,
          useValue: fakeAuthService,
        },
      ],
      controllers: [UsersController],
    }).compile();

    controller = module.get<UsersController>(UsersController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  it('find user returns a single user with given id', async () => {
    const user = await controller.getUser(1);
    expect(user).toBeDefined();
  });

  it('findUser throws an error if user with given id is not found', async () => {
    fakeUsersService.findOne = () => null;
    await expect(controller.getUser(1)).rejects.toThrow(NotFoundException);
  });

  it('signin update session object and returns user', async () => {
    const session = {
      userId: undefined,
    };
    const user = await controller.signin(
      {
        email: 'asd@123',
        pwd: 'asdf',
      },
      session,
    );

    expect(user.id).toEqual(1);
    expect(session.userId).toEqual(1);
  });
});
