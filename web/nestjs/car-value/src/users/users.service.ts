import { Injectable, NotFoundException } from '@nestjs/common';
import { Repository } from 'typeorm';
import { User } from './users.entity';
import { InjectRepository } from '@nestjs/typeorm';
import { CreateUserDTO } from './dtos/create-user.dto';

@Injectable()
export class UsersService {
  constructor(@InjectRepository(User) private repo: Repository<User>) {}

  createUser(data: CreateUserDTO) {
    const user = this.repo.create(data);
    return this.repo.save(user);
  }

  findOne(id: number) {
    if (!id) {
      return null;
    }
    return this.repo.findOneBy({ id });
  }
  find(email: string) {
    return this.repo.findBy({ email });
  }
  finAll() {
    return this.repo.find();
  }
  async update(id: number, attr: Partial<User>) {
    const user = await this.repo.findOneBy({ id });

    if (!user) {
      throw new NotFoundException('user not found');
    }
    Object.assign(user, attr);

    return this.repo.save(user);
  }
  async remove(id: number) {
    const user = await this.repo.findOneBy({ id });

    if (!user) {
      throw Error('user not found');
    }

    return this.repo.remove(user);
  }
}
