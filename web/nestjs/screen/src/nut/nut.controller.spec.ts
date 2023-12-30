import { Test, TestingModule } from '@nestjs/testing';
import { NutController } from './nut.controller';

describe('NutController', () => {
  let controller: NutController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [NutController],
    }).compile();

    controller = module.get<NutController>(NutController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
