import { Injectable } from '@nestjs/common';
import { getBuildInformation } from '@techstark/opencv-js';

@Injectable()
export class AppService {
  getHello(): string {
    return getBuildInformation();
  }
}
