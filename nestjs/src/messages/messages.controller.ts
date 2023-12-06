import { Body, Controller, Get, Post, Req } from '@nestjs/common';

@Controller('messages')
export class MessagesController {
  @Get()
  listMessages() {
    return 'list';
  }
  @Post()
  createMessages(@Body('content') content: string) {
    console.log(content);
    return 'create: ' + content;
  }
  @Get('/:id')
  getMessage() {
    return '1000';
  }
}
