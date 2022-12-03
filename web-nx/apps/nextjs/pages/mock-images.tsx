import axios from 'axios';
import Image from 'next/image';
import { Fragment } from 'react';
import useSWR from 'swr';

/* eslint-disable-next-line */
export interface MockImagesProps {}

type Data = typeof DataMock[];
const DataMock = {
  author: 'Alejandro Escamilla',
  download_url: 'https://picsum.photos/id/0/5000/3333',
  height: 3333,
  id: '0',
  url: 'https://unsplash.com/photos/yC-Yzbqy7PY',
  width: 5000,
};

export function MockImages(props: MockImagesProps) {
  const { data, error } = useSWR<Data>('https://picsum.photos/v2/list', (url) =>
    axios.get(url).then((res) => res.data)
  );

  console.log({ data });

  return (
    <div>
      <h1 className="text-2xl">Welcome to MockImages!</h1>
      <div className="grid grid-cols-3 mx-auto gap-4 max-w-[1232px]">
        {data?.map((image) => {
          return (
            <Fragment key={`${image.id}`}>
              <div className="">
                <Image
                  src={image.download_url}
                  unoptimized
                  alt=""
                  width={400}
                  height={200}
                />
              </div>
            </Fragment>
          );
        })}
      </div>
    </div>
  );
}

export default MockImages;
