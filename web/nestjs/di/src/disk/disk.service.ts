import { Injectable } from '@nestjs/common';
import { PowerService } from 'src/power/power.service';

@Injectable()
export class DiskService {
  constructor(private powerService: PowerService) {}

  getDate() {
    console.log('Drawing 3 watts of power from Power Service');
    this.powerService.supplyPower(3);
    return 'data';
  }
}
