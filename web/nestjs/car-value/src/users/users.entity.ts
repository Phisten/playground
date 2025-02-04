import { Report } from 'src/reports/reports.entity';
import {
  AfterInsert,
  AfterUpdate,
  BeforeRemove,
  Column,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  email: string;

  @Column()
  pwd: string;

  @OneToMany(() => Report, (report) => report.user)
  reports: Report[];

  @AfterInsert()
  logInsert() {
    console.log('user insert: ' + this.id);
  }
  @AfterUpdate()
  logUpdate() {
    console.log('user update: ' + this.id);
  }
  @BeforeRemove()
  logRemoveBF() {
    console.log('user remove: ' + this.id);
  }
}
