import { Test } from '@nestjs/testing';
import { AuthService } from './auth.service';
import { UsersService } from './users.service';

it('can create an instance of auth service', async () => {
  // create a fake copy of the user service
  const fakeUsersService = {
    createUser: (email: string, password: string) =>
      Promise.resolve({ id: 1, email, password }),
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

  const service = module.get(AuthService);
  expect(service).toBeDefined();
});
