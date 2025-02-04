import { User } from 'src/users/users.entity';
import { Column, Entity, ManyToOne, PrimaryGeneratedColumn } from 'typeorm';

@Entity()
export class Report {
  @PrimaryGeneratedColumn()
  id: number;

  @Column({ default: false })
  approved: boolean;

  @Column()
  price: number;

  @Column()
  make: string;

  @Column()
  lng: number;

  @Column()
  lat: number;

  @Column()
  mileage: number;

  @Column()
  year: number;

  @Column()
  model: string;

  @ManyToOne(() => User, (user) => user.reports)
  user: User;
}
