import { Body, Controller, Get, Param, Post, Query } from '@nestjs/common';

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
  @Get(':id')
  getMessage(@Param('id') id: string) {
    return 'id: ' + id;
  }
}
