import { IsEmail, IsString } from 'class-validator';

export class CreateUserDTO {
  @IsEmail({}, { message: 'Invalid email format' })
  email: string;

  @IsString()
  pwd: string;
}
