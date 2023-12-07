import { Injectable } from '@nestjs/common';
import { MessagesRepository } from './messages.repository';

@Injectable()
export class MessagesService {
  messagesRepo: MessagesRepository;

  constructor() {
    this.messagesRepo = new MessagesRepository();
  }

  findOne(id: string) {
    return this.messagesRepo.findOne(id);
  }
  findAll() {
    return this.messagesRepo.findAll();
  }
  create(contents: string) {
    return this.messagesRepo.create(contents);
  }
}
