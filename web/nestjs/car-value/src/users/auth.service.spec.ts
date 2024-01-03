import { Test } from '@nestjs/testing';
import { AuthService } from './auth.service';
import { UsersService } from './users.service';
import { User } from './users.entity';
import { CreateUserDTO } from './dtos/create-user.dto';

describe('AuthService', () => {
  let service: AuthService;

  beforeEach(async () => {
    // create a fake copy of the user service
    const fakeUsersService: Partial<UsersService> = {
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
    const [salt, hash] = user.pwd.split('.');
    expect(salt).toBeDefined();
    expect(hash).toBeDefined();
  });
});
