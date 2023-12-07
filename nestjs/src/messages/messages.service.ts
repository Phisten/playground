import { Injectable, NotFoundException } from '@nestjs/common';
import { MessagesRepository } from './messages.repository';

@Injectable()
export class MessagesService {
  messagesRepo: MessagesRepository;

  constructor() {
    this.messagesRepo = new MessagesRepository();
  }

  async findOne(id: string) {
    const message = await this.messagesRepo.findOne(id);

    if (!message) {
      throw new NotFoundException('message not found');
    }
    return message;
  }
  async findAll() {
    const message = await this.messagesRepo.findAll();

    return message;
  }
  async create(contents: string) {
    const messages = await this.messagesRepo.create(contents);

    return messages;
  }
}
