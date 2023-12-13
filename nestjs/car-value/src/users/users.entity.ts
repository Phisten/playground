import {
  AfterInsert,
  AfterRemove,
  AfterUpdate,
  Column,
  Entity,
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

  @AfterInsert()
  logInsert() {
    console.log('user insert: ' + this.id);
  }
  @AfterUpdate()
  logUpdate() {
    console.log('user update: ' + this.id);
  }
  @AfterRemove()
  logRemove() {
    console.log('user remove: ' + this.id);
  }
}
