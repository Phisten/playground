import { Injectable } from '@nestjs/common';
import { mouse, left, right, up, down } from '@nut-tree/nut-js';

@Injectable()
export class NutService {
  demoMouseMove() {
    (async () => {
      await mouse.move(left(500));
      await mouse.move(up(500));
      await mouse.move(right(500));
      await mouse.move(down(500));
    })();
  }
}
