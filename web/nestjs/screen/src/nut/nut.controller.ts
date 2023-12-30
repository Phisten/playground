import { Controller, Get } from '@nestjs/common';
import { NutService } from './nut.service';

@Controller('nut')
export class NutController {
  constructor(private readonly nutService: NutService) {}

  @Get('mouse')
  demoMouse() {
    console.log('demo mouse');
    this.nutService.demoMouseMove();
    return 'demo mouse move';
  }

  @Get()
  demoSS() {
    console.log('demoSS');
    return this.nutService.demoSS();
  }
}
