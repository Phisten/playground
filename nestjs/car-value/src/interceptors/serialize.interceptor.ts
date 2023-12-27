import { CallHandler, ExecutionContext, NestInterceptor } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { Observable, map } from 'rxjs';

export class SerializeInterceptor implements NestInterceptor {
  /**
   *
   */
  constructor(private dto: any) {}

  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    // console.log('Im running before handler', context);

    return next.handle().pipe(
      map((data: any) => {
        console.log('run something before the response is sent out', data);
        return plainToInstance(this.dto, data, {
          excludeExtraneousValues: true,
        });
      }),
    );
  }
}
