import { Test } from '@nestjs/testing';
import { AuthService } from './auth.service';

it('can create an instance of auth service', async () => {
  const module = await Test.createTestingModule({
    providers: [AuthService],
  }).compile();
});
