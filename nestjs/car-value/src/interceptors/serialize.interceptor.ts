import { CallHandler, ExecutionContext, NestInterceptor } from '@nestjs/common';
import { plainToClass } from 'class-transformer';
import { Observable, map } from 'rxjs';
import { UserDto } from 'src/users/dtos/user.dto';

export class SerializeInterceptor implements NestInterceptor {
  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    // console.log('Im running before handler', context);

    return next.handle().pipe(
      map((data: any) => {
        console.log('run something before the response is sent out', data);
        return plainToClass(UserDto, data, { excludeExtraneousValues: true });
      }),
    );
  }
}
