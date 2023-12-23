import { Module } from '@nestjs/common';
import { NutController } from './nut.controller';
import { NutService } from './nut.service';

@Module({
  controllers: [NutController],
  providers: [NutService],
})
export class NutModule {}
