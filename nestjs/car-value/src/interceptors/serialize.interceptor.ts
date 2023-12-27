import { CallHandler, ExecutionContext, NestInterceptor } from '@nestjs/common';
import { Observable, map } from 'rxjs';

export class SerializeInterceptor implements NestInterceptor {
  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    console.log('Im running before handler', context);

    return next.handle().pipe(
      map((data: any) => {
        // run something before the response is sent out
        console.log('run something before the response is sent out', data);
      }),
    );
  }
}
