import { Controller, Get } from '@nestjs/common';
import { NutService } from './nut.service';

@Controller('nut')
export class NutController {
  constructor(private readonly nutService: NutService) {}

  @Get()
  demo() {
    this.nutService.demoMouseMove();
    return 'demo mouse move';
  }
}
