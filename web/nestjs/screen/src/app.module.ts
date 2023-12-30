import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { NutController } from './nut/nut.controller';
import { NutService } from './nut/nut.service';
import { NutModule } from './nut/nut.module';

@Module({
  imports: [NutModule],
  controllers: [AppController, NutController],
  providers: [AppService, NutService],
})
export class AppModule {}
