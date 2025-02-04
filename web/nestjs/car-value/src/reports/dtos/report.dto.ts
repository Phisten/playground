import { Expose, Transform } from 'class-transformer';

export class ReportDTO {
  @Expose()
  id: number;

  @Expose()
  approved: boolean;

  @Expose()
  price: number;

  @Expose()
  make: string;

  @Expose()
  lng: number;

  @Expose()
  lat: number;

  @Expose()
  mileage: number;

  @Expose()
  year: number;

  @Expose()
  model: string;

  @Transform(({ obj }) => obj.user.id)
  @Expose()
  userId: number;
}
