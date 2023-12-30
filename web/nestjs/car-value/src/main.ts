import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';

// eslint-disable-next-line @typescript-eslint/no-var-requires
const cookeSession = require('cookie-session');
// import { CookieSession } from 'cookie-session';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.use(
    cookeSession({
      keys: ['asd'],
    }),
  );
  app.useGlobalPipes(new ValidationPipe());

  await app.listen(3500);
}
bootstrap();
