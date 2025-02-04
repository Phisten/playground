import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateReportDTO } from './dtos/create-report.dto';
import { Report } from './reports.entity';
import { User } from 'src/users/users.entity';

@Injectable()
export class ReportsService {
  constructor(@InjectRepository(Report) private repo: Repository<Report>) {}

  async onModuleInit() {}

  create(reportDto: CreateReportDTO, user: User) {
    const report = this.repo.create({ ...reportDto, user });

    return this.repo.save(report);
  }
}
